package observability

import (
	"sync/atomic"
)

// Metrics holds operational in-memory metrics for RestoreForge.
// Can be exported as Prometheus-compatible metrics.
type Metrics struct {
	BackupJobsCompleted   uint64
	BackupJobsFailed      uint64
	RestoreJobsCompleted  uint64
	RestoreJobsFailed     uint64
	BytesRead             uint64
	BytesWritten          uint64
	UniqueChunksStored    uint64
	ChunksReused          uint64
	VerificationFailures  uint64
}

var DefaultMetrics = &Metrics{}

func (m *Metrics) IncBackupCompleted() {
	atomic.AddUint64(&m.BackupJobsCompleted, 1)
}

func (m *Metrics) IncBackupFailed() {
	atomic.AddUint64(&m.BackupJobsFailed, 1)
}

func (m *Metrics) IncRestoreCompleted() {
	atomic.AddUint64(&m.RestoreJobsCompleted, 1)
}

func (m *Metrics) IncRestoreFailed() {
	atomic.AddUint64(&m.RestoreJobsFailed, 1)
}

func (m *Metrics) AddBytesRead(n uint64) {
	atomic.AddUint64(&m.BytesRead, n)
}

func (m *Metrics) AddBytesWritten(n uint64) {
	atomic.AddUint64(&m.BytesWritten, n)
}

func (m *Metrics) IncUniqueChunks() {
	atomic.AddUint64(&m.UniqueChunksStored, 1)
}

func (m *Metrics) IncChunksReused() {
	atomic.AddUint64(&m.ChunksReused, 1)
}

func (m *Metrics) IncVerificationFailures() {
	atomic.AddUint64(&m.VerificationFailures, 1)
}

func (m *Metrics) DeduplicationRatio() float64 {
	total := atomic.LoadUint64(&m.UniqueChunksStored) + atomic.LoadUint64(&m.ChunksReused)
	if total == 0 {
		return 1.0
	}
	unique := atomic.LoadUint64(&m.UniqueChunksStored)
	if unique == 0 {
		return 1.0
	}
	return float64(total) / float64(unique)
}

