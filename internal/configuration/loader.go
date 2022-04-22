package configuration

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Configuration composed configuration structure: Api, Server, DataBase
type Configuration struct {
	Api
	DataBase
	Server
	Redis
	JWT
}

// Api version structure, Version string
type Api struct {
	Version string `yaml:"version"`
}

// DataBase configuration structure
type DataBase struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

// Server configuration structure, Port string
type Server struct {
	Port string `yaml:"port"`
}

// Reddis configuration structure, Host string Port int
type Redis struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// JWT access secret
type JWT struct {
	Secret string `yaml:"secret"`
}

// Load configuration from configuration.yml
func Load(fileName string) (*Configuration, error) {
	var config Configuration

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	_ = os.Setenv("JWT_ACCESS_SECRET", config.JWT.Secret)

	return &config, nil
}
