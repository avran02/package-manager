package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SSHConfig
}

type SSHConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	// CertPath string
}

func (c SSHConfig) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func New() Config {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configFileDir := home + "/.config/package-manager/.env"
	if err := godotenv.Load(configFileDir); err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	return Config{
		SSHConfig: SSHConfig{
			Host:     os.Getenv("SSH_HOST"),
			Port:     os.Getenv("SSH_PORT"),
			User:     os.Getenv("SSH_USER"),
			Password: os.Getenv("SSH_PASSWORD"),
			// CertPath: os.Getenv("CERT_PATH"),
		},
	}
}
