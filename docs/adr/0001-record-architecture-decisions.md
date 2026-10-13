# ADR 0001: Record Architecture Decisions

## Context
RestoreForge requires content-addressed storage, deduplication, durable background worker queues, and tenant isolation.

## Decisions
1. **Content Addressing**: Use SHA-256 digests over fixed-size blocks (default 1MB).
2. **Object Storage**: Store raw immutable chunks named by their hex SHA-256 hash.
3. **Queueing Mechanism**: Use PostgreSQL row-level locks (`SELECT FOR UPDATE SKIP LOCKED`) for worker leasing before introducing external message brokers.
4. **Tenant Isolation**: Isolate metadata strictly by tenant ID; evaluate tenant-scoped vs global deduplication with explicit privacy boundaries.

