package config

import (
	"github.com/pelletier/go-toml"
	"log"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type RedisConfig struct {
	Host string
	Port string
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{}
	data, err := toml.LoadFile(path)
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
		return nil, err
	}
	err = data.Unmarshal(config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
		return nil, err
	}
	return config, nil
}
