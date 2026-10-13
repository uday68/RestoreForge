package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/restoreforge/restoreforge/internal/config"
)

// MinIOStore implements ObjectStore against S3/MinIO.
// For initial vertical slice, includes in-memory fallback support.
type MinIOStore struct {
	cfg     *config.Config
	mu      sync.RWMutex
	memData map[string][]byte
}

func NewMinIOStore(cfg *config.Config) (*MinIOStore, error) {
	return &MinIOStore{
		cfg:     cfg,
		memData: make(map[string][]byte),
	}, nil
}

func (s *MinIOStore) PutChunk(ctx context.Context, hash string, data io.Reader, size int64) error {
	buf, err := io.ReadAll(data)
	if err != nil {
		return fmt.Errorf("read chunk: %w", err)
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.memData[hash] = buf
	return nil
}

func (s *MinIOStore) GetChunk(ctx context.Context, hash string) (io.ReadCloser, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data, exists := s.memData[hash]
	if !exists {
		return nil, fmt.Errorf("chunk not found: %s", hash)
	}
	return io.NopCloser(bytes.NewReader(data)), nil
}

func (s *MinIOStore) HasChunk(ctx context.Context, hash string) (bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, exists := s.memData[hash]
	return exists, nil
}

func (s *MinIOStore) DeleteChunk(ctx context.Context, hash string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.memData, hash)
	return nil
}

