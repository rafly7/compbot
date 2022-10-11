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

		// log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		// log.Printf("[CATEGORY] %s\n", v.Info.Category)
		// log.Printf("[IS GROUP]%t\n", v.Info.IsGroup)
		// log.Printf("[TYPE]%s\n", v.Info.Type)
		// log.Printf("[USER]%s\n", v.Info.Chat.User)
		// log.Printf("[ID]%s\n", v.Info.ID)
		// log.Printf("[MEDIA_TYPE]%s\n", v.Info.MediaType)
		// log.Printf("[MULTICAST]%t\n", v.Info.Multicast)
		// log.Printf("[PUSHNAME]%s\n", v.Info.PushName)
		// log.Printf("[DEVICE]%d\n", v.Info.Chat.Device)
		// log.Printf("[GET MESSAGE]%s\n", v.Message.GetConversation())
		// log.Printf("[GET MESSAGE]%s\n", v.Info.Sender.User)
		// log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		// log.Println("EXECUTED >>>>>>>>>>")
	}
}
