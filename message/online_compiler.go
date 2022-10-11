package message

import (
	"compbot/configs"
	"compbot/lib"
	"compbot/services"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/go-redis/redis/v8"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func onlineCompilerConversation(msg *events.Message, rdb *redis.Client, infoSender string, l *lib.LibImpl) {
	extendedMsg := msg.Message.GetExtendedTextMessage()
	if extendedMsg != nil {
		if msg.Message.ExtendedTextMessage.ContextInfo != nil && msg.Message.ExtendedTextMessage.Text != nil {
			// log.Print(msg.Message.ExtendedTextMessage)
			val, err := rdb.Get(context.Background(), infoSender).Result()
			if err != redis.Nil {
				if err != nil {
					log.Print(err)
					return
				}
				if val == *msg.Message.ExtendedTextMessage.ContextInfo.StanzaId {
					replyBot := *msg.Message.ExtendedTextMessage.ContextInfo.QuotedMessage.Conversation
					receiveText := *msg.Message.ExtendedTextMessage.Text
					switch replyBot {
					case replyPrepareLanguage("C"):
						{
							body := *payloadCodeCompiler("c", 5, receiveText)
							res, err := services.RunCode(body)
							if err != nil {
								log.Print(err)
							}
							res = strings.ReplaceAll(res, "jdoodle.c", "main.c")
							output := fmt.Sprintf("*Output:*\n%s", res)
							deleteInfoUserOnlineCompiler(rdb, infoSender, func() (whatsmeow.SendResponse, error) {
								return l.SendReplyMessage(output)
							})
						}
					case replyPrepareLanguage("C++"):
						{
							body := *payloadCodeCompiler("cpp17", 1, receiveText)
							res, err := services.RunCode(body)
							if err != nil {
								log.Print(err)
							}
							res = strings.ReplaceAll(res, "jdoodle.cpp", "main.cpp")
							output := fmt.Sprintf("*Output:*\n%s", res)
							deleteInfoUserOnlineCompiler(rdb, infoSender, func() (whatsmeow.SendResponse, error) {
								return l.SendReplyMessage(output)
							})
						}
					case replyPrepareLanguage("Python 2"):
						{
							body := *payloadCodeCompiler("python2", 3, receiveText)
							res, err := services.RunCode(body)
							if err != nil {
								log.Print(err)
							}
							res = strings.ReplaceAll(res, "jdoodle.py", "main.py")
							output := fmt.Sprintf("*Output:*\n%s", res)
							deleteInfoUserOnlineCompiler(rdb, infoSender, func() (whatsmeow.SendResponse, error) {
								return l.SendReplyMessage(output)
							})
						}
					case replyPrepareLanguage("Python 3"):
						{
							body := *payloadCodeCompiler("python3", 4, receiveText)
							res, err := services.RunCode(body)
							if err != nil {
								log.Print(err)
							}
							res = strings.ReplaceAll(res, "jdoodle.py", "main.py")
							output := fmt.Sprintf("*Output:*\n%s", res)
							deleteInfoUserOnlineCompiler(rdb, infoSender, func() (whatsmeow.SendResponse, error) {
								return l.SendReplyMessage(output)
							})
						}
					case replyPrepareLanguage("Nodejs"):
						{
							body := *payloadCodeCompiler("nodejs", 4, receiveText)
							res, err := services.RunCode(body)
							if err != nil {
								log.Print(err)
							}
							res = strings.ReplaceAll(res, "jdoodle.js", "main.js")
							output := fmt.Sprintf("*Output:*\n%s", res)
							deleteInfoUserOnlineCompiler(rdb, infoSender, func() (whatsmeow.SendResponse, error) {
								return l.SendReplyMessage(output)
							})
						}
					}
				}
			}

		}
	}
}

func deleteInfoUserOnlineCompiler(rdb *redis.Client, infoSender string, cb func() (whatsmeow.SendResponse, error)) {
	_, errCB := cb()
	if errCB != nil {
		log.Print(errCB)
		return
	}
	_, errRDB := rdb.Del(context.Background(), infoSender).Result()
	if errRDB != nil {
		log.Print(errRDB)
		return
	}
}

func payloadCodeCompiler(language string, versionIndex int, receiveText string) *map[string]interface{} {
	return &map[string]interface{}{
		"language":     language,
		"versionIndex": fmt.Sprintf("%d", versionIndex),
		"clientId":     configs.ClientID(),
		"clientSecret": configs.ClientSecret(),
		"script":       receiveText,
	}
}
