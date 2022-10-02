package event

import (
	"go.mau.fi/whatsmeow"
)

func ClientImpl(c *whatsmeow.Client) *client {
	return &client{WAClient: c}
}

type client struct {
	WAClient *whatsmeow.Client
}
