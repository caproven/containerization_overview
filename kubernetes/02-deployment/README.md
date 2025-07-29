# 02-deployment

Example showing horizontal scaling with a Deployment resource.

## Running

```sh
# Note: only includes 1 replica
kubectl apply -f deployment.yaml
```

Most fields within a Pod are immutable. A Deployment however is mostly mutable and applies a rolling update strategy to underlying Pods when changes are made.

```sh
# Note: call out that this is an "imperative" command. Could equally modify
# deployment.yaml and set replicas field.
k edit deployment/http-server # set spec.replicas > 1
```

Then access the http server.

```sh
kubectl port-forward deployment/http-server 8080:8080 &
```

Run curl a few times against it - it should give the same instance name every time. This is because `kubectl port-forward` on a Deployment will match a pod at command execution time and port-forward it, not anything dynamic. If you delete the selected Pod (so new one with different name is created), the port-forward will break.

To load balance, we need a Service resource.

```sh
kubectl apply -f service.yaml
```

So we can show load balancing working, create a new Pod we can exec into.

```sh
kubectl run nginx --image nginx
# Wait for it to start
```

```sh
kubectl exec -it nginx -- sh
# From within the container..
curl http://http-server-svc:8080
```
