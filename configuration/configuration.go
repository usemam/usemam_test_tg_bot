package configuration

import (
	"os"
)

// Configuration storage
type Configuration struct {
	Port     string
	BotToken string
	URL      string
}

// New - factory method for reading current configuration
func New() *Configuration {
	return &Configuration{
		Port:     os.Getenv("PORT"),
		BotToken: os.Getenv("TOKEN"),
		URL:      os.Getenv("URL"),
	}
}
