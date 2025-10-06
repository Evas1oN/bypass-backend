package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Url string `mapstructure:"url"`
}

type ServerConfig struct {
	Address string `mapstructure:"address"`
	Port    string `mapstructure:"port"`
}

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Server   ServerConfig   `mapstructure:"server"`
}

func LoadConfig() (*Config, error) {
	var config Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(".")
	viper.AddConfigPath("config/")

	viper.SetEnvPrefix("BYPASS")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config was not found. Trying to use environment variables")
		} else {
			return nil, err
		}
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
