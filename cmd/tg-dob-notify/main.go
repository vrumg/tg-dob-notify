package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	botLib "github.com/vrumg/tg-dob-notify/internal/bot"
	repo2 "github.com/vrumg/tg-dob-notify/internal/repo"
	"github.com/vrumg/tg-dob-notify/internal/server"
	"log"
)

func main() {

	// read flags
	// set default values
	configPath, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}

	// load configuration file
	config, err := loadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// init telegram bot
	bot, err := botLib.InitTelegramBot(config.Telegram.URL, config.Telegram.Token)
	if err != nil {
		log.Fatal(err)
	}

	// init db connection
	conn, err := connect()
	if err != nil {
		log.Fatal(err)
	}

	// init repository
	repo := repo2.InitRepo(conn)
	_ = repo

	// init server instance
	serv, err := server.InitServer(bot)
	if err != nil {
		log.Fatal(err)
		return
	}

	// register server handlers
	serv.RegisterHandlers()

	// start application and bot
	log.Println("Starting application")
	bot.Start()
}

func connect() (*sqlx.DB, error) {
	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("postgres", "user=postgres password=admin dbname=tgbot sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}
