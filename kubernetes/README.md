# Kubernetes Overview

## Container Limitations

Containers are awesome, but `docker run` is only for my own machine. What if I need to scale? Example:

> When launching a service, it receives 1000 req/sec. Then my startup or whatever pops off, and now I receive 1m req/sec. If the service running on my machine can only handle half a million, I need a second machine to serve all my users. I'd need a second machine in the first place, and then I'd have to `docker run` there as well. Then I'd need some sort of networking abstraction, probably a load balancer, between them.

Docker and containers are great for quickly deploying but leave a lot to be desired if you're scaling to meet potentially global production needs. They lack adequate abstractions for running on multiple machines, and doesn't help us if we need to gracefully update services. Say I have v1 of my app then need to roll out v2. Do I deploy v2 on one machine, delete v1, then deploy v2 everywhere? Sounds like a lot of work.

## What is Kubernetes?

"Container orchestration platform". It provides a set of abstractions which help us run and manage containers. Of interest, it lets us easily deploy containers to multiple machines, lets us dynamically scale applications, and lets us monitor health and logs for entire application suites.

### Cluster Architecture

Kubernetes is a platform composed of different nodes or machines acting together to form a cluster. You can think of the cluster as two separate parts: the control plane and the data plane.

The control plane is the set of processes that govern the entire cluster. This is things like an API server we can use to interact with the cluster, or ETCD which is a database (k/v store) used to hold the state of everything in the cluster. The data plane is then the set of nodes where our containers run, as well as extra tools to make everything work together such as internal proxies.

### K8s API

Kubernetes is interacted with through an API. This can be through REST calls but primarily you interact with a cluster using the CLI, `kubectl` (no official pronunciation).

The API is meant to be used declaratively as opposed to imperatively, but it does support both. Really important to understand what these mean so I'll expand on them. Declarative commands are when you tell a system (or *declare*) a desired state and it is responsible for determining what that means and how to do it. Imperative commands are instructions where you give specific actions; the system executes it on your behalf

SQL statements are an example of imperative commands. You tell a database to *create a user* or *delete these rows from a table*. You are specifying the exact action to be done and any applicable targets for the action. For example if you want to create rows in a table, you specify the table you care about and provide the exact rows you'd like to insert. You don't care about what's already in the table, you just want your new rows added. Declarative commands in contrast would be like you updating table rows by providing the entire set of rows, along with your updates.

Before running below commands, have state:

```sql
SELECT * FROM mytable;
 id
----
  1
  2
(2 rows)
```

```sql
-- imperative
INSERT INTO mytable VALUES (3);

SELECT * FROM mytable;
-- note how our insertion was purely additive
 id 
----
  1
  2
  3
(3 rows)
```

```sql
-- declarative (not actual sql syntax)
mytable (
  (3)
)

SELECT * FROM mytable;
-- note how the update erased existing rows
 id
----
  3
(1 row)
```

Benefit of declarative is idempotency, meaning you can rerun the same instruction and always get the same output. In the declarative example, you provide the *entire* table. That way you don't care about existing rows, hell the table may not even exist yet. But once you issue that command, you know you have a table with exactly that one row. Outside the SQL example, a major benefit of idempotent design is that you never end up in corrupted states. No data is merged automatically, each write operation must specify the full specification of what it'd like.

### Controllers

API broken up into different groups which contain their own resources. For each resource, there is a controller whose sole responsibility is *reconciling*. In other words, when a resource of a controller's type is created, modified, or deleted, the controller performs any necessary steps to make the system's actual state reflect the desired state. This is the backbone of Kubernetes' declarative model.

## Running an App

Basic pod, explain manifest & show running. Reiterate declarative nature of `k apply -f <file>`

* With just pod, can't access from host machine (http://localhost:80)
* Create service, now can access
    * NOT the same as opening a new port. Lot of rules & ways to get network traffic into the cluster
* What if I want more than one instance? Show deployment
    * Update service port then access (http://localhost:80)
    * Scale from 1 to 3
    * Show inherent load balancing
    * Wait we created a deployment and ended up with pods, what's going on?

### Scalability

Look up multi-node kind cluster

```bash
while; do
    output=$(curl -ks "http://localhost:80")
    echo "$(date +%T) $output"
    sleep 1
done
```

## Other Abstractions

networking (kinda already covered)

storage

config maps & secrets

CI/CD, gitops

helm & kustomize
