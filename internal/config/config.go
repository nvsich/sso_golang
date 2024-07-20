package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"dev"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTtl    time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoadConfig() *Config {
	configPath := fetchConfigPath()

	if configPath == "" {
		panic("config file path is empty")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file not found: " + configPath)
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &config
}

// fetching config path from cmd flag or env variable
// priority: flag > env > default
func fetchConfigPath() string {
	var configPath string

	flag.StringVar(&configPath, "config", "", "config file path")
	flag.Parse()

	if configPath == "" {
		configPath = os.Getenv("CONFIG_PATH")
	}

	return configPath
}
