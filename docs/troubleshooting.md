# Troubleshooting Guide

## Common Issues & Diagnoses

### 1. Worker Cannot Connect to PostgreSQL
- Ensure PostgreSQL container is running: `docker compose ps`
- Verify database port `5432` is reachable.

### 2. MinIO S3 Connection Refused
- Confirm MinIO is up and healthy: `curl http://localhost:9000/minio/health/live`
- Check S3 bucket exists or auto-initialization is active.

### 3. Stale or Abandoned Jobs
- Workers lease jobs with an expiration timestamp (`lease_expires_at`).
- If a worker crashes, the sweeper process recovers jobs past their lease timeout.

