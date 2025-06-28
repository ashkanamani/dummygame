package telegram

import (
	"github.com/ashkanamani/dummygame/internal/entity"
	"gopkg.in/telebot.v4"
	"time"
)

var (
	DefaultInputTimeout = 5 * time.Minute
	DefaultTimeoutText  = "We were waiting for your message. Please send message when you come back."
)

func GetAccount(c telebot.Context) entity.Account {
	return c.Get("account").(entity.Account)
}
