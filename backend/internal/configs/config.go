package configs

import (
	"log"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

const (
	// Prefix indicates environment variables prefix.
	Prefix = "AEP_"
)

// Config stores the application parameters.
type Config struct {
	Port    int     `koanf:"int"`
	Metrics Metrics `koanf:"metrics"`
	JWT     JWT     `koanf:"jwt"`
	Logger  Logger  `koanf:"logger"`
	Storage Storage `koanf:"storage"`
}

// Load returns the config struct from the given YAML file.
func Load(path string) Config {
	var instance Config

	k := koanf.New(".")

	// load default
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Printf("error loading deafult: %v\n", err)
	}

	// load configs file
	if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
		log.Printf("error loading config.yaml file: %v\n", err)
	}

	// load environment variables
	if err := k.Load(env.Provider(Prefix, ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(
			strings.TrimPrefix(s, Prefix)), "__", ".")
	}), nil); err != nil {
		log.Printf("error loading environment variables: %s", err)
	}

	// unmarshalling
	if err := k.Unmarshal("", &instance); err != nil {
		log.Printf("error unmarshalling config: %v\n", err)
	}

	return instance
}
