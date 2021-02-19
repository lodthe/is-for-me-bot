package migration

import (
	"github.com/jinzhu/gorm"

	"github.com/lodthe/is-for-me-bot/tg/state"
)

func Migrate(db *gorm.DB) error {
	db.AutoMigrate(
		state.StateDB{},
	)

	return nil
}
