package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	APIPort     string
	APIHost     string
	Environment string
	LogLevel    string

	// Postgres
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresSSLMode  string

	// MinIO / S3
	S3Endpoint  string
	S3AccessKey string
	S3SecretKey string
	S3Bucket    string
	S3UseSSL    bool
	S3Region    string

	// Worker
	WorkerConcurrency   int
	WorkerPollInterval  time.Duration
	WorkerLeaseDuration time.Duration
	WorkerMaxRetries    int

	// Chunking
	ChunkSizeBytes int64
}

func Load() *Config {
	return &Config{
		APIPort:             getEnv("API_PORT", "8080"),
		APIHost:             getEnv("API_HOST", "0.0.0.0"),
		Environment:         getEnv("ENVIRONMENT", "development"),
		LogLevel:            getEnv("LOG_LEVEL", "info"),
		PostgresHost:        getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:        getEnv("POSTGRES_PORT", "5432"),
		PostgresUser:        getEnv("POSTGRES_USER", "restoreforge"),
		PostgresPassword:    getEnv("POSTGRES_PASSWORD", "restoreforge_secret"),
		PostgresDB:          getEnv("POSTGRES_DB", "restoreforge"),
		PostgresSSLMode:     getEnv("POSTGRES_SSLMODE", "disable"),
		S3Endpoint:          getEnv("S3_ENDPOINT", "localhost:9000"),
		S3AccessKey:         getEnv("S3_ACCESS_KEY", "minioadmin"),
		S3SecretKey:         getEnv("S3_SECRET_KEY", "minioadmin"),
		S3Bucket:            getEnv("S3_BUCKET", "restoreforge-chunks"),
		S3UseSSL:            getEnvAsBool("S3_USE_SSL", false),
		S3Region:            getEnv("S3_REGION", "us-east-1"),
		WorkerConcurrency:   getEnvAsInt("WORKER_CONCURRENCY", 4),
		WorkerPollInterval:  getEnvAsDuration("WORKER_POLL_INTERVAL", 2*time.Second),
		WorkerLeaseDuration: getEnvAsDuration("WORKER_LEASE_DURATION", 30*time.Second),
		WorkerMaxRetries:    getEnvAsInt("WORKER_MAX_RETRIES", 3),
		ChunkSizeBytes:      int64(getEnvAsInt("CHUNK_SIZE_BYTES", 1024*1024)),
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
	}
	return defaultVal
}

func getEnvAsBool(key string, defaultVal bool) bool {
	if val := os.Getenv(key); val != "" {
		if boolVal, err := strconv.ParseBool(val); err == nil {
			return boolVal
		}
	}
	return defaultVal
}

func getEnvAsDuration(key string, defaultVal time.Duration) time.Duration {
	if val := os.Getenv(key); val != "" {
		if dur, err := time.ParseDuration(val); err == nil {
			return dur
		}
	}
	return defaultVal
}

