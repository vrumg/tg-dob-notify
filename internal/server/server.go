package server

import (
	"github.com/vrumg/tg-dob-notify/internal/birthdate_service"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Server representation of main application server
// works with telegram bot pointer
// contains handlers and
type Server struct {
	bot     *tb.Bot
	service *birthdate_service.Service
}

func InitServer(bot *tb.Bot, service *birthdate_service.Service) (*Server, error) {
	serv := &Server{
		bot:     bot,
		service: service,
	}

	return serv, nil
}

func (s *Server) RegisterHandlers() {

	// start help dialog
	s.bot.Handle("/start", s.Start)
	s.bot.Handle("/", s.Start)

	// common messages parsing
	s.bot.Handle(tb.OnText, s.ReadMsg)
}
