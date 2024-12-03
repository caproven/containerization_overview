# 03-curl

Straightforward container running curl against a URL every couple seconds and
printing the return code (0 is success, else failure).

## Building

```bash
# relative to repo root
docker build ./03-curl -t 03-curl
```

## Running

Assumes that [02-angular](../02-angular/README.md) has already been built.

Container is used to show networking within Docker's virtual networks. First create
a network for containers to be attached to.

```bash
docker network create my-network
```

Then spin up the containers using the new network.

```bash
# spin up 02-angular example (runs in background)
docker run --rm -d \
  --network my-network \
  --name 02-angular \
  02-angular:latest

# spin up 03-curl
docker run --rm -it \
  --network my-network \
  --name 03-curl \
  --env "URL=02-angular:8080" \
  03-curl:latest
```

If working correctly, should log `hitting url 02-angular:8080 gave result 0`. Any
other result indicates something went wrong.

Stop with CTRL-C. To clean up other resources:

```bash
docker stop 02-angular
docker network rm my-network
```
