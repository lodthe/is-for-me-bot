package callback

import (
	"github.com/petuhovskiy/telegram"
)

// Button creates a new inline keyboard button with the given label.
// Clb is encoded to the callback_data field.
// Clb can be restored from the callback_data field using callback.Unmarshal.
func Button(label string, clb interface{}) telegram.InlineKeyboardButton {
	return telegram.InlineKeyboardButton{
		Text:         label,
		CallbackData: Marshal(clb),
	}
}
