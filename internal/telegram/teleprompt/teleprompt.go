package teleprompt

import (
	"fmt"
	"gopkg.in/telebot.v4"
	"log/slog"
	"sync"
	"time"
)

type Prompt struct {
	TeleCtx telebot.Context
}

type TelePrompt struct {
	accountPrompts sync.Map
}

func NewTelePrompt() *TelePrompt {
	return &TelePrompt{}

}

func (t *TelePrompt) Register(userID int64) <-chan Prompt {
	ch := make(chan Prompt, 1)

	if preChannel, loaded := t.accountPrompts.LoadAndDelete(userID); loaded {
		close(preChannel.(chan Prompt))
	}

	t.accountPrompts.Store(userID, ch)

	return ch
}

func (t *TelePrompt) Dispatch(userID int64, c telebot.Context) bool {
	ch, loaded := t.accountPrompts.LoadAndDelete(userID)
	if !loaded {
		slog.Info("channel of currect user did not loaded in Dispatch function.")
		return false
	}
	slog.Info("channel of currect user loaded in Dispatch function.")
	select {
	case ch.(chan Prompt) <- Prompt{TeleCtx: c}:
		fmt.Println("data sended:")
	default:
		return false
	}
	return true
}

func (t *TelePrompt) AsMessage(userID int64, timeout time.Duration) (*telebot.Message, bool) {
	ch := t.Register(userID)
	select {
	case val := <-ch:
		fmt.Println("data received:", val)
		return val.TeleCtx.Message(), false
	case <-time.After(timeout):
		return nil, true
	}
}
