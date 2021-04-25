package server

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

// Server representation of main application server
// works with telegram bot pointer
// contains handlers and
type Server struct {
	bot *tb.Bot
}

func InitServer(bot *tb.Bot) (*Server, error) {
	serv := &Server{
		bot: bot,
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
