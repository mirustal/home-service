package config

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type GRPCConfig struct {
	PortAuth       string `yaml:"portauth"`
	PortHomes      string `yaml:"porthomes"`
	TimeOut string `yaml:"timeout"`
}

type Config struct {
	ModeLog    string      `yaml:"modelog" env-default:"debug"`
	ServerPort int         `yaml:"serverport"`
	GRPC       *GRPCConfig `yaml:"grpc"`
	DB         *DBConfig   `yaml:"db"`
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

func fetchConfig() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	fmt.Println(res)
	flag.Parse()

	return res
}
