package main

import (
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vrumg/tg-dob-notify/internal/birthdate_service"
	botLib "github.com/vrumg/tg-dob-notify/internal/bot"
	"github.com/vrumg/tg-dob-notify/internal/repo"
	"github.com/vrumg/tg-dob-notify/internal/server"
)

func main() {

	// init logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// read flags
	// set default values
	configPath, err := parseFlags()
	if err != nil {
		log.Fatal().Str("Failed to parse flags", err.Error())
	}

	// load configuration file
	config, err := loadConfig(configPath)
	if err != nil {
		log.Fatal().Str("Failed to load configuration", err.Error())
	}

	// init telegram bot
	bot, err := botLib.InitTelegramBot(config.Telegram.URL, config.Telegram.Token)
	if err != nil {
		log.Fatal().Str("Failed to init bot", err.Error())
	}

	// init db connection
	conn, err := initDatabaseConnection(
		config.Database.Driver,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
		config.Database.SSL,
	)
	if err != nil {
		log.Fatal().Str("Failed to init db connection", err.Error())
	}

	// init repository
	repository := repo.InitRepo(conn)

	// init service
	service := birthdate_service.InitService(repository)

	// init server instance
	serv, err := server.InitServer(bot, service)
	if err != nil {
		log.Fatal().Str("Failed to init server", err.Error())
	}

	// register server handlers
	serv.RegisterHandlers()

	// start application and bot
	log.Info().Msg("Starting the application")
	bot.Start()
}
