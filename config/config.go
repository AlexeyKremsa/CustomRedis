package config

import (
	"github.com/koding/multiconfig"
	"github.com/sirupsen/logrus"
)

// CRConfig describes settings for CustomRedis
type CRConfig struct {
	Port            string `default:"8282"`
	LogLevel        string `default:"debug"`
}

// LoadConfig load configuration
func LoadConfig() *CRConfig {
	config := &CRConfig{}
	m := multiconfig.NewWithPath("./config/config.toml")
	logrus.Infof("Loading configuration...")
	err := m.Load(config)
	if err != nil {
		logrus.Fatalf("Failed to load configuration. %v", err)
	}

	err = m.Validate(config)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("%+v\n", config)

	return config
}