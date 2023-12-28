package config

import "fmt"

type (
	Config struct{}
)

func Init(path string) (*Config, error) {
	var conf Config

	return &conf, fmt.Errorf("not implemented")
}
