package internal

import (
	"sync"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	Omada struct {
		OmadaURL     string `env:"OMADA_URL,required"`
		SiteName     string `env:"OMADA_SITE_NAME,required"`
		ClientID     string `env:"OMADA_CLIENT_ID,required"`
		ClientSecret string `env:"OMADA_CLIENT_SECRET,required"`
		Username     string `env:"OMADA_USERNAME,required"`
		Password     string `env:"OMADA_PASSWORD,required"`
	}
	Prometheus struct {
		MetricsPath string `env:"METRICS_PATH" envDefault:"/metrics"`
		MetricsPort string `env:"METRICS_PORT" envDefault:"8080"`
	}
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		var err error
		instance, err = loadConfig()
		if err != nil {
			panic("Failed to load configuration: " + err.Error())
		}
	})
	return instance
}

func loadConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
