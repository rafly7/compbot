package lib

import (
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

type LibImpl struct {
	WAClient *whatsmeow.Client
	Message  *events.Message
}

func LiblImpl(client *whatsmeow.Client, message *events.Message) *LibImpl {
	return &LibImpl{
		WAClient: client,
		Message:  message,
	}
}
