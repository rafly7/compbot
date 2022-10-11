package events

import (
	"github.com/go-redis/redis/v8"
	"go.mau.fi/whatsmeow"
)

func ClientImpl(c *whatsmeow.Client, r *redis.Client) *client {
	return &client{WAClient: c, RClient: r}
}

type client struct {
	WAClient *whatsmeow.Client
	RClient  *redis.Client
}
