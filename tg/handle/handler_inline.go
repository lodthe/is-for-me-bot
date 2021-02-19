package handle

import (
	"github.com/petuhovskiy/telegram"

	"github.com/lodthe/is-for-me-bot/tg"
	"github.com/lodthe/is-for-me-bot/tg/tgview"
)

type InlineHandler struct {
}

func (h *InlineHandler) CanHandleInlineQuery(s *tg.Session, query string) bool {
	return true
}

func (h *InlineHandler) HandleInlineQuery(s *tg.Session, query *telegram.InlineQuery) {
	tgview.InlineQuery{}.Send(s, query)
}
