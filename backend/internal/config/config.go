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
	ConnectionString string        `yaml:"connection_string"`
	Grpc             Grpc          `yaml:"grpc"`
	TokenTLL         time.Duration `yaml:"token_tll"`
	LogLevel         string        `yaml:"log_level"`
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
