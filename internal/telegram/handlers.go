package telegram

import (
	"fmt"
	"gopkg.in/telebot.v4"
	"time"
)

func (t *Telegram) setupHandlers() {
	// middlewares
	t.bot.Use(t.registerMiddleware)

	// handlers
	t.bot.Handle("/start", t.start)

	t.bot.Handle(telebot.OnText, func(c telebot.Context) error {
		if t.TelePrompt.Dispatch(c.Sender().ID, c) {
			return nil
		}
		return c.Reply("I did not understand that text!")
	})
}

func (t *Telegram) start(c telebot.Context) error {
	isJustCreated := c.Get("is_just_created").(bool)

	_ = c.Reply("Enter your name")

	ch := t.TelePrompt.Register(c.Sender().ID)

	message, isTimeout := t.TelePrompt.AsMessage(c.Sender().ID, time.Minute)
	if isTimeout {
		_ = c.Reply("Timeout!")
		return nil
	}
	select {
	case val := <-ch:
		_ = c.Reply(fmt.Sprintln("You said: ", val.TeleCtx.Text()))
	case <-time.After(time.Second * 15):
	default:

	}
	return c.Reply(fmt.Sprintln(isJustCreated))

}
