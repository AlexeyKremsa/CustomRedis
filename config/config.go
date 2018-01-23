package config

import (
	"github.com/koding/multiconfig"
	log "github.com/sirupsen/logrus"
)

// CRConfig describes settings for CustomRedis
type CRConfig struct {
	Port              string `default:"8282"`
	LogLevel          string `default:"debug"`
	CleanupTimeoutSec int64  `default:"60"`
}

// LoadConfig load configuration
func LoadConfig() *CRConfig {
	config := &CRConfig{}
	m := multiconfig.NewWithPath("./config/config.toml")
	log.Infof("Loading configuration...")
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
