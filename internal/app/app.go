package app

import (
	"context"
	"flag"
	"log"

	"github.com/arthurshafikov/tg-notebot/internal/config"
	"github.com/arthurshafikov/tg-notebot/internal/repository"
	"github.com/arthurshafikov/tg-notebot/internal/repository/mongodb"
	"github.com/arthurshafikov/tg-notebot/internal/services"
	"github.com/arthurshafikov/tg-notebot/internal/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	envPath          string
	configFolderPath string
)

func init() {
	flag.StringVar(&envPath, "env", "./", "Path to .env file folder")
	flag.StringVar(&configFolderPath, "cfgFolder", "./configs", "Path to configs folder")
}

func Run() {
	ctx := context.Background()
	config := config.NewConfig(envPath, configFolderPath)

	botApi, err := tgbotapi.NewBotAPI(config.TelegramBotConfig.APIKey)
	if err != nil {
		log.Fatal(err)
	}

	mongo, err := mongodb.NewMongoDB(ctx, mongodb.Config{
		Host:     config.DatabaseConfig.Host,
		Port:     config.DatabaseConfig.Port,
		Username: config.DatabaseConfig.Username,
		Password: config.DatabaseConfig.Password,
	})
	if err != nil {
		log.Fatalln(err)
	}

	repository := repository.NewRepository(mongo)

	services := services.NewServices(services.Deps{
		Repository: repository,
	})

	telegramBot := telegram.NewTelegramBot(ctx, botApi, services, config.Messages)

	if err := telegramBot.Start(); err != nil {
		log.Fatalln(err)
	}
}
