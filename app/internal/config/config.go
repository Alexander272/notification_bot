package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Environment string     `yaml:"environment" env:"APP_ENV" env-default:"dev"`
		Http        HttpConfig `yaml:"http"`
		Limiter     LimiterConfig
		Most        MostConfig
	}

	HttpConfig struct {
		Host           string        `yaml:"host" env:"HOST" env-default:"localhost"`
		Port           string        `yaml:"port" env:"PORT" env-default:"8080"`
		ReadTimeout    time.Duration `yaml:"read_timeout" env:"READ_TIMEOUT" env-default:"10s"`
		WriteTimeout   time.Duration `yaml:"write_timeout" env:"WRITE_TIMEOUT" env-default:"10s"`
		MaxHeaderBytes int           `yaml:"max_header_bytes" env-default:"1"`
	}

	LimiterConfig struct {
		RPS   int           `yaml:"rps" env:"RPS" env-default:"10"`
		Burst int           `yaml:"burst" env:"BURST" env-default:"20"`
		TTL   time.Duration `yaml:"ttl" env:"TTL" env-default:"10m"`
	}

	MostConfig struct {
		ServerLink string `env:"MOST_SERVER"`
		Token      string `env:"MOST_TOKEN"`
		BotName    string `env:"MOST_BOT_NAME"`
	}
)

func Init(path string) (*Config, error) {
	var conf Config

	if err := cleanenv.ReadConfig(path, &conf); err != nil {
		return nil, fmt.Errorf("failed to read config file. error: %w", err)
	}

	// logger.Info(conf)

	// if err := cleanenv.ReadEnv(&conf); err != nil {
	// 	return nil, fmt.Errorf("failed to read env variable. error: %w", err)
	// }

	return &conf, nil
}
