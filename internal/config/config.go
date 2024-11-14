package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Grpc struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}
type Config struct {
	Grpc     Grpc          `yaml:"grpc"`
	TokenTTL time.Duration `yaml:"token_ttl"`
	LogLevel string        `yaml:"log_level"`
}

func InitConfig(path string) (*Config, error) {
	var config Config

	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(fileContent, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
