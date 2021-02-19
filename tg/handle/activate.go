package handle

import (
	"reflect"

	"github.com/petuhovskiy/telegram"
	log "github.com/sirupsen/logrus"

	"github.com/lodthe/is-for-me-bot/tg"
	"github.com/lodthe/is-for-me-bot/tg/callback"
	"github.com/lodthe/is-for-me-bot/tg/state"
)

type methodState interface {
	// State returns the required state for handling.
	State() state.ID
}

type methodCallback interface {
	// Callback returns a callback with the required type for handling.
	Callback() interface{}
}

type methodCanHandle interface {
	// CanHandle returns true if the handler can handle the update.
	CanHandle(s *tg.Session, msg *telegram.Message, clb *telegram.CallbackQuery) bool
}

type methodCanHandleInlineQuery interface {
	CanHandleInlineQuery(s *tg.Session, query string) bool
}

type methodHandleMessage interface {
	HandleMessage(s *tg.Session, msgText string)
}

type methodHandleCallback interface {
	HandleCallback(s *tg.Session, clb interface{})
}

type methodHandleInlineQuery interface {
	HandleInlineQuery(s *tg.Session, query *telegram.InlineQuery)
}

// activateHandlers goes through the given list of handlers (the same order as they are given)
// and stops when the current handler can handle the given update.
// Then it handles the update with the found handler, and search stops.
func activateHandler(s *tg.Session, update telegram.Update, handlers ...interface{}) {
	activate := func(handler interface{}) {
		logger := log.WithFields(log.Fields{
			"update":      update,
			"telegram_id": s.TelegramID,
			"handler":     handler,
		})

		switch {
		case update.Message != nil:
			handler, ok := handler.(methodHandleMessage)
			if !ok {
				logger.Error("missed HandleMessage method")
			} else {
				handler.HandleMessage(s, update.Message.Text)
			}

		case update.CallbackQuery != nil:
			handler, ok := handler.(methodHandleCallback)
			if !ok {
				logger.Error("missed HandleCallback method")
			} else {
				handler.HandleCallback(s, callback.Unmarshal(update.CallbackQuery.Data))
			}

		case update.InlineQuery != nil:
			handler, ok := handler.(methodHandleInlineQuery)
			if !ok {
				logger.Error("missed HandleInlineQuery method")
			} else {
				handler.HandleInlineQuery(s, update.InlineQuery)
			}

		default:
			logger.Error("the update can be handled, but a callback method is not provided")
		}
	}

	// State-ID-triggered conditions are more valuable than callback-, inline-query- or canHandle-triggered conditions.

	for i := range handlers {
		handler, ok := handlers[i].(methodState)
		if !ok || handler.State() != s.State.StateBefore {
			continue
		}

		activate(handlers[i])

		return
	}

	for i := range handlers {
		var canHandle bool

		handlerByCanHandleInlineQuery, ok := handlers[i].(methodCanHandleInlineQuery)
		if !canHandle && ok && update.InlineQuery != nil {
			canHandle = handlerByCanHandleInlineQuery.CanHandleInlineQuery(s, update.InlineQuery.Query)
		}

		handlerByCallback, ok := handlers[i].(methodCallback)
		if !canHandle && ok && update.CallbackQuery != nil {
			canHandle = reflect.TypeOf(callback.Unmarshal(update.CallbackQuery.Data)) == reflect.TypeOf(handlerByCallback.Callback())
		}

		handlerByCanHandle, ok := handlers[i].(methodCanHandle)
		if !canHandle && ok {
			canHandle = handlerByCanHandle.CanHandle(s, update.Message, update.CallbackQuery)
		}

		if !canHandle {
			continue
		}

		activate(handlers[i])

		return
	}
}
