package config

import (
	"log"

	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	DBHost     string `yaml:"db_host"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBName     string `yaml:"db_name"`
	APIURL     string `yaml:"api_url"`
}

var AppConfig Config

func LoadConfig() {
	file, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	err = yaml.Unmarshal(file, &AppConfig)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}
}
