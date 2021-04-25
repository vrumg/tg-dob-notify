package bot

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)

// InitTelegramBot initialize telegram bot with token and url
// return bot object
func InitTelegramBot(url string, token string) (*tb.Bot, error) {
	// init middleware
	poller := &tb.LongPoller{Timeout: 10 * time.Second}
	groupOnly := tb.NewMiddlewarePoller(poller, func(upd *tb.Update) bool {
		if upd.Message != nil {
			if !upd.Message.FromGroup() {
				return false
			}
		}

		if upd.Message == nil {
			return false
		}

		return true
	})

	// init telegram bot
	bot, err := tb.NewBot(tb.Settings{
		URL:    url,
		Token:  token,
		Poller: groupOnly,
	})
	if err != nil {
		return nil, err
	}

	return bot, nil
}
