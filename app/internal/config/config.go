package config

import (
	"fmt"
	"time"
)

type (
	Config struct {
		Environment string `yaml:"environment" env:"APP_ENV" env-default:"dev"`
		Http        HttpConfig
		Limiter     LimiterConfig
	}

	HttpConfig struct {
		Host           string        `yaml:"host" env:"HOST" env-default:"localhost"`
		Port           string        `yaml:"post" env:"PORT" env-default:"8080"`
		ReadTimeout    time.Duration `yaml:"read_timeout" env:"READ_TIMEOUT" env-default:"10s"`
		WriteTimeout   time.Duration `yaml:"write_timeout" env:"WRITE_TIMEOUT" env-default:"10s"`
		MaxHeaderBytes int           `yaml:"max_header_bytes" env-default:"1"`
	}

	LimiterConfig struct {
		RPS   int           `yaml:"rps" env:"RPS" env-default:"10"`
		Burst int           `yaml:"burst" env:"BURST" env-default:"20"`
		TTL   time.Duration `yaml:"ttl" env:"TTL" env-default:"10m"`
	}
)

func Init(path string) (*Config, error) {
	var conf Config

	return &conf, fmt.Errorf("not implemented")
}
