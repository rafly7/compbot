package message

import (
	"bytes"
	"compbot/configs"
	"compbot/lib"
	"compbot/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

const infoBot string = `*Compbot online compiler bot*
_List command yang tersedia:_

!compbot@c				gcc compiler
!compbot@cpp			g++ compiler
!compbot@python2	python2 interpreted
!compbpt@python3	python3 interpreted
!compbot@nodejs		node interpreted
			`

func replyPrepareLanguage(language string) string {
	return fmt.Sprintf("Ok, beri saya beberapa kode %s untuk dieksekusi", language)
}

func Message(client *whatsmeow.Client, msg *events.Message, rdb *redis.Client) {
	l := lib.LiblImpl(client, msg)

	if msg.Info.IsGroup {
		// var messageExtended = make(chan proto.Message)
		if msg.Message.GetConversation() == "!compbot" {
			l.SendInfoBotMessage(infoBot, "Bot: compbot TPLE 09")
		} else if msg.Message.GetConversation() == "!compbot@c" {
			l.SendReplyMessage(replyPrepareLanguage("C"))
		} else if msg.Message.GetConversation() == "!compbot@cpp" {
			l.SendReplyMessage(replyPrepareLanguage("C++"))
		} else if msg.Message.GetConversation() == "!compbot@python2" {
			l.SendReplyMessage(replyPrepareLanguage("Python 2"))
		} else if msg.Message.GetConversation() == "!compbot@python3" {
			_, err := l.SendReplyMessage(replyPrepareLanguage("Python 3"))
			if err != nil {
				log.Print(err)
			}
			comp := fmt.Sprintf("%s@compbot", msg.Info.Sender.User)
			val, err := rdb.Get(context.Background(), comp).Result()
			if err == redis.Nil {
				log.Print(err)
			}
			if val == "" {
				_, err = rdb.SetNX(context.Background(), comp, "test", 5*time.Minute).Result()
				if err == redis.Nil {
					log.Print(err)
				}
			} else {
				_, err := rdb.Del(context.Background(), comp).Result()
				if err == redis.Nil {
					// log.Print(err)
				}
				// log.Println(val)
				_, err = rdb.SetNX(context.Background(), comp, "testwkwkwk", 5*time.Minute).Result()
				if err == redis.Nil {
					// log.Print(err)
				}
			}
			// log.Println(set)
		} else if msg.Message.GetConversation() == "!compbot@nodejs" {
			l.SendReplyMessage(replyPrepareLanguage("Nodejs"))
		} else if msg.Message.ExtendedTextMessage != nil {
			replyBot := *msg.Message.ExtendedTextMessage.ContextInfo.QuotedMessage.Conversation
			receiveText := *msg.Message.ExtendedTextMessage.Text
			switch replyBot {
			case replyPrepareLanguage("C"):
				{
					body := *payloadCodeCompiler("c", 5, receiveText)
					res, err := onlineCompiler(body)
					if err != nil {
						log.Print(err)
					}
					res = strings.ReplaceAll(res, "jdoodle.c", "main.c")
					output := fmt.Sprintf("*Output:*\n%s", res)
					l.SendReplyMessage(output)
				}
			case replyPrepareLanguage("C++"):
				{
					body := *payloadCodeCompiler("cpp17", 1, receiveText)
					res, err := onlineCompiler(body)
					if err != nil {
						log.Print(err)
					}
					res = strings.ReplaceAll(res, "jdoodle.cpp", "main.cpp")
					output := fmt.Sprintf("*Output:*\n%s", res)
					l.SendReplyMessage(output)
				}
			case replyPrepareLanguage("Python 2"):
				{
					body := *payloadCodeCompiler("python2", 3, receiveText)
					res, err := onlineCompiler(body)
					if err != nil {
						log.Print(err)
					}
					res = strings.ReplaceAll(res, "jdoodle.py", "main.py")
					output := fmt.Sprintf("*Output:*\n%s", res)
					l.SendReplyMessage(output)
				}
			case replyPrepareLanguage("Python 3"):
				{
					body := *payloadCodeCompiler("python3", 4, receiveText)
					res, err := onlineCompiler(body)
					if err != nil {
						log.Print(err)
					}
					res = strings.ReplaceAll(res, "jdoodle.py", "main.py")
					output := fmt.Sprintf("*Output:*\n%s", res)
					l.SendReplyMessage(output)
				}
			case replyPrepareLanguage("Nodejs"):
				{
					body := *payloadCodeCompiler("nodejs", 4, receiveText)
					res, err := onlineCompiler(body)
					if err != nil {
						log.Print(err)
					}
					res = strings.ReplaceAll(res, "jdoodle.js", "main.js")
					output := fmt.Sprintf("*Output:*\n%s", res)
					l.SendReplyMessage(output)
				}
			}
		}
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

func onlineCompiler(body map[string]interface{}) (string, error) {
	const url = "https://api.jdoodle.com/v1/execute"

	jsonValue, _ := json.Marshal(body)

	r, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}

	defer r.Body.Close()

	onlineCompiler := &models.OnlineCompiler{}
	err = json.NewDecoder(r.Body).Decode(onlineCompiler)

	if err != nil {
		return "", err
	}

	if r.StatusCode == http.StatusOK {
		return onlineCompiler.Output, nil
	}

	return "", errors.New("something went wrong")
}
