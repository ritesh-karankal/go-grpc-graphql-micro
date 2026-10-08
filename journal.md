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
| [INC-012](#inc-012--signoz-ui-not-reachable-load-balancer-created-as-internal) | 2026-10-09 | SigNoz / AWS LB Controller | SigNoz UI not reachable: load balancer created as `internal` | Medium | Fix applied |
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

## INC-012 – SigNoz UI not reachable: load balancer created as `internal`

| | |
|---|---|
| **Date** | 2026-10-09 |
| **Where** | `k8s/monitoring/signoz-ui-service.yaml`; Recreate guide Step 16 (SigNoz) |
| **Severity** | Medium: the SigNoz UI was unreachable from outside the VPC; SigNoz itself was running |
| **Status** | Fix applied: Service recreated with the `internet-facing` annotation; waiting for confirmation that the UI opens |

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
6. Added the annotation to the manifest and recreated the Service.

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
