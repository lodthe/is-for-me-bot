package tg

import (
	"github.com/petuhovskiy/telegram"
	log "github.com/sirupsen/logrus"

	"github.com/lodthe/is-for-me-bot/tg/state"
)

type Session struct {
	General

	TelegramID int
	LastUpdate *telegram.Update

	State *state.State
}

func NewSession(telegramID int, general *General, update *telegram.Update) (*Session, error) {
	st, err := state.LoadState(general.DB, telegramID)
	if err != nil {
		log.WithField("telegram_id", telegramID).WithError(err).Error("failed to load the state")
		return nil, err
	}

	return &Session{
		General:    *general,
		TelegramID: telegramID,
		LastUpdate: update,
		State:      st,
	}, nil
}

func (s *Session) SaveState() error {
	return s.State.Save(s.DB)
}
