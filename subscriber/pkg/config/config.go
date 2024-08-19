package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type JetConfig struct {
	Address string `yaml:"address"`
	Subject string `yaml:"subsject"`
	Name    string `yaml:"name"`
}

type Config struct {
	Jet *JetConfig `yaml"jet"`
}

func LoadConfig(fileName, fileType string) (*Config, error) {
	var cfg *Config
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(fileName)
	v.AddConfigPath("./configs")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return nil, err
	}
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to decode into struct, %w", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return cfg, nil
}
