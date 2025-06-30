package configtypes

// URLEntry represents a URL entry in the configuration.
type URLEntry struct {
	ShortenedID string `koanf:"shortened_id"`
	Address     string `koanf:"address"`
	Description string `koanf:"description"`
}
