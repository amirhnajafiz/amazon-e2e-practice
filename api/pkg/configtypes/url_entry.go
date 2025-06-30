package configtypes

// URLEntry represents a URL entry in the configuration.
type URLEntry struct {
	Name        string `koanf:"name"`
	URL         string `koanf:"url"`
	Description string `koanf:"description"`
}
