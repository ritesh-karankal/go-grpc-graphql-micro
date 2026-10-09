# Selected postmortems

Five incidents from building and operating **go-micro-shop** (Go gRPC microservices + GraphQL gateway on AWS EKS, OpenTelemetry → SigNoz). Each was found through testing, load, or a planned fault drill, and is written up in full, with commands and output, in the [incident journal](../journal.md). The SLOs they relate to are in [SLO.md](SLO.md).

| # | Incident | Found by | Status |
|---|---|---|---|
| 1 | [Fault drill: a service kept working 11 s after its caller gave up](#1-fault-drill-a-service-kept-working-11-s-after-its-caller-gave-up) | planned fault drill | root cause reproduced; fix designed |
| 2 | [Every deploy dropped in-flight requests](#2-every-deploy-dropped-in-flight-requests) | load test during a rollout | fix deployed; rollout-under-load re-test pending |
| 3 | [gRPC traffic pinned to one replica](#3-grpc-traffic-pinned-to-one-replica) | 110-user load test | fix deployed; load re-test pending |
| 4 | [One bad request could crash the catalog service](#4-one-bad-request-could-crash-the-catalog-service) | code review while defining SLIs | fixed, regression test |
| 5 | [The monitoring stack was the riskiest workload in the cluster](#5-the-monitoring-stack-was-the-riskiest-workload-in-the-cluster) | load test + database introspection | OOM risk fixed and verified; CPU fix ready |

---

## 1. Fault drill: a service kept working 11 s after its caller gave up

**Summary.** A planned SLO drill took the orders database down under simulated shopper traffic. The burn-rate alert fired and the root cause was found from the alert → trace → log. The drill also exposed a hidden bug: order-service kept a request alive for **11.3 s** after the gateway had given up at its **3 s** timeout.

- **Impact (drill):** ~23 min of failed checkouts = **≈ 11% of the 28-day checkout error budget**; browse & search stayed at 100% (failure contained). In production, the bug would pile up goroutines and connections during any database outage and slow recovery.
- **Detection:** SigNoz trace: gateway span 3.0 s (`DeadlineExceeded`), order-service span 11.33 s (`INTERNAL`). Logs showed instant `connection refused` errors, which didn't explain 11 s.
- **Root cause:** reproduced locally against a real Postgres. A cleanly stopped database fails in **1 ms**; a *hung* one (connections open, no answers, like a deleted pod) made the request **ignore its 3 s deadline and block for 40 s+**. The `lib/pq` driver handles cancellation by asking the server to cancel; a dead server never answers, so the read on a stale pooled connection keeps waiting.
- **Fix (designed):** move to the `pgx` driver, which interrupts I/O on cancellation; retire idle pooled connections; turn the reproduction into a regression test; re-run the drill (target: server span ≈ 3 s).
- **What it shows:** SLOs and burn-rate alerting that work end to end; fault injection; reasoning from traces; reproducing a production-only failure mode locally. **"Database down" isn't one failure mode.**

## 2. Every deploy dropped in-flight requests

**Summary.** A load test that happened to overlap a rollout saw `Connection reset by peer`. Every deploy was cutting off requests, silently burning error budget.

- **Impact:** a handful of failed requests on **every** deploy, invisible in normal monitoring and attributed to "random" network errors.
- **Detection:** matched the error times to the image push and Argo CD sync (12:48–12:52), then read the shutdown path.
- **Root cause:** on SIGTERM the services flushed telemetry and called `os.Exit(0)` immediately: no draining, while Kubernetes and the ALB were still routing traffic to the pod. There was no `preStop` delay or grace period either.
- **Fix:** a shared `lifecycle` package: stop accepting, drain in-flight gRPC/HTTP requests (≤ 10 s), then flush telemetry. Manifests: `preStop` sleep 15 s, `terminationGracePeriodSeconds` 30, ALB pod readiness gates, deregistration delay 300 s → 30 s.
- **Result:** tests prove an in-flight HTTP request and gRPC call complete during shutdown, and the same tests **fail with `EOF`/`Unavailable` against a hard stop**; the real gateway binary drains and exits cleanly on SIGTERM. Deployed; a load-during-rollout re-test is the remaining step.
- **What it shows:** graceful shutdown done properly (app *and* platform side), and treating deploys as a source of SLO burn.

## 3. gRPC traffic pinned to one replica

**Summary.** At ~70 requests/s (110 simulated users), one `order-service` pod used **55m** CPU while its twin used **1m**. The second replica added failover but almost no capacity.

- **Impact:** effective capacity of each gRPC service ≈ one pod; scaling out (or an HPA) would barely help; latency SLO at risk under load.
- **Detection:** per-pod CPU (`kubectl top`, after adding metrics-server) during the load test; per-pod span counts in SigNoz.
- **Root cause:** a Kubernetes ClusterIP Service balances **connections**, and gRPC keeps **one long-lived HTTP/2 connection**, so each client stuck to the first pod it reached.
- **Fix:** client-side load balancing: headless Services (DNS returns every pod IP) + gRPC `round_robin`. Rolled out as new `*-headless` Services next to the old ones (expand → migrate → contract), because changing a Service to headless in place would briefly break DNS.
- **Result:** deployed; existing gRPC tests pass through the new client configuration; a load re-test to confirm even per-pod CPU is the remaining step.
- **What it shows:** finding capacity problems through load testing and per-replica metrics; understanding L4 vs L7 load balancing; zero-downtime migration.

## 4. One bad request could crash the catalog service

**Summary.** While mapping errors to gRPC status codes for the SLIs, a review found that an order containing an **unknown product ID** would make catalog-service dereference a nil pointer and crash. Browse, search and checkout all depend on it.

- **Impact (potential):** a single crafted request could crash-loop catalog pods, a denial of service. Caught before it happened.
- **Root cause:** Elasticsearch's multi-get returns `found: false` with no `_source` for missing IDs; the code read `_source` unconditionally, and Go gRPC servers don't recover handler panics.
- **Fix:** skip documents that weren't found; reject orders whose products don't exist with `InvalidArgument` instead of storing an empty $0 order.
- **Result:** a regression test with a fake Elasticsearch **panics on the old code and passes on the fix**.
- **What it shows:** defensive handling of external data; proving a test actually catches the bug; reliability work driven by SLO design.

## 5. The monitoring stack was the riskiest workload in the cluster

**Summary.** Under load, one node ran at **78–88% CPU** while the app used ~500m in total. The observability backend itself was the problem, and it had a latent crash.

- **Findings:**
  - **ZooKeeper** (used by SigNoz's ClickHouse) had a **1024 MB Java heap** inside a **512Mi** memory limit, already at 432Mi. One busy period away from an OOM kill, which would stop all telemetry storage.
  - **ClickHouse** used ~860m CPU even at rest. `system.part_log` showed why: in 5 minutes ClickHouse wrote **402k rows to its own profiler log** (plus 190k to other internal logs) against **~28k rows of real telemetry**. It was logging about itself ~20× more than our data, and merging all of it.
- **Fix:** ZooKeeper heap 256 MB (**verified: 432Mi → 203Mi**); ClickHouse requests sized from measurements; unused ClickHouse system logs disabled via a config file (ready to apply; expected to cut idle CPU substantially).
- **What it shows:** measuring before sizing; checking JVM heap vs container limits; database introspection to explain resource use; recognising that monitoring has its own cost and failure modes.

---

*Each incident has a full postmortem in the [journal](../journal.md): timeline, exact errors, commands and their output, root cause, fix, lessons, and how to approach it next time.*
