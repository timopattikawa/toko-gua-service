package config

import (
	"github.com/go-yaml/yaml"
	"log"
	"os"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Type string `yaml:"type"`
	} `yaml:"client"`
	GRPCMaster struct {
		Target string `yaml:"target"`
	} `yaml:"grpc"`
	Midtrans struct {
		ServerKey string `yaml:"server-key"`
		ClientKey string `yaml:"client-key"`
	} `yaml:"midtrans"`
}

func InitConfiguration() *Config {
	file, err := os.Open("config/config.yml")
	if err != nil {
		log.Fatal("Failed to open config.yml")
	}

	var config = &Config{}
	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(config); err != nil {
		log.Fatalln("Failed to decode config")
	}
	return config
}
