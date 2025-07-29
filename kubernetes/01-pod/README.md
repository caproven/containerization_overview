# 01-pod

Simple http server responding with the server's instance name (driven by `INSTANCE` env var).

## Running

Build the image:

```sh
docker build -t httpserver .
```

Deploy the pod:

```sh
kind load docker-image httpserver:latest
kubectl apply -f pod.yaml
```

Port forward:

```sh
kubectl port-forward pod/http-server 8080:8080 &
```

Hit locally:

```sh
curl -isk http://localhost:8080/
```

Then can check logs and show that we indeed hit this container.
