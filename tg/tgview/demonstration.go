package tgview

import (
	"github.com/lodthe/is-for-me-bot/tg"
)

type Demonstration struct {
}

func (Demonstration) Send(s *tg.Session) {
	_ = s.SendInlinePhoto(`Введи @isformebot в поле для сообщения в любом чате.

~~~

Just type @isformebot in the text input field in any chat. 
`, "demonstration.png", nil)
}
