package tgview

import (
	"github.com/petuhovskiy/telegram"

	"github.com/lodthe/is-for-me-bot/tg"
)

const emojis = "👉👈\U0001F97A"

type InlineQuery struct {
}

func (InlineQuery) Send(s *tg.Session, query *telegram.InlineQuery) {
	content := telegram.InputTextMessageContent{
		MessageText:           emojis,
		ParseMode:             tg.ParseMode,
		DisableWebPagePreview: true,
	}

	inlineResults := []telegram.InlineQueryResult{
		telegram.InlineQueryResultArticle{
			Type:                "article",
			ID:                  "send_is_for_me_ru",
			Title:               emojis,
			InputMessageContent: content,
			Description:         "Нажми сюда, чтобы отправить " + emojis,
		},
		telegram.InlineQueryResultArticle{
			Type:                "article",
			ID:                  "send_is_for_me_en",
			Title:               emojis,
			InputMessageContent: content,
			Description:         "Tap here to send " + emojis,
		},
	}

	_ = s.AnswerInlineQuery(inlineResults)
}
