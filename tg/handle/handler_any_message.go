package handle

import (
	"github.com/petuhovskiy/telegram"

	"github.com/lodthe/is-for-me-bot/tg"
	"github.com/lodthe/is-for-me-bot/tg/tgview"
)

type AnyMessageHandler struct {
}

func (h *AnyMessageHandler) CanHandle(s *tg.Session, msg *telegram.Message, clb *telegram.CallbackQuery) bool {
	return msg != nil
}

func (h *AnyMessageHandler) HandleMessage(s *tg.Session, msgText string) {
	tgview.Demonstration{}.Send(s)
}
