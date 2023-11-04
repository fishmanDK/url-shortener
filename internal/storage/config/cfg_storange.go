package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)



type Config_Postgres struct {
	Postgres struct {
		Host     string `yaml:"host"`
		Name     string `yaml:"name"`
		Password string `yaml:"password"`
		User     string `yaml:"user"`
		Port     string `yaml:"port"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"postgres"`
}

func NewConfig_Postgres() *Config_Postgres {
	configPath := "internal/storage/config/storage.yaml"

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config_Postgres
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot reading config: %s", configPath)
	}

	return &cfg
}
