package state

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type StateDB struct { // nolint
	gorm.Model
	TelegramID int
	State      postgres.Jsonb
}

func (s StateDB) TableName() string {
	return "states"
}

func ToJSON(st *State) (postgres.Jsonb, error) {
	buffer, err := json.Marshal(st)
	return postgres.Jsonb{RawMessage: json.RawMessage(buffer)}, err
}

func FromJSON(j postgres.Jsonb) (*State, error) {
	st := &State{}
	err := json.Unmarshal(j.RawMessage, st)
	return st, err
}
