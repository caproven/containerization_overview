# 04-postgres

Sample of running PostgreSQL in a container with data written to a mounted host
volume. Illustrates the purpose of volumes, decoupling data persistence from
the runtime.

## Running

```bash
docker run --rm -it \
  -p 5432:5432 \
  --env "POSTGRES_PASSWORD=foo123" \
  -v ./04-postgres/data:/var/lib/postgresql/data \
  postgres:latest
```

Stop with CTRL-C.

**Note:** A separate tool such as `psql` would be required to connect to the running
database on `localhost:5432`.

```bash
psql -h localhost -U postgres
# when prompted for password, use value of POSTGRES_PASSWORD above
```

**Note:** PostgreSQL container can be deleted & recreated as many times as you'd like,
but the volume (if remounted) lets us access the same data.
