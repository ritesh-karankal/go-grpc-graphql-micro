"""Gentle, realistic shopper traffic against the go-micro-shop GraphQL API.

Generates the mix a few real users would: browsing, searching, viewing accounts and
orders, signing up, placing orders, plus ~3% deliberately bad requests (unknown
account or product) so error rates and error logs have something to show.
Each simulated user waits 0.3-1.5 s between actions, so 4 users ~= 4-6 requests/s.

Usage:
    python3 scripts/loadgen.py <base_url> [duration_seconds=300] [users=4]

    python3 scripts/loadgen.py http://localhost:8000 60 2          # local docker compose
    python3 scripts/loadgen.py http://<ALB address> 300 4           # EKS

Requests carry "X-Synthetic: true", so the gateway records them as synthetic=true and
SLOs can exclude them. Creates real data: accounts and products named "lt-...". Stdlib only.
Prints a summary every 60 s and at the end.
"""
import json
import random
import sys
import threading
import time
import urllib.error
import urllib.request
from collections import Counter

if len(sys.argv) < 2:
    sys.exit(__doc__)

BASE = sys.argv[1].rstrip("/")
DURATION = int(sys.argv[2]) if len(sys.argv) > 2 else 300
USERS = int(sys.argv[3]) if len(sys.argv) > 3 else 4
GQL = BASE + "/graphql"

SEARCH_TERMS = ["wireless", "usb", "monitor", "laptop", "keyboard", "headphones", "ssd", "webcam", "stand", "hub"]
NEW_PRODUCTS = [
    ("lt-Desk Lamp", "LED desk lamp with dimmer", 29.99),
    ("lt-Mouse Pad XL", "Extended desk mat", 19.99),
    ("lt-Cable Organizer", "Magnetic cable clips, 6 pack", 9.99),
    ("lt-Laptop Stand", "Aluminium adjustable stand", 44.99),
]

stats = Counter()
latencies = {}
lock = threading.Lock()
known = {"products": [], "accounts": []}


def gql(query, variables=None, op="?"):
    """Send one GraphQL request and record its outcome and latency under op."""
    body = json.dumps({"query": query, "variables": variables or {}}).encode()
    req = urllib.request.Request(GQL, data=body, headers={
        "Content-Type": "application/json",
        "X-Synthetic": "true",  # lets SLOs exclude this traffic
    })
    start = time.time()
    try:
        with urllib.request.urlopen(req, timeout=15) as r:
            data = json.loads(r.read())
        ok = not data.get("errors")
    except (urllib.error.URLError, TimeoutError, ValueError):
        data, ok = {}, False
    elapsed = time.time() - start
    with lock:
        stats[op + (" ok" if ok else " err")] += 1
        latencies.setdefault(op, []).append(elapsed)
    return data.get("data") or {}


def load_ids():
    d = gql("{ products(pagination:{skip:0,take:50}) { id } accounts(pagination:{skip:0,take:50}) { id } }", op="seed")
    known["products"] = [p["id"] for p in d.get("products") or []]
    known["accounts"] = [a["id"] for a in d.get("accounts") or []]


def browse():
    skip = random.choice([0, 0, 0, 10, 20])
    gql("query($s:Int){ products(pagination:{skip:$s,take:12}) { id name price } }", {"s": skip}, "browse")


def search():
    gql("query($q:String){ products(query:$q, pagination:{skip:0,take:10}) { id name price } }",
        {"q": random.choice(SEARCH_TERMS)}, "search")


def view_product():
    if known["products"]:
        gql("query($id:String){ products(id:$id) { id name description price } }",
            {"id": random.choice(known["products"])}, "view_product")


def view_account():
    if known["accounts"]:
        gql("query($id:String){ accounts(id:$id) { id name orders { id createdAt totalPrice products { name quantity price } } } }",
            {"id": random.choice(known["accounts"])}, "view_account")


def sign_up():
    name = "lt-" + "".join(random.choices("abcdefghijklmnopqrstuvwxyz", k=6))
    d = gql("mutation($n:String!){ createAccount(account:{name:$n}) { id } }", {"n": name}, "create_account")
    if d.get("createAccount"):
        with lock:
            known["accounts"].append(d["createAccount"]["id"])


def place_order():
    if not known["accounts"] or not known["products"]:
        return
    picks = random.sample(known["products"], k=min(len(known["products"]), random.randint(1, 3)))
    items = [{"id": pid, "quantity": random.randint(1, 3)} for pid in picks]
    gql("mutation($a:String!,$p:[OrderProductInput!]!){ createOrder(order:{accountId:$a, products:$p}) { id totalPrice } }",
        {"a": random.choice(known["accounts"]), "p": items}, "create_order")


def add_product():
    name, desc, price = random.choice(NEW_PRODUCTS)
    d = gql("mutation($n:String!,$d:String!,$p:Float!){ createProduct(product:{name:$n,description:$d,price:$p}) { id } }",
            {"n": name, "d": desc, "p": price}, "create_product")
    if d.get("createProduct"):
        with lock:
            known["products"].append(d["createProduct"]["id"])


def bad_request():
    """The caller's mistakes: should count as client_error, not against the SLO."""
    if random.random() < 0.5:
        gql('mutation($p:[OrderProductInput!]!){ createOrder(order:{accountId:"does-not-exist", products:$p}) { id } }',
            {"p": [{"id": "does-not-exist", "quantity": 1}]}, "bad_order")
    else:
        gql('{ accounts(id:"does-not-exist") { id name } }', op="bad_account")


ACTIONS = [(browse, 35), (search, 15), (view_product, 10), (view_account, 15),
           (sign_up, 8), (place_order, 12), (add_product, 2), (bad_request, 3)]
FUNCS, WEIGHTS = zip(*ACTIONS)


def user(stop_at):
    while time.time() < stop_at:
        random.choices(FUNCS, WEIGHTS)[0]()
        time.sleep(random.uniform(0.3, 1.5))


def report(final=False):
    with lock:
        total = sum(stats.values())
        errors = sum(v for k, v in stats.items() if k.endswith(" err"))
        print(("FINAL" if final else "progress") + f": {total} requests, {errors} errors", flush=True)
        for op in sorted(latencies):
            xs = sorted(latencies[op])
            p95 = xs[max(int(len(xs) * 0.95) - 1, 0)]
            print(f"  {op:15s} ok={stats[op + ' ok']:4d} err={stats[op + ' err']:3d} p95={p95 * 1000:6.0f}ms", flush=True)


if __name__ == "__main__":
    load_ids()
    print(f"seeded: {len(known['products'])} products, {len(known['accounts'])} accounts; "
          f"{USERS} users for {DURATION}s against {GQL}", flush=True)
    stop_at = time.time() + DURATION
    threads = [threading.Thread(target=user, args=(stop_at,), daemon=True) for _ in range(USERS)]
    for t in threads:
        t.start()
    next_report = time.time() + 60
    while any(t.is_alive() for t in threads):
        time.sleep(1)
        if time.time() >= next_report:
            report()
            next_report += 60
    report(final=True)
