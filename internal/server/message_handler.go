package server

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//ReadMsg
func (s Server) ReadMsg(m *tb.Message) {

	// messages start with botName mention
	r, _ := regexp.Compile("^@birthday_ntf_bot\\s\\d\\d-\\d\\d")

	// split and take only text
	splittedBySpace := strings.Split(m.Text, " ")
	messageLen := len(splittedBySpace)

	// reroute message to start handler if login only provided
	if messageLen == 1 {
		s.Start(m)
		return
	} else if messageLen != 2 {
		_, _ = s.bot.Send(m.Chat,
			"@"+m.Sender.Username+" Я не смог прочитать запись. Проверь, чтобы была только дата через дефис. Год не нужен. Для примера жми /start")
		return
	}

	// validate string
	if !r.MatchString(m.Text) {
		_, _ = s.bot.Send(m.Chat,
			"@"+m.Sender.Username+" Я не смог прочитать запись. Проверь, чтобы была только дата через дефис. Год не нужен. Для примера жми /start")
		return
	}

	// split date to day and month
	splittedDate := strings.Split(splittedBySpace[1], "-")

	// convert date
	day, err := strconv.Atoi(splittedDate[0])
	if err != nil || !(day >= 1 && day <= 31) {
		_, _ = s.bot.Send(m.Chat,
			"@"+m.Sender.Username+" Обманываешь меня, день не тот. Попробуй еще раз. Для примера жми /start")
		return
	}

	// convert month
	month, err := strconv.Atoi(splittedDate[1])
	if err != nil || !(month >= 1 && month <= 12) {
		_, _ = s.bot.Send(m.Chat,
			"@"+m.Sender.Username+" Обманываешь меня, месяц не тот. Попробуй еще раз. Для примера жми /start")
		return
	}

	// call service to add record
	err = s.service.SetBirthdate(m.Chat.ID, m.Sender.Username, month, day)
	if err != nil {
		_, _ = s.bot.Send(m.Chat,
			"@"+m.Sender.Username+" Я не смог добавить тебя в рассылку. Попробуй еще раз. Для примера жми /start")
		return
	}

	// get month name
	monthName := time.Month(month)

	// reply
	message := fmt.Sprintf("@%s Я все записал. Буду поздравлять тебя %d %s",
		m.Sender.Username, day, monthName)
	_, _ = s.bot.Send(m.Chat, message)
}
