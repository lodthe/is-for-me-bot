package tgview

import (
	"github.com/lodthe/is-for-me-bot/tg"
)

type Start struct {
}

func (Start) Send(s *tg.Session) {
	_ = s.SendInlinePhoto(`Привет! Я умею отправлять сообщение «🥺👉👈». Для этого введи @isformebot в поле для сообщения в любом чате.

~~~

Hi! I can send '🥺👉👈' for you. Just type @isformebot in the text input field in any chat. 
`, "is-for-me.png", nil)
}
