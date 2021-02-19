package tg

import (
	"errors"
	"strconv"

	"github.com/petuhovskiy/telegram"
	"github.com/petuhovskiy/telegram/markup"
	log "github.com/sirupsen/logrus"

	"github.com/lodthe/is-for-me-bot/static"
)

const ParseMode = "HTML"

var errBotWasBlocked = "Forbidden: bot was blocked by the user"
var errUserIsDeactivated = "Forbidden: user is deactivated"

func (s *Session) AnswerOnLastCallback() {
	if s.LastUpdate == nil || s.LastUpdate.CallbackQuery == nil {
		return
	}

	_, _ = s.Executor.Execute(func() (interface{}, error) {
		return s.Bot.AnswerCallbackQuery(&telegram.AnswerCallbackQueryRequest{
			CallbackQueryID: s.LastUpdate.CallbackQuery.ID,
		})
	})
}

func (s *Session) sendMessage(text string, keyboard telegram.AnyKeyboard) error {
	if s.State.CannotReceiveMessages {
		return nil
	}
	_, err := s.Executor.Execute(func() (interface{}, error) {
		return s.Bot.SendMessage(&telegram.SendMessageRequest{
			ChatID:                strconv.Itoa(s.TelegramID),
			Text:                  text,
			ParseMode:             ParseMode,
			DisableWebPagePreview: true,
			ReplyMarkup:           keyboard,
		})
	})
	if err != nil {
		log.WithFields(log.Fields{
			"telegram_id":  s.TelegramID,
			"message_text": text,
		}).WithError(err).Error("failed to send the message")
	}

	s.onTelegramError(err)

	return err
}

func (s *Session) editInlineMessage(text string, keyboard *telegram.InlineKeyboardMarkup) error {
	if s.State.CannotReceiveMessages {
		return nil
	}
	_, err := s.Executor.Execute(func() (interface{}, error) {
		return s.Bot.EditMessageText(&telegram.EditMessageTextRequest{
			ChatID:                strconv.Itoa(s.TelegramID),
			MessageID:             s.LastUpdate.CallbackQuery.Message.MessageID,
			InlineMessageID:       s.LastUpdate.CallbackQuery.InlineMessageID,
			Text:                  text,
			ParseMode:             ParseMode,
			DisableWebPagePreview: true,
			ReplyMarkup:           keyboard,
		})
	})

	if err != nil {
		log.WithFields(log.Fields{
			"telegram_id":    s.TelegramID,
			"message_text":   text,
			"callback_query": s.LastUpdate.CallbackQuery,
		}).WithError(err).Error("failed to edit the message")
	}

	s.onTelegramError(err)

	return err
}

func (s *Session) SendText(text string, keyboard ...telegram.AnyKeyboard) error {
	if len(keyboard) == 0 {
		return s.sendMessage(text, nil)
	}

	switch buttons := keyboard[0].(type) {
	case [][]telegram.InlineKeyboardButton:
		return s.sendMessage(text, markup.InlineKeyboard(buttons))

	case [][]telegram.KeyboardButton:
		return s.sendMessage(text, telegram.ReplyKeyboardMarkup{
			Keyboard:        buttons,
			ResizeKeyboard:  true,
			OneTimeKeyboard: true,
			Selective:       true,
		})

	case telegram.ReplyKeyboardRemove:
		return s.sendMessage(text, telegram.ReplyKeyboardRemove{})

	default:
		err := errors.New("unknown keyboard type")
		log.WithField("keyboard", keyboard).WithError(err).Error("failed to send a telegram message")
		return err
	}
}

// SendEditText edits the message received with a callback query.
// Only InlineKeyboardButton keyboard is supported.
func (s *Session) SendEditText(text string, keyboard [][]telegram.InlineKeyboardButton, edit bool) error {
	if !edit || s.LastUpdate.CallbackQuery == nil {
		return s.SendText(text, keyboard)
	}
	return s.editInlineMessage(text, markup.InlineKeyboardMarkup(keyboard))
}

func (s *Session) AnswerInlineQuery(results []telegram.InlineQueryResult) error {
	if s.State.CannotReceiveMessages {
		return nil
	}

	_, err := s.Executor.Execute(func() (interface{}, error) {
		return s.Bot.AnswerInlineQuery(&telegram.AnswerInlineQueryRequest{
			InlineQueryID: s.LastUpdate.InlineQuery.ID,
			Results:       results,
			CacheTime:     1,
			IsPersonal:    true,
		})
	})
	if err != nil {
		log.WithFields(log.Fields{
			"telegram_id":  s.TelegramID,
			"inline_query": s.LastUpdate.InlineQuery.Query,
		}).WithError(err).Error("failed to answer inline query")
	}

	s.onTelegramError(err)

	return err
}

func (s *Session) SendInlinePhoto(text string, file string, keyboard telegram.AnyKeyboard) error {
	if s.State.CannotReceiveMessages {
		return nil
	}
	_, err := s.Executor.Execute(func() (interface{}, error) {
		return s.Bot.SendPhoto(&telegram.SendPhotoRequest{
			ChatID:      strconv.Itoa(s.TelegramID),
			Photo:       static.NewFileReader(file),
			Caption:     text,
			ParseMode:   ParseMode,
			ReplyMarkup: keyboard,
		})
	})
	if err != nil {
		log.WithFields(log.Fields{
			"telegram_id":  s.TelegramID,
			"message_text": text,
			"file":         file,
		}).WithError(err).Error("failed to send inline photo")
	}

	s.onTelegramError(err)

	return err
}

func (s *Session) onTelegramError(err error) {
	if err == nil {
		return
	}

	switch err.Error() {
	case errBotWasBlocked:
		s.State.CannotReceiveMessages = true

	case errUserIsDeactivated:
		s.State.CannotReceiveMessages = true
	}
}

func (s *Session) DeleteLastMessage() error {
	if s.State.CannotReceiveMessages {
		return nil
	}
	if s.LastUpdate.CallbackQuery == nil {
		return errors.New("last update is not a callback query")
	}

	msg := s.LastUpdate.CallbackQuery.Message
	_, err := s.Executor.Execute(func() (interface{}, error) {
		return s.Bot.DeleteMessage(&telegram.DeleteMessageRequest{
			ChatID:    strconv.Itoa(msg.Chat.ID),
			MessageID: msg.MessageID,
		})
	})
	if err != nil {
		log.WithFields(log.Fields{
			"telegram_id": s.TelegramID,
			"message":     msg,
		}).WithError(err).Error("failed to delete the last message")
	}

	s.onTelegramError(err)

	return err
}
