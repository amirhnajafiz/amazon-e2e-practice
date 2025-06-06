package configs

import (
	"fmt"
)

// JWT configuration for JSON Web Tokens.
type JWT struct {
	PrivateKey string `koanf:"private_key"`
	ExpireTime int    `koanf:"expire_time"`
}

// Logger configuration for zap logger.
type Logger struct {
	Level string `koanf:"level"`
}

// Metrics configuration for Prometheus metrics.
type Metrics struct {
	Enabled bool `koanf:"enabled"`
	Port    int  `koanf:"port"`
}

// Storage configuration for PostgreSQL database.
type Storage struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	User     string `koanf:"user"`
	Pass     string `koanf:"pass"`
	Database string `koanf:"database"`
	SSL      bool   `koanf:"ssl"`
}

func (s Storage) URI() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%t", s.User, s.Pass, s.Host, s.Port, s.Database, s.SSL)
}
