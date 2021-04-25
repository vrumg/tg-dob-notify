package server

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

// start greeting message handler
func (s Server) start(m *tb.Message) {
	const startMessage = `
Я бот напоминающий о твоем дне рождения!
Доступные команды:

Справка
/start

Установить поздравление на день рождения в формате день.месяц
/bday 01.12

Связаться с разработчиком
/feedback текст сообщения`

	_, _ = s.bot.Send(m.Chat, startMessage)
}
