package server

import tb "gopkg.in/tucnak/telebot.v2"

// feedback send message to developer
func (s Server) feedback(m *tb.Message) {
	const responseMessage = `Я отправил твое сообщение разработчику`

	_, _ = s.bot.Send(m.Chat, responseMessage)
}
