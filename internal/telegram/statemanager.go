package telegram

import "gopkg.in/telebot.v4"

type StateKey string

type State func(c telebot.Context) (StateKey, error)

type StateManager struct {
	states map[StateKey]State
}

func (s *StateManager) Register(key StateKey, f State) {
	s.states[key] = f
}
