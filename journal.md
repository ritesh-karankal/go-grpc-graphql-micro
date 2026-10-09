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
| [INC-017](#inc-017--connection-reset-by-peer-during-a-rollout-pods-exit-without-draining) | 2026-10-09 | Kubernetes / Go services | `Connection reset by peer` during a rollout: pods exit without draining | Medium | Open |
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

## INC-017 – `Connection reset by peer` during a rollout: pods exit without draining

| | |
|---|---|
| **Date** | 2026-10-09 |
| **Where** | `telemetry/telemetry.go` › `ShutdownOnSignal`; the four Go services' `main`; Deployments in `k8s/base/` (no `preStop`, no grace period); ALB target group |
| **Severity** | Medium: some requests fail on **every deploy**, which burns the SLO error budget each time Jenkins ships |
| **Status** | Open: cause identified; fix (graceful shutdown + `preStop` + ALB readiness) proposed |

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

### Code / config change
- Done: `scripts/loadgen.py` (exception handling and reporting).
- Pending: `telemetry/telemetry.go`, the services' `main` / `ListenGRPC`, `k8s/base/*` Deployments, `k8s/base/ingress/ingress.yaml`, namespace label.

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
   Result: build #5 passed the gate, pushed `account/catalog/order/graphql:5`, committed the tag, and Argo CD deployed it. The live API now answers an unknown account with `code = NotFound` (before: `code = Unknown`).

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
