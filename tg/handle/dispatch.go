package handle

import (
	"reflect"
	"runtime/debug"
	"time"

	"github.com/petuhovskiy/telegram"
	log "github.com/sirupsen/logrus"

	"github.com/lodthe/is-for-me-bot/tg"
	"github.com/lodthe/is-for-me-bot/tg/callback"
)

func dispatchUpdate(general *tg.General, sessionTelegramID int, update telegram.Update) {
	defer func() {
		if r := recover(); r != nil {
			log.WithFields(log.Fields{
				"recovered":           r,
				"session_telegram_id": sessionTelegramID,
				"stacktrace":          string(debug.Stack()),
				"update":              update,
			}).Error("recovered from panic")
		}
	}()

	s, err := tg.NewSession(sessionTelegramID, general, &update)
	if err != nil {
		log.WithFields(log.Fields{
			"session_telegram_id": sessionTelegramID,
			"update":              update,
		}).WithError(err).Error("failed to create the session")
		return
	}

	if update.CallbackQuery != nil {
		clb := callback.Unmarshal(update.CallbackQuery.Data)
		log.WithFields(log.Fields{
			"telegram_id": sessionTelegramID,
			"type_name":   reflect.TypeOf(clb).Name(),
		}).Info("unpack a callback")
	}

	s.State.StateBefore = s.State.State
	s.State.LastActionAt = time.Now()

	// If we receive an update from the user, they can receive our messages.
	s.State.CannotReceiveMessages = false

	updateUserInfo(s, update)

	s.AnswerOnLastCallback()
	activateHandler(s, update,
		&StartHandler{},
		&InlineHandler{},

		&AnyMessageHandler{},
	)

	err = s.SaveState()
	if err != nil {
		log.WithField("telegram_id", sessionTelegramID).WithError(err).Error("failed to save the state")
	}
}

// updateUserInfo saves the user's name, username and language code.
func updateUserInfo(s *tg.Session, update telegram.Update) {
	var user *telegram.User

	switch {
	case update.Message != nil && update.Message.From != nil:
		user = update.Message.From

	case update.CallbackQuery != nil && update.CallbackQuery.From != nil:
		user = update.CallbackQuery.From

	case update.InlineQuery != nil && update.InlineQuery.From != nil:
		user = update.InlineQuery.From
	}

	if user == nil {
		return
	}

	s.State.Username = user.Username
	s.State.FirstName = user.FirstName
	s.State.LastName = user.LastName
	s.State.LanguageCode = user.LanguageCode
}
