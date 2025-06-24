package telegram

import (
	"github.com/ashkanamani/dummygame/internal/service"
	"github.com/ashkanamani/dummygame/internal/telegram/teleprompt"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v4"
	"time"
)

type Telegram struct {
	App *service.App
	bot *telebot.Bot

	TelePrompt *teleprompt.TelePrompt
}

func NewTelegram(app *service.App, apiToken string) (*Telegram, error) {
	pref := telebot.Settings{
		Token:  apiToken,
		Poller: &telebot.LongPoller{Timeout: 60 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		logrus.WithError(err).Error("could not connect to telegram servers")
		return nil, err
	}
	t := &Telegram{
		App:        app,
		bot:        bot,
		TelePrompt: teleprompt.NewTelePrompt(),
	}

	t.setupHandlers()

	return t, nil
}

func (t *Telegram) Start() {
	t.bot.Start()
}
