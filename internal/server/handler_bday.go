package server

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	tb "gopkg.in/tucnak/telebot.v2"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type bday struct {
	day       int
	month     int
	monthName time.Month
}

// addBirthdate save user birthdate and set notification
func (s Server) addBirthdate(m *tb.Message) {

	const errMessage = "Я не смог прочитать запись. Проверь запрос по справке /start"
	const successMessage = "Я все записал. Жди поздравления"

	// set user login tag to reply
	appealToUser := "@" + m.Sender.Username + " "

	text := s.stripAddBirthdateRequest(m.Text)
	if text == "" {
		_, _ = s.bot.Send(m.Chat, appealToUser+errMessage)
		return
	}

	if !s.isValidAddBirthdateText(text) {
		_, _ = s.bot.Send(m.Chat, appealToUser+errMessage)
		return
	}

	date, err := s.convertToDayAndMonth(text)
	if err != nil {
		_, _ = s.bot.Send(m.Chat, appealToUser+errMessage)
		return
	}

	err = s.service.SaveUserBirthdate(m.Chat.ID, m.Sender.Username, date.month, date.day)
	if err != nil {
		_, _ = s.bot.Send(m.Chat, appealToUser+errMessage)
		return
	}

	message := fmt.Sprintf("%s %d %s", successMessage, date.day, date.monthName)
	_, _ = s.bot.Send(m.Chat, message)
}

// stripAddBirthdateRequest strip incoming request from metainfo to text
func (s Server) stripAddBirthdateRequest(request string) string {
	// split and take only text
	splitBySpace := strings.Split(request, " ")
	messageLen := len(splitBySpace)

	// request length is exact 2 blocks
	if messageLen != 2 {
		log.Warn().Msg("parsed message is not 2 blocks")
		return ""
	}

	return splitBySpace[1]
}

// isValidAddBirthdateText validate text to contain only formatted date DD.MM
func (s Server) isValidAddBirthdateText(text string) bool {
	r, _ := regexp.Compile("^\\d\\d.\\d\\d")

	if !r.MatchString(text) {
		log.Warn().Msg("validation failed")
		return false
	}

	return true
}

// convertToDayAndMonth split and convert text to day number, month number, month name
func (s Server) convertToDayAndMonth(text string) (*bday, error) {
	var err error
	date := &bday{}

	// split date to day and month
	splittedDate := strings.Split(text, ".")

	// convert date to int
	date.day, err = strconv.Atoi(splittedDate[0])
	if err != nil || !(date.day >= 1 && date.day <= 31) {
		log.Warn().
			Str("day", splittedDate[0]).
			Msg("day cannot be converted to int")
		return nil, errors.New("")
	}

	// convert month to int
	date.month, err = strconv.Atoi(splittedDate[1])
	if err != nil || !(date.month >= 1 && date.month <= 12) {
		log.Warn().
			Str("month", splittedDate[1]).
			Msg("month cannot be converted to int")
		return nil, errors.New("")
	}

	// get month name
	mStr := time.Month(date.month)
	date.monthName = mStr

	return date, nil
}
