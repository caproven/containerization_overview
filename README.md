# Containerization & Docker Overview

Learning objectives

* Understand what a container is
    * Layered architecture
    * Benefits over traditional virtual machines
* Know difference between containers & Docker
* Understand benefits of containerizing applications
* Basic understanding of Dockerfile instructions & syntax
* Working knowledge of Docker CLI (`build`, `run`, `exec` subcommands)
* Container networking
* Accessing storage within containers

## Reset local state

Run before working through examples

```bash
# actually start docker desktop :)

# relative to examples dir
docker rm -f $(docker ps -a -q)
docker image prune -af
docker network prune -f
docker volume prune -af
sudo rm -rf ./04-postgres/data
```

## What are containers

Containers can be thought of as a bundle of an executable program along with a static & reproducible environment for it to execute in. This software can be a full-blown application, a web server, or even just a shell script.

From that brief definition this probably sounds like a virtual machine, however they're two completely different concepts. A virtual machine has much more overhead as it virtualizes hardware and emulates an operating system along with its kernel, whereas containers simply reuse the host OS & kernel. This is kept secure using Linux cgroups to isolate container space from host processes. I can find references to read more about this if anyone's interested, but basically containers are minimized to your running processes and dependencies such as C libraries, Linux packages, TLS certs, etc.

You'll also hear the term "images" used. Images are the on-disk bundle containing the aforementioned items. Containers are the running instance of an image. *You build & publish images, and you run containers*. You can think of this as the difference between an executable and a process running that executable.

### What is Docker

Docker is an implementation & toolchain for building images along with running containers. This will make more sense when we get to examples later on.

### Benefits

Solves the age-old issue of "it works on my machine". With a container, you effectively ship "the machine" alongside your app, minimizing variables and promoting consistent runtime environments. If I have a Java application I can build an image with my intended Java version which is a drastic improvement over relying on the correct version being installed on my host and worrying about API compatibility. This is similar to the compatibility layer offered by the JVM, where compiled Java byte-code (jar file) can run on any machine with the JVM installed.

Unlike the JVM, we have the ability to bundle *all* app dependencies. Continuing the JVM comparison, if I have a Java web app and need TLS support then I not only have to copy my jar file around but also my TLS certs. If my app relies on OS packages being installed, a startup bash script, or any other tools, you can't guarantee the runtime environment has everything required. Ultimately you'd end up documenting requirements and relying on people, the most common source of errors.

To reiterate, containers bundle *all* dependencies so those concerns are eliminated (not necessarily including configuration). We'd describe this as *portability*; I wrap my software in an image, and suddenly I can run it on any machine with a container runtime, e.g., Docker. Practically, this opens the way for dead simple deployments to environments. I can throw my image at Azure, AWS, Google Cloud, Heroku, VMWare, DigitalOcean, & many other providers or even your own on-prem servers. A huge benefit is the amount of tooling available for containers handling modern software needs such as security, scalability, observability, and automation.

## Image Architecture

Images are a series of layers or instructions which "build" the runtime environment and include your software. You'll commonly see a "base" layer stemming from common Linux distributions. This isn't a full emulation of the OS but more about establishing familiar tools & configuration to build upon. For example, if I use an Ubuntu or Debian base image then I can install packages using `apt`, create system users with standard Linux commands, and have a predictable file system.

After your base image you'll have instructions for tasks like creating a system user, installing OS packages, and building or otherwise copying in your executable code. The idea is that these independent layers can be cached & reused. For instance if you have 2 images downloaded that both use the same base image, it's only stored once. This also aids image build times since you don't rebuild un-impacted layers.

```bash
# 01-python example

# walk through code
docker build ./01-python -t numbers
docker run --rm -it numbers

docker ps # mention names
# mention 10 second default grace period
docker stop <container_name>
```

```bash
# 02-angular example

# walk through code. Explain "builder" images/stages (copy select artifacts from stages into final image), mention dockerignore
docker build ./02-angular -t 02-angular
docker run --rm -it -p 8080:8080 02-angular

# access on localhost:8080
```

## Networking

Similar to process isolation between containers and the host, networking is also isolated. Unlike process isolation though where we really never want to bridge a container and its host (or another container), networking is fundamentally a way to connect machines.

When running a Docker container, networking is blocked by default at container boundaries. So 2 processes within a container can communicate, but they can't connect to the host or to another container. But commonly when running a container, we need to expose it to the outside world. In the previous Angular example, we explicitly opened a port to allow web traffic.

```bash
# (revisit) 02-angular example

# show nginx.conf, server listening on 8080 within container. Can change external mapping to any port
docker run --rm -it -p 12345:8080 02-angular
```

