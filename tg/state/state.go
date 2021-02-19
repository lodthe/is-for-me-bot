package state

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type State struct {
	State       ID // Conversation state.
	StateBefore ID // State before handling the current update.

	TelegramID   int
	Username     string
	FirstName    string
	LastName     string
	LanguageCode string

	CannotReceiveMessages bool // It's true, for example, when the user blocked the bot

	LastActionAt time.Time
}

func LoadState(db *gorm.DB, telegramID int) (*State, error) {
	var st StateDB
	err := db.Where(&StateDB{
		TelegramID: telegramID,
	}).Take(&st).Error

	if err == gorm.ErrRecordNotFound {
		var j postgres.Jsonb
		j, err = ToJSON(&State{
			TelegramID: telegramID,
		})
		if err != nil {
			return nil, err
		}

		st = StateDB{
			TelegramID: telegramID,
			State:      j,
		}
		err = db.Create(&st).Error

		log.WithField("telegram_id", telegramID).Info("created a new state entry")
	}

	if err != nil {
		return nil, err
	}
	return FromJSON(st.State)
}

func (s *State) Save(db *gorm.DB) error {
	j, err := ToJSON(s)
	if err != nil {
		return errors.Wrap(err, "failed to marshal state")
	}

	return db.Model(&StateDB{}).
		Where(&StateDB{
			TelegramID: s.TelegramID,
		}).
		Update("state", j).
		Error
}
