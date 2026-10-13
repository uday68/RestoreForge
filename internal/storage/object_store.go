package storage

import (
	"context"
	"io"
)

// ObjectStore defines the interface for interacting with S3-compatible chunk storage.
type ObjectStore interface {
	// PutChunk uploads a chunk identified by its SHA-256 hash.
	PutChunk(ctx context.Context, hash string, data io.Reader, size int64) error

	// GetChunk downloads a chunk identified by its SHA-256 hash.
	GetChunk(ctx context.Context, hash string) (io.ReadCloser, error)

	// HasChunk checks if a chunk already exists in storage.
	HasChunk(ctx context.Context, hash string) (bool, error)

	// DeleteChunk removes a chunk from storage (for GC).
	DeleteChunk(ctx context.Context, hash string) error
}

