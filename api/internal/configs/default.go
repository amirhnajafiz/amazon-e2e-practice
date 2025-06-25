package configs

// Default returns the default configuration.
func Default() Config {
	return Config{
		Port:    8080,
		Migrate: false,
		Storage: Storage{
			Port:     5432,
			Host:     "127.0.0.1",
			User:     "postgres",
			Pass:     "password",
			Database: "postgres",
			SSL:      "disable",
		},
		URLs: make([]URLEntry, 0),
	}
}
