package tgview

import (
	"github.com/petuhovskiy/telegram"

	"github.com/lodthe/is-for-me-bot/tg"
)

const emojis = "👉👈\U0001F97A"
const reversedEmojis = "\U0001F97A👉👈"

type InlineQuery struct {
}

func (InlineQuery) Send(s *tg.Session, query *telegram.InlineQuery) {
	getContent := func(msg string) telegram.InputTextMessageContent {
		return telegram.InputTextMessageContent{
			MessageText:           msg,
			ParseMode:             tg.ParseMode,
			DisableWebPagePreview: true,
		}
	}

	inlineResults := []telegram.InlineQueryResult{
		telegram.InlineQueryResultArticle{
			Type:                "article",
			ID:                  "send_is_for_me_ru",
			Title:               emojis,
			InputMessageContent: getContent(emojis),
			Description:         "Нажми сюда, чтобы отправить " + emojis,
		},
		telegram.InlineQueryResultArticle{
			Type:                "article",
			ID:                  "send_is_for_me_en",
			Title:               reversedEmojis,
			InputMessageContent: getContent(reversedEmojis),
			Description:         "Tap here to send " + reversedEmojis,
		},
	}

	_ = s.AnswerInlineQuery(inlineResults)
}
