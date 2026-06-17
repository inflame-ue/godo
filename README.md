# godo

A simple TODO list API built as a learning exercise for multi-container Docker applications. The project follows the [Multi-Container Application](https://roadmap.sh/projects/multi-container-service) specification from roadmap.sh.

## What it does

A barebones, unauthenticated TODO API with five endpoints:

| Method | Endpoint           | Description       |
|--------|--------------------|-------------------|
| GET    | `/todos`           | List all todos    |
| POST   | `/todos`           | Create a new todo |
| GET    | `/todos/{id}`      | Get one todo      |
| PUT    | `/todos/{id}`      | Update a todo     |
| DELETE | `/todos/{id}`      | Delete a todo     |

Todos are stored in MongoDB. The API and the database run in separate Docker containers.

## Architecture

The project is split into three requirements, each building on the last:

### Requirement #1 — Dockerize

- **Dockerfile** — Multi-stage build: compile the Go binary in a `golang:alpine` image, then copy the static binary into a minimal `alpine` runtime image. The final image is ~15MB.
- **docker-compose.yml** — Two services: `mongo` (MongoDB 7, data persisted via a named volume) and `godo` (the Go API, pulled from Docker Hub).
- **Environment** — The API connects to MongoDB using Docker's internal DNS (`mongodb://mongo:27017`).

### Requirement #2 — Remote server (GCP)

- **Terraform** — Provisions a Compute Engine VM (`e2-micro`, Ubuntu 24.04) and a firewall rule opening ports 80 and 8080 to the internet. SSH access is configured via a key pair injected into the instance metadata.
- **Ansible** — Configures the blank VM: installs Docker and Docker Compose, copies the `docker-compose.yml`, and starts the containers.d

### Requirement #3 — CI/CD (GitHub Actions)

On every push to `main`:
1. Logs into Docker Hub
2. Builds and pushes the Docker image (`inflameue/godo:latest`)
3. SSH's into the VM and runs `docker compose pull && docker compose up -d`

## Running locally

```bash
docker compose up -d --build
curl http://localhost:8080/todos
```

## What I learned

- Docker Compose networking and multi-stage builds
- The difference between `depends_on` (container start) and actual readiness (database accepting connections)
- Terraform for infrastructure-as-code on GCP
- Ansible for post-provisioning configuration
- Wiring a GitHub Actions pipeline end-to-end
- Why contexts are passed, not stored, in Go
- That GCP firewall rules need `source_ranges = ["0.0.0.0/0"]` and not `source_tags` if you want the whole internet to reach you