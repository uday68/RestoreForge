# Backup & Restore Lifecycle

## Backup Lifecycle

1. **Submission**: Client submits backup request for dataset.
2. **Job Enqueueing**: API creates a durable job row in PostgreSQL with `queued` status.
3. **Leasing**: Worker claims job using `FOR UPDATE SKIP LOCKED`, setting `lease_expires_at`.
4. **Chunking & Hashing**: Files are partitioned into chunks; SHA-256 hashes are calculated.
5. **Deduplication Check**: Existing chunks in object store are reused. Unseen chunks are uploaded.
6. **Manifest Construction**: File manifest containing ordered chunk hashes and file metadata is built.
7. **Two-Phase Commit**: Manifest and chunk references are saved to PostgreSQL; job marked `succeeded`.
8. **Audit & Metrics**: Job completion metrics recorded and audit record emitted.

## Restore Lifecycle

1. **Validation**: Client specifies target backup ID and clean destination directory.
2. **Manifest Loading**: Manifest is fetched and validated from PostgreSQL.
3. **Chunk Fetching & Integrity**: Chunks are retrieved from S3 and checked against SHA-256 hash.
4. **Reassembly**: Files are reconstructed in isolated destination preventing path traversal.
5. **Post-Restore Verification**: Reconstructed files have their full SHA-256 checksums verified against the manifest.

