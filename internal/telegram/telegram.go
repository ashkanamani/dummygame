package telegram

import (
	"github.com/ashkanamani/dummygame/internal/service"
	"github.com/ashkanamani/dummygame/internal/telegram/teleprompt"
	"gopkg.in/telebot.v4"
	"log/slog"
	"time"
)

type Telegram struct {
	App *service.App
	bot *telebot.Bot

	TelePrompt *teleprompt.TelePrompt
}

func NewTelegram(app *service.App, apiToken string) (*Telegram, error) {
	t := &Telegram{
		App:        app,
		TelePrompt: teleprompt.NewTelePrompt(),
	}
	pref := telebot.Settings{
		Token:   apiToken,
		Poller:  &telebot.LongPoller{Timeout: 60 * time.Second},
		OnError: t.OnError,
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		slog.Error("could not connect to telegram servers", "error", err)
		return nil, err
	}
	slog.Info("bot created and connected to telegram servers")
	t.bot = bot
	
	t.setupHandlers()

	return t, nil
}

func (t *Telegram) Start() {
	t.bot.Start()
}
