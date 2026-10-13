# RestoreForge 🛡️

A resilient, local-first, multi-tenant backup and disaster recovery platform focused on deduplicated storage, incremental snapshots, and durable restore workflows.

RestoreForge is designed for teams and operators who want reliable backup primitives with strong integrity guarantees, clear tenant isolation, and a simple path to deploy on-prem or in a private cloud environment.

## Why RestoreForge?

Modern backup systems often struggle with a few recurring problems:

- duplicate data is stored repeatedly across snapshots and tenants
- backups can become difficult to verify during restore operations
- worker coordination can fail under concurrent execution
- multi-tenant workloads need stricter isolation and clearer operational boundaries

RestoreForge addresses these issues by combining:

- content-addressed chunk storage
- deduplication and incremental snapshotting
- object-store backends such as S3 or MinIO
- durable background job coordination with PostgreSQL lease-based worker locking
- integrity checks before and after restore
- tenant-scoped authorization and metadata boundaries

---

## Core Features

- Content-addressed storage using SHA-256 digests for immutable chunk identity
- Deduplication across files, backups, and repeated content blocks
- Incremental backup workflows to reduce storage and bandwidth overhead
- S3 / MinIO-compatible object storage backend
- Durable worker leases using PostgreSQL transactions and lock expiration recovery
- Multi-tenant isolation across metadata, jobs, storage paths, and API access
- End-to-end integrity verification for backups and restores
- REST API and worker-driven execution model for scalable orchestration

---

## Architecture Overview

```text
restoreforge/
├── .github/workflows/       # CI, Docker, and release automation
├── cmd/
│   ├── api/                 # REST API server
│   ├── worker/              # Background job worker
│   └── restoreforge/        # CLI for operators and developers
├── internal/
│   ├── api/                 # HTTP routes, handlers, and auth middleware
│   ├── backup/              # chunking, manifest generation, and backup orchestration
│   ├── restore/             # restore engine and integrity validation
│   ├── storage/             # S3 / MinIO abstraction layer
│   ├── repository/          # PostgreSQL access and job lease coordination
│   ├── config/              # configuration and environment settings
│   └── observability/       # logging, metrics, and health diagnostics
├── migrations/              # database schema migrations
├── testdata/                # synthetic datasets for testing and benchmarking
├── deployments/              # Docker and Compose deployment assets
├── docs/                    # design notes, diagrams, and ADRs
├── Dockerfile               # container definition
├── docker-compose.yml       # local deployment setup
├── go.mod                   # Go module manifest
├── Makefile                 # common dev tasks (if present)
├── LICENSE                  # project license
└── README.md                # project documentation
```

---

## Quick Start

### Prerequisites

- Go 1.24+
- Docker & Docker Compose
- PostgreSQL (when running the full stack locally)
- Optional: MinIO or S3-compatible storage for object backend tests

### 1. Start supporting services

```bash
docker compose -f deployments/docker/docker-compose.yml up -d
```

### 2. Run the API

```bash
go run ./cmd/api/main.go
```

### 3. Run the worker

```bash
go run ./cmd/worker/main.go
```

### 4. Run the test suite

```bash
go test -v ./...
```

---

## Typical Workflow

1. Create a backup job or tenant-scoped backup configuration
2. Chunk and hash file content for content-addressed deduplication
3. Upload unique chunks to the storage backend
4. Record manifests and metadata in PostgreSQL
5. Run periodic verification and repair procedures
6. Restore files using stored manifests and integrity validation

---

## Project Goals

RestoreForge aims to provide a practical, understandable backup platform for:

- local-first disaster recovery scenarios
- managed service provider workloads
- multi-tenant storage environments
- integrity-focused backup operations with verifiable restoration

---

## Roadmap

- improve API coverage and authentication flow
- expand restore verification and repair workflows
- add richer policy and quota controls for tenants
- improve observability, metrics, and health dashboards
- enhance CLI usability for on-call operational tasks
- add benchmarking and test fixtures for large restore scenarios

---

## Contributing

Contributions are welcome. For development work, we recommend:

- keeping the backup and restore semantics explicit and test-driven
- validating lock and lease behavior under concurrent workers
- documenting tenant boundaries and security assumptions
- maintaining backward compatibility for schema and manifest formats

---

## License

This project is open source. Please check the repository license file for the current terms.

---

## Status

RestoreForge is a platform-focused project intended to demonstrate robust backup architecture patterns, data integrity checks, and multi-tenant operational design in a Go-based system.
