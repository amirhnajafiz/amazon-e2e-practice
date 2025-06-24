package configs

import (
	"fmt"
)

// JWT configuration for JSON Web Tokens.
type JWT struct {
	PrivateKey string `koanf:"private_key"`
	ExpireTime int    `koanf:"expire_time"`
}

// Storage configuration for PostgreSQL database.
type Storage struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	User     string `koanf:"user"`
	Pass     string `koanf:"pass"`
	Database string `koanf:"database"`
	SSL      string `koanf:"ssl"`
}

func (s Storage) URI() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", s.User, s.Pass, s.Host, s.Port, s.Database, s.SSL)
}
