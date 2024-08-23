package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env      string         `yaml:"env" env-required:"true"`
	Postgres PostgresConfig `yaml:"postgres" env-required:"true"`
	TokenTTL time.Duration  `yaml:"token_ttl" env-required:"true"`
	GRPC     GRPCConfig     `yaml:"grpc" env-required:"true"`
}

type PostgresConfig struct {
	DB       string `yaml:"POSTGRES_DB" env-required:"true"`
	User     string `yaml:"POSTGRES_USER" env-required:"true"`
	Password string `yaml:"POSTGRES_PASSWORD" env-required:"true"`
	Host     string `yaml:"POSTGRES_HOST" env-required:"true"`
	Port     int    `yaml:"POSTGRES_PORT" env-required:"true"`
}

type GRPCConfig struct {
	Port    int
	Timeout time.Duration
}

func MustLoad() *Config {
	path := fetchConfigPath()

	if path == "" {
		panic("config file path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not found")
	}

	var config Config

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		panic("failed to parse config" + err.Error())
	}

	return &config
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "Path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
