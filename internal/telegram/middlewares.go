package telegram

import (
	"context"
	"github.com/ashkanamani/dummygame/internal/entity"
	"gopkg.in/telebot.v4"
	"time"
)

func (t *Telegram) registerMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		acc := entity.Account{
			ID:        c.Sender().ID,
			FirstName: c.Sender().FirstName,
			Username:  c.Sender().Username,
		}
		if err := t.App.Account.CreateOrUpdate(context.Background(), acc); err != nil {
			return err
		}
		return next(c)
	}
}