### Inter-Container Networking

For inter-container networking, Docker lets us create virtual networks which containers can be attached to. Containers within the virtual network can communicate with any other containers in it as we'll see in a second.

```bash
# 03-curl example - SPLIT TERMINAL

# show script
docker build ./03-curl -t 03-curl # LEFT
# start angular (in detached mode)
docker run --rm -d -p 8080:8080 02-angular # LEFT

# start curl. Replace <CONTAINER_NAME> with 02-angular name from docker ps. Explain DNS
docker run --rm -it --env "URL=<CONTAINER_NAME>:8080" 03-curl # LEFT
# call out above command giving 6 err code

# create virtual network - RIGHT
docker network create my-network
docker network connect my-network <ANGULAR_CONTAINER>
docker network connect my-network <CURL_CONTAINER>
# curl container should now be giving result 0 (success)
```

## Persistent Storage

Next we'll talk about persistent storage, working with the example of a database. So let's run Postgres.

```bash
# 04-postgres example
docker run --rm -it -e POSTGRES_PASSWORD=foo123 -d -p 5432:5432 postgres
psql -h localhost -U postgres # pw: foo123

# sql
create table nums (id integer);
insert into nums (id) select generate_series(1, 1000000);
select count(*) from nums;

# now will shut postgres down and restart it
docker stop <CONTAINER_NAME>
# restart using original run cmd
# reconnect using original cmd. Call out that we have no data (\dt shows no nums table)
docker stop <CONTAINER_NAME>
```

What we've been missing is *persistent* storage. Containers have their own storage, that's how we were able to write to Postgres before, but it's tied to the container and therefore lost when the container is deleted. If we need storage to persist, we'd want to use volumes mounted into the container. You can think of this as us giving a directory in the container where files written are synced with the host. This is actually a *mount* so it's not copying or anything, but that might be an easier way to understand what we're accomplishing.

```bash
# (continuing) 04-postgres
# explain -v arg
docker run --rm -it -d \
	-e POSTGRES_PASSWORD=foo123 \
	-p 5432:5432 \
	-v ./04-postgres/data:/var/lib/postgresql/data \
	postgres

ls ./04-postgres/data
# fails with permission denied. Explain permissions (postgres runs as root, show with docker exec)
docker exec -it <CONTAINER_NAME> whoami
sudo ls ./04-postgres/data
# call out that we're seeing postgres data files (weren't there before)

# in side terminal
docker exec -it <CONTAINER_NAME> du -sh /var/lib/postgresql/data
# in orig terminal
psql -h localhost -U postgres # pw: foo123
# sql
create table nums (id integer);
insert into nums (id) select generate_series(1, 1000000);
select count(*) from nums;
# re-run du -sh in side terminal. Should be larger than before

# because storage is in this mount... delete container & recreate
docker rm <CONTAINER_NAME>
# re-run docker run from before with volume
# connect with psql & show table still present
```

What we did is called "host volumes" since you're mounting a host directory into the container, but Docker itself can also create and manage volumes for us. I used host volumes since they're easier to inspect & show off but for completeness, these extras approaches are

* Named volumes - managed by Docker (`docker volume create` + `docker run -v VOLUME_NAME:/target/dir postgres`)
* Anonymous volumes - also managed by Docker but are given an ID instead of a name (`docker run -v /target/dir postgres`)
    * Not really sure what use case is

### Container Debugging

Used it before but want to be explicit. `docker exec` is a really powerful command for debugging purposes. If Postgres crashed or file permissions got screwed up, I can exec in and run whatever commands I'd like to troubleshoot. Before I was running direct commands using `docker exec ... <command>`  but if a container includes a shell (most will), I can just start a shell session.

```bash
docker exec -it <CONTAINER_NAME> bash
# within postgres container
cd /var/lib/postgresql/data
vi pg_hba.conf # command not found, mentioned later
cat pg_hba.conf
exit
```

This is limited however by tools available within the container. As we saw before, the Postgres container doesn't include vim (but if I had to guess it probably has another editor like nano). If my container doesn't stem from a Linux base image and is only an executable, you won't have a shell or really any tools at your disposal. There are some nifty workarounds to this, I'm more familiar with Kubernetes solutions to that very problem, but I don't have anything to really show off. If containers house Linux distros & debugging tools, that's great for troubleshooting & development but that will also bloat image size. As with most things, I think there's a sweet spot of tools to provide.

## Microservice Design

Learning objectives

* Understand microservice architecture & motivations
    * Why industry adopting the pattern, monolith = bad
    * Scalability
