package configs

// Default returns the default configuration.
func Default() Config {
	return Config{
		Port: 8080,
		Metrics: Metrics{
			Enabled: false,
			Port:    8081,
		},
		JWT: JWT{
			PrivateKey: "secret",
			ExpireTime: 30, // in miutes
		},
		Logger: Logger{
			Level: "debug",
		},
		Storage: Storage{
			Port:     5432,
			Host:     "127.0.0.1",
			User:     "user",
			Pass:     "password",
			Database: "caaas-db",
			SSL:      false,
		},
	}
}
