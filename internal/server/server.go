package server

import (
	"github.com/vrumg/tg-dob-notify/internal/birthdate_service"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Server main server class. Serves endpoints for telegram events
type Server struct {
	bot     *tb.Bot
	service *birthdate_service.Service
}

// InitServer initialize and populate new Server object
func InitServer(bot *tb.Bot, service *birthdate_service.Service) (*Server, error) {
	serv := &Server{
		bot:     bot,
		service: service,
	}

	serv.registerHandlers()

	return serv, nil
}

// Serve start bot and serve incoming messages
func (s *Server) Serve() {
	s.bot.Start()
}

// registerHandlers registers server handlers in telegram bot
func (s *Server) registerHandlers() {
	// handlers with tags
	s.bot.Handle("/", s.start)
	s.bot.Handle("/help", s.start)
	s.bot.Handle("/start", s.start)
	s.bot.Handle("/bday", s.addBirthdate)
	s.bot.Handle("/feedback", s.feedback)
	// other messages
	s.bot.Handle(tb.OnText, s.readMessage)
}
