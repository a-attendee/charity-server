package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
    Env string                          `yaml:"env" env-required: "true"`
    StoragePath string                  `yaml: "storage_path" env-required: "true"` 
    HttpServer HttpServerConfig         `yaml: "http_server"`
}

type HttpServerConfig struct {
    Address string                      `yaml: "address" env-default: "localhost"`
    Port int                            `yaml: "port" env-default: "3000"`
    RequestTimeout time.Duration        `yaml: "request_timeout" env-default: "3s"`
    MaxRequests int                     `yaml: "max_requests" env-default: "20"`
    MaxRequestsExperation time.Duration `yaml: "max_requests_experation:" "30s"`
}

func MustLoad() Config {
    configPath := os.Getenv("CONFIG_PATH")
    if configPath == "" {
        log.Fatalf("CONFIG_PATH is not defined\n")
        os.Exit(2)
    }

    if _, err := os.Stat(configPath); os.IsNotExist(err) {
        log.Fatalf("config file is not exist: ", configPath) 
        os.Exit(2)
    }

    var cfg Config

    if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
        log.Fatalf("cannot read config")
        os.Exit(2)
    }
    return cfg
}
