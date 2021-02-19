package tg

import (
	"github.com/jinzhu/gorm"
	"github.com/petuhovskiy/telegram"

	"github.com/lodthe/is-for-me-bot/conf"
	"github.com/lodthe/is-for-me-bot/tg/tglimiter"
)

// General stores the general fields for a session.
type General struct {
	Bot      *telegram.Bot
	Executor *tglimiter.Executor

	DB     *gorm.DB
	Config conf.Config
}
