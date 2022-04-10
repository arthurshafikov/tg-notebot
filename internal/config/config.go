package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppConfig      `mapstructure:",squash"`
	DatabaseConfig `mapstructure:",squash"`
}

type AppConfig struct {
	HTTPPort string `mapstructure:"HTTP_PORT"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"MONGODB_HOST"`
	Port     string `mapstructure:"MONGODB_PORT"`
	Username string `mapstructure:"MONGODB_USER"`
	Password string `mapstructure:"MONGODB_PASSWORD"`
}

func NewConfig(envPath string) *Config {
	viper.AddConfigPath("./")

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalln(err)
	}

	return &config
}
