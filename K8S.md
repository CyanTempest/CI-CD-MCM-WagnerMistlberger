### Task 4: Production Readiness

**Scaling Screenshot:** Scale the API deployment to 3 replicas and verify all pods are running:

![image-20260521142041109](/images/screenshot1)

![image-20260521142112839](/images/screenshot2)

**Health Checks questions:**

**What is the difference between a readiness and a liveness probe?**

- **Liveness Probe:** Checks if the container is still alive and running. If a container deadlocks or crashes internally without terminating the process, Kubernetes uses this probe to detect the failure.

- **Readiness Probe:** Checks if the container is ready to accept network traffic. A container might be running (alive), but still busy loading large configuration files, establishing database connections, or running initial migrations.

**What happens if the readiness probe fails? What about the liveness probe?**

- **If Readiness Fails:** Kubernetes removes the specific Pod's IP address from the endpoints of the associated Service. No network traffic/HTTP requests will be routed to this Pod until the probe passes again. The Pod is *not* restarted.

- **If Liveness Fails:** Kubernetes considers the container broken, kills it, and automatically restarts a fresh instance of the container according to its restart policy.

**Why are different `initialDelaySeconds` values used?**

- Different containers take different amounts of time to initialize. For example, a heavy database or a framework-heavy API needs more time to run startup scripts than a tiny utility container. By adjusting `initialDelaySeconds`, we prevent Kubernetes from prematurely killing a container (via liveness) or spamming it with requests (via readiness) before its internal processes have even had a chance to start up.

**Resource Limits Questions:**

**What happens if a pod exceeds its memory limit?**

- If a container tries to allocate more RAM than allowed by its `limits.memory`, it triggers an Out-Of-Memory event. The Linux kernel immediately kills the process, and Kubernetes marks the Pod status as **`OOMKilled`**. The Pod is then restarted.

**What happens if it exceeds its CPU limit?**

- CPU is a compressible resource. If a container reaches its `limits.cpu`, Kubernetes (via cgroups) will throttle the container's CPU usage. The Pod **does not die**, but its performance drops significantly, causing response times/latency to increase.

**Why are requests and limits both specified?**

- **Requests (Guaranteed Resources):** Used by the Kubernetes Scheduler to find a suitable Node that has enough free space to run the Pod. It ensures the minimum required resources are always reserved for the container.
- **Limits (Maximum Allowed):** Prevents a single misbehaving or leaking container from consuming all available resources on the entire host Node, which would otherwise starve or crash neighboring applications (Noisy Neighbor effect).