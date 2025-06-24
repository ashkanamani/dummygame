package telegram

import (
	"context"
	"github.com/ashkanamani/dummygame/internal/entity"
	"gopkg.in/telebot.v4"
)

func (t *Telegram) registerMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		acc := entity.Account{
			ID:        c.Sender().ID,
			FirstName: c.Sender().FirstName,
			Username:  c.Sender().Username,
		}
		acc, created, err := t.App.Account.CreateOrUpdate(context.Background(), acc)
		if err != nil {
			return err
		}
		c.Set("account", acc)
		c.Set("is_just_created", created)
		return next(c)
	}
}
