package events

import (
	"compbot/message"
	_ "database/sql"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow/types/events"
)

func (c *client) EventHandler(evt interface{}) {
	switch v := evt.(type) {

	case *events.Message:
		go message.Message(c.WAClient, v, c.RClient)
	}
}
