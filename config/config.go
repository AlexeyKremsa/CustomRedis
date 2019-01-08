package config

import (
	"github.com/koding/multiconfig"
	log "github.com/sirupsen/logrus"
)

// Config describes settings for CustomRedis
type Config struct {
	Port              string `default:"8282"`
	LogLevel          string `default:"debug"`
	CleanupTimeoutSec int    `default:"60"`
	ShardCount        int    `default:"1"`
}

// LoadConfig load configuration
func LoadConfig(path string) *Config {
	config := &Config{}
	m := multiconfig.NewWithPath(path)
	log.Infof("Loading configuration from: %s", path)
	err := m.Load(config)
	if err != nil {
		log.Fatalf("Failed to load configuration. %v", err)
	}

	err = m.Validate(config)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("%+v\n", config)

	if config.CleanupTimeoutSec < 0 {
		log.Fatal("CleanupTimeoutSec can`t be less than 0")
	}

	return config
}
