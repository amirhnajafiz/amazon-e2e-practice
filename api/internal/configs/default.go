package configs

import "github.com/amirhnajafiz/aep/backend/pkg/configtypes"

// Default returns the default configuration.
func Default() Config {
	return Config{
		Port:     8080,
		Revision: 1,
		AdminKey: "admin",
		Storage: configtypes.Storage{
			Port:     5432,
			Host:     "127.0.0.1",
			User:     "postgres",
			Pass:     "password",
			Database: "postgres",
			SSL:      "disable",
		},
		URLs: make([]configtypes.URLEntry, 0),
	}
}
