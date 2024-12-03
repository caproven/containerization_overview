# 01-python

Straightforward example of a container running a Python script printing incrementing
numbers every couple seconds.

## Building

```bash
# relative to repo root
docker build ./01-python -t 01-python
```

## Running

```bash
docker run --rm -it 01-python:latest
```

Stop with CTRL-C.
