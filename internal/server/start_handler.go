package server

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
)

const startMessage = `Я бот напоминающий о твоем дне рождения
Нажми "ответить" и отправь мне дату в формате день-месяц и я поздравлю тебя в этот день)
Пример: 05-12`

func (s Server) Start(m *tb.Message) {
	_, _ = s.bot.Send(m.Chat, startMessage)
}

func (s Server) ButtonPrevious(c *tb.Callback) {
	// ...
	// Always respond!
	log.Println(c.Data)
	log.Println(c.Message.Text)
	_ = s.bot.Respond(c, &tb.CallbackResponse{CallbackID: c.ID, Text: "callback text"})
}
