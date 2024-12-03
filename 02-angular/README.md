# 02-angular

Fairly standard container serving an Angular application (client-side rendering)
through nginx. Introduces notion of stages in a Dockerfile.

First stage is the "builder", building the Angular app with a production configuration.

Second stage is what's executed, serving the Angular app's static files with nginx.

## Building

```bash
# relative to repo root
docker build ./02-angular -t 02-angular
```

## Running

```bash
docker run --rm -it -p 8080:8080 02-angular:latest
```

Access in browser at `http://localhost:8080`. Stop with CTRL-C.
