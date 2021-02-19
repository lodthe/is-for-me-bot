package handle

import (
	"strings"

	"github.com/petuhovskiy/telegram"

	"github.com/lodthe/is-for-me-bot/tg"
	"github.com/lodthe/is-for-me-bot/tg/tgview"
)

type StartHandler struct {
}

func (h *StartHandler) CanHandle(s *tg.Session, msg *telegram.Message, clb *telegram.CallbackQuery) bool {
	return msg != nil && strings.HasPrefix(msg.Text, "/start")
}

func (h *StartHandler) HandleMessage(s *tg.Session, msgText string) {
	tgview.Start{}.Send(s)
}
