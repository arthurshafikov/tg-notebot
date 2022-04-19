package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseConfig    `mapstructure:",squash"`
	TelegramBotConfig `mapstructure:",squash"`
	Messages
}

type TelegramBotConfig struct {
	APIKey string `mapstructure:"BOT_API_KEY"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"MONGODB_HOST"`
	Port     string `mapstructure:"MONGODB_PORT"`
	Username string `mapstructure:"MONGODB_USER"`
	Password string `mapstructure:"MONGODB_PASSWORD"`
}

type Messages struct {
	Start string
}

func NewConfig(envPath, configFolder string) *Config {
	var config Config

	// Read from yml
	viper.AddConfigPath(configFolder)
	viper.SetConfigName("main")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalln(err)
	}

	// Read from env
	viper.AddConfigPath(envPath)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalln(err)
	}

	return &config
}
