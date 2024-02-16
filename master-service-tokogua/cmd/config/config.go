package config

import (
	"github.com/go-yaml/yaml"
	"log"
	"os"
)

type Config struct {
	Database struct {
		Name     string `yaml:"name"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`

	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Type string `yaml:"type"`
	} `yaml:"server"`
}

func InitConfig() *Config {
	f, err := os.Open("config/config.yml")
	if err != nil {
		if err != nil {
			log.Fatal(err)
		}
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf(err.Error())
		}
	}(f)

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}
