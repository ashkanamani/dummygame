package telegram

import (
	"errors"
	"gopkg.in/telebot.v4"
)

var (
	ErrInputTimeout = errors.New("input timeout")
)

type Confirm struct {
	ConfirmText func(msg *telebot.Message) string
}
type InputConfig struct {
	Prompt         any
	PromptKeyboard [][]string
	OnTimeout      any
	Confirm        Confirm
}

func (t *Telegram) Input(c telebot.Context, config InputConfig) (*telebot.Message, error) {
	if config.Prompt == nil {
		_ = c.Reply(config.Prompt)
	}
	response, isTimeout := t.TelePrompt.AsMessage(c.Sender().ID, DefaultInputTimeout)
	if isTimeout {
		if config.OnTimeout != nil {
			_ = c.Reply(config.OnTimeout)
		} else {
			_ = c.Reply(DefaultTimeoutText)
		}
		return nil, ErrInputTimeout
	}

	// client has to confirm
	if config.Confirm.ConfirmText != nil {
		configText := config.Confirm.ConfirmText(response)
		c.Reply(configText)
	}
	return response, nil
}

//func (t *Telegram) generateKeyboard(rows [][]string) *telebot.ReplyMarkup {
//	mu := &telebot.ReplyMarkup{
//		ResizeKeyboard:  true,
//		OneTimeKeyboard: true,
//	}
//
//	for _, row := range rows {
//		mu.Reply(mu.Row())
//	}
//}
