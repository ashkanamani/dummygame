package telegram

import (
	"context"
	"fmt"
	"gopkg.in/telebot.v4"
	"log/slog"
)

func (t *Telegram) startHandler(c telebot.Context) error {
	isJustCreated := c.Get("is_just_created").(bool)
	account := GetAccount(c)
	if !isJustCreated {
		return t.myInfoHandler(c)
	}
	_ = c.Reply("Welcome to the game. Enter your name")

	msg, err := t.Input(c, InputConfig{
		Prompt:    "👋Welcome to Battle of Kings\nPlease enter your display name?",
		OnTimeout: "Timeout!!!",
	})
	if err != nil {
		return err
	}

	displayName := msg.Text
	//TODO: Validation
	account.DisplayName = displayName
	if err := t.App.Account.Update(context.Background(), account); err != nil {
		slog.Error("error while updating account", "error", err)
		return err
	}
	c.Set("account", account)
	_ = c.Reply(fmt.Sprintf("✅ We call you %s from now.", displayName))
	return c.Reply(fmt.Sprintf("👑 King <<%s>>\nWelcome to the Battle of Kings game.\nHow can I help you?", displayName))

}

func (t *Telegram) myInfoHandler(c telebot.Context) error {
	account := GetAccount(c)
	return c.Reply(fmt.Sprintf("👑 King <<%s>>\nWelcome to the Battle of Kings game.\nHow can I help you?", account.DisplayName))

}
