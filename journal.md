# Incident Journal

Every error hit while building, deploying or operating this project, written up as a postmortem. Each entry explains, in plain language:

- **where** it happened and **what** we saw (the exact error);
- the **background** you need to understand it;
- **how** we got into that state, and **why** it really happened (root cause, as a chain of cause and effect);
- what it **affected**;
- **how** it was resolved, step by step, with what each command does and what output to expect;
- what we **learned**, and how to approach the same kind of problem next time.

Entries are newest first. Self-inflicted mistakes (tooling slips, wrong docs) are logged the same way as any other incident.

> **Before committing:** this file is public on GitHub. Never paste real IPs, AWS account IDs, load balancer hostnames, cluster IDs, tokens or passwords. Use placeholders: `<YOUR_IP>`, `<ACCOUNT_ID>`, `<LB_HOST>`, `<CLUSTER_ID>`, `<TOKEN>`.

## Index

| ID | Date | Area | Title | Severity | Status |
|---|---|---|---|---|---|
| [INC-023](#inc-023--fault-drill-order-service-keeps-working-11-s-after-the-gateway-gave-up-at-3-s) | 2026-10-09 | order-service / Postgres | Fault drill: order-service keeps working 11 s after the gateway gave up at 3 s | Medium | Root cause found |
| [INC-022](#inc-022--clickhouse-busy-logging-itself-internal-system-logs-outweigh-real-telemetry-20x) | 2026-10-09 | SigNoz / ClickHouse | ClickHouse busy logging itself: internal system logs outweigh real telemetry ~20x | Medium | Fix ready |
| [INC-021](#inc-021--quality-gate-error-again-go-only-credits-coverage-to-the-package-under-test) | 2026-10-09 | Jenkins / SonarQube | Quality gate `ERROR` again: Go only credits coverage to the package under test | Medium | Resolved |
| [INC-020](#inc-020--grpc-traffic-not-balanced-across-replicas) | 2026-10-09 | gRPC / Kubernetes Services | gRPC traffic not balanced across replicas | Medium | Deployed, verifying |
| [INC-019](#inc-019--signoz-sizing-zookeeper-heap-larger-than-its-memory-limit-clickhouse-under-requested) | 2026-10-09 | SigNoz / capacity | SigNoz sizing: ZooKeeper heap larger than its memory limit, ClickHouse under-requested | Medium | Resolved |
| [INC-018](#inc-018--kubectl-top-fails-metrics-api-not-available) | 2026-10-09 | Kubernetes / EKS add-ons | `kubectl top` fails: `Metrics API not available` | Low | Fix ready |
| [INC-017](#inc-017--connection-reset-by-peer-during-a-rollout-pods-exit-without-draining) | 2026-10-09 | Kubernetes / Go services | `Connection reset by peer` during a rollout: pods exit without draining | Medium | Deployed, verifying |
| [INC-016](#inc-016--backend-quality-gate-error-sonarqube-saw-0-coverage-on-new-code) | 2026-10-09 | Jenkins / SonarQube | Backend quality gate `ERROR`: SonarQube saw 0% coverage on new code | Medium | Resolved |
| [INC-015](#inc-015--sonarqube-ui-on-port-9000-keeps-loading-ip-allowlist-out-of-date) | 2026-10-09 | Jenkins / AWS SG | SonarQube UI on port 9000 keeps loading: IP allowlist out of date | Low | Resolved |
| [INC-014](#inc-014--catalog-service-would-crash-on-an-order-with-an-unknown-product-id) | 2026-10-09 | Catalog / Elasticsearch | catalog-service would crash on an order with an unknown product ID | High | Resolved |
| [INC-013](#inc-013--signoz-shows-youre-not-sending-any-data-yet) | 2026-10-09 | SigNoz / OpenTelemetry | SigNoz shows "You're not sending any data yet" | Medium | Resolved |
| [INC-012](#inc-012--signoz-ui-not-reachable-load-balancer-created-as-internal) | 2026-10-09 | SigNoz / AWS LB Controller | SigNoz UI not reachable: load balancer created as `internal` | Medium | Resolved |
| [INC-011](#inc-011--update-deployment-file-fails-nothing-added-to-commit) | 2026-10-08 | Jenkins / GitOps | `Update Deployment file` fails: "nothing added to commit" | Low | Resolved |
| [INC-010](#inc-010--backend-pipeline-aborted-after-60-minutes-in-owasp-dependency-check) | 2026-10-08 | Jenkins / OWASP | `backend` pipeline aborted after 60 minutes in OWASP Dependency-Check | Medium | Resolved |
| [INC-009](#inc-009--argo-cd-ui-not-reachable-in-the-browser) | 2026-10-08 | Argo CD / Browser | Argo CD UI "not reachable" in the browser | Low | Resolved |
| [INC-008](#inc-008--lb-controller-iam-role-create-fails-entityalreadyexists) | 2026-10-08 | AWS / IAM | LB Controller IAM role create fails: `EntityAlreadyExists` | Low | Explained |
| [INC-007](#inc-007--su-user-ec2-user-does-not-exist-on-the-bastion) | 2026-10-08 | Bastion / SSM | `su: user ec2-user does not exist` on the "bastion" | Low | Resolved |
| [INC-006](#inc-006--ecr-create-repository-fails-repositoryalreadyexistsexception) | 2026-10-08 | AWS / ECR | ECR `create-repository` fails: `RepositoryAlreadyExistsException` | Low | Explained |
| [INC-005](#inc-005--first-eks-cluster-build-fails-value-must-not-contain-) | 2026-10-08 | Jenkins / Terraform | First `eks-cluster` build fails: `Value must not contain "//"` | Low | Resolved |
| [INC-004](#inc-004--local-docker-compose-stack-recreated-by-a-broken-heredoc) | 2026-10-08 | Local / tooling | Local Docker Compose stack recreated by a broken heredoc | Low | Resolved |
| [INC-003](#inc-003--signoz-local-docker-compose-instructions-out-of-date) | 2026-10-08 | Docs / SigNoz | SigNoz local Docker Compose instructions out of date | Low | Resolved |
| [INC-002](#inc-002--signoz-would-not-filter-services-by-environment) | 2026-10-08 | Observability | SigNoz would not filter services by environment | Medium | Resolved |
| [INC-001](#inc-001--kubectl-dry-run-fails-with-no-such-host) | 2026-10-08 | Kubernetes | `kubectl` dry-run fails with "no such host" | Low | Explained |
| [INC-044](#inc-044--teardown-deleting-the-argo-cd-application-would-not-have-deleted-the-app) | 2026-10-04 | Teardown / Argo CD | Teardown: deleting the Argo CD Application would not have deleted the app | Medium | Resolved |
| [INC-043](#inc-043--docsarchitecturedrawio-looks-empty-on-github) | 2026-10-04 | Docs / draw.io | `docs/architecture.drawio` looks empty on GitHub | Low | Resolved |
| [INC-042](#inc-042--commit-messages-contain-the-pasted-command-git-commit--m-) | 2026-10-04 | Git | Commit messages contain the pasted command (`git commit -m "…"`) | Low | Workaround |
| [INC-041](#inc-041--storefront-says-make-sure-the-backend-is-running-at-localhost8000-while-the-backend-was-up) | 2026-10-04 | Frontend | Storefront says "Make sure the backend is running at localhost:8000" while the backend was up | Low | Explained |
| [INC-040](#inc-040--grafana-ui-loads-with-curl-but-not-in-the-browser) | 2026-10-03 | Grafana / Browser | Grafana UI loads with `curl` but not in the browser | Low | Resolved |
| [INC-039](#inc-039--helm-warns-this-chart-is-deprecated-for-grafanagrafana) | 2026-10-03 | Helm / Grafana | Helm warns `this chart is deprecated` for `grafana/grafana` | Low | Resolved |
| [INC-038](#inc-038--kustomize-puts-a-namespace-on-the-cluster-scoped-clustersecretstore) | 2026-10-03 | Kustomize / External Secrets | Kustomize puts a `namespace` on the cluster-scoped `ClusterSecretStore` | Low | Resolved |
| [INC-037](#inc-037--jenkinsfile-sed-would-write-newtag-42-without-quotes-groovy-escaping) | 2026-10-03 | Jenkins / Groovy | Jenkinsfile `sed` would write `newTag: 42` without quotes (Groovy escaping) | Low | Resolved |
| [INC-036](#inc-036--argo-cds-registry-override-would-have-dropped-the-image-tag-kustomize-set-image) | 2026-10-03 | Kustomize / Argo CD | Argo CD's registry override would have dropped the image tag (Kustomize `set image`) | Medium | Resolved |
| [INC-035](#inc-035--trivy-gate-blocks-the-frontend-image-cve-2026-31789-openssl-critical) | 2026-10-03 | Jenkins / Trivy | Trivy gate blocks the frontend image: CVE-2026-31789 (OpenSSL, CRITICAL) | Medium | Resolved |
| [INC-034](#inc-034--jenkins-permission-denied--varrundockersock) | 2026-10-03 | Jenkins / Docker | Jenkins: `permission denied … /var/run/docker.sock` | Medium | Resolved |
| [INC-033](#inc-033--sonarqube-scan-http-connect-timed-out-server-url-was-the-public-ip) | 2026-10-03 | Jenkins / SonarQube | SonarQube scan: `HTTP connect timed out` (server URL was the public IP) | Medium | Resolved |
| [INC-032](#inc-032--ecrcreaterepository-accessdenied-on-the-bastion) | 2026-10-03 | AWS / ECR / IAM | `ecr:CreateRepository` AccessDenied on the bastion | Low | Explained |
| [INC-031](#inc-031--argo-cd-service-got-a-classic-load-balancer-not-one-from-the-lb-controller) | 2026-10-03 | Argo CD / AWS LB | Argo CD Service got a Classic Load Balancer, not one from the LB Controller | Low | Explained |
| [INC-030](#inc-030--new-alb-curl-6-could-not-resolve-host-and-dns_probe_possible-in-the-browser) | 2026-10-03 | AWS ALB / DNS | New ALB: `curl: (6) Could not resolve host` (and `DNS_PROBE_POSSIBLE` in the browser) | Low | Explained |
| [INC-029](#inc-029--aws-iam-create-role-argument---role-name-expected-one-argument) | 2026-10-03 | Bastion / Shell | `aws iam create-role`: `argument --role-name: expected one argument` | Low | Resolved |
| [INC-028](#inc-028--bastion-accessdenied-iamcreatepolicy-permission-in-git-but-not-applied-yet) | 2026-10-03 | Bastion / IAM / Terraform | Bastion `AccessDenied iam:CreatePolicy`: permission in git, but not applied yet | Low | Resolved |
| [INC-027](#inc-027--bastion-curl-23--permission-denied-and-unable-to-load-paramfile) | 2026-10-03 | Bastion / SSM | Bastion: `curl: (23) … Permission denied` and `Unable to load paramfile` | Low | Resolved |
| [INC-026](#inc-026--jenkins-eks-cluster-1-error-aws_account_id) | 2026-10-03 | Jenkins / Credentials | Jenkins `eks-cluster` #1: `ERROR: AWS_ACCOUNT_ID` | Low | Resolved |
| [INC-025](#inc-025--terraform-invalid-for_each-argument-for-eks-access-entries) | 2026-10-03 | Terraform | Terraform: `Invalid for_each argument` for EKS access entries | Low | Resolved |
| [INC-024](#inc-024--planned-eks-version-130-and-al2_x86_64-nodes-no-longer-available) | 2026-10-03 | Terraform / EKS | Planned EKS version 1.30 and `AL2_x86_64` nodes no longer available | Medium | Resolved |

**Severity:** High = outage or data loss · Medium = a feature silently broken · Low = annoyance, docs, or local only.
**Status:** Open · Resolved (fixed) · Workaround (unblocked, real fix pending) · Explained (expected behaviour, nothing to fix).

---

## Template (copy for each new incident)

```markdown
## INC-00X – <short title that names the symptom>

| | |
|---|---|
| **Date** | YYYY-MM-DD |
| **Where** | Guide step, tool, job/stage, file (e.g. Recreate guide Step 7, Jenkins `backend` › Docker Image Build, `k8s/base/...`) |
| **Severity** | High / Medium / Low, with one line on why |
| **Status** | Open / Resolved / Workaround / Explained |

### Summary
Two or three sentences a newcomer can follow: what broke, why, and whether it is fixed.

### Background
The concepts you need to understand this error (how the tool works, what the setting means).
Skip only if the error is self-explanatory.

### Timeline
1. What we were doing.
2. What we ran.
3. What we saw.
4. What we checked, what we found, what we changed.

### What happened (symptom)
The exact error, copied from the output (redacted):

    <error output>

What it looked like from the outside (UI, logs, pod status, etc.).

### How we got there
The command or click that triggered it, and the state the system was in at that moment.

### Why (root cause)
The real reason, as a chain: A happened, which caused B, which made C fail.
Check it against docs or source code, and say where you checked.

### Impact
What was affected (and what was NOT): data, cost, other services, time lost.

### Resolution (step by step)
1. Step: the command.
   What it does, and the output you should expect.
2. ...
3. Verify: a command or check that proves it is fixed.

### Terminal commands
The commands actually run, tagged with where they ran (laptop / bastion / Jenkins UI). Mark anything only recommended as `# suggested, not run`. Redact hostnames, IPs, instance and account IDs.

**Finding it** (command, then its real output)
    laptop$ <command that showed the error>
    <output as printed, trimmed with … if long>

**Fixing it / trying to fix it**
    bastion$ <command>
    <output>      (or "(output not captured)"; "# suggested, not run" for steps not executed)

### Code / config change
The diff (or "none"), and which commit it went into.

### Lessons learned
What this taught us about the tool or our setup.

### How to approach it next time
A short debugging playbook: what to look at first, which commands to run, what each result means.

### Prevention / follow-up
What stops it happening again (code, docs, checks), and anything still to do.

### References
Docs, source files or issues used to confirm the cause.
```

---

## INC-023 – Fault drill: order-service keeps working 11 s after the gateway gave up at 3 s

| | |
|---|---|
| **Date** | 2026-10-09 |
| **Where** | SLO fault drill (`order-db` scaled to 0); trace of a failed `CreateOrder`; `order-service` → Postgres |
| **Severity** | Medium: during a database outage, order-service does work nobody waits for, which can pile up connections/goroutines and slow recovery |
| **Status** | Root cause found (reproduced locally): `lib/pq` can't abort a query on a hung connection; fix proposed |

### Summary
During the planned SLO fault drill (orders database taken down on purpose), a failing checkout's trace showed the gateway giving up after **3.00 s** (its timeout) while **order-service's `PostOrder` kept running for 11.33 s** before failing with `INTERNAL: could not post order`. gRPC passes the deadline to the server, so order-service *knew* the caller had given up, but something in the database path didn't stop.

### Background
- The gateway calls order-service with a **3-second timeout** (`context.WithTimeout` in the resolvers). gRPC sends that deadline to the server (`grpc-timeout` header), so the server's request context is cancelled at the same moment.
- Code that respects the context stops when it's cancelled. Code that ignores it keeps running: **wasted work**, holding a goroutine and possibly a database connection.
- During an outage, wasted work multiplies: every retrying user adds another stuck request, which is a classic way for a dependency failure to cascade.

### Timeline
1. Drill: `order-db` scaled to 0 (Argo CD self-heal paused first; see the drill notes in `docs/SLO.md`).
2. 22:07:26 (UTC+5:30): outage confirmed: order history → `DeadlineExceeded` after 3.3 s.
3. ≤ 22:15:07: checkout fast-burn alert (drill rule) **Firing**; the real-user rule stayed OK.
4. 22:13:47: failed `CreateOrder` trace (`1d9cf9fb…`) inspected in SigNoz:

   | Span | Service | Duration | Status |
   |---|---|---|---|
   | `POST /graphql` | gateway | 3.01 s | HTTP **200** |
   | `CreateOrder` | gateway | 3.01 s | error |
   | `pb.OrderService/PostOrder` (client) | gateway | **3.00 s** | error (deadline) |
   | `pb.AccountService/GetAccount` | order → account | short | ok |
   | `pb.OrderService/PostOrder` (server) | order-service | **11.33 s** | `INTERNAL`, "could not post order" |

### What happened (symptom)
The server span outlived its caller by ~8.3 s. (The same trace also shows HTTP 200 for a failed checkout, which is the reason the SLIs use the GraphQL `outcome` metric instead of HTTP status codes.)

### Why (root cause)
Evidence so far: order-service's error logs during the drill (`Failed to post order`, `exception.message: dial tcp <order-db>:5432: connect: connection refused`, `*net.OpError`) show each connection attempt failing **immediately** (refused, not timing out). An 11 s span made of instant failures points to **repeated attempts** inside one request rather than one long wait. The trace doesn't show child spans for these attempts (failed connects aren't traced by the current `otelsql` span options), so the exact loop is still *to be confirmed*. Hypotheses:
- the Postgres **connection attempt** doesn't honour the cancelled context (driver or `otelsql` connector path), and runs until its own TCP/connect timeout;
- or `database/sql` retries a bad connection (it does up to 2 retries on `driver.ErrBadConn`), each attempt waiting on a dead endpoint.

### Impact
- During the drill: extra latency only on the server side; users had already got an error after 3 s.
- At real traffic: stuck goroutines and connections in order-service for every failed checkout, for the whole outage, and a slower recovery afterwards.

### Reproduction (local, real Postgres)
A throwaway test connected `order`'s repository to a Postgres container, inserted one order (so the pool holds an open connection), broke the database in two different ways, then called `PutOrder` with the gateway's **3-second deadline**:

| Database state | How | Result |
|---|---|---|
| stopped cleanly | `docker stop` (Postgres closes its connections, port closed) | returns in **~1 ms**: `dial tcp …: connect: connection refused` |
| hung | `docker pause` (connections stay open, nothing answers) | **ignores the 3 s deadline: still blocked after 40 s** (both attempts) |

The first matches the many instant `connection refused` log lines during the drill (new connections to a Service with no pods). The second matches the 11.33 s span: a request that picked up an **existing pooled connection to the deleted pod**. In the cluster it ended after ~11 s, most likely when the network reported the vanished pod IP as unreachable *(inferred, not measured)*; locally the frozen container never answers, so it blocks indefinitely.

**Confirmed cause:** `lib/pq` handles context cancellation by sending a *cancel request to the server* over a new connection, and keeps reading from the original connection until the server responds. When the server is gone or hung, nobody responds: the deadline reaches the driver, but the blocked read isn't interrupted. (`lib/pq` is in maintenance mode; `pgx` instead closes or interrupts the connection when the context is cancelled.)

### Resolution (step by step)
1. Identify where the time goes. *(done: reproduction above)*
2. **Fix (proposed):** switch the Postgres driver from `lib/pq` to **`pgx`** (`github.com/jackc/pgx/v5/stdlib`, through the same `database/sql` + `otelsql` setup), which interrupts I/O when the context is cancelled. `order/repository.go` uses `pq.CopyIn` for `order_products`, which has to become a plain multi-row `INSERT` (or pgx's own COPY). Also bound the pool: `SetConnMaxIdleTime`, so stale connections to a replaced pod are retired.
3. **Prove it** with the same reproduction as a test: with a hung database, `PutOrder` must return within ~3 s with `context deadline exceeded`.
4. Re-run the drill: the order-service span should end at about the same time as the gateway's (≈ 3 s).

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**The drill**
```console
laptop$ python3 scripts/loadgen.py http://$ALB 2400 10
bastion$ kubectl -n argocd patch application go-micro-shop-dev --type merge -p '{"spec":{"syncPolicy":{"automated":{"prune":true,"selfHeal":false}}}}'
bastion$ kubectl -n go-micro-shop scale statefulset order-db --replicas=0
(output not captured)

laptop$ curl -s -X POST -H 'Content-Type: application/json' -H 'X-Synthetic: true' \
  -d '{"query":"query Accounts { accounts(pagination:{skip:0,take:1}) { id orders { id } } }"}' http://$ALB/graphql
laptop$                                  # (no output: $ALB was empty in this terminal)
laptop$ curl -sS --max-time 20 -w '\nHTTP %{http_code} in %{time_total}s\n' -X POST -H 'Content-Type: application/json' -H 'X-Synthetic: true' \
  -d '{"query":"query Accounts { accounts(pagination:{skip:0,take:1}) { id orders { id } } }"}' "http://<ALB_HOST>/graphql"
{"errors":[{"message":"rpc error: code = DeadlineExceeded desc = context deadline exceeded","path":["accounts",0,"orders"],"locations":[{"line":1,"column":60}]}],"data":null}
HTTP 200 in 3.326880s

# SigNoz UI: drill alert Firing (<= 22:15:07 IST); real-user alert OK
# SigNoz UI: trace 1d9cf9fb…: gateway PostOrder 3.00 s; order-service PostOrder 11.33 s, INTERNAL "could not post order"
# SigNoz UI: Logs (order-service, ERROR): "Failed to post order", exception.message "dial tcp <order-db>:5432: connect: connection refused"

laptop$ # probe order history every 10 s, log state changes
22:29:00 order history DOWN: code = Internal
22:29:14 order history DOWN: code = DeadlineExceeded
22:29:24 order history DOWN: code = Internal
22:29:45 order history UP
23:14:03 watcher finished (last state: UP)

[ec2-user@ip-10-10-20-x bin]$ date '+%H:%M:%S restore'
kubectl -n go-micro-shop scale statefulset order-db --replicas=1
kubectl -n argocd patch application go-micro-shop-dev --type merge -p '{"spec":{"syncPolicy":{"automated":{"prune":true,"selfHeal":true}}}}'
kubectl -n go-micro-shop get events --sort-by=.lastTimestamp | grep -i order-db | tail -8
16:59:35 restore
statefulset.apps/order-db scaled
application.argoproj.io/go-micro-shop-dev patched
23m         Normal   Killing            pod/order-db-0                    Stopping container postgres
23m         Normal   SuccessfulDelete   statefulset/order-db              Delete Pod order-db-0 in StatefulSet order-db successful
1s          Normal   Scheduled          pod/order-db-0                    Successfully assigned go-micro-shop/order-db-0 to ip-10-10-x-x.eu-north-1.compute.internal
1s          Normal   SuccessfulCreate   statefulset/order-db              Create Pod order-db-0 in StatefulSet order-db successful
```
**Reproducing the 11 s locally**
```console
laptop$ docker run -d --name inc023-pg -e POSTGRES_USER=u -e POSTGRES_PASSWORD=p -e POSTGRES_DB=d -p 55433:5432 \
          -v $PWD/order/up.sql:/docker-entrypoint-initdb.d/up.sql:ro postgres:16-alpine

laptop$ go test ./zz_inc023/ -run TestPutOrderWhenDatabaseGoesAway -v -count=1        # database broken with: docker stop
    main_test.go:23: database stopped: inc023-pg
    main_test.go:31: attempt 1: returned after 1ms, err=dial tcp 127.0.0.1:55433: connect: connection refused, ctxErr=context canceled
    main_test.go:31: attempt 2: returned after 0s, err=dial tcp 127.0.0.1:55433: connect: connection refused, ctxErr=context canceled
    main_test.go:31: attempt 3: returned after 0s, err=dial tcp 127.0.0.1:55433: connect: connection refused, ctxErr=context canceled
ok  	github.com/ritesh-karankal/go-grpc-graphql-micro/zz_inc023	0.347s

laptop$ go test ./zz_inc023/ -run TestPutOrderWhenDatabaseHangs -v -count=1 -timeout 150s   # database broken with: docker pause
    main_test.go:26: database paused (connections open, no answers)
    main_test.go:38: attempt 1: STILL BLOCKED after 40s (deadline was 3s)
    main_test.go:38: attempt 2: STILL BLOCKED after 40s (deadline was 3s)
ok  	github.com/ritesh-karankal/go-grpc-graphql-micro/zz_inc023	80.278s

laptop$ rm -rf zz_inc023; docker rm -f inc023-pg
```

### Code / config change
Pending.

### Lessons learned
- **"Database down" isn't one failure mode.** A cleanly closed port fails in 1 ms; a hung peer (deleted pod, network partition) can block far past every timeout. Test both.
- Passing `ctx` down is necessary but not sufficient: the driver must be able to *act* on cancellation.
- A fault drill tests more than the alert: the traces show how each service behaves *while* failing.
- Every service in a call chain should stop when its caller stops. Timeouts are only half of it; honouring cancellation is the other half.

### How to approach it next time
1. In a failed trace, compare each server span's duration with its caller's timeout.
2. A server span much longer than its client span means cancellation is ignored somewhere below it: follow the longest child span.

### Prevention / follow-up
- Add the drill (with this check) to the regular release checklist in `docs/SLO.md`.

### References
- Go: [`database/sql` — contexts and cancellation](https://go.dev/doc/database/cancel-operations).
- gRPC: [Deadlines](https://grpc.io/docs/guides/deadlines/).

---

## INC-022 – ClickHouse busy logging itself: internal system logs outweigh real telemetry ~20x

| | |
|---|---|
| **Date** | 2026-10-09 |
| **Where** | SigNoz's ClickHouse (`chi-signoz-clickhouse-cluster-0-0-0`); `k8s/monitoring/signoz-values.yaml` (`clickhouse.files`) |
| **Severity** | Medium: ClickHouse used ~770–860m CPU constantly, most of one node, to write data nobody reads |
| **Status** | Fix ready: noisy system logs disabled via a ClickHouse config file; waiting for `helm upgrade` |

### Summary
After fixing ZooKeeper (INC-019), ClickHouse still used **~860m CPU** with light traffic. Its `system.part_log` showed what it was writing: in 5 minutes, **402k rows to `system.trace_log`** (ClickHouse's own sampling profiler), 82k to `asynchronous_metric_log`, 75k to `zookeeper_log`, 33k to `processors_profile_log`, against only ~20–28k rows per table of actual SigNoz telemetry. Every insert creates parts that must be merged, so most of ClickHouse's CPU went into **logging about itself**. SigNoz doesn't use these tables, so a small config file disables them.

### Background
- ClickHouse keeps internal **system log tables** (`system.*_log`): query history, profiler samples, server metrics, ZooKeeper requests, and more. They're enabled by default and are flushed to disk every few seconds as regular MergeTree tables.
- In a MergeTree table, every insert creates a **part**, and background **merges** combine parts. Lots of small inserts means lots of merge work, which shows up as constant CPU.
- `trace_log` is filled by the **query profiler** (stack samples, by default every second per thread) and by memory profiling, so it's the biggest by far.
- Config files in `config.d/` are merged in **alphabetical order**. An element with `remove="1"` deletes that section from the config.
- The SigNoz chart passes extra files through `clickhouse.files` to the ClickHouse operator, which mounts them into `config.d/`. The operator's own defaults are named `01-clickhouse-*.xml`.

### Timeline
1. INC-019 fixed: ZooKeeper at 203Mi. ClickHouse still at 863m CPU / 1151Mi.
2. Asked ClickHouse what it was doing (bastion, `clickhouse-client` inside the pod):
   - `system.merges`: merging `system.metric_log`;
   - active parts per table: the top 8 tables were all `system.*` (trace_log 17, metric_log 14, zookeeper_log 12, …);
   - rows written in the last 5 minutes (`system.part_log`): table below.
3. Checked the chart: system logs come from the operator's `01-clickhouse-*.xml` files; custom files can be added through `clickhouse.files`.
4. Added `config.d/z_disable_noisy_system_logs.xml`; rendered the chart: the file is included alongside the chart's defaults.

### What happened (symptom)

    signoz   chi-signoz-clickhouse-cluster-0-0-0   863m   1151Mi

Rows written in 5 minutes (`system.part_log`, `event_type = 'NewPart'`):

| Table | Rows | Needed by SigNoz? |
|---|---|---|
| `system.trace_log` | **402,662** | no (ClickHouse profiler) |
| `system.asynchronous_metric_log` | 82,494 | no |
| `system.zookeeper_log` | 74,612 | no |
| `system.processors_profile_log` | 33,401 | no |
| `signoz_traces.signoz_index_v3` | 27,944 | **yes** (spans) |
| `signoz_metrics.metadata` / `samples_v4` / aggregates | 18–23k each | **yes** (metrics) |

### How we got there
The chart's ClickHouse runs with ClickHouse's default system logging, which is tuned for debugging ClickHouse itself, not for a small cluster running an observability backend.

### Why (root cause)
1. ClickHouse's default system logs are on, including the per-second query profiler (`trace_log`).
2. They generate far more rows than our telemetry volume.
3. Every flush creates parts, and merging them costs CPU continuously, whether or not anyone uses the data.

### Impact
- ~0.8 cores spent constantly on unused data: most of node A's CPU, and the main reason that node ran at 57–88%.
- Extra disk writes and storage on the 20Gi ClickHouse volume (bounded by the logs' TTLs).

### Resolution (step by step)
1. **Ask ClickHouse what it's doing** (bastion):
   ```bash
   CH="kubectl -n signoz exec chi-signoz-clickhouse-cluster-0-0-0 -- clickhouse-client -q"
   $CH "SELECT database, table, round(elapsed,1), round(progress,2) FROM system.merges ORDER BY elapsed DESC LIMIT 10"
   $CH "SELECT database, table, count() FROM system.parts WHERE active GROUP BY database, table ORDER BY count() DESC LIMIT 10"
   $CH "SELECT database, table, sum(rows) FROM system.part_log WHERE event_type='NewPart' AND event_time > now() - INTERVAL 5 MINUTE GROUP BY database, table ORDER BY sum(rows) DESC LIMIT 10"
   ```
2. **Disable the noisy logs** in `signoz-values.yaml` (diff below). Keep `query_log` and `part_log`: cheap, and `part_log` is what found this.
3. **Render and check** that the file is included:
   `helm template signoz signoz/signoz --version 0.145.0 -f k8s/monitoring/signoz-values.yaml` → `ClickHouseInstallation.spec.configuration.files` contains `config.d/z_disable_noisy_system_logs.xml`.
4. **Apply** (bastion): `helm upgrade signoz signoz/signoz -n signoz --version 0.145.0 -f k8s/monitoring/signoz-values.yaml`. ClickHouse restarts; telemetry pauses briefly (exporters retry).
5. **Verify** after ~10 minutes: re-run the `part_log` query → no new rows for the disabled tables; `kubectl -n signoz top pods` → ClickHouse CPU clearly lower; node A's CPU down.
6. Optional cleanup of old data: `$CH "TRUNCATE TABLE system.trace_log"` (and the others). Otherwise their TTLs remove it.

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
bastion$ kubectl -n signoz top pods
chi-signoz-clickhouse-cluster-0-0-0           863m         1151Mi
…
bastion$ CH="kubectl -n signoz exec chi-signoz-clickhouse-cluster-0-0-0 -- clickhouse-client -q"
bastion$ $CH "SELECT database, table, round(elapsed,1) AS secs, round(progress,2) AS progress FROM system.merges ORDER BY elapsed DESC LIMIT 10"
Defaulted container "clickhouse" out of: clickhouse, signoz-clickhouse-udf-init (init)
system	metric_log	0.1	0
bastion$ $CH "SELECT database, table, count() AS parts FROM system.parts WHERE active GROUP BY database, table ORDER BY parts DESC LIMIT 10"
system	trace_log	17
system	metric_log	14
system	zookeeper_log	12
system	processors_profile_log	11
system	part_log	9
system	query_log	9
system	query_views_log	8
system	asynchronous_metric_log	8
signoz_metrics	metadata	7
signoz_traces	dependency_graph_minutes_v2	7
bastion$ $CH "SELECT database, table, sum(rows) AS rows FROM system.part_log WHERE event_type='NewPart' AND event_time > now() - INTERVAL 5 MINUTE GROUP BY database, table ORDER BY rows DESC LIMIT 10"
system	trace_log	402662
system	asynchronous_metric_log	82494
system	zookeeper_log	74612
system	processors_profile_log	33401
signoz_traces	signoz_index_v3	27944
signoz_metrics	metadata	22659
signoz_metrics	samples_v4	21731
signoz_metrics	samples_v4_agg_5m	18936
signoz_traces	tag_attributes_v2	18108
signoz_metrics	samples_v4_agg_30m	18007
```
**Fixing it**
```console
laptop$ helm template signoz signoz/signoz --version 0.145.0 -n signoz -f k8s/monitoring/signoz-values.yaml > r4.yaml   # (checked with a small script)
CHI config files: ['config.d/crash.xml', 'config.d/formatting.xml', 'config.d/load.xml', 'config.d/system_log.xml', 'config.d/z_disable_noisy_system_logs.xml', 'events.proto']

# suggested, not run yet: apply and re-check where the writes go
bastion$ helm upgrade signoz signoz/signoz -n signoz --version 0.145.0 -f k8s/monitoring/signoz-values.yaml
bastion$ $CH "SELECT database, table, sum(rows) FROM system.part_log WHERE event_type='NewPart' AND event_time > now() - INTERVAL 5 MINUTE GROUP BY database, table ORDER BY sum(rows) DESC LIMIT 10"
```

### Code / config change
`k8s/monitoring/signoz-values.yaml`:
```yaml
clickhouse:
  files:
    config.d/z_disable_noisy_system_logs.xml: |
      <clickhouse>
        <trace_log remove="1"/>
        <asynchronous_metric_log remove="1"/>
        <zookeeper_log remove="1"/>
        <processors_profile_log remove="1"/>
        <metric_log remove="1"/>
      </clickhouse>
```

### Lessons learned
- An observability backend has its own observability overhead. Measure what *it* spends resources on, not just your app.
- Defaults are tuned for the software's own debugging, not for your workload. Check what's on.
- `system.part_log` answers "where do my writes go?" in one query.

### How to approach it next time
1. High idle CPU on a database → look at background work first (merges, compactions, vacuum), not just queries.
2. ClickHouse specifically: `system.merges`, `system.parts`, `system.part_log` (writes per table), `system.query_log` (heavy queries).
3. Compare the volume of internal tables with the volume of real data.

### Prevention / follow-up
- Re-check node A's CPU after the change. If ClickHouse is still the top consumer under load, consider a dedicated node for SigNoz.

### References
- ClickHouse: [System tables](https://clickhouse.com/docs/en/operations/system-tables); [`trace_log`](https://clickhouse.com/docs/en/operations/system-tables/trace_log); [Configuration files — `remove` attribute](https://clickhouse.com/docs/en/operations/configuration-files).
- Altinity ClickHouse operator: [custom configuration files](https://github.com/Altinity/clickhouse-operator/blob/master/docs/chi-examples/05-settings-05-files-nested.yaml).

---

## INC-021 – Quality gate `ERROR` again: Go only credits coverage to the package under test

| | |
|---|---|
| **Date** | 2026-10-09 |
| **Where** | Jenkins `backend` › **Quality Check**; `jenkins/Jenkinsfile-Backend` (Go Unit Tests stage); commit `fix: graceful shutdown and gRPC load balancing` |
| **Severity** | Medium: the gate blocked the graceful-shutdown and load-balancing fixes from deploying |
| **Status** | Resolved: build #8 passed the quality gate and deployed |

### Summary
The build for the graceful-shutdown and gRPC load-balancing fixes (INC-017, INC-020) failed the quality gate again, this time with a coverage report present. A local estimate of the new code's coverage gave **42%**, for two reasons. **(1)** `go test -coverprofile` only credits coverage to the *package being tested*. `order`'s tests run `account.NewClient` and `catalog.NewClient` (including the new `round_robin` lines), but that counted for nothing, so the client files showed 0%. **(2)** The new `ListenGRPC` wiring and some `lifecycle` branches really had no tests. Fixed with `-coverpkg` (cross-package credit) plus targeted tests. Duplicated helper code in the four `main.go` files was also consolidated into `telemetry.Flush`.

### Background
- `go test -coverprofile` instruments **only the packages being tested** by default. Code in package B that a test in package A executes counts as covered only with `-coverpkg=<list>`, which instruments the listed packages for every test binary.
- SonarQube applies the 80% condition to **new code**. Small changes spread across many files (a few wiring lines each) easily drop below it when each line lives in a package whose own tests don't execute it.
- SonarQube's *Sonar way* gate also limits **duplicated lines on new code** (≤ 3%). Copy-pasting the same helper into several files counts against it.

### Timeline
1. Commit `fix: graceful shutdown and gRPC load balancing` pushed and built.
2. Quality Check: `Quality gate is 'ERROR'`.
3. Local estimate of new-code coverage (`coverage.out` blocks matched to the commit's `git diff -U0` lines, test and `main.go` files excluded): **21/50 = 42%**.

   | File | Covered / new lines |
   |---|---|
   | `account/client.go`, `catalog/client.go`, `order/client.go` | 0/3 each |
   | `account/server.go`, `catalog/server.go` | 0/2 each |
   | `order/server.go` | 0/4 |
   | `lifecycle/lifecycle.go` | 21/33 |
4. The client lines *are* executed by `order`'s `PostOrder` tests, which construct real `account` and `catalog` clients. Not credited because of per-package instrumentation.
5. Added `-coverpkg`, tests for `ListenGRPC` (start, cancel, clean return), `lifecycle` error branches and SIGTERM handling, and `telemetry.Flush`. Moved the four copies of `flush()` into `telemetry.Flush`.
6. New estimate: **49/56 = 88%**.

### What happened (symptom)

    SonarQube task '<task>' status is 'SUCCESS'
    SonarQube task '<task>' completed. Quality gate is 'ERROR'

### How we got there
A change that touched many packages with one or two lines each, with the coverage set-up from INC-016 (per-package `-coverprofile`).

### Why (root cause)
1. Per-package coverage didn't credit cross-package test execution, so the client changes counted as untested.
2. The new `ListenGRPC` signatures/wiring had no direct tests.
3. Some `lifecycle` branches (error returns, signal handling) weren't exercised.

### Impact
- The deploy of INC-017/INC-020 was blocked until fixed. Nothing running was affected.

### Resolution (step by step)
1. Estimate new-code coverage locally first (as in INC-016), so you know *which* files are short.
2. Check whether "uncovered" lines are executed by tests elsewhere. If yes, the fix is `-coverpkg`, not more tests.
3. Add tests where code is genuinely untested:
   - `account/listen_test.go`, `catalog/listen_test.go`, `order/listen_test.go`: `ListenGRPC` with an already-cancelled context starts and returns `nil` (drained).
   - `lifecycle_test.go`: `ServeGRPC` on a closed listener and `ServeHTTP` on a busy address return errors; `SignalContext` is cancelled by a real `SIGTERM` to the test process.
   - `telemetry/telemetry_test.go`: `Flush` calls shutdown with a deadline and doesn't fail on a flush error.
4. Remove duplication: `telemetry.Flush` replaces four identical `flush()` functions.
5. Pipeline (diff below), then push and run `backend`. Expect `Quality gate is 'OK'`.
   Result: build #8 passed the gate, pushed `:8` images, and Argo CD deployed them.
   SonarQube's own figure for the new code: **84.9% coverage on 126 new lines, 0.0% duplication**, all ratings A (the local estimate was 88%; SonarQube counts lines slightly differently).

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
# Jenkins console:
SonarQube task '<task>' status is 'PENDING'
SonarQube task '<task>' status is 'SUCCESS'
SonarQube task '<task>' completed. Quality gate is 'ERROR'

laptop$ go test -coverprofile=cov2.out ./account/... ./catalog/... ./order/... ./graphql/... ./telemetry/... ./lifecycle/...
laptop$ # changed lines of the commit (git diff -U0 9c8235b^ origin/main) matched against cov2.out
  account/client.go              0/  3
  account/server.go              0/  2
  catalog/client.go              0/  3
  catalog/server.go              0/  2
  lifecycle/lifecycle.go        21/ 33
  order/client.go                0/  3
  order/server.go                0/  4
COVERAGE on new code: 21/50 = 42%  (gate >= 80%)
```
**Fixing it**
```console
laptop$ PKGS="./account/... ./catalog/... ./order/... ./graphql/... ./telemetry/... ./lifecycle/..."
laptop$ go test -coverpkg=$(echo $PKGS | tr ' ' ',') -coverprofile=cov3.out $PKGS
ok  	github.com/ritesh-karankal/go-grpc-graphql-micro/account	0.086s	coverage: 1.3% of statements in ./account/..., …
…
laptop$ # re-estimate
  account/client.go              3/  3
  account/server.go              2/  2
  catalog/client.go              3/  3
  catalog/server.go              2/  2
  lifecycle/lifecycle.go        29/ 33
  order/client.go                0/  3
  order/server.go                4/  4
  telemetry/telemetry.go         6/  6
COVERAGE on new code: 49/56 = 88%  (gate >= 80%)
# Jenkins UI: backend -> build #8 -> gate OK, images :8 pushed and deployed
```

### Code / config change
`jenkins/Jenkinsfile-Backend`:
```diff
-                sh 'go test -coverprofile=coverage.out ./account/... ./catalog/... ./order/... ./graphql/... ./telemetry/...'
+                sh '''
+                    PKGS="./account/... ./catalog/... ./order/... ./graphql/... ./telemetry/... ./lifecycle/..."
+                    go test -coverpkg=$(echo $PKGS | tr ' ' ',') -coverprofile=coverage.out $PKGS
+                '''
 ...
-                          -Dsonar.sources=account,catalog,order,graphql,telemetry \
+                          -Dsonar.sources=account,catalog,order,graphql,telemetry,lifecycle \
```

### Lessons learned
- In Go, "covered" means "covered by tests *of this package*" unless you pass `-coverpkg`. For a multi-package service, cross-package coverage is the more truthful number.
- Many small cross-cutting changes are the hardest case for a new-code coverage gate. Add tests for the new wiring in the same commit.
- Shared helpers belong in a shared package; copies in every `main.go` add duplication and untestable code.

### How to approach it next time
1. Gate `ERROR` → SonarQube failed conditions (coverage? duplication?).
2. Coverage: list uncovered new lines per file; for each, is it executed by any test? Yes → `-coverpkg`. No → write a test.
3. Re-estimate locally before pushing again.

### Prevention / follow-up
- Done: `-coverpkg` for all backend packages; tests for the new wiring.
- Optional: a test for `lifecycle`'s forced stop after `DrainTimeout` (needs the timeout to be configurable).

### References
- Go: [`go help testflag` — `-coverpkg`](https://pkg.go.dev/cmd/go#hdr-Testing_flags).
- SonarQube: [Go test coverage](https://docs.sonarsource.com/sonarqube-server/latest/analyzing-source-code/test-coverage/go-test-coverage/).

---

## INC-020 – gRPC traffic not balanced across replicas

| | |
|---|---|
| **Date** | 2026-10-09 |
| **Where** | gRPC clients (`account/client.go`, `catalog/client.go`, `order/client.go`) and the ClusterIP Services `account-service`, `catalog-service`, `order-service` |
| **Severity** | Medium: the second replica of each gRPC service adds failover but little or no capacity |
| **Status** | Fix ready: headless Services + `round_robin` clients; waiting for build + deploy |

### Summary
During a 110-user load test, `kubectl top pods` showed one `order-service` pod at **55m CPU** and the other at **1m**: idle for the whole test. `catalog-service` was uneven too (82m vs 37m). gRPC sends all requests from one client over **one long-lived HTTP/2 connection**, and a Kubernetes ClusterIP Service balances **connections, not requests**. Each client therefore sticks to whichever pod it reached first.

### Background
- A ClusterIP Service is one virtual IP. kube-proxy picks a backend pod **when a TCP connection opens**. Every request on that connection goes to the same pod.
- HTTP/1.1 clients open many short connections, so traffic spreads naturally. **gRPC (HTTP/2) multiplexes all requests over one connection** and keeps it open, so there's no re-balancing.
- Our clients dial `dns:///order-service:8080`. DNS returns the single ClusterIP, and gRPC's default policy (`pick_first`) uses one connection.
- Standard fix: a **headless Service** (`clusterIP: None`) makes DNS return **every pod IP**, and the client policy **`round_robin`** keeps a connection to each pod and spreads requests across them. Alternatives: a service mesh or an L7 proxy.

### Timeline
1. Load test: 110 users, 3 minutes, ~70 requests/s, 99.99% success for real operations.
2. `kubectl -n go-micro-shop top pods` during the test: `order-service-…7585x 55m`, `order-service-…78kzj 1m`; `catalog-service-…46w6r 82m`, `…vmttf 37m`.
3. Matched against the client code: `grpc.NewClient("dns:///"+url, …)` with no load-balancing config, which means `pick_first` over a single ClusterIP.

### What happened (symptom)

    NAME                               CPU(cores)
    catalog-service-78dfc9bbff-46w6r   82m
    catalog-service-78dfc9bbff-vmttf   37m
    order-service-7ff57b9f69-7585x     55m
    order-service-7ff57b9f69-78kzj     1m      <- idle under load

### How we got there
The standard Kubernetes Service, used with gRPC clients on default settings.

### Why (root cause)
1. Each gRPC client keeps one HTTP/2 connection to `order-service:8080`.
2. kube-proxy balances that connection once, at connect time.
3. All requests from that client go to one pod. With only two gateway pods as clients, both can land on the same `order-service` pod.

### Impact
- Effective capacity of each gRPC service is about **one pod**, not two. Scaling replicas or adding an HPA would barely help.
- One pod takes all the load, so its latency degrades first.
- Not visible in error rates. It only shows up as uneven CPU or uneven span counts per pod.

### Resolution (step by step) — proposed
1. **Confirm in SigNoz:** Traces → `service.name = order-service` → group by `k8s.pod.name`. Expect nearly all spans on one pod.
2. **Headless Services** for `account-service`, `catalog-service` and `order-service`: `clusterIP: None`. (A headless Service can't change in place from ClusterIP, so Argo CD must replace it; the gateway's Service stays as it is, since the ALB uses pod IPs anyway.)
3. **Client policy:** add `grpc.WithDefaultServiceConfig(`{"loadBalancingConfig":[{"round_robin":{}}]}`)` to the three `NewClient` functions.
4. **Verify:** re-run the load test → both replicas of each service show similar CPU and span counts.

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
laptop$ python3 scripts/loadgen.py http://$ALB 180 110
…
FINAL: 12574 requests, 362 errors
  network errors (no GraphQL response): ConnectionResetError=1
  bad_account     ok=   0 err=181 p95=   840ms
  bad_order       ok=   0 err=180 p95=   985ms
  browse          ok=4381 err=  1 p95=  1132ms
  …
bastion$ kubectl -n go-micro-shop top pods          # during the run (trimmed)
NAME                               CPU(cores)   MEMORY(bytes)
catalog-service-78dfc9bbff-46w6r   82m          12Mi
catalog-service-78dfc9bbff-vmttf   37m          13Mi
graphql-gateway-5bddc7d95c-9s8g8   76m          14Mi
graphql-gateway-5bddc7d95c-ddvt5   74m          13Mi
order-service-7ff57b9f69-7585x     55m          13Mi
order-service-7ff57b9f69-78kzj     1m           10Mi
```
**Fixing it**
```console
laptop$ go test -count=1 ./account/... ./catalog/... ./order/... ./graphql/... ./telemetry/... ./lifecycle/...
ok  	github.com/ritesh-karankal/go-grpc-graphql-micro/account	0.015s
ok  	github.com/ritesh-karankal/go-grpc-graphql-micro/catalog	0.040s
ok  	github.com/ritesh-karankal/go-grpc-graphql-micro/order	0.056s      # PostOrder through round_robin clients
ok  	github.com/ritesh-karankal/go-grpc-graphql-micro/graphql	0.010s
ok  	github.com/ritesh-karankal/go-grpc-graphql-micro/lifecycle	0.867s
laptop$ kubectl kustomize k8s/overlays/prod > k-prod.yaml      # (checked with a small script)
Service account-service-headless: clusterIP=None selector={'app': 'account-service'}
Service catalog-service-headless: clusterIP=None selector={'app': 'catalog-service'}
Service order-service-headless: clusterIP=None selector={'app': 'order-service'}

# suggested, not run yet (verification), during a load run:
bastion$ kubectl -n go-micro-shop top pods | grep -E "NAME|order|catalog|account"
```

### Code / config change
- `account/client.go`, `catalog/client.go`, `order/client.go`:
  ```go
  grpc.WithDefaultServiceConfig(`{"loadBalancingConfig":[{"round_robin":{}}]}`),
  ```
- New Services `account-service-headless`, `catalog-service-headless`, `order-service-headless` (`clusterIP: None`, same selector and port), next to the existing ClusterIP Services.
- Gateway and order-service env: `*_SERVICE_URL` → `<name>-headless:8080`.
- Why **new** Services rather than changing the existing ones: `clusterIP` is immutable. Switching a Service to headless means delete + recreate, a window in which the name doesn't resolve. Adding new Services and moving clients over avoids it (expand → migrate → contract). The old ClusterIP Services can be removed once nothing uses them.
- The `PostOrder` tests now run through `round_robin` clients against real gRPC servers and still pass.

### Lessons learned
- Kubernetes Services load-balance **connections**. For long-lived protocols (gRPC, WebSockets, database pools) that is not request balancing.
- Uneven CPU across identical replicas is the tell-tale sign. Check per-pod numbers, not only per-service averages.

### How to approach it next time
1. Compare CPU per replica (`kubectl top pods`) and span counts per `k8s.pod.name` in SigNoz.
2. Check the client: DNS target, LB policy (`pick_first` default), connection reuse.
3. Check the Service: ClusterIP vs headless.

### Prevention / follow-up
- Apply the fix, then add an HPA for the gateway (now that metrics-server exists, INC-018).

### References
- Kubernetes: [Headless Services](https://kubernetes.io/docs/concepts/services-networking/service/#headless-services).
- gRPC: [Load balancing](https://grpc.io/docs/guides/load-balancing/); [Kubernetes blog: gRPC load balancing on Kubernetes without tears](https://kubernetes.io/blog/2018/11/07/grpc-load-balancing-on-kubernetes-without-tears/).

---

## INC-019 – SigNoz sizing: ZooKeeper heap larger than its memory limit, ClickHouse under-requested

| | |
|---|---|
| **Date** | 2026-10-09 |
| **Where** | `k8s/monitoring/signoz-values.yaml` (`clickhouse.resources`, `clickhouse.zookeeper`) |
| **Severity** | Medium: ZooKeeper could be OOM-killed (SigNoz then stops storing telemetry); one node ran at up to 88% CPU |
| **Status** | Resolved: `helm upgrade` applied; ZooKeeper runs with a 256 MB heap at ~203Mi (was 432Mi of 512Mi) |

### Summary
The 110-user load test showed one node at **78–88% CPU** while the other sat at ~25%. The app pods only used ~500m in total. `kubectl top pods -A` showed the heaviest workload is **SigNoz's ClickHouse**: **767m CPU and 1.2Gi memory even after the test ended**, while its request in our values was only 200m / 1Gi. The same look found a real bug: **ZooKeeper's Java heap is 1024 MB by default, but our values capped its container at 512Mi**. It was at 432Mi (84%), so one busy period away from being killed. Fixed in `signoz-values.yaml`: `heapSize: 256`, and ClickHouse requests raised to what was measured.

### Background
- **Requests vs limits:** the scheduler places pods by their **requests**; **limits** cap actual use. Requests far below real use make the scheduler overpack a node. A memory limit below what the process may use means an **OOM kill**.
- **ClickHouse** (SigNoz's database) writes incoming telemetry as small "parts", then **merges** them in the background. After a burst of ingestion it keeps using CPU for a while, which is why it was still at 767m after the test.
- **ZooKeeper** coordinates ClickHouse. It's a Java app: `ZOO_HEAP_SIZE` sets the heap (`-Xmx`). The JVM also needs non-heap memory (metaspace, threads, buffers), so the container limit must be comfortably above the heap.
- Per-span telemetry cost scales with traffic. SLIs come from **metrics** (pre-aggregated, cheap), which is one reason not to compute them from traces.

### Timeline
1. Load test with 110 users. `kubectl top nodes`: node A 1.5–1.7 cores (78–88%), node B 0.4–0.5 cores (21–27%).
2. App pods at peak: gateway ~75m ×2, catalog 82m+37m, order 55m, catalog-db 137m, so roughly 500m in total. Not enough to explain node A.
3. After the test, `kubectl top pods -A`: `chi-signoz-clickhouse… 767m 1228Mi`, `signoz-zookeeper-0 13m 432Mi`; everything else small.
4. Compared with `signoz-values.yaml`: ClickHouse requests 200m/1Gi, limit 2Gi; ZooKeeper limit 512Mi.
5. Rendered the chart: `ZOO_HEAP_SIZE=1024` (chart default `clickhouse.zookeeper.heapSize: 1024`). That heap can't fit in 512Mi.
6. Updated the values; re-rendered: `ZOO_HEAP_SIZE=256`; ClickHouse 500m / 1536Mi request, 3Gi limit.

### What happened (symptom)

    NAMESPACE  NAME                                  CPU(cores)  MEMORY(bytes)
    signoz     chi-signoz-clickhouse-cluster-0-0-0   767m        1228Mi      <- after the load test
    signoz     signoz-zookeeper-0                    13m         432Mi       <- limit 512Mi, heap allowed 1024 MB

    ip-10-10-x-a ...   1705m   88%     <- node running ClickHouse during the test
    ip-10-10-x-b ...    523m   27%

### How we got there
The SigNoz values were sized by guesswork before there was real traffic. The ZooKeeper limit was set without checking the chart's heap default.

### Why (root cause)
1. ClickHouse's real usage (~0.8+ cores under ingestion and merging) was far above its 200m request, so the scheduler treated its node as having free capacity it didn't have.
2. ZooKeeper's container limit (512Mi) was **smaller than its configured heap** (1024 MB). As the heap grows, the container hits the limit and gets OOM-killed.

### Impact
- No outage yet. Risks: ZooKeeper OOM → ClickHouse can't coordinate → telemetry not stored; a hot node → throttling or evictions of other pods on it (including app pods).
- Telemetry costs more CPU than the app at this traffic level: an important capacity-planning fact.

### Resolution (step by step)
1. **Find the heavy pods:** `kubectl top pods -A --sort-by=cpu | head` (needs metrics-server, INC-018).
2. **Compare usage with requests/limits** in the values file.
3. **Check JVM heap vs container limit** for Java workloads: render the chart and look for heap settings:
   ```bash
   helm template signoz signoz/signoz --version 0.145.0 -f k8s/monitoring/signoz-values.yaml | grep -A1 ZOO_HEAP_SIZE
   ```
4. **Fix the values** (diff below) and render again to confirm.
5. **Apply** (bastion):
   ```bash
   cd ~/go-grpc-graphql-micro && git pull
   helm upgrade signoz signoz/signoz -n signoz --version 0.145.0 -f k8s/monitoring/signoz-values.yaml
   kubectl -n signoz get pods -w      # zookeeper and clickhouse restart; ~2-5 min
   ```
   Telemetry pauses briefly while they restart; the services' exporters retry (see INC-013).
6. **Verify:** `kubectl -n signoz top pods` → ZooKeeper well under 512Mi; SigNoz shows new data.
   Result after the upgrade:
   ```
   $ kubectl -n signoz get pod signoz-zookeeper-0 -o jsonpath='{...ZOO_HEAP_SIZE...}'
   256
   signoz-zookeeper-0                    35m   203Mi      # was 432Mi
   chi-signoz-clickhouse-cluster-0-0-0   863m  1151Mi     # CPU still high: see follow-up
   ```

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
bastion$ kubectl top pods -A          # after the 110-user test (trimmed)
NAMESPACE     NAME                                  CPU(cores)   MEMORY(bytes)
signoz        chi-signoz-clickhouse-cluster-0-0-0   767m         1228Mi
signoz        signoz-zookeeper-0                    13m          432Mi
go-micro-shop catalog-db-0                          3m           790Mi
…             (everything else 1-15m)

laptop$ helm template signoz signoz/signoz --version 0.145.0 -n signoz -f k8s/monitoring/signoz-values.yaml | grep -nE "ZOO_HEAP_SIZE" -A1
1982:            - name: ZOO_HEAP_SIZE
1983-              value: "1024"
laptop$ grep -n "heapSize" signoz/charts/clickhouse/charts/zookeeper/values.yaml
165:## @param heapSize Size (in MB) for the Java Heap options (Xmx and Xms)
168:heapSize: 1024
```
**Fixing it**
```console
laptop$ helm template signoz signoz/signoz --version 0.145.0 -n signoz -f k8s/monitoring/signoz-values.yaml > r3.yaml && grep -n "ZOO_HEAP_SIZE" -A1 r3.yaml | tail -1
1983-              value: "256"
StatefulSet signoz-zookeeper cpu ['50m'] mem ['512Mi', '384Mi']
ClickHouseInstallation signoz-clickhouse cpu ['500m'] mem ['3Gi', '1536Mi']

bastion$ helm upgrade signoz signoz/signoz -n signoz --version 0.145.0 -f k8s/monitoring/signoz-values.yaml
(output not captured)
bastion$ kubectl -n signoz get pod signoz-zookeeper-0 -o jsonpath='{.spec.containers[0].env[?(@.name=="ZOO_HEAP_SIZE")].value}'; echo
256
bastion$ kubectl -n signoz top pods
NAME                                          CPU(cores)   MEMORY(bytes)
chi-signoz-clickhouse-cluster-0-0-0           863m         1151Mi
signoz-0                                      3m           40Mi
signoz-clickhouse-operator-55fb8db98d-jhftc   1m           23Mi
signoz-otel-collector-5cfc9ccdd8-5blp7        9m           79Mi
signoz-zookeeper-0                            35m          203Mi
```

### Code / config change
`k8s/monitoring/signoz-values.yaml`:
```diff
 clickhouse:
   resources:
     requests:
-      cpu: 200m
-      memory: 1Gi
+      cpu: 500m
+      memory: 1536Mi
     limits:
-      memory: 2Gi
+      memory: 3Gi
   zookeeper:
+    heapSize: 256
     resources:
       requests:
         cpu: 50m
-        memory: 256Mi
+        memory: 384Mi
       limits:
         memory: 512Mi
```

### Lessons learned
- Size from **measurements under load**, not guesses. A load test is a capacity test for the observability stack too.
- For Java workloads, the container memory limit must exceed the configured heap plus overhead; check the chart's defaults.
- Telemetry has a cost. At high volume, consider trace **sampling** (SLIs stay accurate because they come from unsampled metrics) and giving SigNoz its own node.

### How to approach it next time
1. Node hot? `kubectl top pods -A --sort-by=cpu` and `kubectl get pods -A -o wide` to see which pods sit on that node.
2. Compare top consumers' usage with their requests (`kubectl get pod <p> -o jsonpath='{.spec.containers[*].resources}'`).
3. Java pods: compare heap flags (`-Xmx`, `*_HEAP_SIZE`) with the memory limit.

### Prevention / follow-up
- Optional: a dedicated node (or node group with a taint) for SigNoz, so observability load can't starve the app.
- Optional: trace sampling (`OTEL_TRACES_SAMPLER=parentbased_traceidratio`, e.g. 0.25) if traffic grows.
- Open question: ClickHouse still used ~860m CPU after the fix. Check whether that's ingestion during a load run or a constant background cost (merges, ClickHouse's own `system.*_log` tables).

### References
- Kubernetes: [Resource requests and limits](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/).
- ClickHouse: [MergeTree background merges](https://clickhouse.com/docs/en/engines/table-engines/mergetree-family/mergetree).
- SigNoz chart: `clickhouse.zookeeper.heapSize` (Bitnami ZooKeeper chart, `ZOO_HEAP_SIZE`).

---

## INC-018 – `kubectl top` fails: `Metrics API not available`

| | |
|---|---|
| **Date** | 2026-10-09 |
| **Where** | Bastion, `kubectl top pods`; `terraform/eks-cluster/addons.tf` |
| **Severity** | Low: no pod CPU/memory from `kubectl`; it would also block HorizontalPodAutoscalers |
| **Status** | Fix ready: metrics-server EKS add-on added to Terraform; waiting for the `eks-cluster` pipeline apply |

### Summary
Before a load ramp, `kubectl top pods` was going to be used to watch CPU throttling. It failed with `error: Metrics API not available`. EKS does not install **metrics-server** by default, and the cluster's Terraform only installed the VPC CNI, kube-proxy, CoreDNS, Pod Identity agent and EBS CSI add-ons. AWS publishes metrics-server as a managed EKS add-on, so it was added to `addons.tf` next to the others.

### Background
- `kubectl top` and HorizontalPodAutoscalers (HPA) read CPU/memory from the **Metrics API** (`metrics.k8s.io`). That API is served by **metrics-server**, which polls each node's kubelet. Without it, the API doesn't exist.
- This is separate from Prometheus: Prometheus scrapes the same kubelet/cAdvisor data for its own storage (Grafana dashboards), but doesn't serve the Metrics API.
- EKS installs only the essentials. Everything else is opt-in, as an **EKS add-on** (managed by AWS, upgraded with the cluster), a Helm chart, or plain manifests.

### Timeline
1. Planning a stepped load test (10 → 25 → 50 → 110 users), watching CPU per pod between steps.
2. Bastion: `kubectl -n go-micro-shop top pods` → `error: Metrics API not available`; the same for `-n signoz`.
3. `aws eks list-addons` → no metrics-server.
4. `aws eks describe-addon-versions --addon-name metrics-server --kubernetes-version 1.36` → available, publisher `eks`, default `v0.9.0-eksbuild.11`.
5. Added `aws_eks_addon.metrics_server` to `addons.tf`; `terraform fmt -check` and `terraform validate` pass.

### What happened (symptom)

    [ec2-user@bastion]$ kubectl -n go-micro-shop top pods
    error: Metrics API not available
    [ec2-user@bastion]$ kubectl -n signoz top pods
    error: Metrics API not available

### How we got there
The cluster was built with the minimum add-ons. Nothing needed the Metrics API until now.

### Why (root cause)
1. `kubectl top` depends on the Metrics API.
2. The Metrics API is provided by metrics-server.
3. metrics-server isn't part of a default EKS cluster and wasn't in `addons.tf`.

### Impact
- No quick CPU/memory view from `kubectl` during load tests.
- HPA (autoscaling on CPU) would not work: a prerequisite for scaling the gateway.
- Workaround meanwhile: Prometheus + Grafana (dashboard `17375`) show per-pod CPU/memory from cAdvisor.

### Resolution (step by step)
1. **Confirm it's missing:** `aws eks list-addons --region eu-north-1 --cluster-name go-microservices-dev`.
2. **Check that the managed add-on exists** for the cluster version:
   `aws eks describe-addon-versions --region eu-north-1 --addon-name metrics-server --kubernetes-version 1.36`.
3. **Add it to Terraform** (diff below), next to the other add-ons, so it's created and destroyed with the cluster.
4. **Apply through the pipeline:** commit and push, then Jenkins `eks-cluster` → Build with Parameters → `dev` + `apply`. The plan should show `1 to add, 0 to change, 0 to destroy` → Proceed.
5. **Verify** (bastion, about a minute after apply):
   ```bash
   kubectl -n kube-system get deploy metrics-server     # READY
   kubectl top nodes
   kubectl -n go-micro-shop top pods
   ```

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
[ec2-user@ip-10-10-20-x bin]$ kubectl -n go-micro-shop top pods
error: Metrics API not available
[ec2-user@ip-10-10-20-x bin]$ kubectl -n signoz top pods
error: Metrics API not available

laptop$ aws eks list-addons --region eu-north-1 --cluster-name go-microservices-dev --output text
ADDONS	aws-ebs-csi-driver
ADDONS	coredns
ADDONS	eks-pod-identity-agent
ADDONS	kube-proxy
ADDONS	vpc-cni
laptop$ aws eks describe-addon-versions --region eu-north-1 --addon-name metrics-server --kubernetes-version 1.36 --query '…' --output json
{
    "name": "metrics-server",
    "publisher": "eks",
    "versions": [ "v0.9.0-eksbuild.11", "v0.9.0-eksbuild.10", "v0.9.0-eksbuild.9" ],
    "default": "v0.9.0-eksbuild.11"
}
```
**Fixing it**
```console
laptop$ cd terraform/eks-cluster && terraform fmt -check addons.tf && echo fmt-ok; terraform init -backend=false -input=false >/dev/null 2>&1 && terraform validate -no-color
fmt-ok
Success! The configuration is valid.
# Jenkins UI: eks-cluster -> dev + apply -> Plan: 1 to add -> Proceed

[ec2-user@ip-10-10-20-x go-grpc-graphql-micro]$ kubectl -n kube-system get deploy metrics-server
NAME             READY   UP-TO-DATE   AVAILABLE   AGE
metrics-server   2/2     2            2           3m33s
[ec2-user@ip-10-10-20-x go-grpc-graphql-micro]$ kubectl top nodes
NAME                                          CPU(cores)   CPU(%)   MEMORY(bytes)   MEMORY(%)
ip-10-10-x-x.eu-north-1.compute.internal      865m         44%      4289Mi          60%
ip-10-10-x-x.eu-north-1.compute.internal      162m         8%       2514Mi          35%
```

### Code / config change
`terraform/eks-cluster/addons.tf`:
```hcl
# Metrics Server: CPU/memory per pod and node for `kubectl top` and for
# HorizontalPodAutoscalers. Needs nodes to schedule onto.
resource "aws_eks_addon" "metrics_server" {
  cluster_name                = aws_eks_cluster.eks.name
  addon_name                  = "metrics-server"
  resolve_conflicts_on_create = "OVERWRITE"
  resolve_conflicts_on_update = "OVERWRITE"

  depends_on = [aws_eks_node_group.nodes]
}
```

### Lessons learned
- A "minimal" EKS cluster lacks things people assume are built in (metrics-server, an ingress controller, a default StorageClass before the EBS CSI add-on).
- Prefer the **managed EKS add-on** over a hand-installed Helm chart when one exists: AWS keeps it compatible with the cluster version, and Terraform owns it.

### How to approach it next time
- `Metrics API not available` → `kubectl get apiservice v1beta1.metrics.k8s.io` (missing, or `False`) → install or fix metrics-server.
- If it is installed but unavailable: `kubectl -n kube-system logs deploy/metrics-server` (often kubelet TLS or network-policy issues).

### Prevention / follow-up
- Done: add-on in Terraform.
- Next: an HPA for `graphql-gateway` (CPU-based) once the load ramp shows the bottleneck.

### References
- Amazon EKS: [Metrics Server add-on](https://docs.aws.amazon.com/eks/latest/userguide/metrics-server.html); [EKS add-ons](https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html).
- Kubernetes: [Resource metrics pipeline](https://kubernetes.io/docs/tasks/debug/debug-cluster/resource-metrics-pipeline/).

---

## INC-017 – `Connection reset by peer` during a rollout: pods exit without draining

| | |
|---|---|
| **Date** | 2026-10-09 |
| **Where** | `telemetry/telemetry.go` › `ShutdownOnSignal`; the four Go services' `main`; Deployments in `k8s/base/` (no `preStop`, no grace period); ALB target group |
| **Severity** | Medium: some requests fail on **every deploy**, which burns the SLO error budget each time Jenkins ships |
| **Status** | Fix ready: graceful shutdown (`lifecycle` package), `preStop` + grace period, ALB readiness gates and deregistration delay; tested locally, waiting for build + deploy |

### Summary
During a load test, a few requests failed with `ConnectionResetError: [Errno 104] Connection reset by peer`. The run overlapped with Argo CD rolling out backend version 6. When Kubernetes replaces a pod, it sends `SIGTERM` while it is still removing the pod from the Service and the ALB. Our shutdown handler (`telemetry.ShutdownOnSignal`) flushes telemetry and calls `os.Exit(0)` straight away, without finishing in-flight requests or waiting for traffic to stop. Requests in that window were cut off. The same run also exposed a bug in the load generator: an unhandled reset crashed the simulated user's thread.

### Background
- **What Kubernetes does when it replaces a pod** (rolling update): at the same moment it (a) sends `SIGTERM` to the container (after running any `preStop` hook) and (b) starts removing the pod from Service endpoints. For an ALB with `target-type: ip`, it also starts **deregistering** the pod's IP from the target group. (b) takes seconds to spread (kube-proxy on every node, the AWS Load Balancer Controller, the ALB itself), so **new requests can still arrive after `SIGTERM`**.
- The expected app behaviour on `SIGTERM`: keep serving briefly, stop accepting new connections, **finish in-flight requests**, then exit within `terminationGracePeriodSeconds` (default 30 s).
- A `preStop` hook (e.g. sleep 15 s) delays `SIGTERM`, giving the ALB and kube-proxy time to stop routing to the pod first. Kubernetes has a built-in `sleep` action for this, so no shell is needed in the image.
- The **AWS Load Balancer Controller's pod readiness gates** make a new pod count as Ready only once the ALB reports it healthy, so a rollout doesn't remove old pods before the new ones actually receive traffic.

### Timeline
1. 12:48:53 IST: the `backend` job pushed images `:6`; 12:48:55, the deploy commit landed on `main`.
2. ~12:49–12:52: Argo CD synced and rolled every backend Deployment (2 replicas each, one pod at a time).
3. A load-generator run (4 users) was in progress. Several requests failed with `Connection reset by peer`. Two of the simulated users' threads crashed with a traceback (load-generator bug, see below).
4. Matched the timing: the errors fell in the rollout window.
5. Read the shutdown path: `ShutdownOnSignal` → flush telemetry → `os.Exit(0)`. No graceful stop of the HTTP or gRPC servers. Manifests: no `preStop`, no `terminationGracePeriodSeconds`.

### What happened (symptom)

    File ".../scripts/loadgen.py", line 58, in gql
      with urllib.request.urlopen(req, timeout=15) as r:
    ...
    ConnectionResetError: [Errno 104] Connection reset by peer

Load generator summary afterwards: `progress: 1907 requests, 57 errors` (mostly the deliberate bad requests), with fewer active users after the crashed threads.

### How we got there
A deploy during live traffic, the normal case for a CI/CD pipeline. It hadn't shown before because earlier load runs didn't overlap a rollout.

### Why (root cause)
1. A rolling update terminates old pods while traffic is still routed to them for a few seconds.
2. On `SIGTERM`, `ShutdownOnSignal` calls `os.Exit(0)` right after flushing telemetry: in-flight requests are dropped and new ones hit a closed socket, so the connection is reset.
3. No `preStop` delay, so `SIGTERM` arrives before the ALB and kube-proxy have stopped sending traffic.
4. (Load generator) `gql()` caught `URLError` but not a reset raised while *reading* the response (`ConnectionResetError`, an `OSError`), so the user thread died.

### Impact
- A handful of failed requests per rollout under load. With real users, every deploy would burn some error budget, invisible except as "random" resets.
- Load test: two simulated users stopped early, so the run produced less traffic than configured.

### Resolution (step by step)
Done:
1. **Load generator** catches `OSError` and `http.client.HTTPException`, counts them as errors, and lists network errors on their own line. Tested against a local server that resets every connection: no tracebacks, users keep running, correct totals (`10 requests, 10 errors` / `ConnectionResetError=8, URLError=2`).

Proposed (to implement):
2. **Graceful shutdown in the apps:** on `SIGTERM` → stop accepting (`http.Server.Shutdown` in the gateway, `grpc.Server.GracefulStop` in the services) → wait for in-flight requests → flush telemetry → exit.
3. **Manifests:** `lifecycle.preStop.sleep.seconds: 15` and `terminationGracePeriodSeconds: 30` on the four Go Deployments.
4. **ALB:** enable pod readiness gates (label the namespace `elbv2.k8s.aws/pod-readiness-gate-inject=enabled`) and shorten the target group's deregistration delay (e.g. 30 s) via the Ingress annotation `alb.ingress.kubernetes.io/target-group-attributes`.
5. **Verify:** run the load generator, trigger `kubectl -n go-micro-shop rollout restart deploy`, and expect **0 network errors** and no `server_error` spike in SigNoz.

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
laptop$ python3 scripts/loadgen.py http://$ALB 180 110
…
  File "/usr/lib/python3.14/http/client.py", line 297, in _read_status
    line = str(self.fp.readline(_MAXLINE + 1), "iso-8859-1")
  File "/usr/lib/python3.14/socket.py", line 725, in readinto
    return self._sock.recv_into(b)
ConnectionResetError: [Errno 104] Connection reset by peer
progress: 1907 requests, 57 errors
…
laptop$ git log --format='%h %ci %s' origin/main -4
b936b20 2026-10-09 07:18:55 +0000 chore(deploy): update backend images to 6
d065223 2026-10-09 12:43:08 +0530 feat(observability): SLI labels on root span; mark synthetic traffic
…
laptop$ aws ecr describe-images --region eu-north-1 --repository-name graphql --query 'sort_by(imageDetails,&imagePushedAt)[].[imageTags[0],imagePushedAt]' --output text
…
6	2026-10-09T12:48:53.764000+05:30
laptop$ sed -n '/func ShutdownOnSignal/,/^}/p' telemetry/telemetry.go
func ShutdownOnSignal(shutdown func(context.Context) error) {
	…
		<-ch
		…
		if err := shutdown(ctx); err != nil { … }
		os.Exit(0)
	}()
}
laptop$ grep -n "terminationGracePeriod\|preStop\|lifecycle" k8s/base/gateway/graphql-gateway.yaml k8s/base/microservices/*.yaml || echo "no preStop / grace period configured"
no preStop / grace period configured
```
**Fixing it**
```console
laptop$ # load generator against a local server that resets every connection
exit 0 tracebacks: 0
FINAL: 10 requests, 10 errors
  network errors (no GraphQL response): ConnectionResetError=8, URLError=2

laptop$ go test ./lifecycle/ -v
--- PASS: TestServeHTTPDrainsInFlightRequests (0.56s)
--- PASS: TestServeGRPCDrainsInFlightRPCs (0.30s)
ok  	github.com/ritesh-karankal/go-grpc-graphql-micro/lifecycle	0.870s
laptop$ # same tests with a hard stop swapped in (what the old code effectively did)
    lifecycle_test.go:54: in-flight request got "error: Get \"http://127.0.0.1:35179\": EOF", want "done"
--- FAIL: TestServeHTTPDrainsInFlightRequests (0.02s)
    lifecycle_test.go:101: in-flight RPC failed: rpc error: code = Unavailable desc = error reading from server: EOF
--- FAIL: TestServeGRPCDrainsInFlightRPCs (0.00s)

laptop$ go build -o gw ./graphql && ACCOUNT_SERVICE_URL=127.0.0.1:1 CATALOG_SERVICE_URL=127.0.0.1:1 ORDER_SERVICE_URL=127.0.0.1:1 METRICS_ADDR=127.0.0.1:19464 ./gw &
laptop$ curl -s -o /dev/null -w 'before: %{http_code}\n' localhost:8080/playground
before: 200
laptop$ kill -TERM $(pgrep -fx "./gw")
exited after SIGTERM in ~200ms
laptop$ cat gw.log
2026/10/09 15:53:46 INFO Serving Prometheus metrics addr=127.0.0.1:19464 path=/metrics
2026/10/09 15:53:46 INFO OTLP export disabled, OTEL_EXPORTER_OTLP_ENDPOINT is not set
2026/10/09 15:53:57 Stopped

laptop$ for s in account catalog order graphql; do docker build -q -t lc-check/$s -f $s/app.dockerfile . >/dev/null 2>&1 && echo "$s OK" || echo "$s FAIL"; done
account OK
catalog OK
order OK
graphql OK
```

### Code / config change
- `scripts/loadgen.py`: exception handling and reporting.
- **New package `lifecycle`:** `SignalContext()` (cancelled on SIGTERM / Ctrl+C), `ServeGRPC(ctx, server, listener)` (`GracefulStop`, forced `Stop` after `DrainTimeout` = 10 s), `ServeHTTP(ctx, server)` (`Shutdown` with the same timeout).
- **The four `main` functions:** signal context → serve until cancelled → drain → `flush(shutdown)` (telemetry, 5 s timeout) → exit. `telemetry.ShutdownOnSignal` (flush, then `os.Exit(0)`) removed.
- **Gateway:** its own `http.Server` (with `ReadHeaderTimeout: 10s`) instead of the global `http.ListenAndServe`.
- **order-service:** closes its account and catalog clients *after* draining, because in-flight RPCs still use them.
- **Dockerfiles** (catalog, order, graphql): `COPY lifecycle`.
- **Deployments** (account, catalog, order, gateway): `terminationGracePeriodSeconds: 30` and `lifecycle.preStop.sleep.seconds: 15`. Budget: 15 s preStop + ≤10 s drain + ≤5 s flush = 30 s.
- **Namespace:** label `elbv2.k8s.aws/pod-readiness-gate-inject: enabled`.
- **Ingress:** `alb.ingress.kubernetes.io/target-group-attributes: deregistration_delay.timeout_seconds=30` (default 300).

Tests (`lifecycle/lifecycle_test.go`): a request (HTTP) and an RPC (gRPC) still running when shutdown starts must complete. **Checked against the old behaviour:** with a hard stop instead of draining, both fail (`EOF`, `Unavailable … EOF`); with draining both pass. The real gateway binary, sent SIGTERM, logs `Stopped` and exits within ~200 ms when idle.

Note for the first deploy of this fix: pods being replaced still run the **old** code and the old pod spec (no `preStop`), so that one rollout can still reset a few requests. From the next rollout on, the fix applies.

### Lessons learned
- `os.Exit` in a signal handler is almost never right for a server: it skips draining.
- Graceful shutdown in Kubernetes needs both sides: the app drains, and the platform (`preStop`, readiness gates) stops sending traffic first.
- Deploys are a major source of SLO burn. Test "load during rollout" deliberately.
- A load tester has to survive the failures it is meant to measure.

### How to approach it next time
1. Resets or 502s during a deploy? Line up the error times with the rollout (`kubectl rollout history`, the Argo CD sync time, image push time).
2. Read the app's `SIGTERM` handling: does it drain, or exit?
3. Check the Deployment for `preStop` and `terminationGracePeriodSeconds`, and the ALB target group's deregistration delay.
4. Reproduce on purpose: load + `kubectl rollout restart`.

### Prevention / follow-up
- Implement steps 2–4, then make "load during rollout" part of the fault drill (C9).

### References
- Kubernetes: [Pod termination](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-termination); [Container lifecycle hooks — sleep action](https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/).
- AWS Load Balancer Controller: [Pod readiness gate](https://kubernetes-sigs.github.io/aws-load-balancer-controller/latest/deploy/pod_readiness_gate/).
- Go: [`http.Server.Shutdown`](https://pkg.go.dev/net/http#Server.Shutdown); [`grpc.Server.GracefulStop`](https://pkg.go.dev/google.golang.org/grpc#Server.GracefulStop).

---

## INC-016 – Backend quality gate `ERROR`: SonarQube saw 0% coverage on new code

| | |
|---|---|
| **Date** | 2026-10-09 |
| **Where** | Jenkins job `backend` › stage **Quality Check** (`waitForQualityGate abortPipeline: true`); SonarQube project `go-grpc-graphql-micro-backend`; `jenkins/Jenkinsfile-Backend` |
| **Severity** | Medium: the gate blocked the build (by design), so the SLI/metrics change could not ship |
| **Status** | Resolved: build #5 passed the quality gate with the coverage report; images `:5` deployed |

### Summary
The first backend build after the "SLI-ready errors and GraphQL operation metrics" commit stopped at **Quality Check**: SonarQube's quality gate returned `ERROR`. Its per-file view showed every new line as **uncovered**, 92 lines across the six changed files, although the commit added unit tests. The pipeline ran `go test` but **never produced or uploaded a coverage report**, so SonarQube counted new-code coverage as 0%, below the *Sonar way* gate's 80%. Earlier builds had passed only because there was no "new code" yet. Fix: generate `coverage.out` in the test stage, pass it to the scanner, exclude startup wiring from coverage, and add the missing tests for `order-service`'s `PostOrder`.

### Background
- A **quality gate** is a set of conditions SonarQube checks after each analysis. The default gate, *Sonar way*, checks **new code only** (code changed since the baseline), including **coverage ≥ 80%**, duplications ≤ 3%, A ratings, and hotspots reviewed.
- On a project's **first analysis** everything is "overall code" and nothing is "new", so new-code conditions pass trivially. That's why the first backend builds passed with no tests at all.
- SonarQube **doesn't run tests**. It only reads a coverage report the build produces. For Go that is `go test -coverprofile=coverage.out`, passed as `-Dsonar.go.coverage.reportPaths=coverage.out`. Without it, every new executable line counts as uncovered.
- `waitForQualityGate abortPipeline: true` fails the Jenkins stage when the gate is `ERROR`, so the images are never built or pushed. The gate does what it's meant to do.

### Timeline
1. Pushed commit `feat: SLI-ready errors and GraphQL operation metrics` (gRPC status codes, gateway operation metrics, a catalog crash fix, and the repo's first unit tests).
2. Ran `backend`. Tests passed; then **Quality Check**:
   `SonarQube task '…' completed. Quality gate is 'ERROR'`.
3. The SonarQube UI didn't load (see INC-015, IP allowlist). After fixing access, the project's new-code coverage view listed uncovered lines per file: `operation_metrics.go 50`, `order/server.go 20`, `catalog/server.go 8`, `account/server.go 8`, `graphql/main.go 4`, `catalog/repository.go 2`. Every changed file, nothing covered.
4. Read the Jenkinsfile: `go test` without `-coverprofile`, and no `sonar.go.coverage.reportPaths` → SonarQube had no coverage data.
5. Estimated locally what SonarQube would see with a report (coverage of changed lines only): **71%**, still below 80%. Biggest gap: `order/server.go`'s `PostOrder`, 8 of 29 lines covered, because it calls the account and catalog services and had no test.
6. Added a `PostOrder` test with fake account and catalog gRPC servers on localhost (7 cases), excluded `main.go`/`cmd/` from coverage, and added the coverage report to the pipeline. Local estimate: **87%**.

### What happened (symptom)

    Checking status of SonarQube task '<task>' on server 'sonar-server'
    SonarQube task '<task>' status is 'IN_PROGRESS'
    SonarQube task '<task>' status is 'SUCCESS'
    SonarQube task '<task>' completed. Quality gate is 'ERROR'

SonarQube, uncovered lines on new code:

| File | Uncovered new lines |
|---|---|
| `graphql/operation_metrics.go` | 50 |
| `order/server.go` | 20 |
| `catalog/server.go` | 8 |
| `account/server.go` | 8 |
| `graphql/main.go` | 4 |
| `catalog/repository.go` | 2 |
| all other files | 0 (not changed) |

### How we got there
The first commit with real changes to existing code since the project's first analysis. The CI pipeline had been set up before there were any tests, so it never produced coverage.

### Why (root cause)
1. *Sonar way* requires **≥ 80% coverage on new code**.
2. The commit added about 150 new executable lines.
3. Jenkins ran `go test` **without** `-coverprofile`, and the scanner had no `sonar.go.coverage.reportPaths`, so SonarQube had **no coverage data** and treated every new line as uncovered (0%).
4. Even with a report, coverage would have been ~71%: `PostOrder`'s error paths had no tests, and `graphql/main.go` (startup wiring) counted as coverable.

### Impact
- The backend build stopped before Docker build and ECR push, so the SLI and metrics change wasn't deployed. Nothing running was affected.
- About 30 minutes (including INC-015).

### Resolution (step by step)
1. **Read the gate result** in SonarQube → project → **Overview** (failed conditions) and the coverage measures per file. All changed files uncovered → no coverage report.
2. **Check the pipeline:** no `-coverprofile`, no `sonar.go.coverage.reportPaths`.
3. **Estimate new-code coverage locally** before changing anything, so you know whether a report alone is enough:
   ```bash
   go test -coverprofile=coverage.out ./account/... ./catalog/... ./order/... ./graphql/... ./telemetry/...
   go tool cover -func=coverage.out | tail -1        # overall
   ```
   (Per-changed-line coverage was computed by matching `coverage.out` blocks to `git diff -U0` line ranges.) Result: 71%, not enough.
4. **Test the real gap:** `order/post_order_test.go` runs fake account and catalog gRPC servers on `127.0.0.1:0` and drives `PostOrder` through the real clients: order placed, empty order, unknown account, account service failing, catalog unreachable, no existing products, save fails. All 7 pass.
5. **Exclude startup wiring** from coverage (`**/main.go`, `**/cmd/**`): it only connects components and is exercised by running the service, not by unit tests.
6. **Pipeline change** (diff below). Local estimate of new-code coverage: **87%** (≥ 80%).
7. **Verify:** push, run `backend` → Quality Check `Quality gate is 'OK'`; SonarQube shows a non-zero coverage on new code.
   Result: build #5 passed the gate, pushed `account/catalog/order/graphql:5`, committed the tag, and Argo CD deployed it. (Later, after INC-021, SonarQube showed 84.9% coverage on new code.) The live API now answers an unknown account with `code = NotFound` (before: `code = Unknown`).

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
# Jenkins console, Quality Check:
Checking status of SonarQube task '<task>' on server 'sonar-server'
SonarQube task '<task>' status is 'IN_PROGRESS'
SonarQube task '<task>' status is 'SUCCESS'
SonarQube task '<task>' completed. Quality gate is 'ERROR'

# SonarQube UI -> uncovered lines on new code, per file:
graphql/operation_metrics.go 50 · order/server.go 20 · catalog/server.go 8 · account/server.go 8 · graphql/main.go 4 · catalog/repository.go 2 · all others 0

laptop$ go test -coverprofile=coverage.out ./account/... ./catalog/... ./order/... ./graphql/... ./telemetry/...
laptop$ # changed lines (git diff -U0 f94e9a1 HEAD) matched against coverage.out
account/server.go                  9/ 10 covered
catalog/repository.go              3/  3 covered
catalog/server.go                  9/ 11 covered
graphql/main.go                    0/  6 covered
graphql/operation_metrics.go      78/ 91 covered
order/server.go                    8/ 29 covered
TOTAL new executable lines: 107/150 = 71%
```
**Fixing it**
```console
laptop$ go test ./order/ -run TestPostOrder -v
=== RUN   TestPostOrder/order_placed
=== RUN   TestPostOrder/empty_order
=== RUN   TestPostOrder/unknown_account
2026/10/09 11:37:06 WARN Failed to get account err="rpc error: code = NotFound desc = account not found"
=== RUN   TestPostOrder/account_service_failing
2026/10/09 11:37:06 ERROR Failed to get account err="rpc error: code = Internal desc = db down"
=== RUN   TestPostOrder/catalog_unreachable
2026/10/09 11:37:06 ERROR Failed to get products err="rpc error: code = Unavailable desc = es down"
=== RUN   TestPostOrder/none_of_the_products_exist
=== RUN   TestPostOrder/saving_the_order_fails
--- PASS: TestPostOrder (0.03s)
ok  	github.com/ritesh-karankal/go-grpc-graphql-micro/order	0.047s
laptop$ # re-estimate with main.go/cmd excluded
order/server.go                   26/ 29
TOTAL new-code coverage: 125/144 = 87%  (gate: >= 80%)
# Jenkins UI: backend -> build #5 -> Quality gate is 'OK'
```

### Code / config change
`jenkins/Jenkinsfile-Backend`:
```diff
         stage('Go Unit Tests') {
             steps {
-                sh 'go test ./account/... ./catalog/... ./order/... ./graphql/... ./telemetry/...'
+                // coverage.out is read by the SonarQube scan (sonar.go.coverage.reportPaths)
+                sh 'go test -coverprofile=coverage.out ./account/... ./catalog/... ./order/... ./graphql/... ./telemetry/...'
             }
         }
 ...
                           -Dsonar.exclusions=**/pb/**,**/generated.go,**/models_gen.go \
+                          -Dsonar.go.coverage.reportPaths=coverage.out \
+                          -Dsonar.coverage.exclusions=**/main.go,**/cmd/** \
```
New test: `order/post_order_test.go`.

### Lessons learned
- SonarQube only knows coverage the build tells it about. "Tests pass" and "coverage reported" are separate things.
- A quality gate that passed on day one may only have passed because there was no new code yet. The first real change is the first real test of the gate.
- Fix a coverage gate by **testing the untested logic**, not by lowering the threshold. Excluding pure wiring (`main.go`) is fine; excluding business logic is not.
- Estimating locally first avoids a slow cycle of "push → wait 10 minutes → gate fails again".

### How to approach it next time
1. Quality gate `ERROR` → SonarQube project → Overview → read the **failed conditions** (coverage, duplication, ratings, hotspots).
2. Coverage at exactly 0% on new code → the report is missing or not found (check `sonar.go.coverage.reportPaths` and the scanner log for "coverage").
3. Coverage low but non-zero → open **Measures → Coverage → Uncovered lines on new code**, and write tests for the biggest files first.
4. Paths in `coverage.out` are Go import paths; SonarQube maps them with `go.mod`. If coverage stays 0 with a report present, check the scanner log for "unable to resolve" warnings.

### Prevention / follow-up
- Done: coverage report in the pipeline; `PostOrder` tests.
- Follow-up: the same for the frontend pipeline when it gets tests (`lcov` report → `sonar.javascript.lcov.reportPaths`).
- Follow-up: archive `coverage.out` as a Jenkins artifact, or publish an HTML report.

### References
- SonarQube: [Test coverage for Go](https://docs.sonarsource.com/sonarqube-server/latest/analyzing-source-code/test-coverage/go-test-coverage/); [Quality gates — Sonar way](https://docs.sonarsource.com/sonarqube-server/latest/instance-administration/analysis-functions/quality-gates/).
- Go: [`go test -coverprofile`](https://pkg.go.dev/cmd/go#hdr-Testing_flags).

---

## INC-015 – SonarQube UI on port 9000 keeps loading: IP allowlist out of date

| | |
|---|---|
| **Date** | 2026-10-09 |
| **Where** | Browser → `http://<JENKINS_IP>:9000` (SonarQube); `terraform/jenkins-server` security group `jenkins_sg` (`admin_cidrs`) |
| **Severity** | Low: blocked access to the SonarQube UI (needed to read a failed quality gate); Jenkins and the pipelines were unaffected |
| **Status** | Resolved: `terraform apply` updated the 9000 rule in place; SonarQube reachable again |

### Summary
The SonarQube UI stopped loading: the browser just spun until it timed out. Jenkins on port 8080 on the same server answered immediately. The server's security group lets **anyone** reach 8080 (GitHub webhooks need it) but only **one admin IP** reach 9000. The laptop's public IP had **changed** since the server was created (home ISPs reassign IPs), so its connections to 9000 were silently dropped. The fix is to put the new IP in `terraform.tfvars` and apply, so the code and AWS stay in sync.

### Background
- The Jenkins server's security group (`terraform/jenkins-server`) has two inbound rules:

  | Port | Service | Allowed from | Why |
  |---|---|---|---|
  | 8080 | Jenkins | `0.0.0.0/0` | GitHub webhooks come from GitHub's IPs |
  | 9000 | SonarQube | `admin_cidrs` (our IP `/32`) | SonarQube's UI should not be public |

- A security group **drops** traffic it doesn't allow. It doesn't reject it, so the client gets no answer and waits until it times out. "Takes forever to load" is the typical symptom; "connection refused" would mean something else.
- Home and mobile internet IPs are **dynamic**: the ISP can assign a new one after a router restart or after some hours or days. Every IP allowlist (this security group, and the `loadBalancerSourceRanges` of the Argo CD, Prometheus, Grafana and SigNoz load balancers) then stops matching.

### Timeline
1. Day 1: Jenkins server created with `admin_cidrs = ["<OLD_IP>/32"]`. SonarQube reachable.
2. Day 2: the backend pipeline failed its quality gate. Opened SonarQube to see which condition failed; the page kept loading.
3. From the laptop: `curl :9000` → no answer after 10 s; `curl :8080` → `403` in 0.5 s (Jenkins up; 403 is its login page for anonymous requests).
4. Read the security group: port 9000 allows only `<OLD_IP>/32`. Current laptop IP: different (same first two octets, different last two).
5. Updated `terraform.tfvars`; `terraform plan` → `0 to add, 1 to change, 0 to destroy` (in-place update of the 9000 rule).

### What happened (symptom)
Browser: `http://<JENKINS_IP>:9000` loads forever, then times out.

    $ curl -s -o /dev/null -w '%{http_code} in %{time_total}s' --max-time 10 http://<JENKINS_IP>:9000/
    000 in 10.0s        # curl exit 28: timeout, nothing answered
    $ curl -s -o /dev/null -w '%{http_code} in %{time_total}s' --max-time 10 http://<JENKINS_IP>:8080/
    403 in 0.54s        # Jenkins answers (anonymous → 403)

### How we got there
The laptop's ISP assigned a new public IP between yesterday's setup and today.

### Why (root cause)
1. Port 9000 is allowlisted to a single `/32`, our IP at creation time.
2. The laptop's public IP changed.
3. The security group dropped packets from the new IP, so the browser waited with no response.

### Impact
- SonarQube UI unreachable, so the quality-gate failure couldn't be inspected.
- Jenkins, its connection to SonarQube (`localhost:9000`), and the pipelines were unaffected. Jenkins reaches SonarQube locally, not through the security group.
- Likely the same for the other allowlisted UIs (Argo CD, Prometheus, Grafana, SigNoz) once they're used from the new IP.

### Resolution (step by step)
1. **Tell "dropped" from "down":** a `curl` timeout on one port while another port on the same host answers points at a firewall/allowlist, not at the service.
2. **Compare the IPs** (laptop):
   ```bash
   curl -s https://checkip.amazonaws.com                         # current IP
   aws ec2 describe-security-groups --region eu-north-1 --group-ids <SG_ID> \
     --query 'SecurityGroups[0].IpPermissions[?FromPort==`9000`].IpRanges[].CidrIp'
   ```
   Different → allowlist out of date.
3. **Update Terraform**, not the console, so the code stays the source of truth:
   ```bash
   cd terraform/jenkins-server
   MY=$(curl -s https://checkip.amazonaws.com)
   sed -i -E "s|^admin_cidrs *=.*|admin_cidrs = [\"${MY}/32\"]|" terraform.tfvars    # git-ignored file
   terraform plan      # expect: 0 to add, 1 to change, 0 to destroy
   terraform apply
   ```
4. **Update the Kubernetes load balancer allowlists the same way** (bastion), for each UI that hangs:
   ```bash
   kubectl -n <ns> patch svc <svc> -p '{"spec":{"loadBalancerSourceRanges":["<NEW_IP>/32"]}}'
   ```
5. **Verify:** `curl -s -o /dev/null -w '%{http_code}\n' http://<JENKINS_IP>:9000/` → `200` (or a redirect to the login page).

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
laptop$ curl -s https://checkip.amazonaws.com
<YOUR_NEW_IP>
laptop$ curl -s -o /dev/null -w '%{http_code} in %{time_total}s\n' --max-time 10 http://<JENKINS_IP>:9000/ || echo "no answer (rc=$?)"
000 in 10.002467s
no answer (rc=28)
laptop$ curl -s -o /dev/null -w '%{http_code} in %{time_total}s\n' --max-time 10 http://<JENKINS_IP>:8080/
403 in 0.541389s
laptop$ aws ec2 describe-security-groups --region eu-north-1 --group-ids <SG_ID> --query 'SecurityGroups[0].IpPermissions[].[FromPort,ToPort,IpRanges[].CidrIp]' --output text
9000	9000
<OLD_IP>/32
8080	8080
0.0.0.0/0
# -> 9000 allows current IP: NO
```
**Fixing it**
```console
laptop$ cd terraform/jenkins-server
laptop$ sed -i -E "s|^admin_cidrs *=.*|admin_cidrs = [\"$(curl -s https://checkip.amazonaws.com)/32\"]|" terraform.tfvars
laptop$ terraform plan -input=false -no-color
  # aws_security_group.jenkins_sg will be updated in-place
  ~ resource "aws_security_group" "jenkins_sg" {
      ~ ingress                = [
          - { cidr_blocks = [ "<OLD_IP>/32" ], description = "SonarQube UI", from_port = 9000, … },
          + { cidr_blocks = [ "<NEW_IP>/32" ], description = "SonarQube UI", from_port = 9000, … },
Plan: 0 to add, 1 to change, 0 to destroy.
laptop$ terraform apply
(output not captured; afterwards SonarQube loaded again)
```

### Code / config change
Only the git-ignored `terraform/jenkins-server/terraform.tfvars` (`admin_cidrs`). No code change.

### Lessons learned
- A timeout is a firewall problem until proven otherwise; a refusal or error page means the service itself answered.
- Allowlisting a home IP is fragile. It's fine for a short-lived demo, but expect to update it.
- Change the allowlist through Terraform rather than the AWS console, otherwise the next `terraform apply` silently puts the old IP back.

### How to approach it next time
1. Something IP-restricted "loads forever"? Run `curl https://checkip.amazonaws.com` first and compare it with the allowlist.
2. Check another port on the same host to separate "host down" from "port blocked".
3. Fix the source of truth (tfvars or the Service manifest), then apply.

### Prevention / follow-up
- Optional: reach admin UIs without public allowlists at all: **SSM port forwarding** through the bastion (`aws ssm start-session --document-name AWS-StartPortForwardingSessionToRemoteHost ...`), or a VPN. No IP to keep updated, and nothing exposed to the internet.
- Optional: a small script that updates all allowlists (Terraform + the four Kubernetes Services) from `checkip` in one go.

### References
- AWS: [Security group rules](https://docs.aws.amazon.com/vpc/latest/userguide/security-group-rules.html); [Session Manager port forwarding](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-sessions-start.html#sessions-remote-port-forwarding).
- Kubernetes: [`loadBalancerSourceRanges`](https://kubernetes.io/docs/concepts/services-networking/service/#aws-nlb-support).

---

## INC-014 – catalog-service would crash on an order with an unknown product ID

| | |
|---|---|
| **Date** | 2026-10-09 |
| **Where** | `catalog/repository.go` › `ListProductsWithIDs` (called by order-service for every checkout and order-history view) |
| **Severity** | High: one request with a bad product ID would crash a catalog-service pod; the shop's browse, search and checkout all depend on it. Caught before it happened |
| **Status** | Resolved (fix + regression test) |

### Summary
While adding gRPC status codes for SLOs, a review of catalog-service found that `ListProductsWithIDs` dereferenced each Elasticsearch document's `_source` without checking whether the document exists. For an unknown product ID, Elasticsearch returns `found: false` and **no `_source`**, so the code would hit a **nil pointer dereference**. In a gRPC handler that is an unrecovered panic, and the whole catalog-service process exits. A regression test reproduces the panic on the old code, and the fix (skip documents that weren't found) makes it pass.

### Background
- order-service calls catalog's `GetProducts(ids...)` when placing an order and when showing order history. Catalog answers with an Elasticsearch **multi-get** (`_mget`): one entry per requested ID.
- For an ID that doesn't exist, `_mget` still returns an entry, with `"found": false` and **no `_source` field**. In elastic.v5 that is `doc.Found == false` and `doc.Source == nil` (a `*json.RawMessage`).
- `json.Unmarshal(*doc.Source, ...)` dereferences that pointer. With `nil` it panics.
- **grpc-go does not recover panics in handlers.** A panic in a request handler crashes the process. Kubernetes restarts the pod, but in-flight requests on it fail, and repeated bad requests can crash-loop both replicas.

### Timeline
1. Adding SLI support: services must return proper gRPC codes (`NotFound` vs `Internal`), so every error path in account, catalog and order was reviewed.
2. In `catalog/repository.go`, `ListProductsWithIDs` looped over `res.Docs` and called `json.Unmarshal(*doc.Source, &p)` for every document, found or not.
3. Checked elastic.v5's `GetResult`: `Source *json.RawMessage` and `Found bool`. Missing documents leave `Source` nil.
4. Why it hadn't happened yet: the load test's bad orders used an unknown *account*, so order-service stopped before calling catalog. The frontend only sends IDs of products it has just listed.
5. Wrote a test with a fake Elasticsearch (`httptest`) returning one found and one missing document. On the **old** code it panicked; with the fix it passes.

### What happened (symptom)
No production occurrence. Reproduced by the regression test against the old code:

    --- FAIL: TestListProductsWithIDsSkipsMissing (0.00s)
    panic: runtime error: invalid memory address or nil pointer dereference [recovered, repanicked]

In the cluster this would have shown as: a catalog-service pod restarting (`RESTARTS` going up), the triggering checkout failing, and in SigNoz a broken trace ending in catalog with no response.

### How we got there
The original code assumed every requested ID exists. Nothing validated product IDs before the multi-get, so any client (or a stale cart holding a deleted product) could send an unknown ID.

### Why (root cause)
1. Elasticsearch represents "not found" in a multi-get as an entry with `found: false` and no `_source`, not as an error.
2. The loop didn't check `doc.Found` or `doc.Source != nil` before dereferencing.
3. gRPC servers in Go don't recover handler panics, so one bad document crashes the whole service.

### Impact
- None in practice (found in review).
- Potential: a single crafted `createOrder` (or an order history containing a since-deleted product) would crash a catalog pod. Repeated, it could take down browse, search and checkout for everyone, which is a denial-of-service via one bad input.

### Resolution (step by step)
1. **Fix:** skip documents that weren't found:
   ```go
   if !doc.Found || doc.Source == nil {
       continue
   }
   ```
2. **Regression test** (`catalog/server_test.go`, `TestListProductsWithIDsSkipsMissing`): a fake Elasticsearch returns one found and one missing document; the test expects only the found product, and no panic.
3. **Prove the test catches the bug:** restore the old `repository.go`, run the test → panic; restore the fix → pass.
4. **Handle the effect upstream:** order-service now rejects an order whose products all don't exist with `InvalidArgument` ("none of the ordered products exist"), instead of storing an empty $0 order.
5. **Verify after deploy** (laptop):
   ```bash
   curl -s -X POST -H 'Content-Type: application/json' http://<ALB_HOST>/graphql -d \
     '{"query":"mutation{createOrder(order:{accountId:\"<real account id>\",products:[{id:\"does-not-exist\",quantity:1}]}){id}}"}'
   # expect an error "none of the ordered products exist"; catalog pods' RESTARTS unchanged
   ```

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Proving the bug with the new regression test**
```console
laptop$ git show HEAD:catalog/repository.go > catalog/repository.go      # old code
laptop$ go test ./catalog/ -run TestListProductsWithIDsSkipsMissing -v
--- FAIL: TestListProductsWithIDsSkipsMissing (0.00s)
panic: runtime error: invalid memory address or nil pointer dereference [recovered, repanicked]
panic({0xb2ae60?, 0x1245d10?})
	/usr/local/go/src/runtime/panic.go:860 +0x13a
```
**Fixing it**
```console
laptop$ cp repository.go.fixed catalog/repository.go                      # fix restored
laptop$ go test ./catalog/ -run TestListProductsWithIDsSkipsMissing
ok  	github.com/ritesh-karankal/go-grpc-graphql-micro/catalog	0.011s
laptop$ go test ./account/... ./catalog/... ./order/... ./graphql/... ./telemetry/...
ok  	github.com/ritesh-karankal/go-grpc-graphql-micro/account	0.038s
ok  	github.com/ritesh-karankal/go-grpc-graphql-micro/catalog	0.023s
ok  	github.com/ritesh-karankal/go-grpc-graphql-micro/order	0.026s
ok  	github.com/ritesh-karankal/go-grpc-graphql-micro/graphql	0.010s
?   	github.com/ritesh-karankal/go-grpc-graphql-micro/telemetry	[no test files]
```

### Code / config change
```diff
 	for _, doc := range res.Docs {
+		// Unknown IDs come back with Found=false and no _source; skip them
+		if !doc.Found || doc.Source == nil {
+			continue
+		}
 		p := productDocument{}
 		if err = json.Unmarshal(*doc.Source, &p); err == nil {
```
Plus `order/server.go`: an order with no existing products → `codes.InvalidArgument`.

### Lessons learned
- Check every pointer that comes from external data before dereferencing it, especially "optional" fields in API responses.
- In Go gRPC services, a panic is an outage, not just an error. Consider a recovery interceptor (`grpc-ecosystem/go-grpc-middleware/recovery`) as defence in depth.
- Reviewing error paths for one purpose (SLOs) surfaces bugs in others. Error handling is where latent bugs hide.
- A regression test should be shown to **fail on the old code**; otherwise it may not test anything.

### How to approach it next time
1. Pod `RESTARTS` increasing → `kubectl logs <pod> --previous` shows the panic and stack trace of the crashed container.
2. Find the request that triggered it: SigNoz traces with errors around the crash time, ending in the crashed service.
3. Reproduce with a unit test using a fake dependency (`httptest` for HTTP APIs such as Elasticsearch).

### Prevention / follow-up
- Done: fix + regression test.
- Follow-up: add a gRPC **recovery interceptor** to all three services, so an unexpected panic becomes an `Internal` error instead of a crash.
- Follow-up: validate input at the gateway (product IDs format, quantities), as already done for quantity.

### References
- Elasticsearch: [Multi get API — response for missing documents](https://www.elastic.co/guide/en/elasticsearch/reference/5.6/docs-multi-get.html).
- grpc-go: panics in handlers are not recovered; see [go-grpc-middleware recovery](https://github.com/grpc-ecosystem/go-grpc-middleware/tree/main/interceptors/recovery).

---

## INC-013 – SigNoz shows "You're not sending any data yet"

| | |
|---|---|
| **Date** | 2026-10-09 |
| **Where** | SigNoz UI (first login); the services' OTLP exporters (`telemetry/telemetry.go`) → `signoz-otel-collector.signoz.svc.cluster.local:4317`; Recreate guide Steps 14 and 16 |
| **Severity** | Medium: no traces, metrics or logs reached SigNoz for a while; the app itself was unaffected |
| **Status** | Resolved: exporters recovered, pods restarted, SigNoz receiving data |

### Summary
After SigNoz was installed, its UI said "You're not sending any data yet". Every part of the setup checked out (new images, correct environment variables, collector reachable, SigNoz healthy), except the **timing**. The app pods had started at 18:30, **a few minutes before SigNoz was installed**. At that moment the collector's DNS name did not exist yet, so the OpenTelemetry exporters failed with `name resolver error: produced zero addresses`, every 10 seconds. gRPC kept re-resolving the name, and by the time we checked the exporters had already stopped failing. The pods were restarted anyway, so every pod started with the collector already in DNS, and SigNoz then showed data.

### Background
- Each Go service sends traces, metrics and logs over **OTLP/gRPC** to the SigNoz collector at `signoz-otel-collector.signoz.svc.cluster.local:4317` (from the `otel-config` ConfigMap).
- `<service>.<namespace>.svc.cluster.local` names are answered by the cluster's DNS (CoreDNS) **only while that Service exists**. Before `helm install signoz`, the name doesn't resolve.
- The OTel exporters use a gRPC client that connects **lazily** and keeps retrying. If DNS returns nothing, gRPC reports `name resolver error: produced zero addresses`, then re-resolves with exponential backoff. The app keeps running normally while this happens; only telemetry export fails (by design, see `telemetry/telemetry.go`).
- Even with no user traffic, every service pushes Go runtime metrics every 60 seconds. So "no data at all" means nothing reaches the collector, not just "nobody clicked yet".
- SigNoz shows its onboarding message ("not sending any data yet") until the first data arrives.

### Timeline
1. **18:30:49 UTC:** app pods started (Argo CD sync, Step 14). SigNoz not installed yet.
2. **18:31:00 onwards:** pod logs show `exporter export timeout: rpc error: code = Unavailable desc = name resolver error: produced zero addresses` every 10 s.
3. **~18:33:** SigNoz installed with Helm (Step 16). Collector Service and pods came up.
4. Later (after INC-012): first login to the SigNoz UI showed "You're not sending any data yet".
5. Ran a 4-part check on the bastion: images, environment, network path, SigNoz health (results below). Only the startup log showed a problem.
6. Counted export errors in the last 2 minutes: **0**. The exporters had already recovered on their own.
7. Restarted all app deployments anyway (`kubectl rollout restart`), then counted again: **0** errors.
8. Generated traffic in the shop and refreshed SigNoz: data showed up, services listed.
9. **Second phase, seen in SigNoz's own Logs view:** a burst of a *different* error between **18:48:52 and 18:50:46 UTC** (00:18–00:20 IST in the UI):
   `dial tcp <collector ClusterIP>:4317: connect: connection refused`. The collector's own log showed `Starting health_check extension` at **18:48:41 UTC**, 11 s before the burst.
10. SigNoz Logs, *Last 15 minutes*: last error at 18:50:46 UTC; afterwards only normal lines (e.g. `Listening on port 8080...` from the restarted pods at 18:53:07). The burst stopped on its own about 2 minutes before the restart.

### What happened (symptom)
SigNoz UI:

    You're not sending any data yet.
    SigNoz is so much better with your data ⎯ start by sending your telemetry data to SigNoz.

Pod log (`kubectl -n go-micro-shop logs deploy/order-service`):

    2026/10/08 18:30:49 INFO Serving Prometheus metrics addr=:9464 path=/metrics
    time=2026-10-08T18:31:00.823Z level=INFO msg="exporter export timeout: rpc error: code = Unavailable desc = name resolver error: produced zero addresses"
    time=2026-10-08T18:31:10.823Z level=INFO msg="exporter export timeout: rpc error: code = Unavailable desc = name resolver error: produced zero addresses"
    ...

### How we got there
The guide installs SigNoz (Step 16) **after** deploying the app (Step 14), so the app always starts before the collector exists.

### Why (root cause)
1. Argo CD deployed the app before SigNoz existed (install order).
2. The pods' exporters looked up `signoz-otel-collector.signoz.svc.cluster.local` at startup. The Service didn't exist, so DNS returned **no addresses**.
3. Every export attempt failed with `name resolver error: produced zero addresses`, so nothing reached SigNoz, and the UI showed its "no data" message.
4. Once SigNoz was installed, gRPC's re-resolution eventually found the new name and exports started succeeding (0 errors in the 2 minutes before the restart).

*Not fully confirmed:* whether this self-recovery alone would have filled the UI, or whether the restart was also needed. The restart was done before SigNoz was re-checked. Either way, the cause was the pods starting before the collector's DNS name existed.

**Second phase, `connection refused` (18:48:52–18:50:46 UTC).** A different failure from the DNS one:

| Error | What it means |
|---|---|
| `name resolver error: produced zero addresses` | the collector's **name** does not exist (Service not created yet) |
| `connect: connection refused` to the collector's ClusterIP | the name resolves to the Service, but **no ready collector pod** is behind it, so kube-proxy rejects the connection |

5. At 18:48:41 UTC the collector logged `Starting health_check extension`, i.e. it was (re)starting its pipelines. Its pod showed `RESTARTS 0`, so this was an **in-process restart**, not a container restart.
6. While restarting, the collector was not ready and port 4317 was not served, hence about 2 minutes of `connection refused`.
7. *Inferred, not confirmed:* SigNoz manages its collector remotely over **OpAMP** and pushes config updates to it (for example around the first login / initial setup), and the collector restarts its pipelines to apply them. The timing fits; the collector log around 18:48 UTC would confirm it.
8. It recovered on its own at 18:50:46 UTC. The log records written during the outage were buffered by the SDK and delivered afterwards, which is why they appeared in SigNoz itself.

Everything else was ruled out by the checks:

| # | Possible cause | Check (bastion) | Result |
|---|---|---|---|
| 1 | Old images without OpenTelemetry | `kubectl -n go-micro-shop get deploy -o custom-columns=...IMAGE...` | `account/catalog/order/graphql:3` (current) |
| 1b | Instrumentation not running | `kubectl exec deploy/order-service -- wget -qO- localhost:9464/metrics` | `go_goroutine_count`, `rpc_client_call_duration_seconds…` present |
| 2 | Endpoint variable missing | `kubectl exec deploy/order-service -- env \| grep OTEL` | endpoint set; `deployment.environment=dev`, pod and node names |
| 3 | Network path blocked | `kubectl exec deploy/order-service -- nc -zv -w 5 signoz-otel-collector.signoz.svc.cluster.local 4317` | `open` (172.20.x.x:4317) |
| 4 | SigNoz not ready | `kubectl -n signoz get pods` | collector, ClickHouse, ZooKeeper, query service Running; migrator Completed |
| 2b | **Exporter errors** | `kubectl logs deploy/order-service \| grep -iE "otlp\|export"` | **`produced zero addresses` from 18:31** |

### Impact
- No telemetry in SigNoz from pod start until the exporters recovered. Nothing is buffered beyond the SDK's small in-memory queue, so data from that window is lost. It was setup time with no real users.
- The app, Prometheus scraping (`:9464/metrics`) and Grafana were unaffected. Prometheus metrics are pulled, not pushed, so they don't depend on SigNoz.

### Resolution (step by step)
1. **Run the 4-part check** (table above) to find which link is broken, rather than guessing.
2. **Read the exporter errors in the pod log:**
   `produced zero addresses` = the name doesn't resolve (yet). Compare the pod start time (first log line) with the SigNoz pods' age (`kubectl -n signoz get pods`).
3. **Check whether it is still failing**, not just whether it failed at startup:
   ```bash
   kubectl -n go-micro-shop logs deploy/order-service --since=2m | grep -c "exporter export timeout"
   ```
   Here: `0` → already recovered.
4. **Restart the app** so all pods connect with the collector present:
   ```bash
   kubectl -n go-micro-shop rollout restart deploy
   kubectl -n go-micro-shop rollout status deploy --timeout=5m
   ```
   Safe: 2 replicas per deployment, rolled one at a time; databases untouched.
5. **Verify:** `--since=1m` error count `0`. Generate traffic, hard-refresh SigNoz → **Services** lists the four services in `dev`.

**Phase 2 checks:**
6. **Read the errors in SigNoz → Logs** with the time range set to *Last 15 minutes*. SigNoz shows local time (IST = UTC+5:30); convert to UTC before comparing with `kubectl` timestamps.
7. **Find the last occurrence** of the error, and what comes after it. Here: last `connection refused` at 18:50:46 UTC, then only normal startup lines. The burst was over.
8. **Correlate with the collector:**
   ```bash
   kubectl -n signoz get pods -l app.kubernetes.io/component=otel-collector          # RESTARTS
   kubectl -n signoz logs deploy/signoz-otel-collector -c collector --since=2h \
     | grep -iE '"msg":"(Starting|Shutdown|Everything is ready)' | tail
   ```

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
bastion$ kubectl -n go-micro-shop get deploy -o custom-columns=NAME:.metadata.name,IMAGE:.spec.template.spec.containers[0].image
NAME              IMAGE
account-service   <ACCT>.dkr.ecr.eu-north-1.amazonaws.com/account:3
catalog-service   <ACCT>.dkr.ecr.eu-north-1.amazonaws.com/catalog:3
frontend          <ACCT>.dkr.ecr.eu-north-1.amazonaws.com/frontend:2
graphql-gateway   <ACCT>.dkr.ecr.eu-north-1.amazonaws.com/graphql:3
order-service     <ACCT>.dkr.ecr.eu-north-1.amazonaws.com/order:3
bastion$ kubectl -n go-micro-shop exec deploy/order-service -- wget -qO- localhost:9464/metrics | grep -m2 -E "^(rpc_|go_goroutine|target_info)"
go_goroutine_count{otel_scope_name="go.opentelemetry.io/contrib/instrumentation/runtime",…} 64
rpc_client_call_duration_seconds_bucket{…,rpc_method="pb.CatalogService/GetProducts",rpc_response_status_code="OK",…,le="0.005"} 0
bastion$ kubectl -n go-micro-shop exec deploy/order-service -- env | grep -E "OTEL|DEPLOYMENT_ENV"
DEPLOYMENT_ENVIRONMENT=dev
OTEL_EXPORTER_OTLP_ENDPOINT=http://signoz-otel-collector.signoz.svc.cluster.local:4317
OTEL_RESOURCE_ATTRIBUTES=deployment.environment=dev,k8s.namespace.name=go-micro-shop,k8s.pod.name=order-service-<pod>,k8s.node.name=ip-10-10-x-x.eu-north-1.compute.internal
bastion$ kubectl -n go-micro-shop logs deploy/order-service --tail=200 | grep -iE "otlp|metrics|telemetry|export" | head -5
2026/10/08 18:30:49 INFO Serving Prometheus metrics addr=:9464 path=/metrics
time=2026-10-08T18:31:00.823Z level=INFO msg="exporter export timeout: rpc error: code = Unavailable desc = name resolver error: produced zero addresses"
time=2026-10-08T18:31:10.823Z level=INFO msg="exporter export timeout: rpc error: code = Unavailable desc = name resolver error: produced zero addresses"
…
bastion$ kubectl -n go-micro-shop exec deploy/order-service -- sh -c 'nc -zv -w 5 signoz-otel-collector.signoz.svc.cluster.local 4317'
signoz-otel-collector.signoz.svc.cluster.local (172.20.x.x:4317) open
bastion$ kubectl -n signoz get pods
NAME                                          READY   STATUS      RESTARTS      AGE
chi-signoz-clickhouse-cluster-0-0-0           1/1     Running     0             14m
signoz-0                                      1/1     Running     0             15m
signoz-clickhouse-operator-55fb8db98d-jhftc   2/2     Running     2 (14m ago)   15m
signoz-otel-collector-5cfc9ccdd8-5blp7        1/1     Running     0             15m
signoz-telemetrystore-migrator-7wm67          0/1     Completed   0             15m
signoz-zookeeper-0                            1/1     Running     0             15m
bastion$ kubectl -n signoz logs deploy/signoz-otel-collector --tail=30 | grep -iE "error|warn|refused|fail" | tail -8
{"level":"info","ts":"2026-10-08T18:48:41.283Z",…,"msg":"Starting health_check extension",…}
```
**Fixing it**
```console
bastion$ kubectl -n go-micro-shop logs deploy/order-service --since=2m | grep -c "exporter export timeout"
Found 2 pods, using pod/order-service-<pod>
0
bastion$ kubectl -n go-micro-shop rollout restart deploy
deployment.apps/account-service restarted
deployment.apps/catalog-service restarted
deployment.apps/frontend restarted
deployment.apps/graphql-gateway restarted
deployment.apps/order-service restarted
bastion$ kubectl -n go-micro-shop rollout status deploy --timeout=5m
Waiting for deployment "account-service" rollout to finish: 1 out of 2 new replicas have been updated...
…
deployment "order-service" successfully rolled out
bastion$ sleep 60; kubectl -n go-micro-shop logs deploy/order-service --since=1m | grep -c "exporter export timeout"
Found 2 pods, using pod/order-service-<pod>
0
# SigNoz UI -> Logs, Last 15 minutes: last "connect: connection refused" at 00:20:46 IST (18:50:46 UTC); then only
# normal lines such as "Listening on port 8080..." at 00:23:07 IST from the restarted pods.
```

### Code / config change
None.

### Lessons learned
- Logs that were **written** during an outage can still **arrive** afterwards (SDK buffering), so an error showing up in SigNoz now doesn't mean it is happening now. Always check the timestamp, and what comes after the last occurrence.
- `zero addresses` and `connection refused` look alike but point at different layers: DNS/Service existence vs. pod readiness behind the Service.
- A "no data" screen is a symptom of the **whole path**: app → env var → DNS → network → collector → storage. Checking each link in order finds the broken one in minutes.
- `head` on logs only shows the past. To know whether something is failing **now**, use `--since=2m` (or `--tail`) and count.
- Push-based telemetry depends on install order: if the receiver doesn't exist when the sender starts, the sender has to recover on its own. Pull-based (Prometheus) doesn't have this problem.

### How to approach it next time
1. Check that data is produced: `/metrics` on the pod.
2. Check that the pod knows where to send: `env | grep OTEL`.
3. Check that the pod can reach it: `nc -zv <collector> 4317`.
4. Check that the receiver is healthy: `kubectl -n signoz get pods`.
5. Read the exporter errors with `--since`:
   - `produced zero addresses` → DNS: the collector Service is missing, or was created after the pod started;
   - `connection refused` → Service exists but no ready collector pod (starting, restarting, or reloading config);
   - `deadline exceeded` → network or collector overloaded;
   - `Unimplemented` / `404` → wrong port or protocol (gRPC 4317 vs HTTP 4318).
6. Confirm the collector is receiving: its own metrics on `:8888` (`otelcol_receiver_accepted_spans`, `..._metric_points`, `..._log_records`).

### Prevention / follow-up
- Install order: install SigNoz **before** creating the Argo CD Application, or restart the app right after installing SigNoz. The guide's Step 16 already says to restart if a service shows no data; moving SigNoz before Step 14 would avoid the issue.
- Optional: lower the noise. The exporter logs a line every 10 s while it can't export; that is useful when debugging but noisy in normal logs.

### References
- OpenTelemetry Go: [OTLP gRPC exporter](https://pkg.go.dev/go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc) (lazy connection, retries).
- gRPC: [Name resolution](https://grpc.io/docs/guides/custom-name-resolution/).
- Kubernetes: [DNS for Services and Pods](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/).
- SigNoz: [Kubernetes install — send data to the collector](https://signoz.io/docs/install/kubernetes/others/).

---

## INC-012 – SigNoz UI not reachable: load balancer created as `internal`

| | |
|---|---|
| **Date** | 2026-10-09 |
| **Where** | `k8s/monitoring/signoz-ui-service.yaml`; Recreate guide Step 16 (SigNoz) |
| **Severity** | Medium: the SigNoz UI was unreachable from outside the VPC; SigNoz itself was running |
| **Status** | Resolved: Service recreated with the `internet-facing` annotation; new NLB active, targets healthy, UI returns HTTP 200 |

### Summary
The Service that exposes the SigNoz UI got a load balancer, but its hostname did not resolve from the internet. AWS showed it as a **Network Load Balancer with scheme `internal`**: private IPs only, reachable from inside the VPC only. The Service manifest (written as part of the SigNoz setup) had no `aws-load-balancer-scheme: internet-facing` annotation. With the **AWS Load Balancer Controller** installed, a newly created `type: LoadBalancer` Service becomes an NLB, and NLBs default to `internal`. Fixed by adding the annotation and recreating the Service, since a load balancer's scheme cannot be changed in place.

### Background
- A Kubernetes Service of `type: LoadBalancer` asks the cloud for a load balancer. On EKS, two different components can answer:

  | Who handles it | When | What it creates | Default scheme |
  |---|---|---|---|
  | **Built-in AWS cloud provider** (legacy, in-tree) | Services the Load Balancer Controller didn't claim | **Classic** Load Balancer (`<hash>-<id>.<region>.elb.amazonaws.com`) | **internet-facing** |
  | **AWS Load Balancer Controller** (installed in Step 7) | New `LoadBalancer` Services (its webhook sets `spec.loadBalancerClass: service.k8s.aws/nlb` at creation) | **Network** Load Balancer (`k8s-<ns>-<name>-<hash>.elb.<region>.amazonaws.com`) | **internal** |

- The controller's webhook only acts when a Service is **created**. Existing Services that are later *patched* to `type: LoadBalancer` keep going to the built-in provider.
- `service.beta.kubernetes.io/aws-load-balancer-scheme: internet-facing` tells the controller to give the NLB public IPs. `loadBalancerSourceRanges` still limits who may connect.
- A load balancer's **scheme is fixed at creation**: going from internal to internet-facing means a new load balancer.

### Timeline
1. Steps 8 and 15: Argo CD, Prometheus and Grafana were exposed by **patching** their existing Services to `type: LoadBalancer`. All three got **Classic, internet-facing** load balancers and worked. (Argo CD also had the scheme annotation, which Classic LBs ignore.)
2. Step 16: applied `signoz-ui-service.yaml`, a **new** Service with `type: LoadBalancer` and no scheme annotation.
3. The hostname `k8s-signoz-signozui-<hash>.elb.eu-north-1.amazonaws.com` did not open in the browser.
4. From the laptop: DNS did not resolve (`curl` exit code 6). AWS showed `type: network`, `scheme: internal`, targets `initial`.
5. Listed all load balancers: only the SigNoz one was an internal NLB; the other three were internet-facing Classic LBs, and the app's ALB was internet-facing.
6. Added the annotation to the manifest, pushed it, and recreated the Service on the bastion.
7. Right after recreating it, `kubectl -n signoz get svc signoz-ui` returned `NotFound`: the delete had run, but the `sed ... | kubectl apply` line had not (lines merged when pasting). Re-ran the apply on its own.
8. A **new** NLB appeared (different name/hash from the old one), scheme `internet-facing`. For the first ~1–2 minutes it was `provisioning`, targets `initial` (`Elb.RegistrationInProgress`), and DNS did not resolve yet.
9. About 3 minutes after creation: state `active`, both targets `healthy`, `curl` → `200`. UI reachable.

### What happened (symptom)
Browser: site not reachable. From the laptop:

    $ curl -s -o /dev/null -w 'http %{http_code}\n' --max-time 8 http://<LB_HOST>/
    http 000           # curl exit code 6: could not resolve host

    $ aws elbv2 describe-load-balancers ...
    name: k8s-signoz-signozui-<hash>   type: network   scheme: internal   state: provisioning

All load balancers at that moment:

| Load balancer | Created by | Type | Scheme |
|---|---|---|---|
| app Ingress (`k8s-gomicros-...`) | LB Controller (Ingress) | application | internet-facing |
| SigNoz UI (`k8s-signoz-signozui-...`) | LB Controller (new Service) | **network** | **internal** |
| Argo CD, Prometheus, Grafana | built-in provider (patched Services) | classic | internet-facing |

### How we got there
`signoz-ui-service.yaml` was written to match how Prometheus and Grafana were exposed (type LoadBalancer + IP allowlist), without accounting for the fact that a **newly created** Service goes to the Load Balancer Controller, which has a different default.

### Why (root cause)
1. The AWS Load Balancer Controller was installed (Step 7) before the SigNoz Service was created (Step 16).
2. When the Service was created, the controller's webhook claimed it, so an NLB was created.
3. The manifest had no scheme annotation, so the NLB got the controller's default: `internal`.
4. An internal NLB has only private IPs and its public DNS name does not resolve from the internet, so the browser couldn't reach it.

The underlying gap: the manifest relied on a default, and the default depends on *which controller* handles the Service.

### Impact
- SigNoz UI unreachable from the laptop until the fix. SigNoz itself (collector, ClickHouse, query service) kept running and receiving telemetry.
- No effect on the app, Argo CD, Prometheus or Grafana.

### Resolution (step by step)
1. **Check the scheme** (laptop):
   ```bash
   aws elbv2 describe-load-balancers --region eu-north-1 \
     --query 'LoadBalancers[].[LoadBalancerName,Type,Scheme,State.Code]' --output table
   ```
   `internal` confirms the cause.
2. **Add the annotation** to `k8s/monitoring/signoz-ui-service.yaml`:
   ```yaml
   metadata:
     annotations:
       service.beta.kubernetes.io/aws-load-balancer-scheme: internet-facing
   ```
3. **Recreate the Service** (bastion). The scheme can't change in place, so delete and re-apply:
   ```bash
   cd ~/go-grpc-graphql-micro && git pull
   kubectl -n signoz delete svc signoz-ui
   MY_IP=<laptop IP>
   sed "s|<YOUR_IP>|${MY_IP}|" k8s/monitoring/signoz-ui-service.yaml | kubectl apply -f -
   kubectl -n signoz get svc signoz-ui          # new hostname
   ```
4. **Verify** (laptop) after 2–4 minutes:
   ```bash
   aws elbv2 describe-load-balancers --region eu-north-1 \
     --query "LoadBalancers[?starts_with(LoadBalancerName,'k8s-signoz')].[Scheme,State.Code]" --output text
   # internet-facing  active
   curl -s -o /dev/null -w '%{http_code}\n' http://<LB_HOST>/     # 200
   ```
   Result here: `internet-facing`, `active`, targets `healthy`, HTTP `200`, about 3 minutes after the Service was created.
5. **If `get svc` says `NotFound` after the delete**, the apply didn't run. Run it as its own command, check that `MY_IP` is set (an empty value gives the invalid CIDR `/32`), and look for `service/signoz-ui created`.

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
bastion$ sed "s|<YOUR_IP>|${MY_IP}|" k8s/monitoring/signoz-ui-service.yaml | kubectl apply -f -
(output not captured)
# browser: http://<LB_HOST>/ -> not reachable

laptop$ aws elbv2 describe-load-balancers --region eu-north-1 --load-balancer-arns <arn> --query 'LoadBalancers[0].{name:LoadBalancerName,type:Type,scheme:Scheme,state:State.Code}' --output table
---------------------------------------------------------------------------
|                          DescribeLoadBalancers                          |
+---------------------------------+-----------+---------------+-----------+
|              name               |  scheme   |     state     |   type    |
+---------------------------------+-----------+---------------+-----------+
|  k8s-signoz-signozui-<id>       |  internal |  provisioning |  network  |
+---------------------------------+-----------+---------------+-----------+
laptop$ aws elbv2 describe-target-health … --query 'TargetHealthDescriptions[].TargetHealth.State' --output text
initial	initial
laptop$ curl -s -o /dev/null -w 'http %{http_code}\n' --max-time 8 http://<LB_HOST>/ || echo "curl rc=$?"
http 000
curl rc=6
laptop$ grep -n "annotations\|scheme" k8s/monitoring/signoz-ui-service.yaml || echo "no scheme annotation in signoz-ui-service.yaml"
no scheme annotation in signoz-ui-service.yaml
laptop$ aws elbv2 describe-load-balancers --region eu-north-1 --query 'LoadBalancers[].[LoadBalancerName,Type,Scheme,State.Code]' --output table
|  k8s-signoz-signozui-<id>      |  network     |  internal         |  provisioning  |
|  k8s-gomicros-gomicros-<id>    |  application |  internet-facing  |  active        |
laptop$ aws elb describe-load-balancers --region eu-north-1 --query 'LoadBalancerDescriptions[].[LoadBalancerName,Scheme]' --output text
<argocd-clb>	internet-facing
<prometheus-clb>	internet-facing
<grafana-clb>	internet-facing
```
**Fixing it**
```console
laptop$ sed "s|<YOUR_IP>|203.0.113.10|" k8s/monitoring/signoz-ui-service.yaml | python3 -c "import sys,yaml; d=yaml.safe_load(sys.stdin); print(d['metadata']['annotations'], d['spec']['loadBalancerSourceRanges'])"
{'service.beta.kubernetes.io/aws-load-balancer-scheme': 'internet-facing'} ['203.0.113.10/32']

bastion$ kubectl -n signoz delete svc signoz-ui
(output not captured)
bastion$ kubectl -n signoz get svc signoz-ui
Error from server (NotFound): services "signoz-ui" not found
bastion$ sed "s|<YOUR_IP>|${MY_IP}|" k8s/monitoring/signoz-ui-service.yaml | kubectl apply -f -   # re-run on its own
(output not captured)

laptop$ aws elbv2 describe-load-balancers --region eu-north-1 --query "LoadBalancers[?starts_with(LoadBalancerName,'k8s-signoz')].[LoadBalancerName,Scheme,State.Code]" --output text
k8s-signoz-signozui-<new-id>	internet-facing	provisioning
laptop$ # poll every 20 s until it answers
READY after 98s: state=active http=200
healthy	None
healthy	None
```

### Code / config change
`k8s/monitoring/signoz-ui-service.yaml`:

```diff
 metadata:
   name: signoz-ui
   namespace: signoz
+  annotations:
+    # The AWS Load Balancer Controller creates an NLB for new LoadBalancer Services,
+    # and NLBs are internal (VPC-only) unless asked otherwise
+    service.beta.kubernetes.io/aws-load-balancer-scheme: internet-facing
 spec:
   type: LoadBalancer
```

### Lessons learned
- On EKS, *which* controller handles a `LoadBalancer` Service decides the load balancer type **and** its defaults. Installing the AWS Load Balancer Controller changes the behaviour of every new LoadBalancer Service.
- Never rely on a default for something security- or reachability-relevant. Say `internet-facing` or `internal` explicitly, plus a source allowlist.
- The hostname format tells you who made the load balancer: `k8s-<ns>-<svc>-...` is the controller's NLB/ALB; `<hash>-<id>.<region>.elb...` is a Classic LB from the built-in provider.

### How to approach it next time
1. Load balancer hostname doesn't resolve or connect? First check its **scheme**: `aws elbv2 describe-load-balancers` (NLB/ALB) or `aws elb describe-load-balancers` (Classic).
2. `internal` → add the scheme annotation and recreate the Service.
3. `internet-facing` but timing out → check `loadBalancerSourceRanges` and your current IP.
4. Targets `unhealthy` → check the pods and the Service's `targetPort` (bastion: `kubectl -n <ns> get pods,endpoints`).

### Prevention / follow-up
- Done: scheme annotation in `signoz-ui-service.yaml`.
- Follow-up: the guide's Prometheus/Grafana steps work only because they *patch* existing Services. Make all exposed Services explicit (annotation, or `loadBalancerClass` + scheme), so behaviour doesn't depend on install order.
- Optional: expose the admin UIs through one ALB Ingress with host or path rules, instead of one load balancer each (cheaper, and one place for TLS and allowlists).

### References
- AWS Load Balancer Controller: [Network Load Balancer (Service) — annotations, `aws-load-balancer-scheme`](https://kubernetes-sigs.github.io/aws-load-balancer-controller/latest/guide/service/annotations/#lb-scheme); [Service mutator webhook / `loadBalancerClass`](https://kubernetes-sigs.github.io/aws-load-balancer-controller/latest/guide/service/nlb/).
- Kubernetes: [Service type LoadBalancer](https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer).

---

## INC-011 – `Update Deployment file` fails: "nothing added to commit"

| | |
|---|---|
| **Date** | 2026-10-08 |
| **Where** | Jenkins job `backend` (build #3) › stage **Update Deployment file**; `jenkins/Jenkinsfile-Backend` (same logic in `Jenkinsfile-Frontend`) |
| **Severity** | Low: images were built and pushed; `main` already pointed at the right tag, so deployment was not blocked |
| **Status** | Resolved (both Jenkinsfiles skip the commit when the tags are already current) |

### Summary
The re-run of `backend` (build #3, after INC-010) built, scanned and pushed all four images as tag `3`, then failed in the last stage. That stage bumps the image tags in `k8s/base/kustomization.yaml` and commits them. The file **already said `"3"`**, left over from the previous, torn-down setup, so the bump changed nothing. `git commit` then refused with "nothing added to commit", and that non-zero exit failed the stage. Harmless in effect: `main` already pointed at tag `3`, which now existed in ECR. Fixed by making the stage skip commit and push when there is nothing to change.

### Background
- **GitOps flow:** Jenkins never deploys. Its last stage writes the new image tag into `k8s/base/kustomization.yaml` and pushes it to `main`. Argo CD sees the commit and deploys that tag.
- The image tag is the **Jenkins build number** (`IMAGE_TAG = "${BUILD_NUMBER}"`). Build #3 produces `account:3`, `catalog:3`, and so on.
- The tag bump is a `sed` that rewrites the `newTag:` line under each service. If the value is already the same, `sed` leaves the file unchanged.
- `git commit` with nothing staged **exits with an error** ("nothing to commit" / "nothing added to commit"). Jenkins runs `sh` steps with `set -e`-like behaviour: any command that exits non-zero fails the step, and with it the stage.
- After the teardown, Jenkins started counting builds from 1 again, but the tags in git still came from the old builds (`"3"` for the backend). The new build numbers and the old tags **can collide**.

### Timeline
1. INC-010: the first backend run was aborted by the timeout. The re-run that got through OWASP was build #3.
2. Build #3: tests, Sonar, Quality Gate, OWASP (fast now, with the database downloaded), Trivy, Docker build, Trivy image gate and ECR push all passed. ECR now had `account:3`, `catalog:3`, `order:3`, `graphql:3`.
3. **Update Deployment file** failed.
4. Checked from the laptop: images present in ECR; `kustomization.yaml` on `main` still had `"3"` for the backend services, and no new backend commit.
5. The console log showed `sed ... newTag: "3"`, an empty `git diff --cached --stat`, then `nothing added to commit`. Confirmed.
6. Changed both Jenkinsfiles to skip commit and push when nothing is staged. Tested the logic in a scratch git repo.

### What happened (symptom)

    + sed -i -E /^  - name: account$/{n;s/newTag: .*/newTag: "3"/} k8s/base/kustomization.yaml
    + sed -i -E /^  - name: catalog$/{n;s/newTag: .*/newTag: "3"/} k8s/base/kustomization.yaml
    + sed -i -E /^  - name: order$/{n;s/newTag: .*/newTag: "3"/} k8s/base/kustomization.yaml
    + sed -i -E /^  - name: graphql$/{n;s/newTag: .*/newTag: "3"/} k8s/base/kustomization.yaml
    + git add k8s/base/kustomization.yaml
    + git diff --cached --stat
    + git commit -m chore(deploy): update backend images to 3
    On branch main
    Your branch is up to date with 'origin/main'.
    Untracked files:
            .scannerwork/
            dependency-check-report.xml
            trivyfs.txt
            trivyimage.txt
    nothing added to commit but untracked files present (use "git add" to track)

The clue is the **empty output** of `git diff --cached --stat`: nothing changed. (The untracked files are scan reports from earlier stages, which is normal. `git add` only adds the kustomization file.)

### How we got there
The first successful backend build after a full rebuild happened to be **#3**, and the tags in git were also `"3"`, from the previous setup's builds.

### Why (root cause)
1. Jenkins build numbers restarted at 1 on the new server, but `kustomization.yaml` kept the old setup's tags.
2. Build #3 tried to set the tags to `"3"`, their current value, so the file did not change.
3. The stage ran `git commit` unconditionally, and `git commit` exits non-zero when there is nothing to commit.
4. The non-zero exit failed the stage, and the build showed as failed even though its real work (pushing images) had succeeded.

Design weakness: the stage assumed every build changes the tags. It was not **idempotent**: a re-run, or a matching build number, breaks it.

### Impact
- Build #3 shown as failed. No real damage:

  | Check | Result |
  |---|---|
  | ECR | `account:3`, `catalog:3`, `order:3`, `graphql:3` (pushed by build #3) |
  | `main` → `k8s/base/kustomization.yaml` | backend tags `"3"` → **already points at the images just pushed** |
  | Deployment | not blocked; Argo CD would deploy tag 3, which exists |

### Resolution (step by step)
1. **Check the artifacts, not just the stage colour** (laptop):
   ```bash
   for r in account catalog order graphql; do
     echo "$r: $(aws ecr list-images --region eu-north-1 --repository-name $r --query 'imageIds[].imageTag' --output text)"
   done
   git fetch origin && git show origin/main:k8s/base/kustomization.yaml | sed -n '/^images:/,$p'
   ```
   Images at `3`, and `main` at `"3"` → nothing to fix for this deploy.
2. **Read the stage's console log:** the empty `git diff --cached --stat` followed by `nothing added to commit` confirms the cause.
3. **Make the stage idempotent** (both Jenkinsfiles, diff below): if nothing is staged, print a message and skip commit, pull and push. Otherwise behave as before.
4. **Test the logic** in a scratch repo: setting the same tag → "already current, skip" (exit 0); a new tag → commit made.
5. **Carry on with the deployment** (Step 14). The next backend build (#4) changes the tag to `"4"` and commits normally.

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
# Jenkins console, backend #3, stage "Update Deployment file":
+ git fetch origin main
 * branch            main       -> FETCH_HEAD
+ git checkout -B main origin/main
Switched to a new branch 'main'
+ sed -i -E /^  - name: account$/{n;s/newTag: .*/newTag: "3"/} k8s/base/kustomization.yaml
+ sed -i -E /^  - name: catalog$/{n;s/newTag: .*/newTag: "3"/} k8s/base/kustomization.yaml
+ sed -i -E /^  - name: order$/{n;s/newTag: .*/newTag: "3"/} k8s/base/kustomization.yaml
+ sed -i -E /^  - name: graphql$/{n;s/newTag: .*/newTag: "3"/} k8s/base/kustomization.yaml
+ git add k8s/base/kustomization.yaml
+ git diff --cached --stat
+ git commit -m chore(deploy): update backend images to 3
On branch main
Your branch is up to date with 'origin/main'.

Untracked files:
  (use "git add <file>..." to include in what will be committed)
	.scannerwork/
	dependency-check-report.xml
	trivyfs.txt
	trivyimage.txt

nothing added to commit but untracked files present (use "git add" to track)
```
**Fixing it (testing the new guard in a scratch repo)**
```console
laptop$ git init -q gt && cd gt && printf 'images:\n  - name: account\n    newTag: "3"\n' > k.yaml && git add . && git commit -qm init
laptop$ run(){ T=$1; sed -i -E "/^  - name: account$/{n;s/newTag: .*/newTag: \"$T\"/}" k.yaml; git add k.yaml;
               if git diff --cached --quiet; then echo "tag $T: already current, skip (rc=0)"; else git commit -qm "bump $T" && echo "tag $T: committed"; fi; }
laptop$ run 3; run 4; git log --oneline | head -2
tag 3: already current, skip (rc=0)
tag 4: committed
c9ab4e5 bump 4
26c552f init
```

### Code / config change
`jenkins/Jenkinsfile-Backend` (same change in `jenkins/Jenkinsfile-Frontend`):

```diff
                         git add k8s/base/kustomization.yaml
                         git diff --cached --stat
-                        git commit -m "chore(deploy): update backend images to ${IMAGE_TAG}"
-
-                        # The frontend pipeline may have pushed meanwhile
-                        git pull --rebase origin ${GIT_BRANCH}
-                        git push https://...@github.com/<owner>/go-grpc-graphql-micro.git HEAD:${GIT_BRANCH}
+
+                        # Nothing to do if main already has these tags (e.g. a re-run, or a
+                        # build number that matches the current tag); git commit would fail.
+                        if git diff --cached --quiet; then
+                            echo "k8s/base/kustomization.yaml already at ${IMAGE_TAG}, nothing to commit"
+                        else
+                            git commit -m "chore(deploy): update backend images to ${IMAGE_TAG}"
+
+                            # The frontend pipeline may have pushed meanwhile
+                            git pull --rebase origin ${GIT_BRANCH}
+                            git push https://...@github.com/<owner>/go-grpc-graphql-micro.git HEAD:${GIT_BRANCH}
+                        fi
```

`git diff --cached --quiet` exits `0` when nothing is staged, and `1` when something is.

### Lessons learned
- Every pipeline step should be **idempotent**: a re-run with the same input should succeed and do nothing, not fail.
- `git commit` failing on "nothing to commit" is a classic CI trap. Guard it with `git diff --cached --quiet`.
- After rebuilding CI from scratch, build numbers restart while git keeps the old tags, so they can collide. With immutable ECR tags this is still safe here, because the new registry is empty. It would not be safe against a registry that still had old images under the same tag (the push would fail on immutability).
- A red stage at the very end doesn't mean the build's work failed. Check what actually happened (ECR, git).

### How to approach it next time
1. Read the last few commands before the error in the console. An **empty** `git diff --cached --stat` means there was nothing to commit.
2. Compare the tag being written (`newTag: "N"`) with what `main` already has.
3. Check whether the stage's real goal (images in ECR, the right tag on `main`) was already met.

### Prevention / follow-up
- Done: both Jenkinsfiles skip commit and push when the tags are already current.
- Optional: use a tag that can't collide with older builds, for example `${BUILD_NUMBER}-${GIT_COMMIT[0..6]}`. That also links each image to its exact commit.
- Optional: after a full teardown, reset the tags in `kustomization.yaml` (the guide's Step 1 already does this for forks: `newTag: "0"`).

### References
- Git: [`git diff --quiet` / `--cached`](https://git-scm.com/docs/git-diff); [`git commit`](https://git-scm.com/docs/git-commit) exit status.
- Jenkins: [`sh` step](https://www.jenkins.io/doc/pipeline/steps/workflow-durable-task-step/#sh-shell-script) (non-zero exit fails the step).

---

## INC-010 – `backend` pipeline aborted after 60 minutes in OWASP Dependency-Check

| | |
|---|---|
| **Date** | 2026-10-08 |
| **Where** | Jenkins job `backend` › stage **OWASP Dependency-Check Scan**; `jenkins/Jenkinsfile-Backend` (`options { timeout(60 MINUTES) }`); Recreate guide Step 13 |
| **Severity** | Medium: no backend images were built or pushed, which blocked the deployment; nothing was broken |
| **Status** | Resolved: the re-run (build #3) got through OWASP quickly and pushed all four images. Why the first download was so slow is unconfirmed; the long-term fix is a follow-up |

### Summary
The first `backend` run on the fresh Jenkins server spent **57 min 36 s** in the OWASP Dependency-Check stage, downloading the vulnerability database. Together with the earlier stages, that crossed the pipeline's **60-minute timeout**, so Jenkins **aborted** the build. Every later stage was cancelled, including Docker build, ECR push and the deployment-tag commit. As a result the four backend ECR repositories stayed empty, while `frontend` (run afterwards, reusing the downloaded database) succeeded. Re-running `backend` with the database already on disk is the workaround.

### Background
- **OWASP Dependency-Check** compares our dependencies (Go modules, npm packages) with the **NVD** (National Vulnerability Database) to find known CVEs. It keeps a local copy of the NVD on the Jenkins server.
- On the **first run** on a new server that copy is empty, so it downloads the whole NVD, a few hundred thousand CVE records, through the NVD API. Later runs only fetch changes and take a minute or two.
- The NVD API is **rate-limited**. With an **API key** the first download usually takes about 10–20 minutes. Without a working key the limit is much stricter, and the download can take an hour or more. NVD is also sometimes slow or returns errors, which Dependency-Check retries.
- Both pipelines use the same Dependency-Check installation (`DP-Check`), so they **share one local database** on the Jenkins server.
- `options { timeout(time: 60, unit: 'MINUTES') }` in a declarative pipeline is a limit for the **whole run**. When it expires, Jenkins **aborts** the build (status *ABORTED*, not *FAILURE*), and the remaining stages show as *aborted* in Stage View.
- Separately, the Quality Check stage showing "paused for Xs" is normal: it waits for SonarQube's webhook (see Step 4.3).

### Timeline
1. Fresh Jenkins server (recreated today), so the Dependency-Check database was empty.
2. Ran `backend` (build #1). Tests, Sonar analysis and Quality Check passed in about 2–3 minutes.
3. OWASP Dependency-Check started the full NVD download and ran for 57 min 36 s.
4. The 60-minute pipeline timeout expired. Jenkins aborted the build, and Trivy, Docker build, Trivy image scan, ECR push and Update Deployment were all cancelled.
5. Ran `frontend`. Its OWASP stage reused the now-populated database and finished; it pushed `frontend:2` and committed the tag.
6. Before creating the Argo CD Application, checked ECR and `main`: `frontend:2` was present, but `account`, `catalog`, `order` and `graphql` were empty, still on the old tag `"3"`.
7. Stage View for `backend` showed the 57 min 36 s OWASP stage followed by aborted stages. Matched it to the 60-minute timeout in the Jenkinsfile.

### What happened (symptom)
`backend` Stage View (per stage):

    Quality Check            899ms (paused for 7s)
    OWASP Dependency-Check   57min 36s   aborted
    Trivy File Scan          171ms       aborted
    Docker Image Build       148ms       aborted
    TRIVY Image Scan         162ms       aborted
    ECR Image Pushing        146ms       aborted
    Update Deployment file   157ms       aborted

Effect outside Jenkins:

| Check | Result |
|---|---|
| `aws ecr list-images` for account, catalog, order, graphql | no images |
| `aws ecr list-images` for frontend | `2` |
| `k8s/base/kustomization.yaml` on `main` | backend tags still `"3"` (from the previous, torn-down setup); frontend `"2"` |
| Commits on `main` | `chore(deploy): update frontend image to 2`, but no backend deploy commit |

Had the Argo CD Application been created at this point, the four Go services would have gone to `ImagePullBackOff` (image tag `3` does not exist in the new registry).

### How we got there
First `backend` run on a brand-new Jenkins server, where the OWASP database had to be downloaded from scratch inside a pipeline limited to 60 minutes.

### Why (root cause)
1. The Dependency-Check database lives on the Jenkins server, and the server was new, so the first scan had to download the full NVD.
2. That download took 57½ minutes, far longer than the usual 10–20 minutes with an API key. *Why it was that slow is still to confirm from the console log*: either the NVD API key wasn't used (not activated, or an invalid credential), or NVD was rate-limiting or erroring.
3. The pipeline-wide timeout is 60 minutes, and the earlier stages had used about 2–3 minutes, so the timeout expired during OWASP.
4. A timeout **aborts** the run, so the build, push and tag-update stages never ran, and nothing reached ECR.
5. `frontend` succeeded only because it ran **after** backend had filled the shared database.

Design weakness: a **one-time, slow, external download** (the NVD) sits inside the normal build pipeline and is bounded by the same timeout as the build.

### Impact
- No backend images in ECR, and no backend tag commit. Deployment blocked until a successful re-run.
- About an hour of pipeline time. No cost beyond the running EC2 instance. Nothing deployed or broken.

### Resolution (step by step)
1. **Read Stage View** to find which stage ran long and where *aborted* starts: OWASP, 57 min 36 s.
2. **Compare with the Jenkinsfile:** `options { timeout(time: 60, unit: 'MINUTES') }` covers the whole run → the timeout explains the abort.
3. **Check the console log of the aborted build for the NVD key** (search `NVD`, `API key`, `403`, `429`):
   - a warning like *"An NVD API Key was not provided"* → re-check the `nvd-api-key` credential and that the key was activated from NVD's email;
   - many `403`/`429`/`503` responses → NVD was rate-limiting or unavailable.
4. **Re-run `backend`** (Build Now). The database is already populated (backend's partial download plus frontend's completed update), so OWASP only fetches updates and the run reaches ECR push and Update Deployment.
5. **Verify:**
   ```bash
   for r in account catalog order graphql; do
     echo "$r: $(aws ecr list-images --region eu-north-1 --repository-name $r --query 'imageIds[].imageTag' --output text)"
   done
   git fetch origin && git log --oneline origin/main -3     # expect: chore(deploy): update backend images to N
   ```
6. Only then create the Argo CD Application (Step 14).

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
# Jenkins UI, backend -> Stage View (per stage):
Quality Check            899ms (paused for 7s)
OWASP Dependency-Check   57min 36s   aborted
Trivy File Scan          171ms       aborted
Docker Image Build       148ms       aborted
TRIVY Image Scan         162ms       aborted
ECR Image Pushing        146ms       aborted
Update Deployment file   157ms       aborted

laptop$ for r in account catalog order graphql frontend; do echo "$r: $(aws ecr list-images --region eu-north-1 --repository-name $r --query 'imageIds[].imageTag' --output text)"; done
account:
catalog:
order:
graphql:
frontend: 2
laptop$ git fetch -q origin && git log --oneline origin/main -5
9313d5b chore(deploy): update frontend image to 2
66f7ee2 docs: update SigNoz installation instructions to use Foundry and provide YAML configuration
…
laptop$ git show origin/main:k8s/base/kustomization.yaml | sed -n '/^images:/,$p'
images:
  - name: account
    newTag: "3"
  …
  - name: frontend
    newTag: "2"
laptop$ cat -n jenkins/Jenkinsfile-Backend | sed -n 20,23p
    20	    options {
    21	        disableConcurrentBuilds()
    22	        timeout(time: 60, unit: 'MINUTES')
    23	    }
```
**Fixing it**
```console
# Jenkins UI: backend -> Build Now (build #3; OWASP now quick)
laptop$ for r in account catalog order graphql frontend; do echo "$r: $(aws ecr list-images --region eu-north-1 --repository-name $r --query 'imageIds[].imageTag' --output text)"; done
account: 3
catalog: 3
order: 3
graphql: 3
frontend: 2
```

### Code / config change
None yet (see follow-up).

### Lessons learned
- **ABORTED is not FAILURE.** An aborted build usually means a timeout or a manual cancel, not a broken step. Look at the duration of the last running stage.
- Before deploying, check the **artifacts** (images in ECR, the tag commit), not just the colour of the job. That is how this was caught before Argo CD tried to pull non-existent images.
- First runs on a fresh CI server are special: caches (the NVD database, Go module cache, Docker layers) are empty, and the first run can take several times longer than usual.

### How to approach it next time
1. Stage View: find the **first** stage that isn't green. Everything after an abort shows *aborted* within milliseconds, which is a side effect, not the cause.
2. Long duration + *aborted* → compare with `timeout(...)` in the Jenkinsfile (both the pipeline `options` and any stage-level `timeout`).
3. For OWASP specifically: console log → is the NVD API key accepted? Is NVD returning errors?
4. Re-run once the database exists. If it is still slow, the key is the likely problem.

### Prevention / follow-up
Options, from simplest to most robust:
- **Warm the database once** after creating a new Jenkins server: run `backend` once and let OWASP finish (temporarily raise the timeout to 120 minutes for that run), or run Dependency-Check's `--updateonly` from a separate job.
- **Give OWASP its own stage-level timeout**, and keep the pipeline timeout for the build itself, so a slow download is reported clearly rather than aborting the whole run.
- **Separate the NVD update from the build:** a nightly job runs `--updateonly`, and the build pipelines run with `--noupdate`, so builds never wait on NVD.
- Make sure the NVD API key is valid and activated (the biggest single factor in download time).

### References
- OWASP Dependency-Check: [NVD API key and data feeds](https://jeremylong.github.io/DependencyCheck/data/index.html); [CLI arguments (`--updateonly`, `--noupdate`, `--nvdApiKey`)](https://jeremylong.github.io/DependencyCheck/dependency-check-cli/arguments.html).
- Jenkins: [Pipeline `options` — `timeout`](https://www.jenkins.io/doc/book/pipeline/syntax/#options).
- NVD: [Request an API key](https://nvd.nist.gov/developers/request-an-api-key).

---

## INC-009 – Argo CD UI "not reachable" in the browser

| | |
|---|---|
| **Date** | 2026-10-08 |
| **Where** | Browser on the laptop → `https://<LB_HOST>/` (Argo CD UI); Recreate guide Step 8 (expose the UI) |
| **Severity** | Low: the UI was blocked for one person; the cluster and Argo CD were healthy |
| **Status** | Resolved: it started working after 2–3 minutes; cause was the new load balancer's DNS not being live yet |

### Summary
After exposing Argo CD through a load balancer restricted to our IP, the URL would not open in the browser. Testing from the same laptop and network with `curl` showed the load balancer working: HTTPS returned `200` and HTTP returned a `307` redirect to HTTPS. So the problem was on the browser side, not in AWS or Kubernetes. After waiting 2–3 minutes it opened normally: the load balancer's DNS name had not finished going live when the browser first looked it up, and the browser kept that failed lookup for a short while. No change was needed.

### Background
- `kubectl patch svc argocd-server -p '{"spec":{"type":"LoadBalancer"}}'` makes AWS create a **load balancer** in front of the Argo CD server. Its DNS name looks like `<hash>-<id>.eu-north-1.elb.amazonaws.com`. That `<region>.elb.amazonaws.com` form is a **Classic Load Balancer**, created by the cluster's built-in AWS cloud provider (the AWS Load Balancer Controller creates NLBs and ALBs instead).
- `loadBalancerSourceRanges: ["<YOUR_IP>/32"]` becomes a **security group rule**: only that IP can connect. Anyone else's connection just times out.
- **A new load balancer's DNS name takes a few minutes to exist.** If a browser looks it up during that time, it gets "name not found", and browsers cache that failure for a while, even after the name starts working.
- Argo CD serves HTTPS with a **self-signed certificate**, so browsers show a "connection is not private" warning that must be accepted manually. Plain HTTP gets a `307` redirect to HTTPS.
- The terminal and the browser can leave the laptop through **different IPs** (a VPN, a proxy extension) and can resolve DNS differently (the browser's "secure DNS" / DNS-over-HTTPS).

### Timeline
1. Step 8: installed Argo CD and patched `argocd-server` to `type: LoadBalancer` with `loadBalancerSourceRanges` set to our IP. The load balancer was created at 09:39:52 UTC.
2. Opened `https://<LB_HOST>/` in the browser: "not reachable".
3. From the laptop terminal (same network), checked DNS, HTTPS, HTTP and the load balancer's configuration (below). Everything worked.
4. Narrowed the cause to the browser.
5. Waited 2–3 minutes and reloaded: the Argo CD login page opened. Confirmed cause: DNS not yet live / cached failed lookup.

### What happened (symptom)
The browser reported the site as not reachable, a few minutes after the load balancer was created. The same URL opened fine 2–3 minutes later, with no changes.

From the terminal on the same laptop, it **worked**:

| Check | Command | Result |
|---|---|---|
| DNS | `getent hosts <LB_HOST>` | resolves to 2 public IPs |
| HTTPS | `curl -sk -o /dev/null -w '%{http_code}' https://<LB_HOST>/` | `200` in 0.8 s |
| HTTP | `curl -s -o /dev/null -w '%{http_code}' http://<LB_HOST>/` | `307` (redirect to HTTPS) |
| Load balancer | `aws elb describe-load-balancers --load-balancer-names <name>` | `internet-facing`, TCP 443 and 80, 2 instances registered |
| Allowlist | (implied by curl succeeding) | the laptop's current IP is allowed |

### How we got there
Opened the URL in the browser right after creating the load balancer, before (or while) its DNS name became resolvable.

### Why (root cause)
1. AWS side is healthy: the load balancer is internet-facing, has listeners, has registered targets, and answers our IP.
2. Kubernetes side is healthy: Argo CD answered with `200` through the load balancer.
3. So the failure was in how the **browser** reached the same name. Candidates were: a cached DNS failure, the certificate warning, a different outgoing IP (VPN/proxy), or secure DNS.
4. **Confirmed:** it fixed itself after 2–3 minutes with no change. That rules out the IP allowlist and VPN (those don't fix themselves) and points to DNS timing:
   - AWS publishes a new load balancer's DNS name a few minutes after creation;
   - the browser's first lookup came back "not found", and browsers and OS resolvers cache negative answers for a short time;
   - once the record was live and the cached failure had expired, the page loaded.

### Impact
- The Argo CD UI was unavailable in the browser for 2–3 minutes. Argo CD itself, the cluster and the load balancer were unaffected.

### Resolution (step by step)
1. **Prove the server side works** (laptop terminal):
   ```bash
   curl -sk -o /dev/null -w '%{http_code}\n' https://<LB_HOST>/     # expect 200
   ```
   `200` means AWS and Argo CD are fine; continue with browser checks. A timeout means an allowlist or IP problem (go to step 4).
2. **Clear the browser's DNS cache:** `chrome://net-internals/#dns` → *Clear host cache*, then `chrome://net-internals/#sockets` → *Flush socket pools*. Or try an incognito window.
3. **Accept the certificate warning:** *Advanced → Proceed to <LB_HOST> (unsafe)*. If there's no Proceed link, type `thisisunsafe` on the warning page.
4. **Compare IPs:** open `https://checkip.amazonaws.com` in the **browser** and run `curl https://checkip.amazonaws.com` in the terminal. If they differ, turn off the VPN/proxy, or add the browser's IP to `loadBalancerSourceRanges`.
5. **Rule out secure DNS:** Settings → Privacy and security → Security → *Use secure DNS* → off, then retry.
6. **What fixed it here:** waiting 2–3 minutes and reloading. Steps 2–5 are for when waiting doesn't help.

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
# browser: https://<LB_HOST>/ -> "not reachable"
laptop$ getent hosts <LB_HOST> | awk '{print $1}'
13.61.x.x
51.21.x.x
laptop$ curl -sk -o /dev/null -w '%{http_code} %{time_total}s\n' --max-time 10 https://<LB_HOST>/
200 0.785122s
laptop$ curl -s -o /dev/null -w '%{http_code} %{time_total}s\n' --max-time 10 http://<LB_HOST>/
307 0.425301s
laptop$ aws elb describe-load-balancers --region eu-north-1 --load-balancer-names <lb-name> \
  --query 'LoadBalancerDescriptions[0].{scheme:Scheme,listeners:ListenerDescriptions[].Listener.[Protocol,LoadBalancerPort,InstancePort],created:CreatedTime,instances:length(Instances)}' --output json
{
    "scheme": "internet-facing",
    "listeners": [ [ "TCP", 443, 32598 ], [ "TCP", 80, 31154 ] ],
    "created": "2026-10-08T09:39:52.020000+00:00",
    "instances": 2
}
```
**Fixing it**
```console
# No commands: waited 2-3 minutes for the new load balancer's DNS name and reloaded the page -> Argo CD login.
```

### Code / config change
None.

### Lessons learned
- Separate "is the server reachable?" from "is my browser reaching it?". `curl` from the same machine answers the first question in seconds.
- New load balancer names need a few minutes. Opening them too early can leave a cached failure that outlives the actual delay.

### How to approach it next time
1. `curl` the URL from the same laptop.
   - `200`/`30x` → server is fine; it's the browser (DNS cache, certificate, proxy).
   - **Timeout** → network path: IP allowlist (`loadBalancerSourceRanges`), wrong IP, or security group.
   - **`Could not resolve host`** → DNS not ready yet (wait 2–5 minutes) or a wrong name.
   - **Connection refused / 503** → load balancer reached, but no healthy targets: check the pods and service on the bastion (`kubectl -n argocd get pods,svc`).
2. Check the load balancer in AWS: `aws elb describe-load-balancers` (Classic) or `aws elbv2 describe-load-balancers` (NLB/ALB), and the target health.

### Prevention / follow-up
- Wait 2–3 minutes after `kubectl get svc` shows the hostname before opening it in a browser (the guide already says so for the ALB).
- Possible improvement: expose Argo CD through the existing ALB Ingress, or use `kubectl port-forward` from the bastion, instead of a separate Classic Load Balancer.

### References
- Kubernetes: [Service type LoadBalancer](https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer); `loadBalancerSourceRanges`.
- Argo CD: [Getting started — access the Argo CD API server](https://argo-cd.readthedocs.io/en/stable/getting_started/#3-access-the-argo-cd-api-server).

---

## INC-008 – LB Controller IAM role create fails: `EntityAlreadyExists`

| | |
|---|---|
| **Date** | 2026-10-08 |
| **Where** | Bastion, `~/lbc`; Recreate guide Step 7.2 (IAM role for the AWS Load Balancer Controller) |
| **Severity** | Low: no impact, the role already existed and was correct |
| **Status** | Explained (no fix needed) |

### Summary
Creating the AWS Load Balancer Controller's IAM role failed with `EntityAlreadyExists`. The step's commands had **already run successfully once**: the policy and the role were created 8 seconds apart a few minutes earlier, and this was a second paste of the same block. The existing role had the correct trust policy and the correct policy attached, so we continued with the next step (the Pod Identity association), which had not been done yet.

### Background
- The **AWS Load Balancer Controller** runs in the cluster and creates AWS load balancers for Kubernetes Ingresses and Services. It needs AWS permissions, which it gets through **EKS Pod Identity**:
  1. an IAM **policy** (`AWSLoadBalancerControllerIAMPolicy`) lists what it may do (Step 7.1);
  2. an IAM **role** (`<cluster>-lb-controller-role`) holds that policy. Its **trust policy** says that only the EKS Pod Identity service (`pods.eks.amazonaws.com`) may assume it (Step 7.2);
  3. a **Pod Identity association** links the role to one Kubernetes service account, `kube-system/aws-load-balancer-controller` (Step 7.3). Only pods running as that service account get the role's credentials.
- `aws iam create-role` and `create-policy` are **not idempotent**: if the name exists, they fail with `EntityAlreadyExists`.
- `aws iam attach-role-policy` **is** idempotent: attaching a policy that is already attached succeeds and prints nothing.
- IAM is **global** (not per region), so a role from any earlier run is still there.

### Timeline
1. On the bastion, ran Step 7.1 (create the policy) and Step 7.2 (create the role and attach the policy). Both succeeded at 09:35:46 and 09:35:54 UTC.
2. Pasted the Step 7.2 block again.
3. `create-role` failed with `EntityAlreadyExists`. `attach-role-policy` ran silently.
4. Checked the role, its trust policy, its attached policies and the Pod Identity associations (below).
5. The role was correct; the association (Step 7.3) was still missing. Continued with 7.3.

### What happened (symptom)

    aws: [ERROR]: An error occurred (EntityAlreadyExists) when calling the CreateRole operation:
    Role with name go-microservices-dev-lb-controller-role already exists.

### How we got there
The same Step 7.2 block was pasted twice. The first run's output had scrolled past or was not noticed, so the step looked "not done".

### Why (root cause)
1. The role already existed from the first, successful run.
2. `create-role` fails on an existing name instead of skipping it.
3. The guide's step is imperative (create commands), so it is not safe to re-run.

This is the same pattern as INC-006 (ECR).

### Impact
None. The existing setup was exactly what the step creates:

| Check | Result |
|---|---|
| Policy `AWSLoadBalancerControllerIAMPolicy` | exists, version `v1`, created 09:35:46 UTC |
| Role `go-microservices-dev-lb-controller-role` | exists, created 09:35:54 UTC |
| Trust policy | principal `pods.eks.amazonaws.com`, actions `sts:AssumeRole` + `sts:TagSession` |
| Attached policies | `AWSLoadBalancerControllerIAMPolicy` |
| Pod Identity associations | `kube-system/ebs-csi-controller-sa`, `external-secrets/external-secrets` (both from Terraform). The controller's association was **not** there yet |

### Resolution (step by step)
1. **Check the role and its trust policy:**
   ```bash
   aws iam get-role --role-name go-microservices-dev-lb-controller-role \
     --query 'Role.{created:CreateDate,trust:AssumeRolePolicyDocument}'
   ```
   Expect `"Service": "pods.eks.amazonaws.com"` and both `sts:AssumeRole` and `sts:TagSession`.
2. **Check the policy is attached:**
   ```bash
   aws iam list-attached-role-policies --role-name go-microservices-dev-lb-controller-role
   ```
   Expect `AWSLoadBalancerControllerIAMPolicy`.
3. **Check which Pod Identity associations exist**, to see how far the setup got:
   ```bash
   aws eks list-pod-identity-associations --region eu-north-1 --cluster-name go-microservices-dev \
     --query 'associations[].[namespace,serviceAccount]' --output text
   ```
   No `kube-system aws-load-balancer-controller` line means Step 7.3 is still to do.
4. **Continue with Step 7.3** (create the association), then 7.4 (Helm install).

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
[ec2-user@ip-10-10-20-x lbc]$ cat > trust.json <<'EOF'
{ … "Principal": { "Service": "pods.eks.amazonaws.com" }, "Action": ["sts:AssumeRole", "sts:TagSession"] … }
EOF
[ec2-user@ip-10-10-20-x lbc]$ aws iam create-role --role-name ${CLUSTER_NAME}-lb-controller-role --assume-role-policy-document file://trust.json
aws: [ERROR]: An error occurred (EntityAlreadyExists) when calling the CreateRole operation: Role with name go-microservices-dev-lb-controller-role already exists.
[ec2-user@ip-10-10-20-x lbc]$ aws iam attach-role-policy --role-name ${CLUSTER_NAME}-lb-controller-role \
  --policy-arn arn:aws:iam::${ACCOUNT_ID}:policy/AWSLoadBalancerControllerIAMPolicy
(no output: attaching an already-attached policy succeeds silently)
```
**Checking it**
```console
laptop$ aws iam get-role --role-name go-microservices-dev-lb-controller-role --query 'Role.{created:CreateDate,trust:AssumeRolePolicyDocument}' --output json
{
    "created": "2026-10-08T09:35:54+00:00",
    "trust": {
        "Version": "2012-10-17",
        "Statement": [ { "Effect": "Allow",
                         "Principal": { "Service": "pods.eks.amazonaws.com" },
                         "Action": [ "sts:AssumeRole", "sts:TagSession" ] } ]
    }
}
laptop$ aws iam list-attached-role-policies --role-name go-microservices-dev-lb-controller-role --query 'AttachedPolicies[].PolicyName' --output text
AWSLoadBalancerControllerIAMPolicy
laptop$ aws iam list-policies --scope Local --query 'Policies[?PolicyName==`AWSLoadBalancerControllerIAMPolicy`].[PolicyName,CreateDate,DefaultVersionId]' --output text
AWSLoadBalancerControllerIAMPolicy	2026-10-08T09:35:46+00:00	v1
laptop$ aws eks list-pod-identity-associations --region eu-north-1 --cluster-name go-microservices-dev --query 'associations[].[namespace,serviceAccount]' --output text
kube-system	ebs-csi-controller-sa
external-secrets	external-secrets
```
**Fixing it (continue with the missing step)**
```console
bastion$ source ~/lbc/env.sh
bastion$ aws eks create-pod-identity-association --region $AWS_REGION --cluster-name $CLUSTER_NAME \
  --namespace kube-system --service-account aws-load-balancer-controller \
  --role-arn arn:aws:iam::${ACCOUNT_ID}:role/${CLUSTER_NAME}-lb-controller-role
(output not captured)
```

### Code / config change
None.

### Lessons learned
- After an `AlreadyExists` error, check the existing resource *and* the next step's resource. That shows exactly where a half-remembered sequence stopped.
- Run multi-command setup blocks once, and read their output before moving on.

### How to approach it next time
1. `EntityAlreadyExists` → `get-role` / `get-policy` and compare with what the command would create.
2. Find the first step whose result is missing (here, the Pod Identity association) and continue from there.
3. If the existing role is wrong (for example a bad trust policy), fix it in place with `aws iam update-assume-role-policy`, rather than deleting a role that something may already use.

### Prevention / follow-up
- Optional: move the controller's policy, role and association into Terraform (`terraform/eks-cluster`). It already manages the EBS CSI and External Secrets equivalents the same way, so the whole thing becomes declarative and re-runnable.

### References
- AWS: [EKS Pod Identity](https://docs.aws.amazon.com/eks/latest/userguide/pod-identities.html); [AWS Load Balancer Controller installation](https://kubernetes-sigs.github.io/aws-load-balancer-controller/latest/deploy/installation/).
- AWS CLI: [`iam create-role`](https://docs.aws.amazon.com/cli/latest/reference/iam/create-role.html), [`iam attach-role-policy`](https://docs.aws.amazon.com/cli/latest/reference/iam/attach-role-policy.html).

---

## INC-007 – `su: user ec2-user does not exist` on the "bastion"

| | |
|---|---|
| **Date** | 2026-10-08 |
| **Where** | SSM session opened from the laptop; Recreate guide Step 6 (use the bastion) |
| **Severity** | Low: only blocked the step until the right machine was used |
| **Status** | Resolved |

### Summary
The guide's first command on the bastion, `sudo su - ec2-user`, failed with "user ec2-user does not exist". The session was not on the bastion at all: it was on the **Jenkins server**, which runs Ubuntu and has no `ec2-user`. Connecting to the bastion's instance, looked up by its Name tag, fixed it.

### Background
- Two kinds of EC2 machines are reachable with **SSM Session Manager** (`aws ssm start-session --target <instance-id>`):

  | Machine | OS | Default login user | Purpose |
  |---|---|---|---|
  | `jenkins-server` | Ubuntu 24.04 | `ubuntu` | CI: Jenkins + SonarQube |
  | `go-microservices-dev-bastion` | Amazon Linux 2023 | `ec2-user` | the only machine allowed to run `kubectl` against the private EKS API |

  (The EKS worker nodes are also SSM-managed, but nobody logs into them.)
- SSM always starts you as `ssm-user`. The guide switches to `ec2-user` on the bastion for a normal home directory and the preconfigured `kubectl`.
- An SSM target is just an **instance ID**. Nothing in the session prompt says which machine you are on, until you look.

### Timeline
1. Step 2 (Jenkins setup) used an SSM session on the Jenkins server to watch its install log.
2. In Step 6, an SSM session was started, most likely with the Jenkins instance ID still at hand from Step 2.
3. `sudo su - ec2-user` failed with "user ec2-user does not exist".
4. Checked from the laptop which OS each instance runs: the bastion is Amazon Linux 2023 (`…d1c0`); the only Ubuntu instance is Jenkins (`…bca3`).
5. Started a new session using the bastion's ID, looked up by Name tag. `sudo su - ec2-user` worked; the prompt became `ec2-user@ip-10-10-…`.

### What happened (symptom)

    su: user ec2-user does not exist or the user entry does not contain all the required fields

### How we got there
`aws ssm start-session --target <id>` with the **Jenkins** instance ID instead of the bastion's.

### Why (root cause)
1. The session was on the Jenkins server (Ubuntu).
2. Ubuntu cloud images create a `ubuntu` user, not `ec2-user`.
3. Instance IDs are opaque and both machines are SSM targets, so it is easy to reuse the wrong ID from earlier output.

### Impact
- None. Only the `su` failed; nothing ran on the Jenkins server.

### Resolution (step by step)
1. **Check where you are** (inside the session):
   ```bash
   hostname; grep PRETTY_NAME /etc/os-release
   ```
   `Ubuntu 24.04` → Jenkins server. `Amazon Linux 2023` → bastion.
2. **Leave the wrong session:** `exit`.
3. **Look up the bastion by its Name tag** (laptop) instead of copying an ID:
   ```bash
   BASTION=$(aws ec2 describe-instances --region eu-north-1 \
     --filters Name=tag:Name,Values=go-microservices-dev-bastion Name=instance-state-name,Values=running \
     --query 'Reservations[0].Instances[0].InstanceId' --output text)
   aws ssm start-session --region eu-north-1 --target $BASTION
   ```
4. **Switch user and verify:**
   ```bash
   sudo su - ec2-user
   kubectl get nodes          # 2 nodes Ready
   ```

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
laptop$ aws ssm start-session --region eu-north-1 --target <instance-id>      # an ID copied from earlier
$ sudo su - ec2-user
su: user ec2-user does not exist or the user entry does not contain all the required fields

laptop$ for n in go-microservices-dev-bastion jenkins-server; do …describe-instances / describe-images…; done
go-microservices-dev-bastion: instance ...d1c0, AMI al2023-ami-2023.12.20260930.0-kernel-6.18-x86_64
jenkins-server: instance ...bca3, AMI ubuntu/images/hvm-ssd-gp3/ubuntu-noble-24.04-amd64-server-20261004
laptop$ aws ssm describe-instance-information --region eu-north-1 --query 'InstanceInformationList[].[InstanceId,PlatformName,PingStatus]' --output text
... 03b1 Amazon Linux Online
... d1c0 Amazon Linux Online
... ec5f Amazon Linux Online
... bca3 Ubuntu Online
```
**Fixing it**
```console
laptop$ BASTION=$(aws ec2 describe-instances --region eu-north-1 \
          --filters Name=tag:Name,Values=go-microservices-dev-bastion Name=instance-state-name,Values=running \
          --query 'Reservations[0].Instances[0].InstanceId' --output text)
laptop$ aws ssm start-session --region eu-north-1 --target $BASTION
$ sudo su - ec2-user
[ec2-user@ip-10-10-20-x ~]$
```

### Code / config change
None. (The bastion command printed by the `eks-cluster` pipeline already has the right ID; this happened because a different, older ID was used.)

### Lessons learned
- With several SSM-managed machines, always confirm which one you are on before running anything: `hostname` and `/etc/os-release`.
- Look instances up by tag rather than copying IDs between steps.

### How to approach it next time
1. "User does not exist" right after connecting → you are probably on a different OS than expected. Check `/etc/os-release`.
2. List what you can connect to and what each one runs:
   ```bash
   aws ssm describe-instance-information --region eu-north-1 \
     --query 'InstanceInformationList[].[InstanceId,PlatformName,PingStatus]' --output table
   ```
3. Match the instance ID to its Name tag with `aws ec2 describe-instances`.

### Prevention / follow-up
- Optional: set a clear shell prompt on each machine at boot (for example `PS1='[bastion] \u@\h \W\$ '` in the bastion's user data), so the machine name is always visible.

### References
- AWS: [Start a Session Manager session](https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-sessions-start.html); [default user names for EC2 AMIs](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/managing-users.html).

---

## INC-006 – ECR `create-repository` fails: `RepositoryAlreadyExistsException`

| | |
|---|---|
| **Date** | 2026-10-08 |
| **Where** | Laptop, AWS CLI; Recreate guide Step 10 (create the ECR repositories) |
| **Severity** | Low: no impact, the repositories were already there and correctly configured |
| **Status** | Explained (no fix needed); guide loop made safe to re-run |

### Summary
Running the guide's loop to create the five ECR repositories printed `RepositoryAlreadyExistsException` for all of them. The loop had **already been run once successfully** earlier in the session, while the EKS cluster was building. The second run failed because `aws ecr create-repository` errors instead of skipping a repository that exists. The existing repositories were checked and had the right settings, so nothing needed fixing. The guide's loop now skips existing repositories, so it can be re-run safely.

### Background
- **ECR** (Elastic Container Registry) holds our Docker images, one repository per image: `account`, `catalog`, `order`, `graphql`, `frontend`. Jenkins pushes to them; EKS pulls from them.
- Each repository is created with two settings that matter:
  - `--image-tag-mutability IMMUTABLE`: a tag such as `account:42` can never be overwritten, so a tag always means exactly one build;
  - `--image-scanning-configuration scanOnPush=true`: AWS scans each pushed image for CVEs.
- `aws ecr create-repository` is **not idempotent**: running it for a name that already exists fails with `RepositoryAlreadyExistsException` instead of reporting "already exists, nothing to do".
- AWS resources are **per region**. Our repositories live in `eu-north-1`. A CLI command without `--region` uses the profile's default region; on this laptop that is `us-east-1`, which has no repositories.
- The registry address `<ACCOUNT_ID>.dkr.ecr.eu-north-1.amazonaws.com` contains the AWS **account ID**. AWS error messages include it too ("registry with id '<ACCOUNT_ID>'"), which is why such output must be redacted before it goes anywhere public.

### Timeline
1. While the `eks-cluster` pipeline was running, the ECR creation loop was run from the laptop. All five repositories were created at 15:00, within 8 seconds of each other.
2. After the cluster was ready, the same loop was run again as the "next step".
3. All five commands failed with `RepositoryAlreadyExistsException`.
4. Checked the repositories read-only (below): all five exist in `eu-north-1` with the right settings, and contain no images yet.
5. Concluded the second run was redundant. Made the guide's loop skip existing repositories.

### What happened (symptom)

    aws: [ERROR]: An error occurred (RepositoryAlreadyExistsException) when calling the
    CreateRepository operation: The repository with name 'account' already exists in the
    registry with id '<ACCOUNT_ID>'

The same error for `catalog`, `order`, `graphql` and `frontend`.

### How we got there
The creation loop ran twice: once early (in parallel with the cluster build, as suggested) and again later, when following the step list in order.

### Why (root cause)
1. The repositories already existed from the first run.
2. `create-repository` treats "exists" as an error rather than a no-op.
3. The guide's loop called `create-repository` unconditionally, so any second run fails, even though the end state is already correct.

### Impact
- None. The existing repositories had exactly the intended configuration:

  | Repository | Tag mutability | Scan on push | Images |
  |---|---|---|---|
  | account, catalog, order, graphql, frontend | `IMMUTABLE` | `True` | 0 (expected before the first pipeline run) |

### Resolution (step by step)
1. **Check what exists, and with which settings:**
   ```bash
   aws ecr describe-repositories --region eu-north-1      --query 'repositories[].[repositoryName,imageTagMutability,imageScanningConfiguration.scanOnPush,createdAt]'      --output table
   ```
   Expect five rows, each `IMMUTABLE` and `True`. Close-together `createdAt` times show they came from one earlier run.
2. **Check they are empty** (so they belong to this rebuild, not an old one):
   ```bash
   for r in account catalog order graphql frontend; do
     echo "$r: $(aws ecr list-images --region eu-north-1 --repository-name $r --query 'length(imageIds)' --output text)"
   done
   ```
   Expect `0` for each until the `backend` and `frontend` pipelines push images.
3. **Carry on.** No change needed in AWS.
4. **Make the guide's loop re-runnable** (diff below): check with `describe-repositories` first and skip existing names.
5. **Set the laptop's default region** so commands without `--region` don't silently look at `us-east-1`:
   ```bash
   aws configure set region eu-north-1
   ```

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
laptop$ for repo in account catalog order graphql frontend; do
          aws ecr create-repository --region eu-north-1 --repository-name $repo \
            --image-scanning-configuration scanOnPush=true --image-tag-mutability IMMUTABLE \
            --query 'repository.repositoryUri' --output text
        done
aws: [ERROR]: An error occurred (RepositoryAlreadyExistsException) when calling the CreateRepository operation: The repository with name 'account' already exists in the registry with id '<ACCOUNT_ID>'
aws: [ERROR]: An error occurred (RepositoryAlreadyExistsException) when calling the CreateRepository operation: The repository with name 'catalog' already exists in the registry with id '<ACCOUNT_ID>'
aws: [ERROR]: An error occurred (RepositoryAlreadyExistsException) when calling the CreateRepository operation: The repository with name 'order' already exists in the registry with id '<ACCOUNT_ID>'
aws: [ERROR]: An error occurred (RepositoryAlreadyExistsException) when calling the CreateRepository operation: The repository with name 'graphql' already exists in the registry with id '<ACCOUNT_ID>'
aws: [ERROR]: An error occurred (RepositoryAlreadyExistsException) when calling the CreateRepository operation: The repository with name 'frontend' already exists in the registry with id '<ACCOUNT_ID>'
```
**Checking it**
```console
laptop$ aws ecr describe-repositories --region eu-north-1 --query 'repositories[].[repositoryName,imageTagMutability,imageScanningConfiguration.scanOnPush,createdAt]' --output table
-----------------------------------------------------------------------
|                        DescribeRepositories                         |
+----------+------------+-------+-------------------------------------+
|  account |  IMMUTABLE |  True |  2026-10-08T15:00:41.099000+05:30   |
|  catalog |  IMMUTABLE |  True |  2026-10-08T15:00:43.049000+05:30   |
|  order   |  IMMUTABLE |  True |  2026-10-08T15:00:44.893000+05:30   |
|  graphql |  IMMUTABLE |  True |  2026-10-08T15:00:46.728000+05:30   |
|  frontend|  IMMUTABLE |  True |  2026-10-08T15:00:48.468000+05:30   |
+----------+------------+-------+-------------------------------------+
laptop$ for r in account catalog order graphql frontend; do echo "$r: $(aws ecr list-images --region eu-north-1 --repository-name $r --query 'length(imageIds)' --output text)"; done
account: 0
catalog: 0
order: 0
graphql: 0
frontend: 0
laptop$ aws configure get region
us-east-1

# suggested, not run:
laptop$ aws configure set region eu-north-1
```

### Code / config change
`docs/RECREATE-GUIDE.md`, Step 10:

```diff
 for repo in account catalog order graphql frontend; do
+  # skip repos that already exist, so the loop is safe to re-run
+  aws ecr describe-repositories --region eu-north-1 --repository-names $repo >/dev/null 2>&1 \
+    && { echo "$repo: already exists"; continue; }
   aws ecr create-repository --region eu-north-1 --repository-name $repo \
```

### Lessons learned
- An "already exists" error from a create command is often good news: check the existing resource's settings before assuming something is broken.
- Setup scripts should be **idempotent** (safe to run twice). Check first, or use a tool that converges to a desired state (Terraform) rather than imperative create calls.
- AWS error messages can leak the account ID. Redact before sharing.

### How to approach it next time
1. On `AlreadyExists`, **describe** the resource and compare its settings with what the command would have created.
2. If the settings match, move on. If they don't, decide whether to update it in place or delete and recreate it (ECR delete also deletes the images; use `--force` only on purpose).
3. If a resource "doesn't exist" but you are sure it does, check the region first: `aws configure get region`.

### Prevention / follow-up
- Done: guide loop skips existing repositories.
- Optional: manage the ECR repositories with Terraform too, so they are created, configured and destroyed with the rest of the infrastructure.

### References
- AWS CLI: [`ecr create-repository`](https://docs.aws.amazon.com/cli/latest/reference/ecr/create-repository.html), [`ecr describe-repositories`](https://docs.aws.amazon.com/cli/latest/reference/ecr/describe-repositories.html).
- Amazon ECR: [Image tag mutability](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-tag-mutability.html).

---

## INC-005 – First `eks-cluster` build fails: `Value must not contain "//"`

| | |
|---|---|
| **Date** | 2026-10-08 |
| **Where** | Jenkins job `eks-cluster` › stage **Init**; file `terraform/eks-cluster/Jenkinsfile`; Recreate guide Step 5 |
| **Severity** | Low: it blocked only the first run, and nothing was created or changed in AWS |
| **Status** | Resolved (unblocked by building with parameters; code fixed so it can't happen again) |

### Summary
The very first build of the new `eks-cluster` Jenkins job failed in the Terraform `init` step. That build ran **before Jenkins knew the job's parameters**, so the `Environment` value was empty and the Terraform state path came out as `eks-cluster//terraform.tfstate`, which current Terraform refuses. We unblocked it with **Build with Parameters**, and changed the Jenkinsfile so that a missing parameter now defaults to a harmless `dev` plan.

### Background
- **Terraform state** is Terraform's record of everything it has created. We keep it in S3 (the *S3 backend*), so it survives laptop loss and is locked during `apply`. Each environment gets its own state file: the backend `key` is `eks-cluster/<env>/terraform.tfstate`, so `dev` and `prod` never overwrite each other.
- The bucket name and key are not in `backend.tf`. The pipeline passes them at `terraform init` time with `-backend-config=...` (a *partial backend configuration*), so the account ID never enters git.
- **Jenkins parameters** (`Environment`: dev/prod, `Terraform_Action`: plan/apply/destroy) are declared **inside the Jenkinsfile** with `properties([parameters([...])])`, not in the Jenkins UI. Jenkins only reads that block **while a build is running**. On a brand-new job, the first build runs before Jenkins has ever seen the parameters, so that build gets none. Only *afterwards* does the job page show **Build with Parameters**.
- In a Jenkins shell step, an unset variable such as `${Environment}` expands to an **empty string**, with no error.

### Timeline
1. Recreating the platform from scratch (everything had been torn down). Jenkins server and credentials were set up (guide Steps 2–4).
2. Created the `eks-cluster` Pipeline job pointing at `terraform/eks-cluster/Jenkinsfile` and clicked **Build Now**, as Step 5 says: "the first build only registers the parameters (and runs a dev plan)".
3. The build failed in **Init** with `Value must not contain "//"`.
4. Read the `+ terraform ... init` line in the console output: the key was `eks-cluster//terraform.tfstate`, with the environment segment missing.
5. Read the Jenkinsfile: the key is built from `${Environment}`, and every later stage also uses the raw parameters.
6. Unblocked with **Build with Parameters** (`dev` + `apply`), and changed the Jenkinsfile to default missing parameters.

### What happened (symptom)

    + terraform -chdir=terraform/eks-cluster init -input=false -reconfigure
        -backend-config=bucket=go-grpc-micro-tfstate-****
        -backend-config=key=eks-cluster//terraform.tfstate
    Initializing the backend...
    Error: Invalid Value
      on backend.tf line 5, in terraform:
       5:   backend "s3" {
    Value must not contain "//"

The build stopped at the **Init** stage; Validate, Plan, Approve and Apply never ran. (`****` is Jenkins masking the `AWS_ACCOUNT_ID` credential, which is expected.)

### How we got there
**Build Now** on a job that had never run before. The Jenkinsfile's `properties([parameters(...)])` had not been applied yet, so `Environment` and `Terraform_Action` did not exist for this build.

### Why (root cause)
1. A new job's first build has **no parameter values**, because the parameters are defined inside the Jenkinsfile and get registered during that same build.
2. The Init stage builds the state key as `eks-cluster/${Environment}/terraform.tfstate`. With `Environment` empty, the shell produced `eks-cluster//terraform.tfstate`.
3. The S3 backend in current Terraform (1.15 on our Jenkins server) **validates the key** and rejects an empty path segment (`//`). Older Terraform versions accepted it, which is likely why the guide's claim ("the first build runs a dev plan") once looked true: it was quietly planning against a state file at a malformed path.
4. **A second, hidden problem in the same code.** The Approve and Apply stages ran when `params.Terraform_Action != 'plan'`. With the parameter missing, the value is `null`, and `null != 'plan'` is **true**. On a Terraform version that accepted the key, a parameter-less first build would have gone on to *ask to apply* (with an empty environment and a missing `.tfvars` file) instead of stopping after a plan. A human still has to click **Proceed**, but a harmless-looking first run should never get that far.

So the design flaw was **trusting parameters that may not exist**, plus **guarding a dangerous branch with "not the safe value" instead of "is the dangerous value"**.

### Impact
- **Nothing created, changed or destroyed in AWS.** Terraform failed before connecting to the state bucket.
- About 5 minutes lost. No cost.

### Resolution (step by step)
1. **Refresh the `eks-cluster` job page.**
   The failed build still registered the parameters, so the left menu now shows **Build with Parameters** instead of **Build Now**.
2. **Build with Parameters** → `Environment = dev`, `Terraform_Action = apply` → **Build**.
   Init now runs with `key=eks-cluster/dev/terraform.tfstate`.
3. **Open Console Output and wait for the Approve stage.**
   Read the plan summary. For `dev` on a fresh account it should be roughly `Plan: 46 to add, 0 to change, 0 to destroy`. Anything to *change* or *destroy* on a fresh build means something unexpected exists. Stop and investigate.
4. **Click Proceed.** Apply takes 15–20 minutes. The last stage prints the bastion command:
   `aws ssm start-session --region eu-north-1 --target i-0...`
5. **Fix the Jenkinsfile** so this can't recur (diff below):
   - two new variables in the `environment` block: `TF_ENV` (the environment, default `dev`) and `TF_ACTION` (the action, default `plan`);
   - every stage now uses those instead of the raw parameters.
6. **Verify the fix** (next time a fresh job is created): **Build Now** on a new job runs Init → Validate → Plan for `dev`, then **skips** Approve, Apply and Bastion Access, because `TF_ACTION` is `plan`.

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
# Jenkins UI: new job eks-cluster -> Build Now. Console output:
+ terraform -chdir=terraform/eks-cluster init -input=false -reconfigure -backend-config=bucket=go-grpc-micro-tfstate-**** -backend-config=key=eks-cluster//terraform.tfstate
Initializing the backend...
╷
│ Error: Invalid Value
│
│   on backend.tf line 5, in terraform:
│    5:   backend "s3" {
│
│ Value must not contain "//"
╵

laptop$ grep -n "Environment\|Terraform_Action" terraform/eks-cluster/Jenkinsfile
5:            name: 'Environment'
9:            name: 'Terraform_Action'
41:                      -backend-config="key=eks-cluster/${Environment}/terraform.tfstate"
55:                    def destroyFlag = (params.Terraform_Action == 'destroy') ? '-destroy' : ''
58:                          -var-file=envs/\${Environment}.tfvars -out=tfplan
66:                expression { params.Terraform_Action != 'plan' }
69:                input message: "Run terraform ${params.Terraform_Action} on the ${params.Environment} EKS cluster?", ok: 'Proceed'
75:                expression { params.Terraform_Action != 'plan' }
85:                expression { params.Terraform_Action == 'apply' }
```
**Fixing it**
```console
# Jenkins UI: Build with Parameters -> Environment=dev, Terraform_Action=apply -> Approve -> Proceed

# after adding TF_ENV / TF_ACTION defaults to the Jenkinsfile:
laptop$ grep -n "Environment\|Terraform_Action" terraform/eks-cluster/Jenkinsfile
5:            name: 'Environment'
9:            name: 'Terraform_Action'
24:        TF_ENV           = "${params.Environment ?: 'dev'}"
25:        TF_ACTION        = "${params.Terraform_Action ?: 'plan'}"
```

### Code / config change
`terraform/eks-cluster/Jenkinsfile`:

```diff
     environment {
         ...
         AWS_ACCOUNT_ID   = credentials('AWS_ACCOUNT_ID')
+        // The first build of a new job runs before the parameters exist (they are
+        // registered by properties() above), so default to a harmless dev plan.
+        TF_ENV           = "${params.Environment ?: 'dev'}"
+        TF_ACTION        = "${params.Terraform_Action ?: 'plan'}"
     }
 ...
-  -backend-config="key=eks-cluster/${Environment}/terraform.tfstate"
+  -backend-config="key=eks-cluster/${TF_ENV}/terraform.tfstate"
 ...
-  def destroyFlag = (params.Terraform_Action == 'destroy') ? '-destroy' : ''
+  def destroyFlag = (env.TF_ACTION == 'destroy') ? '-destroy' : ''
-  -var-file=envs/\${Environment}.tfvars -out=tfplan
+  -var-file=envs/\${TF_ENV}.tfvars -out=tfplan
 ...
-  expression { params.Terraform_Action != 'plan' }     // Approve, Apply
+  expression { env.TF_ACTION != 'plan' }
-  expression { params.Terraform_Action == 'apply' }    // Bastion Access
+  expression { env.TF_ACTION == 'apply' }
```

`?:` is Groovy's "Elvis" operator: use the left value if it is set, otherwise the right one. The `!= 'plan'` checks are safe now, because `TF_ACTION` can no longer be empty.

Committed as: `fix(ci): default eks-cluster parameters so the first build is a safe dev plan`.

### Lessons learned
- Parameters defined in a Jenkinsfile **do not exist on the job's first build**. Any pipeline that declares its own parameters must handle that build.
- Shell variable expansion fails **silently**: an empty value doesn't error, it just produces a malformed path.
- Guard dangerous actions **positively** ("run only if the action *is* apply/destroy"), never negatively ("run unless it is plan"). A negative check treats missing or unexpected values as dangerous ones.
- Tool upgrades can turn a quietly wrong input into a hard error. Here that was a good thing: it surfaced a latent bug.

### How to approach it next time
1. In Jenkins, open **Console Output** and find the `+ <command>` line just above the error. Jenkins prints every shell command after expanding variables.
2. Look for blanks: `//`, `--flag=` with nothing after it, a trailing `/`, `envs/.tfvars`. Each means a variable was empty.
3. Find where that variable comes from (a parameter, a credential, an `environment {}` entry) and why it was empty on *this* build. First build? Typo in the name? Credential stored under a user instead of System → Global?
4. If it is a parameter: does the job page show **Build with Parameters**? If not, the parameters aren't registered yet.

### Prevention / follow-up
- Done: defaults in the Jenkinsfile, so the first build is a safe `dev` plan, as the guide describes.
- Pattern to reuse: in any pipeline with Jenkinsfile-defined parameters, copy them into `environment {}` with safe defaults, and use only those copies.

### References
- Jenkins: [Pipeline syntax — parameters](https://www.jenkins.io/doc/book/pipeline/syntax/#parameters) ("parameters are only available after the first run").
- Terraform: [S3 backend configuration](https://developer.hashicorp.com/terraform/language/backend/s3) (`key`) and [partial configuration](https://developer.hashicorp.com/terraform/language/backend#partial-configuration).

---

## INC-004 – Local Docker Compose stack recreated by a broken heredoc

| | |
|---|---|
| **Date** | 2026-10-08 |
| **Where** | Laptop, repo root, during a scripted edit of `README.md` |
| **Severity** | Low: local only, no data lost, no files created |
| **Status** | Resolved |

### Summary
A README edit was done with a shell script whose quoting broke. Instead of only editing a file, bash ran several of the README's *example* commands for real. One of them, `docker compose up -d`, recreated the five app containers of the running local stack. The databases and their data were untouched, and the edit was then redone safely.

### Background
- A **heredoc** feeds a block of text to a command:
  ```bash
  python3 - <<'EOF'
  ...text...
  EOF
  ```
  Bash ends the block at the **first line that is exactly the delimiter** (`EOF`). It does not understand nesting: if the text itself contains a line `EOF`, the block ends there.
- Everything after that early end is read by bash as **commands to run**.
- `docker compose up -d` recreates any service whose configuration differs from the running container, here because of a new environment variable. It does **not** rebuild images unless `--build` is passed, and it leaves unchanged services (the databases) running.

### Timeline
1. The local SigNoz instructions in the README were outdated (see INC-003). The new text included a `cat > casting.yaml <<'EOF' ... EOF` example.
2. The README edit was wrapped in `python3 - <<'EOF' ... EOF`, so the README text (with its own `EOF`) was *inside* another `EOF` heredoc.
3. Bash ended the outer heredoc at the inner `EOF`. Python received half a program and printed `SyntaxError`.
4. Bash then ran the rest of the README text line by line:
   - `foundryctl cast -f casting.yaml` → `command not found` (harmless)
   - `cd -` → `OLDPWD not set` (harmless)
   - `OTEL_EXPORTER_OTLP_ENDPOINT=http://host.docker.internal:4317 docker compose up -d` → **ran**, in the repo directory
5. Stopped, and checked the damage read-only before doing anything else.

### What happened (symptom)

    File "<stdin>", line 10
    SyntaxError: unterminated triple-quoted string literal (detected at line 22)
    /bin/bash: line 25: foundryctl: command not found
    /bin/bash: line 26: cd: OLDPWD not set
     Container go-grpc-graphql-micro-catalog_db-1 Running
     Container go-grpc-graphql-micro-account-1 Recreate
     Container go-grpc-graphql-micro-account-1 Recreated
     ...
     Container go-grpc-graphql-micro-frontend-1 Started
    /bin/bash: eval: line 30: unexpected EOF while looking for matching ``'

The `Container ... Recreate` lines are the giveaway: a file edit should never touch containers.

### How we got there
Using a shell heredoc to write text that **itself contained a heredoc with the same delimiter**, instead of using a file editor.

### Why (root cause)
1. The outer heredoc and the inner example both used `EOF` as their delimiter.
2. Bash closed the outer heredoc at the first `EOF` line, which belonged to the inner example.
3. Python got a truncated script and failed. That failure was harmless.
4. Bash executed the remaining lines as commands. One of them was a real, state-changing command (`docker compose up -d`) with an environment variable set.

The underlying mistake: **generating text that contains shell commands, through the shell**. Any quoting slip turns documentation into execution.

### Impact
Checked immediately, read-only:

| Check | Command | Result |
|---|---|---|
| Databases | `docker compose ps` | `account_db`, `order_db`, `catalog_db` still **Up 12 minutes (healthy)**, not restarted. **Data intact.** |
| App containers | `docker compose ps` | `account`, `catalog`, `order`, `graphql`, `frontend` **recreated** ("Up 13 seconds") |
| Which code they run | `docker images` | The existing images (2 weeks old, built before the OpenTelemetry changes). No rebuild happened. |
| Extra setting | `docker compose exec order printenv OTEL_EXPORTER_OTLP_ENDPOINT` | `http://host.docker.internal:4317`. The old images ignore it, so behaviour is unchanged. |
| Files | `ls ../signoz-local`, `ls casting.yaml`, `git status` | Nothing created; only intended changes in git |

Net effect: a few seconds of local app downtime. No data lost, nothing in the cloud touched.

### Resolution (step by step)
1. **Stop and assess before acting.** Ran only read-only checks (table above) to see exactly what had executed.
2. **Leave the stack as it was** rather than run more commands on it: it was healthy and behaving as before.
3. **Redo the README edit as a direct file edit** (in an editor, no shell), and change the README example's inner delimiter from `EOF` to `YAML`.
4. **Return the local stack to a clean state** whenever you next run it:
   ```bash
   docker compose up -d --build
   ```
   This rebuilds the images with the current code (including OpenTelemetry) and recreates the containers without the stray environment variable.

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**How it happened**
```console
laptop$ python3 - <<'EOF'
…README text that itself contained:  cat > casting.yaml <<'EOF' … EOF  …
EOF
  File "<stdin>", line 10
    new="""**Traces, metrics and logs locally (optional).** …
SyntaxError: unterminated triple-quoted string literal (detected at line 22)
/bin/bash: line 25: foundryctl: command not found
/bin/bash: line 26: cd: OLDPWD not set
 Container go-grpc-graphql-micro-catalog_db-1 Running
 Container go-grpc-graphql-micro-order_db-1 Running
 Container go-grpc-graphql-micro-account_db-1 Running
 Container go-grpc-graphql-micro-account-1 Recreate
 Container go-grpc-graphql-micro-catalog-1 Recreate
 …
 Container go-grpc-graphql-micro-graphql-1 Started
 Container go-grpc-graphql-micro-frontend-1 Started
/bin/bash: eval: line 30: unexpected EOF while looking for matching ``'
```
**Assessing the damage (read-only)**
```console
laptop$ ls -d ../signoz-local; ls casting.yaml
ls: cannot access '../signoz-local': No such file or directory
ls: cannot access 'casting.yaml': No such file or directory
laptop$ docker compose ps --format 'table {{.Service}}\t{{.Status}}\t{{.Image}}'
SERVICE      STATUS                    IMAGE
account      Up 13 seconds             go-grpc-graphql-micro-account
account_db   Up 12 minutes (healthy)   go-grpc-graphql-micro-account_db
catalog      Up 13 seconds             go-grpc-graphql-micro-catalog
catalog_db   Up 12 minutes (healthy)   docker.elastic.co/elasticsearch/elasticsearch:6.2.4
frontend     Up 12 seconds             go-grpc-graphql-micro-frontend
graphql      Up 12 seconds             go-grpc-graphql-micro-graphql
order        Up 12 seconds             go-grpc-graphql-micro-order
order_db     Up 12 minutes (healthy)   go-grpc-graphql-micro-order_db
laptop$ docker compose exec -T order printenv OTEL_EXPORTER_OTLP_ENDPOINT
http://host.docker.internal:4317
laptop$ docker images --format '{{.Repository}} {{.CreatedSince}}' | grep go-grpc-graphql-micro
go-grpc-graphql-micro-frontend 2 weeks ago
go-grpc-graphql-micro-order 2 weeks ago
go-grpc-graphql-micro-graphql 2 weeks ago
go-grpc-graphql-micro-catalog 2 weeks ago
go-grpc-graphql-micro-account 2 weeks ago
go-grpc-graphql-micro-order_db 2 weeks ago
go-grpc-graphql-micro-account_db 7 weeks ago
```
**Fixing it**
```console
# No shell: the README edit was redone as a direct file edit; inner heredoc delimiter changed to YAML.
# suggested, not run: rebuild with current code and drop the stray env var
laptop$ docker compose up -d --build
```

### Code / config change
- `README.md`: the local SigNoz example now uses `cat > casting.yaml <<'YAML' ... YAML`.
- No application code changed.

### Lessons learned
- Heredocs don't nest with the same delimiter. Bash stops at the first match, whatever surrounds it.
- Write documentation containing commands with a file editor, never by piping it through a shell.
- After an unexpected error, **investigate read-only first**. Don't "fix" by running more commands until you know what already ran.

### How to approach it next time
1. Output that doesn't belong to what you ran (container logs from a file edit, say) means something else executed. Stop.
2. Find out what ran: scroll the output for commands; check `docker compose ps`, `docker ps -a`, `git status`, and new files with `ls -lt | head`.
3. List what changed vs. what didn't, then decide whether anything needs rolling back.
4. When writing scripts, give each heredoc a unique delimiter (`EOF`, `PYEOF`, `YAML`), or use separate files.

### Prevention / follow-up
- Done: README example uses `YAML` as its delimiter.
- Done: later scripted edits use a distinct outer delimiter (`PYEOF`), or a direct file edit.

### References
- Bash manual: [Here Documents](https://www.gnu.org/software/bash/manual/html_node/Redirections.html#Here-Documents).

---

## INC-003 – SigNoz local Docker Compose instructions out of date

| | |
|---|---|
| **Date** | 2026-10-08 |
| **Where** | `README.md`, section "Traces, metrics and logs locally (optional)" |
| **Severity** | Low: documentation only; anyone following it would have hit a missing folder |
| **Status** | Resolved |

### Summary
The README told readers to start SigNoz locally by cloning its repo and running Docker Compose from `deploy/docker`. That folder no longer exists: SigNoz moved local installs to a new installer called **Foundry**. The section was rewritten with the current official steps.

### Background
- SigNoz can run on Kubernetes (our EKS setup uses its Helm chart) or on a single machine with Docker (for local development).
- For years the Docker method was `deploy/docker/docker-compose.yaml` in the SigNoz repo. SigNoz has since **deprecated** those Compose files and the `install.sh` script in favour of **Foundry** (`foundryctl`). Foundry reads a small `casting.yaml` describing the install and generates and runs the Compose setup for you.

### Timeline
1. Added a "run SigNoz locally" section to the README while adding OpenTelemetry, written from memory: `git clone ... && cd signoz/deploy/docker && docker compose up -d`.
2. Later, preparing to run the app locally, checked those commands against the real repo before recommending them.
3. A fresh clone of `SigNoz/signoz` had no `deploy/docker`. `deploy/` contained only `README.md`, `MIGRATION.md` and `install.sh`.
4. `deploy/README.md` says the Compose manifests are deprecated and points to the Foundry docs.
5. Rewrote the README section from the official Docker install page.

### What happened (symptom)
Following the README would have failed at the `cd`:

    cd: ../signoz/deploy/docker: No such file or directory

From SigNoz's `deploy/README.md`:

> The `install.sh` script and the `docker-compose` manifests have been deprecated. SigNoz now installs and runs through Foundry.

### How we got there
The instructions were written from older knowledge of SigNoz, without checking the project's current repository or docs.

### Why (root cause)
1. Third-party install methods change; SigNoz replaced Compose with Foundry.
2. Our docs copied a command instead of linking to, and checking against, the official source.

### Impact
- Nobody ran it (caught before use). A reader would have lost time on a missing folder.

### Resolution (step by step)
1. Read `deploy/README.md` in the SigNoz repo. It points to <https://signoz.io/docs/install/docker/>.
2. Read the official Docker install page. The current method:
   ```bash
   curl -fsSL https://signoz.io/foundry.sh | bash         # installs foundryctl
   mkdir -p ../signoz-local && cd ../signoz-local
   cat > casting.yaml <<'YAML'
   apiVersion: v1alpha1
   kind: Installation
   metadata:
     name: signoz
   spec:
     deployment:
       flavor: compose
       mode: docker
   YAML
   foundryctl cast -f casting.yaml                        # generates and starts the stack
   ```
   - `flavor: compose` + `mode: docker` = run with Docker Compose on this machine.
   - Requirements: Docker Engine 20.10+, Compose v2, about 4 GB of RAM for Docker.
3. Replaced the README section with these steps, plus how to point our services at it:
   `OTEL_EXPORTER_OTLP_ENDPOINT=http://host.docker.internal:4317 docker compose up -d --build`
4. Added how to stop it: `docker compose down` inside `../signoz-local/pours/deployment` (the folder Foundry generates).
5. Verify when running: UI on `http://localhost:8080`; OTLP accepted on ports `4317` (gRPC) and `4318` (HTTP).

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
laptop$ git clone -q --depth 1 --filter=blob:none --sparse https://github.com/SigNoz/signoz.git sz2
laptop$ cd sz2 && git sparse-checkout set deploy/docker && ls deploy/docker
ls: cannot access 'deploy/docker': No such file or directory
laptop$ git ls-tree -r --name-only HEAD deploy
deploy/MIGRATION.md
deploy/README.md
deploy/install.sh
laptop$ git show HEAD:deploy/README.md | head -40
# Deploy
…
> **Note:** The `install.sh` script and the `docker-compose` manifests have been deprecated.

SigNoz now installs and runs through [Foundry](https://signoz.io/docs/install/docker/).
…
```
**Fixing it**
```console
# No commands: README rewritten with the Foundry steps from https://signoz.io/docs/install/docker/
```

### Code / config change
- `README.md` section rewritten (commit `docs: update SigNoz installation instructions to use Foundry ...`).

### Lessons learned
- Check install commands for third-party tools against their **current** docs or repo before writing them down; they go stale quickly.
- Link the official page next to any copied command, so readers can check for themselves.

### How to approach it next time
1. A path or command from someone else's docs doesn't exist? Open the project's repo at that path and read its README or CHANGELOG for deprecation notes.
2. Check the official install page; prefer it over blog posts and old READMEs.

### Prevention / follow-up
- Done: README links to <https://signoz.io/docs/install/docker/>.
- The EKS install uses the Helm chart, which is still supported (`signoz/signoz` 0.145.0).

### References
- SigNoz: [Install with Docker (Foundry)](https://signoz.io/docs/install/docker/); `deploy/README.md` and `deploy/MIGRATION.md` in [SigNoz/signoz](https://github.com/SigNoz/signoz).

---

## INC-002 – SigNoz would not filter services by environment

| | |
|---|---|
| **Date** | 2026-10-08 |
| **Where** | `OTEL_RESOURCE_ATTRIBUTES` in `k8s/base/microservices/{account,catalog,order}-service.yaml` and `k8s/base/gateway/graphql-gateway.yaml` |
| **Severity** | Medium: a feature would have been silently broken, with no error anywhere |
| **Status** | Resolved (caught in review, before the first deploy) |

### Summary
Our services labelled their telemetry with the environment (`dev`/`prod`) using the **new** OpenTelemetry attribute name, `deployment.environment.name`. SigNoz still reads the **old** name, `deployment.environment`. After deploying, SigNoz would have shown every service under environment "default", and its environment filter would have been empty. We switched to the name SigNoz actually reads.

### Background
- Every span, metric and log a service sends carries **resource attributes**: facts about where it came from (`service.name`, the pod name, the environment...). We set them per pod with the standard `OTEL_RESOURCE_ATTRIBUTES` environment variable, filled from the Kubernetes downward API and the `otel-config` ConfigMap:
  `deployment.environment=$(DEPLOYMENT_ENVIRONMENT),k8s.namespace.name=$(K8S_NAMESPACE),...`
- OpenTelemetry's **semantic conventions** define standard attribute names. Newer versions renamed `deployment.environment` to `deployment.environment.name`, and the old name is marked deprecated.
- Backends (SigNoz, Grafana, etc.) don't all follow renames at the same pace. The backend's UI only filters on the names it is coded to read.

### Timeline
1. While adding OpenTelemetry, set `deployment.environment.name=...` in the four Deployments, following the latest OTel spec.
2. Reviewed the whole setup against the OpenTelemetry and SigNoz docs.
3. SigNoz's docs and search results describe filtering Services by `deployment.environment`, which raised a doubt.
4. Checked SigNoz's source: every quick-filter definition uses `deployment.environment`.
5. Checked the SigNoz collector config our Helm chart renders: its span-metrics processor groups by `deployment.environment`, with default `default`.
6. Changed the attribute name in all four manifests.

### What happened (symptom)
No error; this would have been invisible until someone used the filter. Expected behaviour after deploy, had it not been fixed:
- **Services** page → environment dropdown: only `default`, or empty.
- Request-rate and latency charts per environment: everything grouped under `default`.
- The data is stored (under `deployment.environment.name`), but SigNoz's built-in views ignore it.

### How we got there
We followed the newest OpenTelemetry naming without checking which name the backend reads.

### Why (root cause)
1. The OpenTelemetry spec renamed the attribute to `deployment.environment.name`.
2. SigNoz (chart 0.145.0) still reads only `deployment.environment`:
   - quick filters: `pkg/types/quickfiltertypes/filter.go` → `{Name: "deployment.environment", ...}`;
   - the collector's `signozspanmetrics/delta` processor has dimension `name: deployment.environment, default: default`. The RED metrics (rate, errors, duration) behind the Services page are grouped by it.
3. So an attribute that is valid according to the spec was invisible to the backend's UI.

### Impact
- None in practice: caught before any deploy.
- Had it shipped, `dev` and `prod` telemetry would have been indistinguishable in SigNoz's Services view.

### Resolution (step by step)
1. **Confirm which name SigNoz reads.** Searched its source:
   `grep -rn '"deployment.environment' pkg/` → only `deployment.environment` in filters and preferences.
2. **Confirm in our own deployment config.** Rendered the chart with our values:
   `helm template signoz signoz/signoz --version 0.145.0 -f k8s/monitoring/signoz-values.yaml` → `signozspanmetrics/delta` dimensions include `deployment.environment`.
3. **Change the four manifests** from `deployment.environment.name=$(DEPLOYMENT_ENVIRONMENT)` to `deployment.environment=$(DEPLOYMENT_ENVIRONMENT)`, with a comment explaining why the old name is used.
4. **Verify the rendered manifests:**
   `kubectl kustomize k8s/overlays/prod | grep deployment.environment` →
   `deployment.environment=$(DEPLOYMENT_ENVIRONMENT),k8s.namespace.name=...`
   and the `otel-config` ConfigMap shows `DEPLOYMENT_ENVIRONMENT: prod` (dev overlay: `dev`).
5. **Verify after deploy:** SigNoz → **Services** → the environment dropdown lists `dev`.

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
laptop$ git clone -q --depth 1 --filter=blob:none --sparse https://github.com/SigNoz/signoz.git sz
laptop$ cd sz && git sparse-checkout set pkg ee frontend/src/container/ResourceAttributesFilter frontend/src/hooks
laptop$ grep -rn 'deployment\.environment\(\.name\)\?"' --include=*.go --include=*.ts --include=*.tsx . | grep -v _test
pkg/types/quickfiltertypes/filter.go:177:		{Name: "deployment.environment", FieldContext: telemetrytypes.FieldContextResource, FieldDataType: telemetrytypes.FieldDataTypeString},
pkg/types/quickfiltertypes/filter.go:192:		{Name: "deployment.environment", FieldContext: telemetrytypes.FieldContextResource, FieldDataType: telemetrytypes.FieldDataTypeString},
…
pkg/types/preferencetypes/preference.go:188:			DefaultValue:  MustNewValue([]any{"deployment.environment"}, ValueTypeArray),
pkg/sqlmigration/031_update_quick_filters.go:51:			{"key": "deployment.environment", "dataType": "string", "type": "resource"},
…

laptop$ helm template signoz signoz/signoz --version 0.145.0 -n signoz -f k8s/monitoring/signoz-values.yaml > r2.yaml
# otel-collector config in r2.yaml (excerpt):
      signozspanmetrics/delta:
        aggregation_temporality: AGGREGATION_TEMPORALITY_DELTA
        dimensions:
        - default: default
          name: service.namespace
        - default: default
          name: deployment.environment
        - name: signoz.collector.id
```
**Fixing it**
```console
laptop$ sed -i 's|value: "deployment.environment.name=\$(DEPLOYMENT_ENVIRONMENT),|value: "deployment.environment=$(DEPLOYMENT_ENVIRONMENT),|' k8s/base/microservices/*.yaml k8s/base/gateway/graphql-gateway.yaml
laptop$ grep -c 'value: "deployment.environment=' k8s/base/microservices/*.yaml k8s/base/gateway/graphql-gateway.yaml
k8s/base/microservices/account-service.yaml:1
k8s/base/microservices/order-service.yaml:1
k8s/base/microservices/catalog-service.yaml:1
k8s/base/gateway/graphql-gateway.yaml:1
laptop$ kubectl kustomize k8s/overlays/prod | grep -m1 "deployment.environment"
          value: deployment.environment=$(DEPLOYMENT_ENVIRONMENT),k8s.namespace.name=$(K8S_NAMESPACE),k8s.pod.name=$(K8S_POD_NAME),k8s.node.name=$(K8S_NODE_NAME)
```

### Code / config change
In each of the four Deployments:

```diff
-            # Tag every span, metric and log with where it came from (shown in SigNoz)
+            # Tag every span, metric and log with where it came from. SigNoz filters and
+            # groups by deployment.environment (not the newer deployment.environment.name)
             - name: OTEL_RESOURCE_ATTRIBUTES
-              value: "deployment.environment.name=$(DEPLOYMENT_ENVIRONMENT),k8s.namespace.name=..."
+              value: "deployment.environment=$(DEPLOYMENT_ENVIRONMENT),k8s.namespace.name=..."
```

Part of commit `feat(k8s): wire services to SigNoz and Prometheus`.

### Lessons learned
- "Correct per the spec" and "works in the backend" are different questions. Check the attribute names the backend actually reads.
- Silent failures (nothing errors, a feature just doesn't work) are the most expensive kind. Review for them deliberately.

### How to approach it next time
1. Data arrives in SigNoz but a filter or grouping is empty? Open a trace and look at its **resource attributes**: which names did we actually send?
2. Compare those names with what the UI filters on: the SigNoz docs page for that view, or its source (`quickfiltertypes`).
3. Check the collector config (`kubectl -n signoz get cm -o yaml | grep -A5 dimensions`) for the attributes it groups metrics by.

### Prevention / follow-up
- When adding any resource attribute, check it against the backend's docs or source, not just the OTel spec.
- Revisit when SigNoz adopts `deployment.environment.name`. Sending both names is a cheap option then.

### References
- OpenTelemetry semantic conventions: [deployment attributes](https://opentelemetry.io/docs/specs/semconv/resource/deployment-environment/).
- SigNoz: [View Services](https://signoz.io/docs/userguide%2Fmetrics/); source `pkg/types/quickfiltertypes/filter.go` in [SigNoz/signoz](https://github.com/SigNoz/signoz).

---

## INC-001 – `kubectl` dry-run fails with "no such host"

| | |
|---|---|
| **Date** | 2026-10-08 |
| **Where** | Laptop, validating the rendered Kubernetes manifests with `kubectl apply --dry-run=client` |
| **Severity** | Low: a validation command failed; nothing was affected |
| **Status** | Explained (expected behaviour; validated another way) |

### Summary
To check the new Kubernetes manifests, we ran `kubectl apply --dry-run=client` on the laptop. It failed because kubectl tried to contact the EKS cluster in the laptop's kubeconfig, and that cluster had been deleted. This was expected: even when it exists, our cluster's API is private and only reachable from the bastion. We validated the manifests offline instead.

### Background
- **kubeconfig** (`~/.kube/config`) tells `kubectl` which cluster to talk to (its *current context*).
- `--dry-run=client` doesn't change anything, but by default kubectl still **downloads the cluster's OpenAPI schema** to validate fields. So it needs to reach the API server.
- Our EKS API endpoint is **private-only**: reachable only from inside the VPC (the bastion), never from the internet or the laptop.

### Timeline
1. Rendered both overlays: `kubectl kustomize k8s/overlays/{dev,prod}` succeeded.
2. Ran `kubectl apply --dry-run=client -f prod.yaml` for a stricter check.
3. It failed with `no such host` for an `*.eks.amazonaws.com` name.
4. Recognised the hostname as the old, torn-down cluster still in the laptop's kubeconfig.
5. Relied on the offline render (and later the Helm template render) for validation.

### What happened (symptom)

    error: error validating "prod.yaml": error validating data: failed to download openapi:
    Get "https://<CLUSTER_ID>.gr7.us-east-1.eks.amazonaws.com/openapi/v2?timeout=32s":
    dial tcp: lookup <CLUSTER_ID>.gr7.us-east-1.eks.amazonaws.com on 127.0.0.53:53: no such host;
    if you choose to ignore these errors, turn validation off with --validate=false

### How we got there
The laptop's kubeconfig still pointed at a cluster from an earlier session that had since been destroyed.

### Why (root cause)
1. `kubectl apply --dry-run=client` validates against the live cluster's schema.
2. The current context named a cluster that no longer exists, so its DNS name doesn't resolve (`no such host`).
3. Even with a live cluster, the private endpoint would be unreachable from the laptop (it would time out instead).

### Impact
- None. The manifests were validated another way.

### Resolution (step by step)
1. **Validate offline:**
   `kubectl kustomize k8s/overlays/dev > /dev/null && echo OK`
   This renders everything (catching YAML and Kustomize errors) without contacting a cluster.
2. **Or skip schema download:** `kubectl apply --dry-run=client --validate=false -f <file>`.
3. **Run real `kubectl` commands from the bastion**, the only machine with access to the private API.
4. **Optional cleanup** of the stale context:
   ```bash
   kubectl config get-contexts                 # find the old cluster's context
   kubectl config delete-context <old-context>
   ```

### Terminal commands
Commands and their output as they appeared in the session. The prompt shows where each ran (`laptop$`, `bastion$`; UI steps are `#` comments). `# suggested, not run` marks recommended steps that weren't executed. Long output is trimmed with `…` and some multi-line output (JSON, Terraform plans) is condensed; content is as printed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
laptop$ for e in dev prod; do kubectl kustomize k8s/overlays/$e > $e.yaml || exit 1; done
laptop$ grep -n "otel-config" prod.yaml
45:  name: otel-config-kbh6ht7bkc
257:            name: otel-config-kbh6ht7bkc
325:            name: otel-config-kbh6ht7bkc
441:            name: otel-config-kbh6ht7bkc
530:            name: otel-config-kbh6ht7bkc
laptop$ kubectl apply --dry-run=client -f prod.yaml
error: error validating "prod.yaml": error validating data: failed to download openapi: Get "https://<CLUSTER_ID>.gr7.us-east-1.eks.amazonaws.com/openapi/v2?timeout=32s": dial tcp: lookup <CLUSTER_ID>.gr7.us-east-1.eks.amazonaws.com on 127.0.0.53:53: no such host; if you choose to ignore these errors, turn validation off with --validate=false
```
**Working around it (validate offline)**
```console
laptop$ for e in dev prod; do grep -B1 -A8 "^data:" $e.yaml | grep -E "DEPLOYMENT_ENV|OTLP_ENDPOINT"; done
  DEPLOYMENT_ENVIRONMENT: dev
  OTEL_EXPORTER_OTLP_ENDPOINT: http://signoz-otel-collector.signoz.svc.cluster.local:4317
  DEPLOYMENT_ENVIRONMENT: prod
  OTEL_EXPORTER_OTLP_ENDPOINT: http://signoz-otel-collector.signoz.svc.cluster.local:4317

# suggested, not run: remove the stale context of the deleted cluster
laptop$ kubectl config get-contexts
laptop$ kubectl config delete-context <old-context>
```

### Code / config change
None.

### Lessons learned
- Client-side dry-run is not fully offline; it still needs the API server for schema validation.
- With a private EKS endpoint, the laptop is never the place to run `kubectl` against the cluster.

### How to approach it next time
- `no such host` for `*.eks.amazonaws.com`: the cluster is gone, or the context is stale.
- A **timeout** (instead of `no such host`): the cluster exists but its endpoint is private; use the bastion.
- `kubectl config current-context` shows which cluster kubectl is aiming at.

### Prevention / follow-up
- None needed. Validate manifests with `kubectl kustomize`; run cluster commands on the bastion.

### References
- Kubernetes: [`kubectl apply` — dry-run and validate](https://kubernetes.io/docs/reference/kubectl/generated/kubectl_apply/).
- AWS: [EKS cluster endpoint access control](https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html).

---

> **INC-024 to INC-044 were back-filled from the first build (2026-10-03 → 2026-10-04)**, written up afterwards from the session log. They are older than INC-001, so they sit below it. Their **Terminal commands** sections list the commands actually run to hit and to fix each error, with key output, taken from the session log and redacted.

---

## INC-044 – Teardown: deleting the Argo CD Application would not have deleted the app

| | |
|---|---|
| **Date** | 2026-10-04 |
| **Where** | Teardown (PROCESS.md §23, Recreate guide Step 17); bastion; Argo CD `Application/go-micro-shop-dev`; Jenkins `eks-cluster` destroy |
| **Severity** | Medium: following the documented teardown would have left 4 load balancers and 5 EBS volumes behind, and `terraform destroy` would have hung on the VPC |
| **Status** | Resolved (teardown done in a safe order; documentation caveat added here) |

### Summary
The teardown plan said "delete the Argo CD Application (prune removes the ALB and EBS volumes)". A pre-check showed the Application had **no deletion finalizer**, which means deleting it only removes the Application object and leaves everything it deployed running. Separately, the Jenkins server had already been destroyed, so the documented `eks-cluster` → `destroy` pipeline could not run. The teardown was done in a safe order instead: Application first (so it can't recreate anything), then the namespaces themselves, then a check in AWS that every load balancer and volume was gone, then `terraform destroy` from the laptop. It finished with `Destroy complete! Resources: 54 destroyed.`

### Background
- **Argo CD cascade delete.** An Argo CD `Application` only deletes the resources it manages if it carries the finalizer `resources-finalizer.argocd.argoproj.io`. The UI adds it when you delete with "cascade"; an Application created from YAML without it is deleted *non-cascading*: the Deployments, Services, Ingress and PVCs stay.
- **Controller-created AWS resources.** The ALB (from the Ingress), the Classic Load Balancers (from `type: LoadBalancer` Services) and the EBS volumes (from PVCs) are created by in-cluster controllers. Terraform doesn't know about them.
- **Why that matters for `terraform destroy`.** Those load balancers own network interfaces and security groups inside the VPC. While they exist, AWS refuses to delete the subnets and the VPC, so the destroy hangs and eventually fails, and the leftovers keep billing.
- **Deletion order.** Delete the Kubernetes objects while their controllers (Load Balancer Controller, EBS CSI driver) are still running, so the controllers delete the AWS side. Only then destroy the cluster.

### Timeline
1. Cluster cleanup started from the bastion over SSM, read-only first.
2. The pre-check showed `go-micro-shop-dev` with `FINALIZERS <none>`, 3 `LoadBalancer` Services, 1 Ingress (ALB) and 5 PVCs.
3. Order chosen: delete the Application → `helm uninstall` monitoring → delete the Argo CD LB Service → delete the namespaces `go-micro-shop` and `monitoring` (`--wait`).
4. In-cluster check: no LoadBalancer Services, Ingresses, PVCs or PVs left.
5. AWS check, polled until zero: `ALB/NLB=0 CLB=0 PVC-volumes=0`.
6. The Jenkins server turned out to be already destroyed, so the EKS environment was destroyed from the laptop with the same S3 state.
7. `Destroy complete! Resources: 54 destroyed.` A final inventory showed nothing left.


### What happened (symptom)
Nothing failed, because the problem was caught by a pre-check. The would-be symptom was:

    kubectl -n argocd delete application go-micro-shop-dev
    # → application deleted, but Deployments, Ingress (ALB), 3 CLB Services and 5 PVCs/EBS volumes keep running
    terraform destroy
    # → hangs on aws_subnet / aws_vpc: "DependencyViolation ... has dependencies and cannot be deleted"

### How we got there
`k8s/scripts/create-argocd-app.sh` creates the Application with `kubectl apply`, and its YAML has no `metadata.finalizers`. The teardown notes assumed `prune: true` covers deletion. It doesn't: prune removes resources that disappear from **git** during a sync, not resources of a deleted Application.

### Why (root cause)
1. The Application was created from YAML without `resources-finalizer.argocd.argoproj.io`.
2. Without the finalizer, deleting it is non-cascading (checked against the Argo CD docs, "App Deletion").
3. The app's Ingress, LoadBalancer Services and PVCs would therefore survive, and so would their AWS load balancers and EBS volumes.
4. Those load balancers hold ENIs and security groups in the VPC, so `terraform destroy` would fail on the subnets and VPC.

### Impact
- None in practice (caught before acting).
- Avoided: about 4 load balancers (~$18/month each) and 5 EBS volumes left running, plus a stuck destroy.

### Resolution (step by step)
1. **Inventory first** (bastion, read-only):
   ```bash
   kubectl get applications -n argocd -o custom-columns=NAME:.metadata.name,FINALIZERS:.metadata.finalizers
   kubectl get svc -A --field-selector spec.type=LoadBalancer
   kubectl get ingress -A; kubectl get pvc -A; helm list -A
   ```
   `FINALIZERS <none>` means deleting the Application will **not** cascade.
2. **Delete the Application first**, so self-heal can't recreate what you delete next:
   `kubectl -n argocd delete application go-micro-shop-dev`
3. **Remove everything that created AWS resources:**
   ```bash
   helm uninstall grafana prometheus -n monitoring
   kubectl -n argocd delete svc argocd-server
   kubectl delete namespace go-micro-shop monitoring --wait=true --timeout=10m
   ```
   Deleting the namespaces removes the Ingress, Services and PVCs; the LB Controller and the EBS CSI driver then delete the ALB/CLBs and volumes.
4. **Verify in AWS until zero** (laptop):
   ```bash
   aws elbv2 describe-load-balancers --query 'length(LoadBalancers)'
   aws elb   describe-load-balancers --query 'length(LoadBalancerDescriptions)'
   aws ec2 describe-volumes --filters Name=tag-key,Values=kubernetes.io/created-for/pvc/name --query 'length(Volumes)'
   ```
   Expect `0 / 0 / 0`.
5. **Destroy the environment.** From Jenkins (`eks-cluster`, `dev` + `destroy`), or, since Jenkins was gone, from the laptop:
   ```bash
   cd terraform/eks-cluster
   terraform init -reconfigure -backend-config=backend.hcl -backend-config="key=eks-cluster/dev/terraform.tfstate"
   terraform destroy -var-file=envs/dev.tfvars
   ```
   Expect `Destroy complete! Resources: 54 destroyed.`

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it** (pre-teardown inventory, bastion over SSM)
```console
# laptop → run on the bastion through SSM (read-only)
$ aws ssm send-command --region eu-north-1 --instance-ids <BASTION_INSTANCE_ID> \
    --document-name AWS-RunShellScript --parameters 'commands=[
      "export KUBECONFIG=/etc/eks/kubeconfig",
      "kubectl get applications -n argocd -o custom-columns=NAME:.metadata.name,FINALIZERS:.metadata.finalizers",
      "kubectl get svc -A --field-selector spec.type=LoadBalancer",
      "kubectl get ingress -A",
      "kubectl get pvc -A",
      "helm list -A"]'
NAME                FINALIZERS
go-micro-shop-dev   <none>                      # ← non-cascading delete
NAMESPACE    NAME                TYPE           EXTERNAL-IP
argocd       argocd-server       LoadBalancer   <hash>-<id>.eu-north-1.elb.amazonaws.com
monitoring   grafana             LoadBalancer   <hash>-<id>.eu-north-1.elb.amazonaws.com
monitoring   prometheus-server   LoadBalancer   <hash>-<id>.eu-north-1.elb.amazonaws.com
NAMESPACE       NAME            CLASS   ADDRESS
go-micro-shop   go-micro-shop   alb     <APP_LB_HOST>
NAMESPACE       NAME                                STATUS   CAPACITY   STORAGECLASS
go-micro-shop   account-db-pvc                      Bound    5Gi        ebs-csi-default-sc
go-micro-shop   catalog-db-pvc                      Bound    10Gi       ebs-csi-default-sc
go-micro-shop   order-db-pvc                        Bound    5Gi        ebs-csi-default-sc
monitoring      prometheus-server                   Bound    8Gi        ebs-csi-default-sc
monitoring      storage-prometheus-alertmanager-0   Bound    2Gi        ebs-csi-default-sc

$ aws elbv2 describe-load-balancers --region eu-north-1 --query 'LoadBalancers[].[LoadBalancerName,Type]' --output text
k8s-gomicros-gomicros-<hash>   application
$ aws elb describe-load-balancers --region eu-north-1 --query 'LoadBalancerDescriptions[].LoadBalancerName' --output text
<hash>   <hash>   <hash>
```

**Fixing it / trying to fix it** (safe teardown order)
```console
# bastion (through SSM): app first, then everything that created AWS resources
$ kubectl -n argocd delete application go-micro-shop-dev --wait=true
application.argoproj.io "go-micro-shop-dev" deleted from argocd namespace
$ helm uninstall grafana prometheus -n monitoring
release "grafana" uninstalled
release "prometheus" uninstalled
$ kubectl -n argocd delete svc argocd-server
service "argocd-server" deleted from argocd namespace
$ kubectl delete namespace go-micro-shop monitoring --wait=true --timeout=10m
namespace "go-micro-shop" deleted
namespace "monitoring" deleted
$ kubectl get svc -A --field-selector spec.type=LoadBalancer; kubectl get ingress -A; kubectl get pvc -A; kubectl get pv
No resources found
No resources found
No resources found
No resources found

# laptop: wait until AWS really deleted the load balancers and volumes
$ for i in $(seq 12); do
    a=$(aws elbv2 describe-load-balancers --region eu-north-1 --query 'length(LoadBalancers)' --output text)
    c=$(aws elb describe-load-balancers --region eu-north-1 --query 'length(LoadBalancerDescriptions)' --output text)
    v=$(aws ec2 describe-volumes --region eu-north-1 --filters Name=tag-key,Values=kubernetes.io/created-for/pvc/name --query 'length(Volumes)' --output text)
    echo "ALB/NLB=$a  CLB=$c  PVC-volumes=$v"; [ "$a$c$v" = "000" ] && break; sleep 15
  done
ALB/NLB=0  CLB=0  PVC-volumes=0

# the Jenkins server was already destroyed, so destroy EKS from the laptop with the same S3 state
$ aws ec2 describe-instances --region eu-north-1 --filters Name=instance-state-name,Values=running \
    --query 'Reservations[].Instances[].Tags[?Key==`Name`]|[].Value' --output text
go-microservices-dev-node  go-microservices-dev-bastion  go-microservices-dev-node     # no jenkins-server
$ cd terraform/eks-cluster
$ terraform init -reconfigure -input=false -backend-config=backend.hcl \
    -backend-config="key=eks-cluster/dev/terraform.tfstate" >/dev/null
(no output: redirected)
$ terraform destroy -input=false -auto-approve -var-file=envs/dev.tfvars
Plan: 0 to add, 0 to change, 54 to destroy.
...
Destroy complete! Resources: 54 destroyed.

# leftovers outside Terraform
$ aws iam detach-role-policy --role-name go-microservices-dev-lb-controller-role \
    --policy-arn arn:aws:iam::<ACCOUNT_ID>:policy/AWSLoadBalancerControllerIAMPolicy \
  && aws iam delete-role --role-name go-microservices-dev-lb-controller-role && echo "deleted role"
deleted role
$ aws iam delete-policy --policy-arn arn:aws:iam::<ACCOUNT_ID>:policy/AWSLoadBalancerControllerIAMPolicy && echo "deleted policy"
deleted policy
$ for r in account catalog order graphql frontend; do
    aws ecr delete-repository --region eu-north-1 --repository-name $r --force --query 'repository.repositoryName' --output text
  done
account
catalog
order
graphql
frontend
$ aws ecr describe-repositories --region eu-north-1 --query 'repositories[].repositoryName' --output text
(no output: no repositories left)
# state bucket is versioned: delete every object version and delete marker, then the bucket
$ aws s3api list-object-versions --bucket go-grpc-micro-tfstate-<ACCOUNT_ID> ...   # → delete-objects in batches
deleted 65 object versions
$ aws s3api delete-bucket --bucket go-grpc-micro-tfstate-<ACCOUNT_ID> --region eu-north-1 && echo "bucket deleted"
bucket deleted
$ aws s3 ls | grep -c go-grpc
0
$ rm -f terraform/eks-cluster/backend.hcl terraform/jenkins-server/backend.hcl terraform/jenkins-server/terraform.tfvars
(no output)
$ rm -rf terraform/jenkins-server/.terraform && rmdir platform
(no output)
$ git status --short --ignored | grep -v node_modules
!! .env
!! frontend/dist/
!! frontend/tsconfig.tsbuildinfo

# final check: nothing left
$ aws ec2 describe-vpcs --region eu-north-1 --filters Name=is-default,Values=false --query 'Vpcs[].VpcId' --output text
(no output: no VPCs)
$ aws ec2 describe-addresses --region eu-north-1 --query 'Addresses[].PublicIp' --output text
(no output)
$ aws secretsmanager list-secrets --region eu-north-1 --include-planned-deletion --query 'SecretList[].Name' --output text
(no output)
$ aws iam list-roles --query 'Roles[?contains(RoleName,`go-microservices`)||contains(RoleName,`jenkins-server`)].RoleName' --output text
(no output)
```

### Code / config change
None at the time. Options for later: add the finalizer in `create-argocd-app.sh`, or keep the explicit namespace deletion in the teardown steps.

### Lessons learned
- `prune` and cascade-delete are different things in Argo CD.
- Anything a controller creates in AWS must be removed while the controller still runs, before Terraform touches the network.
- Check what still exists before deleting, and verify in AWS (not just in Kubernetes) afterwards.

### How to approach it next time
1. `kubectl get application -o yaml | grep -A2 finalizers`: no finalizer → non-cascading delete.
2. List `LoadBalancer` Services, Ingresses and PVCs: those are your AWS leftovers.
3. If `terraform destroy` hangs on a subnet or VPC: `aws ec2 describe-network-interfaces --filters Name=vpc-id,Values=<vpc>` shows who still holds ENIs (usually an ELB).

### Prevention / follow-up
- Add `finalizers: [resources-finalizer.argocd.argoproj.io]` to the Application in `create-argocd-app.sh`, or keep the namespace-delete step in the teardown docs.
- Destroy the EKS environment **before** the Jenkins server, since Jenkins runs the destroy pipeline.

### References
- Argo CD: [App deletion and cascading](https://argo-cd.readthedocs.io/en/stable/user-guide/app_deletion/).
- AWS: [Troubleshoot VPC deletion: dependencies](https://repost.aws/knowledge-center/troubleshoot-dependency-error-delete-vpc).

---

## INC-043 – `docs/architecture.drawio` looks empty on GitHub

| | |
|---|---|
| **Date** | 2026-10-04 |
| **Where** | GitHub file view of `docs/architecture.drawio`; README link "Editable source" |
| **Severity** | Low: docs only; the diagram file itself was valid |
| **Status** | Resolved |

### Summary
The README linked to the editable diagram, but opening `docs/architecture.drawio` on GitHub showed what looked like an empty file. The file was valid (122 diagram cells), but it had been generated as **one 94 KB line of XML**, which GitHub's viewer doesn't display usefully, and GitHub can't render draw.io diagrams anyway. The file was re-formatted as indented XML (421 lines, verified to render pixel-identically), and the README link now opens the diagram directly in draw.io.

### Background
- A `.drawio` file is XML. GitHub shows it as plain text, with no diagram preview.
- GitHub's code viewer handles very long single lines poorly; a 94 KB one-liner appears blank or truncated.
- draw.io can open a file straight from a URL: `https://app.diagrams.net/#U<url-encoded raw file URL>`.

### Timeline
1. Diagram generated by a script and exported to PNG with the draw.io CLI.
2. The README showed the PNG plus a link to the `.drawio` source.
3. The link "had nothing in it".
4. Checked: the file was on GitHub (`raw.githubusercontent.com` → 200, 94,681 bytes), valid XML, 122 cells, but all on one line.
5. Pretty-printed it, re-exported, compared the images (identical), and changed the README link.

### What happened (symptom)
The GitHub page for `docs/architecture.drawio` showed no readable content.

### How we got there
The generator wrote the whole `<mxfile>` with no newlines.

### Why (root cause)
1. Single-line XML → GitHub's viewer can't show it sensibly.
2. GitHub has no draw.io renderer → even well-formatted XML isn't a diagram there.
3. So the link pointed somewhere that couldn't show the diagram.

### Impact
Docs only. The PNG in the README was fine.

### Resolution (step by step)
1. Confirm the file is intact:
   `python3 -c "import xml.dom.minidom as m; print(len(m.parse('docs/architecture.drawio').getElementsByTagName('mxCell')))"` → `122`.
2. Pretty-print it with `xml.dom.minidom ... toprettyxml(indent="  ")` (dropping the `<?xml?>` header and blank lines): 421 lines.
3. Prove nothing changed: export the new file to PNG and compare with the old export (`ImageChops.difference(...).getbbox() is None` → identical).
4. README link → `https://app.diagrams.net/#Uhttps%3A%2F%2Fraw.githubusercontent.com%2F<user>%2Fgo-grpc-graphql-micro%2Fmain%2Fdocs%2Farchitecture.drawio`.

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
# laptop: generate the diagram XML (one line) and export it
$ python3 gen_drawio.py docs/architecture.drawio
wrote docs/architecture.drawio 122 cells
$ ./drawio.AppImage --no-sandbox -x -f png -s 1.5 -b 20 -o docs/architecture.png docs/architecture.drawio
docs/architecture.drawio -> docs/architecture.png
$ git push      # README links to docs/architecture.drawio → GitHub shows nothing useful
(push output not captured)
```

**Fixing it / trying to fix it**
```console
$ git fetch -q; git ls-tree origin/main docs/ | grep -i arch
100644 blob <sha>  docs/architecture.drawio
100644 blob <sha>  docs/architecture.png
$ curl -s -o /dev/null -w "raw on github: %{http_code} %{size_download} bytes\n" \
    https://raw.githubusercontent.com/<user>/go-grpc-graphql-micro/main/docs/architecture.drawio
raw on github: 200 94681 bytes
$ head -c 120 docs/architecture.drawio
<mxfile host="drawio" agent="generator"><diagram id="arch" name="Architecture"><mxGraphModel dx="1980" ...
$ python3 -c "import xml.dom.minidom as m; d=m.parse('docs/architecture.drawio'); print('valid xml, cells:', len(d.getElementsByTagName('mxCell')))"
valid xml, cells: 122

# pretty-print (same content, 421 lines)
$ python3 - <<'EOF'
import xml.dom.minidom as m
p='docs/architecture.drawio'
x=m.parse(p).toprettyxml(indent="  ").replace('<?xml version="1.0" ?>\n','')
open(p,'w').write("\n".join(l for l in x.splitlines() if l.strip())+"\n")
EOF
(no output)
$ wc -l docs/architecture.drawio
421 docs/architecture.drawio

# prove the diagram is unchanged
$ ./drawio.AppImage --no-sandbox -x -f png -s 1.5 -b 20 -o /tmp/check.png docs/architecture.drawio
docs/architecture.drawio -> /tmp/check.png
$ python3 -c "from PIL import Image, ImageChops; a=Image.open('/tmp/check.png').convert('RGB'); b=Image.open('docs/architecture.png').convert('RGB'); print('identical render:', ImageChops.difference(a,b).getbbox() is None)"
identical render: True

# README: link changed to open the file directly in draw.io:
#   https://app.diagrams.net/#Uhttps%3A%2F%2Fraw.githubusercontent.com%2F<user>%2Fgo-grpc-graphql-micro%2Fmain%2Fdocs%2Farchitecture.drawio
```

### Code / config change
`docs/architecture.drawio` (formatting only) and the README link.

### Lessons learned
Generated text files that people might open should be formatted for humans.

### How to approach it next time
"File looks empty on GitHub": check the raw URL and `wc -l`; a 1-line file of tens of KB is the usual reason.

### Prevention / follow-up
The generator should pretty-print before writing.

### References
- draw.io: [Open a diagram from a URL](https://www.drawio.com/doc/faq/load-file-from-url).

---

## INC-042 – Commit messages contain the pasted command (`git commit -m "…"`)

| | |
|---|---|
| **Date** | 2026-10-04 |
| **Where** | Git history on `main` (commits `ff98439`, `755203b`, `4639278`, `8d9c594`, `c0d643e`) |
| **Severity** | Low: cosmetic, but visible on the repo's front page next to the file list |
| **Status** | Workaround (process fixed; the pushed messages stay unless history is rewritten) |

### Summary
Several commits ended up with messages like `git commit -m "docs: add draw.io architecture diagram"` or `"docs: …"` with literal quotes. The full command (or the quoted message) had been pasted into a tool that **already** wraps the text as the commit message (an editor or a Git GUI commit box), so the command itself became the message.

### Background
- `git commit -m "<msg>"` in a terminal: the shell strips the quotes and Git stores `<msg>`.
- In `git commit` (editor), VS Code's Source Control box, or a GUI, **everything you type is the message**, quotes and all.
- GitHub shows the latest commit message above the file list, and recruiters do look at it.


### Timeline
1. The suggested commands were given as terminal commands (`git commit -m "…" -m "…"`).
2. They were pasted into a commit-message box instead of a terminal.
3. `git log --oneline` showed the pasted text in four messages.

### What happened (symptom)

    8d9c594 git commit -m "docs: add draw.io architecture diagram"
    4639278 "docs: add project documentation, screenshots and portfolio README"
    755203b git commit -m "docs: add project documentation and portfolio README"
    ff98439 git commit -m "docs: add project documentation and README"

### Why (root cause)
Terminal command syntax was pasted into a message field. Nothing validates commit message content.

### Impact
Cosmetic only; no code affected.

### Resolution (step by step)
1. **For the latest unpushed commit:** `git commit --amend -m "docs: …"`.
2. **For several pushed commits** (only your own, nobody else building on them):
   ```bash
   git reset --soft <last-good-commit>     # keeps all files staged
   git commit -m "docs: …" -m "<body>"
   git push --force-with-lease origin main
   ```
   Watch out: on this repo the pipelines also push `chore(deploy)` commits to `main`, so only squash a range that contains none of them.
3. **Going forward:** run commit commands in the terminal, or type just the message text in a commit box.

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
# the commit text was pasted into a commit-message box (not a terminal), so the
# whole "git commit -m ..." line became the message
$ git log --oneline -6
8d9c594 git commit -m "docs: add draw.io architecture diagram"
be6ddd5 chore: update .gitignore to include *.tsbuildinfo and remove tsconfig.tsbuildinfo
4639278 "docs: add project documentation, screenshots and portfolio README"
755203b git commit -m "docs: add project documentation and portfolio README"
ff98439 git commit -m "docs: add project documentation and README"
104962b docs: improve inline comments
```

**Fixing it / trying to fix it**
```console
$ git fetch -q && git status -sb | head -1
## main...origin/main
$ git diff --stat 104962b origin/main | tail -1      # the 3 bad commits only touch docs
 16 files changed, 1024 insertions(+), 845 deletions(-)
$ git log 104962b..origin/main --format='%h %an'
4639278 <author>
755203b <author>
ff98439 <author>

# suggested, not run (the pushed messages are still as above):
# squash them into one clean commit (run in a TERMINAL, not a commit box)
$ git reset --soft 104962b
$ git commit -m "docs: add project documentation, screenshots and portfolio README" \
    -m "Add a process document, a step-by-step recreation guide and interview questions, ..."
$ git push --force-with-lease origin main
(not run, so no output)

# for a single latest commit
$ git reset --soft HEAD~1
$ git commit -m "docs: add draw.io architecture diagram" -m "..."
$ git push --force-with-lease origin main
(not run, so no output)

$ git log --oneline -3      # check before every push
```

### Lessons learned
Know whether you're in a shell or a message field before pasting.

### How to approach it next time
Run `git log --oneline -5` after committing, before pushing.

### Prevention / follow-up
Optional `commit-msg` hook that rejects messages starting with `git ` or `"`:
```bash
#!/bin/sh
grep -qE '^(git |")' "$1" && { echo "commit message looks like a pasted command"; exit 1; } || exit 0
```

### References
- Git: [`git commit` — `-m`](https://git-scm.com/docs/git-commit#Documentation/git-commit.txt--mltmsggt).

---

## INC-041 – Storefront says "Make sure the backend is running at localhost:8000" while the backend was up

| | |
|---|---|
| **Date** | 2026-10-04 |
| **Where** | Browser, storefront home page; `frontend/src/pages/Home.tsx:126`, `frontend/src/components/layout/Footer.tsx:61` |
| **Severity** | Low: a misleading message; the backend and the EKS deployment were healthy |
| **Status** | Explained (the message text still needs fixing) |

### Summary
The shop showed *"Could not load products: Failed to fetch. Make sure the backend is running at localhost:8000."* Everything checked out healthy: 13 pods `Running`, Argo CD `Synced/Healthy`, and the GraphQL API answering through the ALB from the terminal. "Failed to fetch" means the **browser** couldn't complete the request (wrong page URL, a stale tab, or a DNS lookup cached while the ALB was new). The "localhost:8000" part is **hard-coded text** from local Docker Compose development, so it pointed the investigation in the wrong direction.

### Background
- The frontend calls `import.meta.env.VITE_GRAPHQL_URL ?? "/graphql"`, which is relative to the host that served the page. On EKS that's the ALB, which routes `/graphql` to the gateway.
- `Failed to fetch` is the browser's generic network error: no HTTP status came back at all.
- Locally (Compose), the API is at `localhost:8000`, which is where the text came from.

### Timeline
1. Error shown in the browser after the monitoring stack was installed.
2. Checked the API through the ALB with `curl` → products returned (200).
3. Checked pods on the bastion → all Running; catalog restarts only from startup.
4. Checked nodes → memory about 23%; Argo CD `Synced/Healthy`.
5. `grep -rn localhost:8000 frontend/src` → the text is hard-coded in two places.


### What happened (symptom)

    Could not load products: Failed to fetch. Make sure the backend is running at localhost:8000.

### How we got there
The page was opened from another URL or an old tab (for example the local dev site, or a tab loaded before the ALB was live), and its `/graphql` call failed in the browser.

### Why (root cause)
1. The browser's request to `/graphql` failed before any HTTP response (`Failed to fetch`).
2. The server side was verified healthy, so the cause was on the client side: the page's origin or a cached DNS failure (see INC-030, INC-009).
3. The UI message blames `localhost:8000` unconditionally, so it is wrong on every non-local deployment.

### Impact
No outage. Time lost chasing the wrong host.

### Resolution (step by step)
1. Prove the API works: `curl -s -X POST -H 'Content-Type: application/json' -d '{"query":"{ products(pagination:{skip:0,take:2}) { name } }"}' http://<APP_LB_HOST>/graphql` → JSON with products.
2. Open the shop from the **ALB URL** (`http://<APP_LB_HOST>/`), hard refresh (Ctrl+Shift+R), and clear the DNS cache if needed.
3. If it still fails: DevTools → Network → the red `graphql` request → its URL and error.

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
# browser: shop page → "Could not load products: Failed to fetch.
#           Make sure the backend is running at localhost:8000."
```

**Fixing it / trying to fix it**
```console
# laptop: is the API really down? (through the ALB)
$ H=<APP_LB_HOST>
$ curl -s -o /dev/null -w "/ -> %{http_code}\n" --max-time 10 http://$H/
/ -> 200
$ curl -s --max-time 15 -X POST -H 'Content-Type: application/json' \
    -d '{"query":"{ products(pagination:{skip:0,take:2}) { id name } }"}' http://$H/graphql
{"data":{"products":[{"id":"3KC8…","name":"Portable SSD"},{"id":"3KC8…","name":"HD Webcam"}]}}

# bastion (through SSM): pods, events, node capacity, Argo CD status
$ kubectl -n go-micro-shop get pods -o wide
account-db-0                       1/1   Running   0             32m
catalog-service-b6f4479c9-pvhq6    1/1   Running   1 (31m ago)   32m
frontend-6f9d69497f-dh2hz          1/1   Running   0             32m
...                                (13 pods, all Running)
$ kubectl -n go-micro-shop get events --sort-by=.lastTimestamp | tail -5
31m  Warning  Unhealthy  pod/catalog-service-…  Liveness probe failed: dial tcp …:8080: connect: connection refused   # startup only
$ kubectl describe nodes | grep -A6 "Allocated resources" | grep -E "cpu|memory"
  memory   1680Mi (23%)   4320Mi (61%)
$ kubectl -n argocd get application go-micro-shop-dev -o jsonpath='{.status.sync.status}/{.status.health.status}'; echo
Synced/Healthy

# where does "localhost:8000" come from?
$ grep -rn "localhost:8000" frontend/src
frontend/src/components/layout/Footer.tsx:61:          <p>GraphQL API at localhost:8000</p>
frontend/src/pages/Home.tsx:126:            <code className="rounded bg-red-100 px-1">localhost:8000</code>.
$ grep -n "GRAPHQL_URL" frontend/src/lib/graphql.ts
1:const GRAPHQL_URL = import.meta.env.VITE_GRAPHQL_URL ?? "/graphql";

# fix on the client side: open the ALB URL, hard refresh, or clear the cached DNS failure
$ resolvectl flush-caches        # browser: chrome://net-internals/#dns → Clear host cache
(no output)
```

### Code / config change
Proposed, not yet made: show the URL actually called instead of `localhost:8000` in `Home.tsx` and `Footer.tsx`.

### Lessons learned
Error messages that guess at the cause become wrong once the deployment changes. Report what was tried, not a hard-coded host.

### How to approach it next time
`Failed to fetch` → check the address bar's host first, then call the same endpoint with `curl`, then use DevTools → Network.

### Prevention / follow-up
- Fix the two hard-coded strings.

### References
- MDN: [`fetch()` network errors (TypeError: Failed to fetch)](https://developer.mozilla.org/en-US/docs/Web/API/Window/fetch#exceptions).

---

## INC-040 – Grafana UI loads with `curl` but not in the browser

| | |
|---|---|
| **Date** | 2026-10-03 |
| **Where** | Browser → `<GRAFANA_LB_HOST>`; monitoring step (`kubectl patch svc grafana … LoadBalancer`) |
| **Severity** | Low: one person's browser; Grafana was healthy |
| **Status** | Resolved |

### Summary
After exposing Grafana (and Prometheus) through Classic Load Balancers restricted to our IP, the Grafana URL didn't open in the browser, while from the same laptop `curl` returned `302 → /login` and `/login` returned `200`. Grafana serves **HTTP only on port 80**; typed without a scheme, the browser tries **HTTPS first**, and port 443 isn't open on that load balancer, so the attempt just hangs. The first attempts were also made seconds after the load balancer was created, the same DNS-timing effect as INC-030 and INC-009.

### Background
- Modern browsers upgrade bare hostnames to `https://` first ("HTTPS-First" / "Always use secure connections").
- The Classic LB from `kubectl patch … LoadBalancer` only listens on the Service's ports, here port 80 for Grafana.
- `loadBalancerSourceRanges` turns into security-group rules, so traffic from any other IP is silently dropped (a timeout, not a refusal).

### Timeline
1. Installed Prometheus and Grafana with Helm, patched both Services to `LoadBalancer` with `loadBalancerSourceRanges: ["<YOUR_IP>/32"]`.
2. Opened the URLs about 30 seconds later: not reachable.
3. Checked from the laptop: DNS resolving, the LB `internet-facing`, the SG allowing `<YOUR_IP>/32` on 80, backends `InService`, `curl` → 302 / 200.
4. The browser kept failing: the cause was narrowed to the scheme (HTTPS) and the cached DNS.


### What happened (symptom)
The browser showed "This site can't be reached" for the Grafana host; `curl http://…` worked.

### How we got there
The bare hostname was pasted into the browser right after the LB was created.

### Why (root cause)
1. The browser upgraded the request to `https://<GRAFANA_LB_HOST>/`.
2. The LB has no 443 listener, so the connection waits until it times out.
3. Earlier attempts had also cached a DNS failure, made while the LB's name wasn't live yet.

### Impact
A few minutes lost; nothing else affected.

### Resolution (step by step)
1. Server side: `curl -s -o /dev/null -w '%{http_code}\n' http://<GRAFANA_LB_HOST>/login` → `200`.
2. Open it with the scheme spelled out: `http://<GRAFANA_LB_HOST>/login`. If the browser forces HTTPS, turn off "Always use secure connections" or click "Continue to site".
3. Clear the cached lookup: `chrome://net-internals/#dns` → *Clear host cache*, or use an incognito window.
4. If it still fails, check the browser's own egress IP (`https://checkip.amazonaws.com` **in the browser**) against the allowlist, since a VPN changes it.

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
# bastion: monitoring the tutorial way, exposed with an IP allowlist
$ helm repo add prometheus-community https://prometheus-community.github.io/helm-charts && helm repo update
"prometheus-community" has been added to your repositories                       # expected
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "prometheus-community" chart repository
Update Complete. ⎈Happy Helming!⎈
$ helm install prometheus prometheus-community/prometheus -n monitoring --create-namespace --version 29.35.0
NAME: prometheus                                                                  # expected
NAMESPACE: monitoring
STATUS: deployed
REVISION: 1
$ helm install grafana oci://ghcr.io/grafana-community/helm-charts/grafana -n monitoring --version 13.2.7
Pulled: ghcr.io/grafana-community/helm-charts/grafana:13.2.7                      # expected
NAME: grafana
NAMESPACE: monitoring
STATUS: deployed
REVISION: 1
$ kubectl -n monitoring patch svc prometheus-server -p '{"spec":{"type":"LoadBalancer","loadBalancerSourceRanges":["<YOUR_IP>/32"]}}'
service/prometheus-server patched                                                 # expected
$ kubectl -n monitoring patch svc grafana           -p '{"spec":{"type":"LoadBalancer","loadBalancerSourceRanges":["<YOUR_IP>/32"]}}'
service/grafana patched                                                           # expected
$ kubectl -n monitoring get svc prometheus-server grafana
NAME                TYPE           EXTERNAL-IP                                   PORT(S)        AGE
prometheus-server   LoadBalancer   <hash>-<id>.eu-north-1.elb.amazonaws.com      80:32099/TCP   31s
grafana             LoadBalancer   <hash>-<id>.eu-north-1.elb.amazonaws.com      80:32368/TCP   29s
# browser (31 seconds later): "This site can't be reached"
```

**Fixing it / trying to fix it**
```console
# laptop: load balancer, allowlist and backends
$ aws elb describe-load-balancers --region eu-north-1 --load-balancer-names <grafana-lb-name> \
    --query 'LoadBalancerDescriptions[0].[Scheme,SecurityGroups[0]]' --output text
internet-facing   sg-<id>
$ aws ec2 describe-security-groups --region eu-north-1 --group-ids sg-<id> \
    --query 'SecurityGroups[0].IpPermissions[].[FromPort,IpRanges[].CidrIp|join(`,`,@)]' --output text
80    <YOUR_IP>/32
3     <YOUR_IP>/32
$ aws elb describe-instance-health --region eu-north-1 --load-balancer-name <grafana-lb-name> \
    --query 'InstanceStates[].State' --output text
InService   InService
$ curl -s https://checkip.amazonaws.com
<YOUR_IP>                                        # matches the allowlist

# DNS and HTTP from the terminal
$ getent hosts <GRAFANA_LB_HOST>
<LB_IP_1>   <GRAFANA_LB_HOST>
$ curl -s -o /dev/null -w '%{http_code} -> %{redirect_url}\n' --max-time 10 http://<GRAFANA_LB_HOST>/
302 -> http://<GRAFANA_LB_HOST>/login
$ curl -s -o /dev/null -w '%{http_code}\n' --max-time 10 http://<GRAFANA_LB_HOST>/login
200

# bastion: pods healthy?
$ kubectl -n monitoring get pods,svc
pod/grafana-…                                1/1   Running   0   7m42s
pod/prometheus-server-…                      2/2   Running   0   7m44s
...

# browser fix: use the scheme explicitly and drop the cached lookup
#   http://<GRAFANA_LB_HOST>/login   (not https://)
$ resolvectl flush-caches        # or chrome://net-internals/#dns → Clear host cache
(no output)
$ kubectl -n monitoring get secret grafana -o jsonpath="{.data.admin-password}" | base64 -d; echo
<grafana-admin-password>                                                          # redacted
```

### Code / config change
None.

### Lessons learned
`curl` working while the browser fails points at the browser: its scheme, its cache, or its network path.

### How to approach it next time
Compare `curl http://`, `curl https://` and the browser's egress IP. One of them will differ.

### Prevention / follow-up
Always write these URLs as `http://…` in docs; long term, put the UIs behind TLS (needs a domain).

### References
- Chrome: [HTTPS-First Mode](https://blog.chromium.org/2021/07/increasing-https-adoption.html).
- See also INC-009 (same pattern with Argo CD on the rebuild).

---

## INC-039 – Helm warns `this chart is deprecated` for `grafana/grafana`

| | |
|---|---|
| **Date** | 2026-10-03 |
| **Where** | Monitoring step: `helm template/install grafana --repo https://grafana.github.io/helm-charts` |
| **Severity** | Low: it worked, but the chart no longer gets fixes |
| **Status** | Resolved (switched to `grafana-community/grafana`) |

### Summary
Rendering the Grafana chart printed `level=WARN msg="this chart is deprecated"`. The chart's own README said it is being migrated to **grafana-community/helm-charts**. The install switched to `oci://ghcr.io/grafana-community/helm-charts/grafana` (chart 13.2.7, Grafana 13.2.3), and the same values rendered identically.

### Background
- Helm prints this warning when `Chart.yaml` contains `deprecated: true`.
- Deprecated charts stop receiving security and compatibility fixes, so the warning matters even when the install works.

### Timeline
1. `helm pull grafana --repo https://grafana.github.io/helm-charts --version 10.5.15`, then `helm template` → the warning.
2. `tar -xzf grafana-10.5.15.tgz -O grafana/Chart.yaml` → `deprecated: true`; the README pointed to grafana-community.
3. Found the new repository and version: `helm show chart grafana --repo https://grafana-community.github.io/helm-charts` → `version: 13.2.7`.
4. Rendered the same values with the new chart: Service, PVC, data source and dashboards all present.

### What happened (symptom)

    level=WARN msg="this chart is deprecated"

### Why (root cause)
Grafana Labs moved the chart to a community-maintained repository and marked the old one deprecated.

### Impact
None (caught before installing).

### Resolution (step by step)
1. `helm install grafana oci://ghcr.io/grafana-community/helm-charts/grafana -n monitoring --version 13.2.7`
2. Verify: `kubectl -n monitoring get pods` → `grafana-…` `Running`.

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
$ helm pull grafana --repo https://grafana.github.io/helm-charts --version 10.5.15 -d ./charts
(no output)
$ helm template grafana ./charts/grafana-10.5.15.tgz -n monitoring --set 'service.loadBalancerSourceRanges={1.2.3.4/32}' > /dev/null
level=WARN msg="this chart is deprecated"
```

**Fixing it / trying to fix it**
```console
$ tar -xzf ./charts/grafana-10.5.15.tgz -O grafana/Chart.yaml | grep -iE "deprecat|description|sources" -A1
deprecated: true
description: The leading tool for querying and visualizing time series and metrics.
$ tar -xzf ./charts/grafana-10.5.15.tgz -O grafana/README.md | grep -iE "migrated|community"

**Fixing it / trying to fix it** (https://github.com/grafana-community/helm-charts)
$ curl -fsSL https://raw.githubusercontent.com/grafana-community/helm-charts/main/README.md | grep -iE "helm repo add|oci://"
helm install RELEASE-NAME oci://ghcr.io/grafana-community/helm-charts/<chart-name>
helm repo add grafana-community https://grafana-community.github.io/helm-charts
$ helm show chart grafana --repo https://grafana-community.github.io/helm-charts | grep -E "^(version|appVersion|deprecated):"
appVersion: 13.2.3
version: 13.2.7

# same chart from the maintained repo: no warning
$ helm template grafana oci://ghcr.io/grafana-community/helm-charts/grafana --version 13.2.7 -n monitoring > /dev/null
Pulled: ghcr.io/grafana-community/helm-charts/grafana:13.2.7
Digest: sha256:<digest>
(no deprecation warning)
$ helm install grafana oci://ghcr.io/grafana-community/helm-charts/grafana -n monitoring --version 13.2.7
NAME: grafana                                                                     # expected
NAMESPACE: monitoring
STATUS: deployed
```

### Code / config change
Install command and docs now use the grafana-community chart.

### Lessons learned
Read Helm warnings. "Deprecated" means no future fixes, even when everything works today.

### How to approach it next time
`helm show chart <chart> --repo <repo> | grep deprecated` before adopting a chart, and look for a `home`/`sources` link to the new location.

### Prevention / follow-up
Pin chart versions; re-check deprecation when upgrading.

### References
- grafana-community/helm-charts: <https://github.com/grafana-community/helm-charts>.
- Helm: [`Chart.yaml` `deprecated` field](https://helm.sh/docs/topics/charts/#the-chartyaml-file).

---

## INC-038 – Kustomize puts a `namespace` on the cluster-scoped `ClusterSecretStore`

| | |
|---|---|
| **Date** | 2026-10-03 |
| **Where** | `k8s/base/secret-store.yaml`, `k8s/overlays/*/kustomization.yaml` (`namespace: go-micro-shop`) |
| **Severity** | Low: caught while rendering, before any deploy |
| **Status** | Resolved (namespaced `SecretStore` instead) |

### Summary
The first version used an External Secrets **`ClusterSecretStore`** (cluster-scoped). Rendering the overlay with `kustomize build` showed it with `namespace: go-micro-shop`, because the `namespace:` transformer adds a namespace to every resource it doesn't know is cluster-scoped, including custom resources. A cluster-scoped object with a namespace causes errors and endless drift in Argo CD. It was replaced with a namespaced **`SecretStore`**, which does the same job for one namespace.

### Background
- Kustomize knows which **built-in** kinds are cluster-scoped (Namespace, ClusterRole, …). For CRDs it doesn't know unless told, so it adds the namespace.
- `ClusterSecretStore` can be used from any namespace; `SecretStore` only from its own. This app only needs `go-micro-shop`.

### Timeline
1. Added `ClusterSecretStore` to `base/` and `ExternalSecret`s to the overlays.
2. `kustomize build k8s/overlays/dev | grep -B2 -A6 "kind: ClusterSecretStore"` → `namespace: go-micro-shop` present.
3. Switched to `kind: SecretStore` and changed `secretStoreRef.kind` in both ExternalSecrets.
4. Rebuilt: SecretStore namespaced correctly, ExternalSecrets reference it.

### What happened (symptom)

    kind: ClusterSecretStore
    metadata:
      name: aws-secrets-manager
      namespace: go-micro-shop      # ← shouldn't exist on a cluster-scoped object

### Why (root cause)
The overlay's `namespace:` transformer plus a CRD whose scope Kustomize doesn't know about.

### Impact
None; it was caught in local rendering.

### Resolution (step by step)
1. `sed -i 's/kind: ClusterSecretStore/kind: SecretStore/' k8s/overlays/*/external-secrets.yaml`, and change the store's `kind` in `base/secret-store.yaml`.
2. Verify: `kustomize build k8s/overlays/dev | grep -c ClusterSecretStore` → `0`.

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
# base/secret-store.yaml had kind: ClusterSecretStore; overlays set namespace: go-micro-shop
$ kustomize build k8s/overlays/dev | grep -B2 -A6 "kind: ClusterSecretStore" | grep -n namespace
6:  namespace: go-micro-shop          # ← namespace stamped on a cluster-scoped object
```

**Fixing it / trying to fix it**
```console
$ for e in dev prod; do echo "== $e"; kustomize build k8s/overlays/$e | grep -E "^kind:|key: go-micro" | sort | uniq -c; done
== dev
      1       key: go-micro-shop/dev/account-db
      1       key: go-micro-shop/dev/order-db
      1 kind: ClusterSecretStore
      2 kind: ExternalSecret
...
# switch to a namespaced SecretStore (base file rewritten with kind: SecretStore)
$ sed -i 's/kind: ClusterSecretStore/kind: SecretStore/' k8s/overlays/*/external-secrets.yaml
(no output)
$ kustomize build k8s/overlays/dev | grep -B1 -A4 "kind: SecretStore$"
kind: SecretStore
metadata:
  name: aws-secrets-manager
  namespace: go-micro-shop
$ kustomize build k8s/overlays/dev | grep -c ClusterSecretStore
0
```

### Code / config change
`k8s/base/secret-store.yaml`, `k8s/overlays/{dev,prod}/external-secrets.yaml` (commit `9e76177`).

### Lessons learned
Always render Kustomize output before committing. Namespace transformers and CRDs don't mix silently.

### How to approach it next time
For a cluster-scoped CRD in a namespaced Kustomization: keep it in a separate kustomization without `namespace:`, or teach Kustomize its scope with an OpenAPI schema (`openapi:` field).

### Prevention / follow-up
CI step: `kustomize build k8s/overlays/<env>` on every push.

### References
- Kustomize: [`namespace` transformer](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/namespace/).
- External Secrets: [SecretStore vs ClusterSecretStore](https://external-secrets.io/latest/api/clustersecretstore/).

---

## INC-037 – Jenkinsfile `sed` would write `newTag: 42` without quotes (Groovy escaping)

| | |
|---|---|
| **Date** | 2026-10-03 |
| **Where** | `jenkins/Jenkinsfile-Backend` and `-Frontend`, stage *Update Deployment file* |
| **Severity** | Low: caught in review before the first pipeline run |
| **Status** | Resolved |

### Summary
The tag-bump `sed` was written with `\"` around `${IMAGE_TAG}` inside a Groovy `'''…'''` string. Groovy turns `\"` into a plain `"`, so the shell received `newTag: "${IMAGE_TAG}"` *inside* a double-quoted argument. The inner quotes close the shell string, and the result is `newTag: 42`, a YAML **number** where Kustomize expects a string. The fix is to write `\\"` in the Jenkinsfile, so the shell receives `\"`.

### Background
- Jenkins runs `sh '''…'''` after Groovy has processed escapes: `\\` → `\`, `\"` → `"`.
- The shell then parses the result. A `"` inside `"…"` ends the string.
- In `kustomization.yaml`, `newTag` is a string; an unquoted `42` parses as an integer.

### Timeline
1. Wrote the `sed` with `\"`.
2. A review traced the escaping layer by layer (Groovy, then shell).
3. Emulated Groovy's unescaping, ran the resulting line on a copy of the file, and saw the quotes missing.
4. Changed to `\\"` and re-ran the emulation: `newTag: "7"` for the frontend and `"42"` for each backend service.


### What happened (symptom)
(Would have happened.) `newTag: 42` in `k8s/base/kustomization.yaml`, which risks a Kustomize or Argo CD error, or a silent type change.

### Why (root cause)
Two escaping layers, Groovy and then the shell, with escapes written for only one of them.

### Impact
None (pre-merge).

### Resolution (step by step)
1. Jenkinsfile: `sed -i -E "/^  - name: ${svc}$/{n;s/newTag: .*/newTag: \\"${IMAGE_TAG}\\"/}" k8s/base/kustomization.yaml`
2. Verify by emulating: replace `\\` with `\` in the line (what Groovy does), run it with `IMAGE_TAG=7`, then `grep -A1 'name: frontend'` → `newTag: "7"`.

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it** (review of the Jenkinsfile line)
```console
$ grep -n 'newTag: ' jenkins/Jenkinsfile-Backend
132:  sed -i -E "/^  - name: ${svc}$/{n;s/newTag: .*/newTag: \"${IMAGE_TAG}\"/}" k8s/base/kustomization.yaml
# inside a Groovy ''' string, \" becomes " → the shell's double-quoted argument ends early
```

**Fixing it / trying to fix it**
```console
# 1. the sed itself is right when the shell gets \" (direct test on a copy)
$ cp k8s/base/kustomization.yaml /tmp/k.yaml
(no output)
$ IMAGE_TAG=42; for svc in account catalog order graphql; do
    sed -i -E "/^  - name: ${svc}$/{n;s/newTag: .*/newTag: \"${IMAGE_TAG}\"/}" /tmp/k.yaml; done
(no output)
$ sed -n '/^images:/,$p' /tmp/k.yaml | head -4
images:
  - name: account
    newTag: "42"

# 2. write \\" in the Jenkinsfile so Groovy hands \" to the shell
$ sed -i 's/newTag: \\"${IMAGE_TAG}\\"/newTag: \\\\"${IMAGE_TAG}\\\\"/' jenkins/Jenkinsfile-Backend jenkins/Jenkinsfile-Frontend
(no output)
$ grep -n 'newTag: ' jenkins/*
jenkins/Jenkinsfile-Backend:132:  sed -i -E "/^  - name: ${svc}$/{n;s/newTag: .*/newTag: \\"${IMAGE_TAG}\\"/}" k8s/base/kustomization.yaml
jenkins/Jenkinsfile-Frontend:123: sed -i -E "/^  - name: ${AWS_ECR_REPO_NAME}$/{n;s/newTag: .*/newTag: \\"${IMAGE_TAG}\\"/}" k8s/base/kustomization.yaml

# 3. emulate Groovy (\\ → \) and run exactly what the shell will run
$ python3 - <<'EOF'
import subprocess, os, shutil
for jf, env in [('jenkins/Jenkinsfile-Frontend', {'AWS_ECR_REPO_NAME':'frontend','IMAGE_TAG':'7'}),
                ('jenkins/Jenkinsfile-Backend',  {'svc':'order','IMAGE_TAG':'42'})]:
    line=[l for l in open(jf) if 'newTag: ' in l][0].strip().replace('\\\\','\\')
    shutil.copy('k8s/base/kustomization.yaml','/tmp/k.yaml')
    subprocess.run(['bash','-c',line.replace('k8s/base/kustomization.yaml','/tmp/k.yaml')], env={**os.environ, **env}, check=True)
    print(open('/tmp/k.yaml').read().split('- name: '+env.get('svc', env.get('AWS_ECR_REPO_NAME')))[1][:25])
EOF
    newTag: "7"
    newTag: "42"
```

### Code / config change
Both Jenkinsfiles (commit `9e76177`).

### Lessons learned
For every string, count the layers that will interpret it (Groovy, shell, sed, YAML) and escape for each.

### How to approach it next time
Use `echo` in a test pipeline to print the exact command the shell receives, or avoid `sed` and edit YAML with a tool (`yq`, `kustomize edit set image`).

### Prevention / follow-up
Consider `kustomize edit set image name=name:${TAG}` in the pipeline instead of `sed`.

### References
- Jenkins: [String interpolation in `sh` steps](https://www.jenkins.io/doc/book/pipeline/jenkinsfile/#string-interpolation).

---

## INC-036 – Argo CD's registry override would have dropped the image tag (Kustomize `set image`)

| | |
|---|---|
| **Date** | 2026-10-03 |
| **Where** | Design of `k8s/` (base/overlays) and `k8s/scripts/create-argocd-app.sh` (`spec.source.kustomize.images`) |
| **Severity** | Medium: if shipped, every image would have been deployed without its version tag |
| **Status** | Resolved (two-layer design, tested) |

### Summary
To keep the AWS account ID out of git, the plan was: git holds only image **tags** (`newTag`, bumped by Jenkins), and the Argo CD Application injects the **registry** via `kustomize.images`. Argo CD applies that with `kustomize edit set image name=<registry>/name`. A local test showed that this **replaces the whole `images` entry and drops `newTag`**, so the output was `<registry>/account` with no tag. Putting the tag in `base/` and letting Argo CD override the registry on the **overlay** keeps both: `<registry>/account:42`.

### Background
- A Kustomize `images:` entry has `name`, `newName` and `newTag`.
- `kustomize edit set image a=b` rewrites the entry for `a` in **that** kustomization; fields you don't pass don't survive.
- Transformers in different layers compose: the base sets the tag, the overlay renames the image, and the tag carries through.

### Timeline
1. Wrote the plan: tag in git, registry from Argo CD.
2. Before relying on it, tested with the real `kustomize` binary (v5.8.2) in a scratch folder.
3. Single layer: tag lost.
4. Two layers (base + overlay): both kept, and bumping the base tag changes the output (`:42` → `:43`).
5. Restructured `k8s/` into `base/` + `overlays/{dev,prod}`.


### What happened (symptom)

    kustomize edit set image account=<registry>/account
    kustomize build . | grep image:
      - image: <registry>/account          # newTag "42" gone

### Why (root cause)
`set image` without a tag rewrites the entry in place and drops the existing `newTag`.

### Impact
None, because it was tested before being built on. Untested, every Deployment would have used an untagged image, which immutable ECR repos reject (`ImagePullBackOff`), or which would pull `:latest`.

### Resolution (step by step)
1. `base/kustomization.yaml`: `images: [{name: account, newTag: "0"}, …]`, bumped by Jenkins.
2. `overlays/dev/kustomization.yaml`: `resources: [../../base]`. Argo CD points here.
3. Argo CD Application: `kustomize.images: [account=<registry>/account, …]` (registry only).
4. Verify: copy the repo, run `kustomize edit set image` in the overlay for all 5 images, then `kustomize build` → `<registry>/<svc>:<tag>` for all five; the DB images are unchanged.

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it** (testing the design with the real kustomize)
```console
$ curl -fsSL "https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2Fv5.8.2/kustomize_v5.8.2_linux_amd64.tar.gz" | tar xz
(no output)
$ mkdir app && cd app      # d.yaml = Deployment with "image: account"
(no output)
$ cat kustomization.yaml
resources: [d.yaml]
images:
  - name: account
    newTag: "42"
# what Argo CD runs for spec.source.kustomize.images:
$ ../kustomize edit set image account=123456789012.dkr.ecr.eu-north-1.amazonaws.com/account
(no output)
$ cat kustomization.yaml
resources:
- d.yaml
images:
- name: account
  newName: 123456789012.dkr.ecr.eu-north-1.amazonaws.com/account          # newTag is gone
$ ../kustomize build . | grep image:
      - image: 123456789012.dkr.ecr.eu-north-1.amazonaws.com/account
```

**Fixing it / trying to fix it** (two layers)
```console
$ mkdir base overlay && cp app/d.yaml base/
(no output)
$ printf 'resources: [d.yaml]\nimages:\n  - name: account\n    newTag: "42"\n' > base/kustomization.yaml
(no output)
$ printf 'resources: [../base]\n' > overlay/kustomization.yaml
(no output)
$ cd overlay && ../kustomize edit set image account=123456789012.dkr.ecr.eu-north-1.amazonaws.com/account
(no output)
$ ../kustomize build . | grep image:
      - image: 123456789012.dkr.ecr.eu-north-1.amazonaws.com/account:42
$ cd ../base && sed -i 's/"42"/"43"/' kustomization.yaml && cd ../overlay && ../kustomize build . | grep image:
      - image: 123456789012.dkr.ecr.eu-north-1.amazonaws.com/account:43     # tag bumps still flow through

# full repo: simulate Argo CD on k8s/overlays/dev for all 5 images
$ cp -r k8s /tmp/argotest && cd /tmp/argotest/k8s/overlays/dev
(no output)
$ for i in account catalog order graphql frontend; do kustomize edit set image $i=123456789012.dkr.ecr.eu-north-1.amazonaws.com/$i; done
(no output)
$ kustomize build . | grep "image:"
        image: 123456789012.dkr.ecr.eu-north-1.amazonaws.com/account:0
        image: 123456789012.dkr.ecr.eu-north-1.amazonaws.com/catalog:0
      - image: 123456789012.dkr.ecr.eu-north-1.amazonaws.com/frontend:0
        image: 123456789012.dkr.ecr.eu-north-1.amazonaws.com/graphql:0
        image: 123456789012.dkr.ecr.eu-north-1.amazonaws.com/order:0
        image: postgres:16-alpine
        image: docker.elastic.co/elasticsearch/elasticsearch:6.8.23
```
(`123456789012` is a placeholder account ID.)

### Code / config change
`k8s/base`, `k8s/overlays`, `k8s/scripts/create-argocd-app.sh` (commit `9e76177`).

### Lessons learned
Test the exact tool behaviour a design depends on before building on it.

### How to approach it next time
Reproduce the tool's operation locally (here, `kustomize edit set image` is what Argo CD runs) and inspect the rendered output.

### Prevention / follow-up
Keep the two-layer structure; don't put the registry override in `base/`.

### References
- Argo CD: [Kustomize parameter overrides](https://argo-cd.readthedocs.io/en/stable/user-guide/kustomize/).
- Kustomize: [`images` field](https://kubectl.docs.kubernetes.io/references/kustomize/kustomization/images/).

---

## INC-035 – Trivy gate blocks the frontend image: CVE-2026-31789 (OpenSSL, CRITICAL)

| | |
|---|---|
| **Date** | 2026-10-03 |
| **Where** | Jenkins `frontend` #2 › *TRIVY Image Scan*; `frontend/Dockerfile` (`nginx:1.27-alpine`) |
| **Severity** | Medium: the frontend release was blocked (by design); no vulnerable image reached ECR |
| **Status** | Resolved (`frontend` #3 green) |

### Summary
The frontend pipeline's image gate (`trivy image --severity CRITICAL --ignore-unfixed --exit-code 1`) failed: the runtime image `nginx:1.27-alpine` was built on **Alpine 3.21.3** with **OpenSSL 3.3.3**, affected by **CVE-2026-31789** (heap buffer overflow, fixed in 3.3.7). The push and the GitOps tag bump were skipped, which is the gate working as intended. The fix was a supported base, `nginx:1.30-alpine` plus `apk upgrade --no-cache` (Alpine 3.24.2, OpenSSL 3.5.9). It was verified locally with the same Trivy command (0 CRITICAL), and `frontend` #3 passed. The `graphql` runtime was also moved off end-of-life `alpine:3.20`.

### Background
- The gate scans **before** the push, so a failing image never reaches the registry.
- `--ignore-unfixed` fails only when a fix exists, so the action is always "update".
- nginx's odd minor versions are *mainline*, even ones *stable*. Once a branch ends (1.27), its images stop being rebuilt, so old OS packages remain.
- `apk upgrade` during the build pulls fixes published after the base image was built.

### Timeline
1. `frontend` #2: Sonar ✓, quality gate ✓, OWASP ✓, Trivy fs ✓, docker build ✓ → **TRIVY Image Scan** ✗ (exit 1); ECR push and Update Deployment file skipped.
2. Report: 2 CRITICAL (libcrypto3, libssl3), status `fixed`, fixed version `3.3.7-r0`.
3. Checked the nginx tags: stable is now 1.30.x and mainline 1.31.x; Alpine 3.24 is the latest.
4. Changed the Dockerfile; built locally: `cat /etc/alpine-release` → `3.24.2`, `apk info -v | grep libssl3` → `libssl3-3.5.9-r0`.
5. Ran the gate locally: `docker run aquasec/trivy image --severity CRITICAL --ignore-unfixed --exit-code 1 local/frontend:test` → 0, exit 0.
6. Pushed; `frontend` #3 passed all stages and committed `chore(deploy): update frontend image to 3`.


### What happened (symptom)

    Total: 2 (CRITICAL: 2)
    libcrypto3  CVE-2026-31789  CRITICAL  fixed  3.3.3-r0  3.3.7-r0
    libssl3     CVE-2026-31789  CRITICAL  fixed  3.3.3-r0  3.3.7-r0
    script returned exit code 1

### Why (root cause)
1. The frontend Dockerfile pinned an ended nginx branch (`1.27-alpine`).
2. That image's Alpine base (3.21.3) shipped OpenSSL 3.3.3.
3. CVE-2026-31789 affects 3.3.3, and a fixed package (3.3.7) exists, so `--ignore-unfixed` doesn't skip it and the gate fails.

### Impact
- The frontend release was delayed by one fix-and-rebuild cycle.
- **No vulnerable image was pushed or deployed.**

### Resolution (step by step)
1. Read the report: library, CVE, installed version, fixed version.
2. Find the current supported base: Docker Hub tags (`nginx:1.30-alpine`, `alpine:3.24`).
3. Dockerfile runtime stage:
   ```dockerfile
   FROM nginx:1.30-alpine
   RUN apk upgrade --no-cache
   ```
4. Verify locally before pushing:
   ```bash
   docker build -t local/frontend:test frontend
   docker run --rm -v /var/run/docker.sock:/var/run/docker.sock aquasec/trivy:latest \
     image --severity CRITICAL --ignore-unfixed --exit-code 1 --quiet local/frontend:test; echo "exit=$?"
   ```
   Expect `0` vulnerabilities and `exit=0`.
5. Push and re-run `frontend` → all stages green; ECR has `frontend:3`.

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it** (Jenkins `frontend` #2 console)
```console
+ trivy image --severity HIGH,CRITICAL <ACCOUNT_ID>.dkr.ecr.eu-north-1.amazonaws.com/frontend:2
INFO  Detected OS  family="alpine" version="3.21.3"
+ trivy image --severity CRITICAL --ignore-unfixed --exit-code 1 --quiet <ACCOUNT_ID>.dkr.ecr.eu-north-1.amazonaws.com/frontend:2
<ACCOUNT_ID>.dkr.ecr.eu-north-1.amazonaws.com/frontend:2 (alpine 3.21.3)
Total: 2 (CRITICAL: 2)
│ libcrypto3 │ CVE-2026-31789 │ CRITICAL │ fixed │ 3.3.3-r0 │ 3.3.7-r0 │ openssl: Heap buffer overflow ...
│ libssl3    │ CVE-2026-31789 │ CRITICAL │ fixed │ 3.3.3-r0 │ 3.3.7-r0 │
script returned exit code 1          # ECR Image Pushing and Update Deployment file skipped
```

**Fixing it / trying to fix it**
```console
# laptop: which base images are current? (Docker Hub)
$ curl -fsSL "https://hub.docker.com/v2/repositories/library/nginx/tags?page_size=100&name=alpine" \
    | jq -r '.results[].name' | grep -E '^[0-9]+\.[0-9]+(\.[0-9]+)?-alpine$|^stable-alpine$' | sort -V | tail -4
1.30.5-alpine
1.31-alpine
1.31.6-alpine
stable-alpine
$ curl -fsSL "https://hub.docker.com/v2/repositories/library/alpine/tags?page_size=50" | jq -r '.results[].name' | grep -E '^3\.[0-9]+$' | sort -V | tail -2
3.23
3.24
$ grep -n "^FROM" */app.dockerfile frontend/Dockerfile
graphql/app.dockerfile:18:FROM alpine:3.20
frontend/Dockerfile:13:FROM nginx:1.27-alpine

# fix: frontend → nginx:1.30-alpine + "RUN apk upgrade --no-cache"; graphql → alpine:3.22
$ git diff --stat
 frontend/Dockerfile    | 5 ++++-
 graphql/app.dockerfile | 2 +-

# verify locally with the SAME gate before pushing
$ docker build -q -t local/frontend:test frontend
sha256:03dc…
$ docker run --rm local/frontend:test sh -c 'cat /etc/alpine-release; apk info -v | grep -E "^libssl3|^libcrypto3"'
3.24.2
libcrypto3-3.5.9-r0
libssl3-3.5.9-r0
$ docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -v $HOME/.cache/trivy:/root/.cache/ aquasec/trivy:latest \
    image --severity CRITICAL --ignore-unfixed --exit-code 1 --quiet local/frontend:test; echo "exit=$?"
│ local/frontend:test (alpine 3.24.2) │ alpine │ 0 │ - │
exit=0
$ docker rmi -f local/frontend:test >/dev/null && echo removed
removed

$ git add frontend/Dockerfile graphql/app.dockerfile && git commit -m "fix: update frontend and graphql base images" && git push
[main e5af417] fix: update frontend and graphql base images                       # expected
 2 files changed, 5 insertions(+), 2 deletions(-)
# Jenkins: frontend → Build Now → #3 all stages green, commit "chore(deploy): update frontend image to 3"
$ aws ecr list-images --region eu-north-1 --repository-name frontend --query 'imageIds[].imageTag' --output text
3
```

### Code / config change
`frontend/Dockerfile`, `graphql/app.dockerfile` (commit `e5af417`).

### Lessons learned
- A blocking gate catches real problems: this one stopped a critical CVE that would otherwise have shipped.
- Pin base images to **supported** branches; ended branches quietly collect CVEs.

### How to approach it next time
1. Read the "Fixed Version" column: the fix is usually a base-image bump.
2. Check whether the base branch is still maintained.
3. Re-run the same Trivy command locally before pushing.

### Prevention / follow-up
- Dependabot/Renovate for Dockerfile base images.
- A scheduled re-scan of deployed tags (ECR continuous scanning / Inspector).

### References
- Trivy: [`--ignore-unfixed`](https://trivy.dev/latest/docs/configuration/filtering/).
- nginx: [stable vs mainline](https://nginx.org/en/download.html).

---

## INC-034 – Jenkins: `permission denied … /var/run/docker.sock`

| | |
|---|---|
| **Date** | 2026-10-03 |
| **Where** | Jenkins `backend` #2 › *Docker Image Build*; `terraform/jenkins-server/setup.sh` |
| **Severity** | Medium: every pipeline's Docker build failed |
| **Status** | Resolved (Jenkins restarted; `setup.sh` fixed) |

### Summary
The first Docker command in the pipeline failed with `permission denied while trying to connect to the docker API`. The `jenkins` user **was** in the `docker` group, but the **running** Jenkins process wasn't: it had started before `setup.sh` added the user to the group, and a process only gets its groups when it starts. A Jenkins restart fixed it, and `setup.sh` now restarts Jenkins right after the `usermod`.

### Background
- `/var/run/docker.sock` is `root:docker`, mode `660`: only root and members of `docker` can use Docker.
- Group membership is read at login or process start. `usermod -aG` doesn't change processes that are already running.
- `/proc/<pid>/status` → `Groups:` shows a process's actual groups.

### Timeline
1. `backend` #2: Sonar ✓, gate ✓, OWASP ✓ (first NVD download), Trivy fs ✓ → Docker Image Build ✗.
2. On the host: `id jenkins` lists `docker` (988), but the Jenkins PID's `Groups:` shows only `113`.
3. Restarted Jenkins (`/safeRestart`) → `backend` #3 passed the Docker build, push and tag bump.
4. Added `systemctl restart jenkins` after the `usermod` in `setup.sh`, and `ignore_changes = [ami, user_data]` in `ec2.tf`, so editing the script doesn't reboot the live server. Plan: `No changes`.


### What happened (symptom)

    + docker container prune -f
    permission denied while trying to connect to the docker API at unix:///var/run/docker.sock

### How we got there
User-data order in `setup.sh`: install Jenkins (which starts it) → install Docker → `usermod -aG docker jenkins` → restart Docker only.

### Why (root cause)
1. Jenkins started before the group change.
2. Its process kept its original groups (113 only).
3. The socket allows only root and group 988, so the connection was refused.

### Impact
Both app pipelines blocked at the Docker stage until the restart.

### Resolution (step by step)
1. Confirm (Jenkins host):
   ```bash
   id jenkins                                                   # user has docker
   grep Groups /proc/$(systemctl show -p MainPID --value jenkins)/status   # process doesn't
   ls -l /var/run/docker.sock                                   # root:docker 660
   ```
2. Restart: `http://<JENKINS_IP>:8080/safeRestart` (waits for running builds) or `sudo systemctl restart jenkins`.
3. Verify: the `Groups:` line now includes `988`; re-run the pipeline.
4. Prevent: in `setup.sh`, after `usermod -aG docker jenkins`, add `systemctl restart jenkins`.

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it** (Jenkins `backend` #2 console)
```console
+ docker container prune -f
permission denied while trying to connect to the docker API at unix:///var/run/docker.sock
```

**Fixing it / trying to fix it**
```console
# laptop → Jenkins host through SSM
$ aws ssm send-command --region eu-north-1 --instance-ids <JENKINS_INSTANCE_ID> --document-name AWS-RunShellScript \
    --parameters 'commands=["id jenkins","P=$(systemctl show -p MainPID --value jenkins); echo jenkins pid $P; grep Groups /proc/$P/status; getent group docker; ls -l /var/run/docker.sock"]'
uid=111(jenkins) gid=113(jenkins) groups=113(jenkins),988(docker)    # the user IS in docker
jenkins pid 3927
Groups:	113                                                           # the running process is NOT
docker:x:988:jenkins,ubuntu
srw-rw---- 1 root docker 0 Oct  3 09:36 /var/run/docker.sock

# fix now: restart Jenkins so the process picks up the group
#   Jenkins UI: http://<JENKINS_IP>:8080/safeRestart   (or on the host:)
$ sudo systemctl restart jenkins
(no output)
$ grep Groups /proc/$(systemctl show -p MainPID --value jenkins)/status
Groups:	113 988                                                               # expected
# Jenkins: backend → Build Now → #3 Docker Image Build ✓, ECR push ✓, Update Deployment file ✓

# fix for next time: setup.sh restarts Jenkins after usermod; don't reboot the live server
$ grep -n "usermod -aG docker\|systemctl restart" terraform/jenkins-server/setup.sh
113:usermod -aG docker jenkins
116:systemctl restart docker
120:systemctl restart jenkins
$ grep -n ignore_changes terraform/jenkins-server/ec2.tf
43:    ignore_changes = [ami, user_data]
$ cd terraform/jenkins-server && terraform init -reconfigure -backend-config=backend.hcl >/dev/null && terraform plan | grep -E "Plan:|No changes"
No changes. Your infrastructure matches the configuration.
```

### Code / config change
`terraform/jenkins-server/setup.sh`, `terraform/jenkins-server/ec2.tf` (commit `67099b3`).

### Lessons learned
Adding a running service's user to a group needs a service restart.

### How to approach it next time
For a permission error on a socket or file: compare the **process's** groups (`/proc/<pid>/status`) with the file's owner and mode, not just `id <user>`.

### Prevention / follow-up
Done in `setup.sh`.

### References
- Docker: [Post-install: manage Docker as a non-root user](https://docs.docker.com/engine/install/linux-postinstall/).

---

## INC-033 – SonarQube scan: `HTTP connect timed out` (server URL was the public IP)

| | |
|---|---|
| **Date** | 2026-10-03 |
| **Where** | Jenkins `backend` #1 and `frontend` #1 › *Sonarqube Analysis*; Manage Jenkins › System › SonarQube servers |
| **Severity** | Medium: both app pipelines failed at the first scan |
| **Status** | Resolved |

### Summary
The SonarQube scanner couldn't reach SonarQube: `Failed to query server version … http://<JENKINS_IP>:9000 … HTTP connect timed out`. Jenkins and SonarQube run on the **same** EC2 instance, but Jenkins was configured with SonarQube's **public IP**. Port 9000's security group only allows the admin's laptop IP, and traffic from the instance to its own public IP comes from that public IP, which isn't on the list, so it was dropped. Changing the server URL to `http://localhost:9000` fixed it.

### Background
- Jenkins injects the server URL into the scanner via `withSonarQubeEnv('sonar-server')`.
- Security groups match the **source IP**. A connection to your own public IP (hairpin through the internet gateway) arrives with the public IP as its source.
- SonarQube's webhook back to Jenkins uses `http://172.17.0.1:8080/…` (Docker bridge) for the same reason.

### Timeline
1. `backend` #1 and `frontend` #1 failed at Sonarqube Analysis with the timeout.
2. On the Jenkins host: the config file had `<serverUrl>http://<JENKINS_IP>:9000</serverUrl>`, and `curl http://localhost:9000/api/system/status` returned `200`.
3. Changed the URL to `http://localhost:9000`; `backend` #2 got `ANALYSIS SUCCESSFUL`, and the quality gate `OK` came back via the webhook.


### What happened (symptom)

    ERROR Failed to query server version: Call to URL [http://<JENKINS_IP>:9000/api/v2/analysis/version] failed: HTTP connect timed out
    INFO  EXECUTION FAILURE

### How we got there
During setup the SonarQube server URL was entered as the browser address (public IP) instead of `localhost`.

### Why (root cause)
1. Scanner → `<JENKINS_IP>:9000` leaves the instance through the internet gateway.
2. It comes back in with source `<JENKINS_IP>`.
3. The security group on 9000 allows only `<YOUR_IP>/32`, so the packets are dropped and the connection times out.

### Impact
Two failed pipeline runs.

### Resolution (step by step)
1. Manage Jenkins → System → SonarQube servers → `sonar-server` → Server URL `http://localhost:9000` → Save.
2. Verify on the host: `grep -o '<serverUrl>[^<]*' /var/lib/jenkins/hudson.plugins.sonar.SonarGlobalConfiguration.xml` → `http://localhost:9000`.
3. Re-run the job → `ANALYSIS SUCCESSFUL`, `Quality gate is 'OK'`.

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it** (Jenkins `backend` #1 / `frontend` #1 console)
```console
+ /var/lib/jenkins/tools/hudson.plugins.sonar.SonarRunnerInstallation/sonar-scanner/bin/sonar-scanner \
    -Dsonar.projectName=go-grpc-graphql-micro-backend -Dsonar.projectKey=go-grpc-graphql-micro-backend \
    -Dsonar.sources=account,catalog,order,graphql -Dsonar.exclusions=**/pb/**,**/generated.go,**/models_gen.go ...
INFO  SonarScanner CLI 8.1.0.6389
ERROR Failed to query server version: Call to URL [http://<JENKINS_IP>:9000/api/v2/analysis/version] failed: HTTP connect timed out
INFO  EXECUTION FAILURE
```

**Fixing it / trying to fix it**
```console
# fix in the UI: Manage Jenkins → System → SonarQube servers → sonar-server → URL http://localhost:9000 → Save

# verify on the Jenkins host (laptop → SSM)
$ aws ssm send-command --region eu-north-1 --instance-ids <JENKINS_INSTANCE_ID> --document-name AWS-RunShellScript \
    --parameters 'commands=["grep -o \"<serverUrl>[^<]*</serverUrl>\" /var/lib/jenkins/hudson.plugins.sonar.SonarGlobalConfiguration.xml","curl -s -o /dev/null -w \"localhost:9000 -> %{http_code}\\n\" http://localhost:9000/api/system/status"]'
<serverUrl>http://localhost:9000</serverUrl>
localhost:9000 -> 200

# the re-run (backend #2) console
$ aws ssm send-command ... 'commands=["grep -E \"ANALYSIS SUCCESSFUL|status is|Quality gate\" /var/lib/jenkins/jobs/backend/builds/2/log"]'
INFO  ANALYSIS SUCCESSFUL, you can find the results at: http://localhost:9000/dashboard?id=go-grpc-graphql-micro-backend
SonarQube task '<task-id>' status is 'IN_PROGRESS'
SonarQube task '<task-id>' status is 'SUCCESS'
SonarQube task '<task-id>' completed. Quality gate is 'OK'
```

### Code / config change
Jenkins configuration only (not in git).

### Lessons learned
Services on the same host should talk over `localhost` (or the Docker bridge), not the public address.

### How to approach it next time
A **timeout** (not "connection refused") to an AWS IP usually means a security group is dropping the packets: check the source IP the target actually sees.

### Prevention / follow-up
Documented in the Recreate guide (Step 4.5).

### References
- AWS: [Security group rules](https://docs.aws.amazon.com/vpc/latest/userguide/security-group-rules.html).

---

## INC-032 – `ecr:CreateRepository` AccessDenied on the bastion

| | |
|---|---|
| **Date** | 2026-10-03 |
| **Where** | Bastion (SSM session); ECR step (`aws ecr create-repository` loop) |
| **Severity** | Low: wrong machine; expected behaviour |
| **Status** | Explained (ran from the laptop instead) |

### Summary
Creating the five ECR repositories from the **bastion** failed with `AccessDeniedException … ecr:CreateRepository`. That's by design: the bastion's role only allows describing the EKS cluster plus a narrow set of Load Balancer Controller setup actions. AWS setup like ECR belongs on the laptop (`devops-user`). Re-running the same loop there created all five repos (`IMMUTABLE`, `scanOnPush=true`).

### Background
| Machine | What it's for |
|---|---|
| Laptop (`devops-user`) | AWS setup: ECR, IAM, the Jenkins server's Terraform |
| Bastion | Only `kubectl` and `helm` against the private cluster (plus the scoped LBC IAM setup) |
| Jenkins | Builds, pushes to ECR, runs Terraform for EKS |

### Timeline
1. Ran the ECR loop in the bastion session that was still open from the previous step.
2. Got `AccessDeniedException` for `repository/order` (the first repo to fail).
3. Ran it on the laptop: 5 repository URIs printed; `describe-repositories` → all `IMMUTABLE`, `True`.


### What happened (symptom)

    An error occurred (AccessDeniedException) when calling the CreateRepository operation:
    User: arn:aws:sts::<ACCOUNT_ID>:assumed-role/go-microservices-dev-bastion-role/<INSTANCE_ID>
    is not authorized to perform: ecr:CreateRepository on resource: arn:aws:ecr:eu-north-1:<ACCOUNT_ID>:repository/order
    because no identity-based policy allows the ecr:CreateRepository action

### Why (root cause)
The command ran with the bastion's instance-role credentials, which intentionally lack ECR permissions.

### Impact
None.

### Resolution (step by step)
1. Check who you are: `aws sts get-caller-identity --query Arn --output text` → it must be `…:user/devops-user`, not `…assumed-role/…bastion-role/…`.
2. Run the loop on the laptop.
3. Verify: `aws ecr describe-repositories --query 'repositories[].[repositoryName,imageTagMutability,imageScanningConfiguration.scanOnPush]' --output table`.

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it** (bastion)
```console
[ec2-user@bastion ~]$ for repo in account catalog order graphql frontend; do
  aws ecr create-repository --region eu-north-1 --repository-name $repo \
    --image-scanning-configuration scanOnPush=true \
    --image-tag-mutability IMMUTABLE \
    --query 'repository.repositoryUri' --output text
done
An error occurred (AccessDeniedException) when calling the CreateRepository operation: User: arn:aws:sts::<ACCOUNT_ID>:assumed-role/go-microservices-dev-bastion-role/<INSTANCE_ID> is not authorized to perform: ecr:CreateRepository on resource: arn:aws:ecr:eu-north-1:<ACCOUNT_ID>:repository/order because no identity-based policy allows the ecr:CreateRepository action
```

**Fixing it / trying to fix it** (laptop)
```console
$ aws sts get-caller-identity --query Arn --output text
arn:aws:iam::<ACCOUNT_ID>:user/devops-user
$ for repo in account catalog order graphql frontend; do
  aws ecr create-repository --region eu-north-1 --repository-name $repo \
    --image-scanning-configuration scanOnPush=true \
    --image-tag-mutability IMMUTABLE \
    --query 'repository.repositoryUri' --output text
done
<ACCOUNT_ID>.dkr.ecr.eu-north-1.amazonaws.com/account      # expected, one line per repo
...
$ aws ecr describe-repositories --region eu-north-1 \
    --query 'repositories[].[repositoryName,imageTagMutability,imageScanningConfiguration.scanOnPush]' --output text | sort
account    IMMUTABLE  True
catalog    IMMUTABLE  True
frontend   IMMUTABLE  True
graphql    IMMUTABLE  True
order      IMMUTABLE  True
```

### Code / config change
None.

### Lessons learned
The error message names the identity (`assumed-role/…bastion-role`), which is the fastest clue to "wrong machine".

### How to approach it next time
On any `AccessDenied`, read the `User: arn:…` part first: wrong identity, or the right identity missing a permission?

### Prevention / follow-up
The recreate guide labels each command block *(laptop)* or *(bastion)*.

### References
- AWS: [Troubleshoot AccessDenied errors](https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_access-denied.html).

---

## INC-031 – Argo CD Service got a Classic Load Balancer, not one from the LB Controller

| | |
|---|---|
| **Date** | 2026-10-03 |
| **Where** | Bastion; `kubectl patch svc argocd-server -n argocd -p '{"spec":{"type":"LoadBalancer"}}'` |
| **Severity** | Low: worked, but not what the design expected |
| **Status** | Explained (kept; IP-restricted) |

### Summary
With the AWS Load Balancer Controller installed, the Argo CD UI was expected to get an NLB named `k8s-argocd-…`. It got a hostname like `<hash>-<id>.eu-north-1.elb.amazonaws.com`, which is the naming of a **Classic Load Balancer** created by the legacy in-tree AWS integration. The LB Controller's mutating webhook only claims Services **at creation**; this Service already existed as `ClusterIP` and was **patched** to `LoadBalancer`, so the legacy provider handled it. It works (internet-facing, restricted to our IP by `loadBalancerSourceRanges`) and is identical to what the reference setup produces, so it was kept.

### Background
- The LB Controller sets `spec.loadBalancerClass: service.k8s.aws/nlb` on **new** `type: LoadBalancer` Services through a webhook. Services without a class fall back to the cloud provider's built-in controller, which creates a CLB.
- `loadBalancerClass` is immutable once set, so you can't simply add it to an existing LoadBalancer Service.
- The `aws-load-balancer-scheme` annotation is an LB Controller annotation; the legacy provider ignores it and makes CLBs internet-facing by default.

### Timeline
1. Annotated the Service with scheme `internet-facing`, patched it to `LoadBalancer`, and added `loadBalancerSourceRanges`.
2. `kubectl get svc` showed an `a707…elb.amazonaws.com` hostname.
3. `aws elb describe-load-balancers` showed a Classic LB, `internet-facing`, with an SG allowing `<YOUR_IP>/32` on 80 and 443.
4. Waited for DNS; `curl -sk https://<ARGOCD_LB_HOST>/` → `200`, title "Argo CD".


### What happened (symptom)

    argocd-server   LoadBalancer   …   <hash>-<id>.eu-north-1.elb.amazonaws.com   80:31792/TCP,443:31566/TCP

### Why (root cause)
Patching an existing Service bypasses the controller's create-time webhook, so the legacy provider creates a CLB.

### Impact
None functionally. A CLB is the older generation, and its source restriction is enforced through the SG the provider creates.

### Resolution (step by step)
1. Accept it (as done), keeping `loadBalancerSourceRanges`.
2. Or, to get an NLB from the controller: delete the Service, then recreate it as `type: LoadBalancer` with `loadBalancerClass: service.k8s.aws/nlb` and the controller annotations (in a Helm install, set these as values).

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it** (bastion)
```console
$ kubectl create namespace argocd
namespace/argocd created                                                          # expected
$ kubectl apply -n argocd --server-side --force-conflicts \
    -f https://raw.githubusercontent.com/argoproj/argo-cd/v3.5.3/manifests/install.yaml
customresourcedefinition.apiextensions.k8s.io/applications.argoproj.io serverside-applied      # expected
customresourcedefinition.apiextensions.k8s.io/applicationsets.argoproj.io serverside-applied
customresourcedefinition.apiextensions.k8s.io/appprojects.argoproj.io serverside-applied
serviceaccount/argocd-application-controller serverside-applied
...
deployment.apps/argocd-server serverside-applied
statefulset.apps/argocd-application-controller serverside-applied
$ kubectl get pods -n argocd -w
NAME                                               READY   STATUS    RESTARTS   AGE     # expected
argocd-application-controller-0                    1/1     Running   0          70s
argocd-applicationset-controller-…                 1/1     Running   0          71s
argocd-dex-server-…                                1/1     Running   0          71s
argocd-notifications-controller-…                  1/1     Running   0          71s
argocd-redis-…                                     1/1     Running   0          71s
argocd-repo-server-…                               1/1     Running   0          71s
argocd-server-…                                    1/1     Running   0          71s
$ kubectl annotate svc argocd-server -n argocd service.beta.kubernetes.io/aws-load-balancer-scheme=internet-facing
service/argocd-server annotated                                                   # expected
$ kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
service/argocd-server patched                                                     # expected
$ kubectl patch svc argocd-server -n argocd -p '{"spec": {"loadBalancerSourceRanges": ["<YOUR_IP>/32"]}}'
service/argocd-server patched                                                     # expected
$ kubectl get svc argocd-server -n argocd
NAME            TYPE           CLUSTER-IP     EXTERNAL-IP                                  PORT(S)                      AGE
argocd-server   LoadBalancer   172.20.93.30   <hash>-<id>.eu-north-1.elb.amazonaws.com     80:31792/TCP,443:31566/TCP   82s
# expected k8s-argocd-… (LB Controller NLB); got the Classic-LB naming
```

**Fixing it / trying to fix it**
```console
# laptop: which kind of load balancer is it?
$ H=<ARGOCD_LB_HOST>
$ aws elb describe-load-balancers --region eu-north-1 \
    --query "LoadBalancerDescriptions[?DNSName=='$H'].[LoadBalancerName,Scheme,SecurityGroups[0]]" --output text
<hash>   internet-facing   sg-<id>                  # Classic LB
$ aws elbv2 describe-load-balancers --region eu-north-1 --query "LoadBalancers[?DNSName=='$H'].[LoadBalancerName,Type]" --output text
                                                     # nothing: not an ALB/NLB
$ aws ec2 describe-security-groups --region eu-north-1 --group-ids sg-<id> \
    --query 'SecurityGroups[0].IpPermissions[].[FromPort,IpRanges[].CidrIp]' --output json
[ [80, ["<YOUR_IP>/32"]], [3, ["<YOUR_IP>/32"]], [443, ["<YOUR_IP>/32"]] ]      # allowlist applied

# bastion: a quick way to confirm who reconciled it (not run at the time)
$ kubectl get svc argocd-server -n argocd -o jsonpath='{.spec.loadBalancerClass}'; echo
                                                     # empty → legacy in-tree provider, not the LB Controller

# kept as-is; wait for DNS and test
$ until getent hosts $H >/dev/null; do sleep 10; done
(no output: returns once the name resolves)
$ until curl -sk -o /dev/null --max-time 8 -w '%{http_code}' https://$H/ | grep -q 200; do sleep 10; done
(no output: returns once HTTPS answers 200)
$ curl -sk https://$H/ | grep -o '<title>[^<]*</title>'
<title>Argo CD</title>
$ kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo
<argocd-admin-password>                                                           # redacted
$ kubectl -n argocd delete secret argocd-initial-admin-secret       # after changing the password in the UI
secret "argocd-initial-admin-secret" deleted                                      # expected
```

### Code / config change
None.

### Lessons learned
"Which controller reconciled this?" decides the LB type: check the hostname format and `spec.loadBalancerClass`.

### How to approach it next time
`kubectl get svc <name> -o jsonpath='{.spec.loadBalancerClass}'`: empty means legacy provider (CLB); `service.k8s.aws/nlb` means the LB Controller (NLB).

### Prevention / follow-up
For new exposures, create the Service with its type and class from the start (Helm values) instead of patching.

### References
- AWS LB Controller: [Service annotations / `loadBalancerClass`](https://kubernetes-sigs.github.io/aws-load-balancer-controller/latest/guide/service/nlb/).

---

## INC-030 – New ALB: `curl: (6) Could not resolve host` (and `DNS_PROBE_POSSIBLE` in the browser)

| | |
|---|---|
| **Date** | 2026-10-03 |
| **Where** | Laptop `curl` / browser → the test ALB (LB Controller smoke test) and later the app's ALB |
| **Severity** | Low: timing; nothing was broken |
| **Status** | Explained |

### Summary
Right after the controller reported `Successfully reconciled`, `curl` to the new ALB's hostname failed with `Could not resolve host`. AWS showed the ALB in state **`provisioning`**: its DNS name doesn't exist until it becomes `active` (about 2–3 minutes). Once active it returned `200 OK`. The same thing happened with the app's ALB in the browser (`DNS_PROBE_POSSIBLE`), where the browser also **cached** the failed lookup and kept failing after `curl` already worked. This is the same pattern as INC-009 on the rebuild.

### Background
- An Ingress gets its ADDRESS (the ALB's DNS name) as soon as AWS creates the ALB, before the ALB is active and before the DNS record is published.
- Browsers and OS resolvers cache negative answers (NXDOMAIN) for a short time.

### Timeline
1. `kubectl -n alb-test describe ingress echo` → `SuccessfullyReconciled` (17 s old).
2. `curl -I http://<ALB_HOST>` → `Could not resolve host`.
3. `aws elbv2 describe-load-balancers` → `provisioning`.
4. A loop waited for `active`, then for HTTP 200 → `HTTP/1.1 200 OK` about 3 minutes after creation.
5. Later, the app's ALB: `curl` worked, but the browser showed `DNS_PROBE_POSSIBLE` until its DNS cache was cleared.


### What happened (symptom)

    curl: (6) Could not resolve host: k8s-albtest-echo-<id>.eu-north-1.elb.amazonaws.com

### Why (root cause)
1. The ALB was still `provisioning`, so its DNS name wasn't published yet.
2. The first lookups returned "not found", and the browser cached that answer.

### Impact
A few minutes of waiting.

### Resolution (step by step)
1. Check the state: `aws elbv2 describe-load-balancers --query "LoadBalancers[?DNSName=='<LB_HOST>'].State.Code"` → wait for `active`.
2. Then: `getent hosts <LB_HOST>` resolves → `curl -I http://<LB_HOST>/` → `200`.
3. Browser: `chrome://net-internals/#dns` → *Clear host cache*, or use an incognito window.

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
# bastion: smoke-test Ingress → ALB
$ kubectl create namespace alb-test
namespace/alb-test created                                                        # expected
$ kubectl -n alb-test create deployment echo --image=nginxinc/nginx-unprivileged:stable-alpine --port=8080
deployment.apps/echo created                                                      # expected
$ kubectl -n alb-test expose deployment echo --port=80 --target-port=8080
service/echo exposed                                                              # expected
$ kubectl -n alb-test create ingress echo --class=alb --rule="/*=echo:80" \
    --annotation alb.ingress.kubernetes.io/scheme=internet-facing \
    --annotation alb.ingress.kubernetes.io/target-type=ip
ingress.networking.k8s.io/echo created                                            # expected
$ kubectl -n alb-test get ingress echo -w
echo   alb     *       k8s-albtest-echo-<id>.eu-north-1.elb.amazonaws.com   80      6s
$ kubectl -n alb-test describe ingress echo | tail -5
  Normal  SuccessfullyReconciled  17s   ingress  Successfully reconciled

# laptop, about 1 minute later
$ curl -I http://k8s-albtest-echo-<id>.eu-north-1.elb.amazonaws.com
curl: (6) Could not resolve host: k8s-albtest-echo-<id>.eu-north-1.elb.amazonaws.com
```

**Fixing it / trying to fix it**
```console
$ H=k8s-albtest-echo-<id>.eu-north-1.elb.amazonaws.com
$ aws elbv2 describe-load-balancers --region eu-north-1 \
    --query "LoadBalancers[?DNSName=='$H'].[State.Code,Scheme,Type]" --output text
provisioning   internet-facing   application
$ getent hosts $H || echo "DNS not resolving yet"
DNS not resolving yet
$ until [ "$(aws elbv2 describe-load-balancers --region eu-north-1 --query "LoadBalancers[?DNSName=='$H'].State.Code" --output text)" = active ]; do sleep 15; done; echo "ALB active"
ALB active
$ until curl -sI --max-time 8 http://$H/ | grep -q '200 OK'; do sleep 10; done; curl -sI http://$H/ | head -1
HTTP/1.1 200 OK
$ kubectl delete namespace alb-test       # bastion: removes the test ALB
namespace "alb-test" deleted                                                      # expected

# same thing later with the app ALB in the browser (DNS_PROBE_POSSIBLE) while curl already worked
$ H=<APP_LB_HOST>
$ aws elbv2 describe-load-balancers --region eu-north-1 --query "LoadBalancers[?DNSName=='$H'].[State.Code,Scheme,CreatedTime]" --output text
active   internet-facing   ...
$ getent hosts $H
<LB_IP_1>   <APP_LB_HOST>
<LB_IP_2>   <APP_LB_HOST>
$ curl -s -o /dev/null -w "/ -> %{http_code}\n" http://$H/
/ -> 200
$ resolvectl flush-caches        # browser: chrome://net-internals/#dns → Clear host cache (or incognito)
(no output)
```

### Code / config change
None.

### Lessons learned
"The controller is done" isn't the same as "AWS is done". Check the ALB's state.

### How to approach it next time
`Could not resolve host` on a brand-new `*.elb.amazonaws.com` name → check the LB state before debugging anything else.

### Prevention / follow-up
Docs now say to wait 2–3 minutes after the ADDRESS appears.

### References
- AWS: [Application Load Balancer states](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_LoadBalancerState.html).
- See INC-009 (same pattern, Argo CD, rebuild).

---

## INC-029 – `aws iam create-role`: `argument --role-name: expected one argument`

| | |
|---|---|
| **Date** | 2026-10-03 |
| **Where** | Bastion; LB Controller setup step 3 |
| **Severity** | Low |
| **Status** | Resolved |

### Summary
`aws iam create-role --role-name ${CLUSTER_NAME}-lb-controller-role …` failed with `argument --role-name: expected one argument`. The shell variables from step 1 were **empty**, because a new shell had been started (`sudo su - ec2-user` / a new SSM session). `${CLUSTER_NAME}-lb-controller-role` became `-lb-controller-role`, and since it starts with `-`, the CLI parsed it as an **option**, leaving `--role-name` without a value. Re-exporting the variables (kept in `~/lbc/env.sh`) fixed it.

### Background
- `export` only lives in the current shell. `sudo su - user` starts a fresh **login** shell without them.
- Unset variables expand to an empty string, with no error unless `set -u` is on.
- CLIs treat any word starting with `-` as a flag.

### Timeline
1. Step 1 exports were done in an earlier shell.
2. After `sudo su - ec2-user`, ran step 3 → the error.
3. `echo "cluster=$CLUSTER_NAME …"` → all empty.
4. Re-exported; created `~/lbc/env.sh` to `source` after every login; step 3 succeeded, and `list-attached-role-policies` showed the policy.


### What happened (symptom)

    aws: [ERROR]: An error occurred (ParamValidation): argument --role-name: expected one argument

### Why (root cause)
Empty `$CLUSTER_NAME` → the value starts with `-` → it's parsed as an option.

### Impact
None.

### Resolution (step by step)
1. `echo "cluster=$CLUSTER_NAME region=$AWS_REGION account=$ACCOUNT_ID vpc=$VPC_ID"`: all must be set.
2. Save the exports once:
   ```bash
   cat > ~/lbc/env.sh <<'EOF'
   export CLUSTER_NAME=go-microservices-dev
   export AWS_REGION=eu-north-1
   export ACCOUNT_ID=$(aws sts get-caller-identity --query Account --output text)
   export VPC_ID=$(aws eks describe-cluster --name $CLUSTER_NAME --region $AWS_REGION --query cluster.resourcesVpcConfig.vpcId --output text)
   EOF
   source ~/lbc/env.sh
   ```
3. Re-run step 3; verify with `aws iam list-attached-role-policies --role-name ${CLUSTER_NAME}-lb-controller-role`.

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it** (bastion, after `sudo su - ec2-user` started a fresh shell)
```console
[ec2-user@bastion lbc]$ aws iam create-role --role-name ${CLUSTER_NAME}-lb-controller-role \
  --assume-role-policy-document file://trust.json
aws: [ERROR]: An error occurred (ParamValidation): argument --role-name: expected one argument

usage: aws [options] <command> <subcommand> [<subcommand> ...] [parameters]
[ec2-user@bastion lbc]$ aws iam attach-role-policy --role-name ${CLUSTER_NAME}-lb-controller-role \
  --policy-arn arn:aws:iam::${ACCOUNT_ID}:policy/AWSLoadBalancerControllerIAMPolicy
aws: [ERROR]: An error occurred (ParamValidation): argument --role-name: expected one argument
```

**Fixing it / trying to fix it**
```console
[ec2-user@bastion lbc]$ echo "cluster=$CLUSTER_NAME region=$AWS_REGION account=$ACCOUNT_ID vpc=$VPC_ID"
cluster= region= account= vpc=                        # all empty → "-lb-controller-role" parsed as a flag

# keep the variables in a file and source it after every login
[ec2-user@bastion lbc]$ cat > ~/lbc/env.sh <<'EOF'
export CLUSTER_NAME=go-microservices-dev
export AWS_REGION=eu-north-1
export ACCOUNT_ID=$(aws sts get-caller-identity --query Account --output text)
export VPC_ID=$(aws eks describe-cluster --name $CLUSTER_NAME --region $AWS_REGION --query cluster.resourcesVpcConfig.vpcId --output text)
EOF
(no output)
[ec2-user@bastion lbc]$ source ~/lbc/env.sh
(no output)
[ec2-user@bastion lbc]$ echo "cluster=$CLUSTER_NAME region=$AWS_REGION account=$ACCOUNT_ID vpc=$VPC_ID"
cluster=go-microservices-dev region=eu-north-1 account=<ACCOUNT_ID> vpc=vpc-<id>     # expected

[ec2-user@bastion lbc]$ ls trust.json
trust.json
[ec2-user@bastion lbc]$ aws iam create-role --role-name ${CLUSTER_NAME}-lb-controller-role \
  --assume-role-policy-document file://trust.json
{                                                                                 # expected
    "Role": {
        "Path": "/",
        "RoleName": "go-microservices-dev-lb-controller-role",
        "Arn": "arn:aws:iam::<ACCOUNT_ID>:role/go-microservices-dev-lb-controller-role",
        "AssumeRolePolicyDocument": { ... "Service": "pods.eks.amazonaws.com" ... }
    }
}
[ec2-user@bastion lbc]$ aws iam attach-role-policy --role-name ${CLUSTER_NAME}-lb-controller-role \
  --policy-arn arn:aws:iam::${ACCOUNT_ID}:policy/AWSLoadBalancerControllerIAMPolicy
(no output)
[ec2-user@bastion lbc]$ aws iam list-attached-role-policies --role-name ${CLUSTER_NAME}-lb-controller-role
    "PolicyName": "AWSLoadBalancerControllerIAMPolicy"            # expected

# next steps then worked with the same variables
[ec2-user@bastion lbc]$ aws eks create-pod-identity-association --region $AWS_REGION --cluster-name $CLUSTER_NAME \
  --namespace kube-system --service-account aws-load-balancer-controller \
  --role-arn arn:aws:iam::${ACCOUNT_ID}:role/${CLUSTER_NAME}-lb-controller-role
{                                                                                 # expected
    "association": {
        "clusterName": "go-microservices-dev",
        "namespace": "kube-system",
        "serviceAccount": "aws-load-balancer-controller",
        "roleArn": "arn:aws:iam::<ACCOUNT_ID>:role/go-microservices-dev-lb-controller-role",
        "associationId": "a-<id>"
    }
}
[ec2-user@bastion lbc]$ helm install aws-load-balancer-controller eks/aws-load-balancer-controller -n kube-system --version 3.5.0 \
  --set clusterName=$CLUSTER_NAME --set serviceAccount.create=true --set serviceAccount.name=aws-load-balancer-controller \
  --set region=$AWS_REGION --set vpcId=$VPC_ID
NAME: aws-load-balancer-controller                                                # expected
NAMESPACE: kube-system
STATUS: deployed
REVISION: 1
NOTES:
AWS Load Balancer controller installed!
[ec2-user@bastion lbc]$ kubectl get pod -n kube-system -l app.kubernetes.io/name=aws-load-balancer-controller \
  -o jsonpath='{.items[0].spec.containers[0].env[*].name}'; echo
AWS_STS_REGIONAL_ENDPOINTS AWS_DEFAULT_REGION AWS_REGION AWS_CONTAINER_CREDENTIALS_FULL_URI AWS_CONTAINER_AUTHORIZATION_TOKEN_FILE
```

### Code / config change
None (procedure).

### Lessons learned
Echo the variables before commands that depend on them; `set -u` turns empty variables into errors.

### How to approach it next time
`expected one argument` right after a variable usually means the variable is empty.

### Prevention / follow-up
The recreate guide creates `env.sh` up front.

### References
- Bash: [`set -u`](https://www.gnu.org/software/bash/manual/html_node/The-Set-Builtin.html).

---

## INC-028 – Bastion `AccessDenied iam:CreatePolicy`: permission in git, but not applied yet

| | |
|---|---|
| **Date** | 2026-10-03 |
| **Where** | Bastion; LB Controller setup step 2; `terraform/eks-cluster/bastion.tf` |
| **Severity** | Low |
| **Status** | Resolved |

### Summary
`aws iam create-policy` on the bastion was denied. The scoped bastion policy (`…-bastion-lb-controller-setup`) had been written in Terraform and **pushed**, but the Jenkins `eks-cluster` **apply** hadn't run yet, so the role only had its original policy. After the apply (`1 to add`) and a short wait for IAM propagation, the command worked.

### Background
- Code in git changes nothing in AWS until Terraform **applies** it.
- IAM changes can take 10–30 seconds to take effect.

### Timeline
1. Added `aws_iam_role_policy.bastion_lb_controller_setup`; committed and pushed.
2. On the bastion: `create-policy` → AccessDenied.
3. `aws iam list-role-policies --role-name go-microservices-dev-bastion-role` (laptop) → only `…-bastion-eks`.
4. Ran Jenkins `eks-cluster` with `dev` + `apply` → `1 to add`; retried → policy created.


### What happened (symptom)

    An error occurred (AccessDenied) when calling the CreatePolicy operation:
    User: arn:aws:sts::<ACCOUNT_ID>:assumed-role/go-microservices-dev-bastion-role/<INSTANCE_ID>
    is not authorized to perform: iam:CreatePolicy ... because no identity-based policy allows the iam:CreatePolicy action.

### Why (root cause)
The desired state was in git, but the actual state in AWS was unchanged: the apply step was skipped.

### Impact
None.

### Resolution (step by step)
1. Check the actual state: `aws iam list-role-policies --role-name <role>`.
2. Run the Terraform apply (Jenkins `eks-cluster`, `dev` + `apply`); check that the plan says `1 to add`.
3. Wait about 30 seconds and retry.

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it** (bastion)
```console
[ec2-user@bastion lbc]$ aws iam create-policy --policy-name AWSLoadBalancerControllerIAMPolicy \
  --policy-document file://iam_policy.json
aws: [ERROR]: An error occurred (AccessDenied) when calling the CreatePolicy operation: User: arn:aws:sts::<ACCOUNT_ID>:assumed-role/go-microservices-dev-bastion-role/<INSTANCE_ID> is not authorized to perform: iam:CreatePolicy on resource: policy AWSLoadBalancerControllerIAMPolicy because no identity-based policy allows the iam:CreatePolicy action.
```

**Fixing it / trying to fix it**
```console
# laptop: what is actually attached to the bastion role right now?
$ aws iam list-role-policies --role-name go-microservices-dev-bastion-role --output text
POLICYNAMES     go-microservices-dev-bastion-eks           # the new …-lb-controller-setup policy is missing
$ git fetch -q; git log origin/main --format='%h %s' -2
99580b7 feat: allow bastion to set up AWS Load Balancer Controller IAM     # pushed, but not applied
f8a540a feat: add dev/prod EKS environments with private API via bastion

# what the apply will do (local plan)
$ cd terraform/eks-cluster
$ terraform init -reconfigure -input=false -backend-config=backend.hcl -backend-config="key=eks-cluster/dev/terraform.tfstate" >/dev/null
(no output: redirected)
$ terraform plan -input=false -var-file=envs/dev.tfvars | grep -E '^\s+# |Plan:'
  # aws_iam_role_policy.bastion_lb_controller_setup will be created
Plan: 1 to add, 0 to change, 0 to destroy.

# Jenkins: eks-cluster → Build with Parameters → dev + apply → Proceed → "Apply complete! Resources: 1 added"
$ aws iam list-role-policies --role-name go-microservices-dev-bastion-role --output text
POLICYNAMES     go-microservices-dev-bastion-eks                    # expected
POLICYNAMES     go-microservices-dev-bastion-lb-controller-setup    # expected

# bastion: retry after ~30 s (IAM propagation)
[ec2-user@bastion lbc]$ aws iam create-policy --policy-name AWSLoadBalancerControllerIAMPolicy \
  --policy-document file://iam_policy.json
    "PolicyName": "AWSLoadBalancerControllerIAMPolicy",              # expected (policy JSON)
```

### Code / config change
`terraform/eks-cluster/bastion.tf` (commit `99580b7`).

### Lessons learned
Check what's live in AWS, not what's in the repo.

### How to approach it next time
AccessDenied right after adding a permission → (1) was it applied? (2) did IAM propagate? (3) does the resource ARN in the policy match the one in the error?

### Prevention / follow-up
List "run the apply" as an explicit step after every infra commit.

### References
- AWS: [IAM eventual consistency](https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency).

---

## INC-027 – Bastion: `curl: (23) … Permission denied` and `Unable to load paramfile`

| | |
|---|---|
| **Date** | 2026-10-03 |
| **Where** | Bastion (SSM session); LB Controller setup step 2 (download `iam_policy.json`) |
| **Severity** | Low |
| **Status** | Resolved |

### Summary
`curl -O …/iam_policy.json` failed with `Failed to open the file iam_policy.json: Permission denied` / `curl: (23)`, and the next command failed because the file didn't exist. An SSM session starts as **`ssm-user` in a system directory** (`/usr/bin` or `/`), where it can't write. After switching to `ec2-user` and working in `~/lbc`, the download and the `create-policy` worked.

### Background
- `curl -O` saves to the **current directory**.
- SSM Session Manager shells start as `ssm-user`, often in `/usr/bin`.
- `file://iam_policy.json` is relative to the current directory too.

### Timeline
1. Opened an SSM session and ran `curl -O` → Permission denied.
2. `create-policy … file://iam_policy.json` → `No such file or directory`.
3. `whoami; pwd` → `ssm-user`, `/usr/bin`.
4. `sudo su - ec2-user; mkdir -p ~/lbc && cd ~/lbc` → download about 9 KB, then `create-policy` OK (after INC-028's apply).


### What happened (symptom)

    Warning: Failed to open the file iam_policy.json: Permission denied
    curl: (23) client returned ERROR on write of 797 bytes
    aws: [ERROR]: ... Unable to load paramfile file://iam_policy.json: [Errno 2] No such file or directory

### Why (root cause)
The working directory wasn't writable by the session user.

### Impact
None.

### Resolution (step by step)
1. `whoami; pwd`.
2. `sudo su - ec2-user` → `mkdir -p ~/lbc && cd ~/lbc`.
3. Re-run `curl -O …` → `ls -l iam_policy.json` (~9 KB).
4. Re-export the variables (see INC-029).

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it** (fresh SSM session on the bastion)
```console
$ aws ssm start-session --region eu-north-1 --target <BASTION_INSTANCE_ID>

Starting session with SessionId: devops-user-<id>                                 # expected
sh-5.2$ curl -O https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/v3.5.0/docs/install/iam_policy.json
Warning: Failed to open the file iam_policy.json: Permission denied
curl: (23) client returned ERROR on write of 797 bytes
sh-5.2$ aws iam create-policy --policy-name AWSLoadBalancerControllerIAMPolicy \
  --policy-document file://iam_policy.json
aws: [ERROR]: An error occurred (ParamValidation): Error parsing parameter '--policy-document': Unable to load paramfile file://iam_policy.json: [Errno 2] No such file or directory: 'iam_policy.json'
```

**Fixing it / trying to fix it**
```console
sh-5.2$ whoami; pwd
ssm-user                                               # expected
/usr/bin                                               # expected: SSM start directory, not writable
sh-5.2$ sudo su - ec2-user
(no output: new login shell as ec2-user)
[ec2-user@bastion ~]$ mkdir -p ~/lbc && cd ~/lbc
(no output)
[ec2-user@bastion lbc]$ pwd
/home/ec2-user/lbc
[ec2-user@bastion lbc]$ curl -O https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/v3.5.0/docs/install/iam_policy.json
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current     # expected
100  8955  100  8955    0     0  ...
[ec2-user@bastion lbc]$ ls -l iam_policy.json
-rw-r--r-- 1 ec2-user ec2-user 8955 ... iam_policy.json        # expected (~9 KB)
[ec2-user@bastion lbc]$ aws iam create-policy --policy-name AWSLoadBalancerControllerIAMPolicy \
  --policy-document file://iam_policy.json
aws: [ERROR]: An error occurred (AccessDenied) when calling the CreatePolicy operation: User: arn:aws:sts::<ACCOUNT_ID>:assumed-role/go-microservices-dev-bastion-role/<INSTANCE_ID> is not authorized to perform: iam:CreatePolicy ...
# → file problem fixed; the AccessDenied is the next incident (INC-028)
```

### Code / config change
None.

### Lessons learned
Check where you are before downloading anything; SSM drops you somewhere unusual.

### How to approach it next time
`curl (23)` = write error, so check the directory permissions and disk space.

### Prevention / follow-up
The guide starts every bastion block with `sudo su - ec2-user` and `cd ~/lbc`.

### References
- curl: [exit code 23](https://everything.curl.dev/cmdline/exitcode).

---

## INC-026 – Jenkins `eks-cluster` #1: `ERROR: AWS_ACCOUNT_ID`

| | |
|---|---|
| **Date** | 2026-10-03 |
| **Where** | Jenkins job `eks-cluster`, first build; Manage Jenkins › Credentials |
| **Severity** | Low |
| **Status** | Resolved |

### Summary
The first build failed right after checkout with only `ERROR: AWS_ACCOUNT_ID`. The pipeline's `environment { AWS_ACCOUNT_ID = credentials('AWS_ACCOUNT_ID') }` couldn't find the credential, because it had been created in the **user's** credential store (click your name → Credentials), which pipelines can't use. Recreating it under **System → Global** fixed it.

### Background
- Jenkins has a **System** store (Global domain), usable by jobs, and per-**user** stores, usable only in that user's context.
- Credential binding in `environment {}` happens before the first stage, so a missing ID fails the build immediately with just the ID name.
- The `properties([parameters(...)])` call had already run, which is why the job showed *Build with Parameters* afterwards.

### Timeline
1. Created the `eks-cluster` job → Build Now.
2. Console: checkout OK → `ERROR: AWS_ACCOUNT_ID` → `FAILURE`.
3. The user had created their own Jenkins user and stored the credential under it.
4. Recreated it as Secret text, ID `AWS_ACCOUNT_ID`, under System → Global → `plan` OK → `apply` OK.


### What happened (symptom)

    [Pipeline] withEnv
    [Pipeline] End of Pipeline
    ERROR: AWS_ACCOUNT_ID
    Finished: FAILURE

### Why (root cause)
The credential ID was not in any store the pipeline can read.

### Impact
One failed build.

### Resolution (step by step)
1. Manage Jenkins → Credentials → **System** → **Global credentials (unrestricted)** → Add.
2. Kind *Secret text*, Scope *Global*, ID exactly `AWS_ACCOUNT_ID`.
3. Check that the list shows Store = System.
4. Re-run.

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it** (Jenkins `eks-cluster` → Build Now → Console Output)
```console
Started by user Ritesh Karankal
Obtained terraform/eks-cluster/Jenkinsfile from git https://github.com/<user>/go-grpc-graphql-micro.git
[Pipeline] Start of Pipeline
[Pipeline] properties
[Pipeline] node
Running on Jenkins in /var/lib/jenkins/workspace/eks-cluster
[Pipeline] { (Declarative: Checkout SCM)
 > git fetch --tags --force --progress -- https://github.com/<user>/go-grpc-graphql-micro.git +refs/heads/*:refs/remotes/origin/*
Checking out Revision <commit> (refs/remotes/origin/main)
Commit message: "feat: harden EKS with dev/prod environments and private API via bastion"
[Pipeline] withEnv
[Pipeline] End of Pipeline
ERROR: AWS_ACCOUNT_ID
Finished: FAILURE
```
The credential had been added under **(your user) → Credentials**, not System.

**Fixing it / trying to fix it** (Jenkins UI; no shell commands)
```text
Manage Jenkins → Credentials → Stores scoped to Jenkins: System → Global credentials (unrestricted) → + Add Credentials
  Kind:   Secret text
  Scope:  Global
  Secret: <ACCOUNT_ID>
  ID:     AWS_ACCOUNT_ID
→ Create
(optional) delete the copy under <your user> → Credentials
eks-cluster → Build with Parameters → Environment=dev, Terraform_Action=plan → Build
  Init: Terraform has been successfully initialized!
  Plan: 46 to add, 0 to change, 0 to destroy.          (expected)
eks-cluster → Build with Parameters → dev + apply → Approve: Proceed
```

### Code / config change
None.

### Lessons learned
A Jenkins error that's just a credential ID means "credential not found".

### How to approach it next time
Check the **Store** column in Manage Jenkins → Credentials; check the exact ID spelling (case-sensitive).

### Prevention / follow-up
Same for `sonar-token`, `github` and `nvd-api-key`: all under System → Global.

### References
- Jenkins: [Using credentials](https://www.jenkins.io/doc/book/using/using-credentials/).

---

## INC-025 – Terraform: `Invalid for_each argument` for EKS access entries

| | |
|---|---|
| **Date** | 2026-10-03 |
| **Where** | `terraform/eks-cluster/eks.tf` (`aws_eks_access_entry.admins`), `locals.tf` |
| **Severity** | Low: plan-time error, nothing created |
| **Status** | Resolved |

### Summary
After making the bastion the cluster admin, `terraform plan` failed with `Invalid for_each argument`. `for_each = toset(local.cluster_admin_arns)` included the bastion role's ARN, which is unknown until the role is created, and Terraform must know every `for_each` **key** at plan time. Switching to a map with **static keys** (`{ bastion = aws_iam_role.bastion.arn, "user-<name>" = … }`) fixed it: the keys are known, and only the values are unknown. Plan: `46 to add`.

### Background
- `for_each` keys become resource addresses (`aws_eks_access_entry.admins["bastion"]`), so Terraform needs them at plan time.
- A set's elements **are** its keys; a map's keys can be literals while its values are computed.

### Timeline
1. Replaced the Jenkins-role access entry with the bastion role.
2. Plan → the error (`local.cluster_admin_arns is tuple with 1 element`).
3. Rewrote it as a map; plan OK.


### What happened (symptom)

    Error: Invalid for_each argument
      on eks.tf line 100, in resource "aws_eks_access_entry" "admins":
     100:   for_each      = toset(local.cluster_admin_arns)
    The "for_each" set includes values derived from resource attributes that cannot be determined until apply

### Why (root cause)
Unknown values were used as `for_each` keys.

### Impact
None.

### Resolution (step by step)
```hcl
locals {
  cluster_admins = merge(
    { bastion = aws_iam_role.bastion.arn },
    { for u in var.admin_user_names : "user-${u}" => "arn:aws:iam::${local.account_id}:user/${u}" },
  )
}
resource "aws_eks_access_entry" "admins" {
  for_each      = local.cluster_admins
  principal_arn = each.value
  ...
}
```
Verify: `terraform plan -var-file=envs/dev.tfvars` → `aws_eks_access_entry.admins["bastion"] will be created`.

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it**
```console
$ cd terraform/eks-cluster
$ terraform fmt && terraform init -backend=false -input=false >/dev/null && terraform validate
Success! The configuration is valid.                  # validate can't see this error
$ terraform init -reconfigure -input=false -backend-config=backend.hcl -backend-config="key=eks-cluster/dev/terraform.tfstate" >/dev/null
(no output: redirected)
$ terraform plan -input=false -no-color -var-file=envs/dev.tfvars
  # aws_iam_role.bastion will be created
  # aws_instance.bastion will be created
  ...
Error: Invalid for_each argument

  on eks.tf line 100, in resource "aws_eks_access_entry" "admins":
 100:   for_each      = toset(local.cluster_admin_arns)
    ├────────────────
    │ local.cluster_admin_arns is tuple with 1 element

The "for_each" set includes values derived from resource attributes that
cannot be determined until apply, and so Terraform cannot determine the full
set of keys that will identify the instances of this resource.
```

**Fixing it / trying to fix it**
```console
# locals.tf: list of ARNs → map with static keys; eks.tf: for_each = local.cluster_admins
$ grep -n "cluster_admin" *.tf
eks.tf:100:  for_each      = local.cluster_admins
eks.tf:106:  for_each      = local.cluster_admins
locals.tf:19:  cluster_admins = merge(
$ terraform fmt
(no output: nothing to reformat)
$ terraform plan -input=false -no-color -var-file=envs/dev.tfvars | grep -E 'Plan:|Error|admins\['
  # aws_eks_access_entry.admins["bastion"] will be created
  # aws_eks_access_policy_association.admins["bastion"] will be created
Plan: 46 to add, 0 to change, 0 to destroy.
```

### Code / config change
`terraform/eks-cluster/locals.tf`, `eks.tf` (commit `f8a540a`).

### Lessons learned
Use maps with fixed keys whenever `for_each` values come from other resources.

### How to approach it next time
The error message names the expression; ask "which part of it is computed?" and move that part into the map's values.

### Prevention / follow-up
None needed.

### References
- Terraform: [`for_each` — limitations on values](https://developer.hashicorp.com/terraform/language/meta-arguments/for_each#limitations-on-values-used-in-for_each).

---

## INC-024 – Planned EKS version 1.30 and `AL2_x86_64` nodes no longer available

| | |
|---|---|
| **Date** | 2026-10-03 |
| **Where** | `terraform/eks-cluster/eks.tf` (`version = "1.30"`), `node_group.tf` (`ami_type = "AL2_x86_64"`) |
| **Severity** | Medium: the first apply would have failed, or landed on paid extended support |
| **Status** | Resolved (caught in review before the first apply) |

### Summary
The original EKS Terraform pinned **Kubernetes 1.30** and **Amazon Linux 2** nodes. A pre-apply check of supported versions in eu-north-1 showed 1.30 is **no longer offered**, 1.31–1.33 are only on **extended support** (an extra ~$0.60/hour per cluster), and AL2 node AMIs don't exist for 1.33+. The cluster moved to **1.36** with `AL2023_x86_64_STANDARD` nodes and `upgrade_policy { support_type = "STANDARD" }`, so it can never slide into extended-support billing.

### Background
- EKS supports each Kubernetes minor version for about 14 months in standard support, then extended support at a higher price, then removes it.
- Amazon Linux 2 node AMIs are not built for Kubernetes 1.33 and later; AL2023 replaces them.
- `upgrade_policy.support_type = STANDARD` makes EKS auto-upgrade at end of standard support instead of charging for extended support.

### Timeline
1. Reviewed the inherited `eks.tf` before the first apply.
2. Ran `aws eks describe-cluster-versions` (output below).
3. Changed the version to a variable (1.36), switched the AMI type, and added the upgrade policy; the plan was clean.


### What happened (symptom)
(Would have happened.) `terraform apply` would fail creating the cluster with an unsupported-version error, or the node group with an invalid AMI type for the version.

### Why (root cause)
The configuration came from an older template; the platform had moved on.

### Impact
None (caught before apply).

### Resolution (step by step)
1. `aws eks describe-cluster-versions --region eu-north-1 --query 'clusterVersions[].[clusterVersion,versionStatus]' --output text`
2. `variable "cluster_version" { default = "1.36" }`; `ami_type = "AL2023_x86_64_STANDARD"`; `upgrade_policy { support_type = "STANDARD" }`.
3. `terraform plan` → clean.

### Terminal commands
Commands as run in the session, each followed by its output. `$` lines are commands; the machine is noted in comments (laptop, bastion, Jenkins host, Jenkins UI). `(no output)` marks silent commands, `# expected` marks output that wasn't captured in the session log (typical output of that command), and `# suggested, not run` marks recommended steps that weren't executed. Hostnames, IPs, instance and account IDs are redacted.

**Finding it** (pre-apply review)
```console
$ for f in terraform/eks-cluster/*.tf; do echo "== $f"; cat $f; done
== terraform/eks-cluster/eks.tf
...
  version  = "1.30"
...
== terraform/eks-cluster/node_group.tf
...
  ami_type       = "AL2_x86_64"
$ aws eks describe-cluster-versions --region eu-north-1 \
    --query 'clusterVersions[].[clusterVersion,versionStatus,endOfStandardSupportDate]' --output text
1.37    STANDARD_SUPPORT    2027-12-01T05:30:00+05:30
1.36    STANDARD_SUPPORT    2027-08-02T05:30:00+05:30
1.35    STANDARD_SUPPORT    2027-03-27T05:30:00+05:30
1.34    STANDARD_SUPPORT    2026-12-02T05:30:00+05:30
1.33    EXTENDED_SUPPORT    2026-07-29T05:30:00+05:30
1.32    EXTENDED_SUPPORT    2026-03-23T05:30:00+05:30
1.31    EXTENDED_SUPPORT    2025-11-26T05:30:00+05:30
# 1.30 is not offered at all
$ aws eks describe-addon-versions --region eu-north-1 --addon-name aws-ebs-csi-driver \
    --query 'addons[0].addonVersions[0].addonVersion' --output text
v1.66.0-eksbuild.1
```

**Fixing it / trying to fix it**
```console
# eks.tf: version = var.cluster_version (1.36), upgrade_policy { support_type = "STANDARD" }
# node_group.tf: ami_type = "AL2023_x86_64_STANDARD"
$ cd terraform/eks-cluster
$ terraform fmt && terraform init -backend=false -input=false >/dev/null && terraform validate
Success! The configuration is valid.
$ terraform init -reconfigure -input=false -backend-config=backend.hcl -backend-config="key=eks-cluster/dev/terraform.tfstate" >/dev/null
(no output: redirected)
$ terraform plan -input=false -no-color -var-file=envs/dev.tfvars | grep -E 'Plan:|Error'
Plan: 32 to add, 0 to change, 0 to destroy.

# after the apply (bastion)
$ kubectl get nodes -o wide
ip-10-10-xx-xx.eu-north-1.compute.internal   Ready   <none>   3m38s   v1.36.4-eks-…   Amazon Linux 2023…
```

### Code / config change
`terraform/eks-cluster/eks.tf`, `node_group.tf`, `variables.tf` (commits `07a6e38`, `f8a540a`).

### Lessons learned
Check what the provider currently supports before applying configuration from an older template.

### How to approach it next time
Before any EKS apply: check the version status, the AMI type and the add-on versions (`aws eks describe-addon-versions`).

### Prevention / follow-up
Version kept in a variable; the standard-support policy stops surprise costs.

### References
- AWS: [EKS Kubernetes version lifecycle](https://docs.aws.amazon.com/eks/latest/userguide/kubernetes-versions.html).
- AWS: [Amazon Linux 2023 for EKS nodes](https://docs.aws.amazon.com/eks/latest/userguide/al2023.html).
