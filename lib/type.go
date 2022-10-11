package lib

import (
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

type libImpl struct {
	WAClient *whatsmeow.Client
	Message  *events.Message
}

func LiblImpl(client *whatsmeow.Client, message *events.Message) *libImpl {
	return &libImpl{
		WAClient: client,
		Message:  message,
	}
}
