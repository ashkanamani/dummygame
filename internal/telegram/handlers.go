package telegram

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/telebot.v4"
	"log/slog"
)

func (t *Telegram) setupHandlers() {
	// middlewares
	t.bot.Use(t.registerMiddleware)

	// handlers
	t.bot.Handle("/start", t.startHandler)

	t.bot.Handle(telebot.OnText, t.TextHandler)
}

func (t *Telegram) TextHandler(c telebot.Context) error {
	canBeDispatched := t.TelePrompt.Dispatch(c.Sender().ID, c)
	if canBeDispatched {
		return nil
	}

	// per state
	return c.Reply("I did not understand that text!")
}

func (t *Telegram) OnError(err error, c telebot.Context) {
	if errors.Is(err, ErrInputTimeout) {
		return
	}
	errorId := uuid.New().String()

	slog.Error("unhandled telegram error:", "tracing_id", errorId)
	_ = c.Reply(fmt.Sprint("there is an error in processing messages. error id:", errorId))
}
