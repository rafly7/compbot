package event

import (
	"bytes"
	"context"
	_ "database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

const clientId string = "cec952474cf8906c6ef7be47fa4e536"
const clientSecret string = "1b07c01723734284f875715784aac52525f975c7d1389ce8c1a8e3176e46af8e"

type OnlineCompiler struct {
	Output     string `json:"output"`
	StatusCode int    `json:"statusCode"`
	Memory     string `json:"memory"`
	CpuTime    string `json:"cpuTime"`
}

func (c *client) EventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		conversation := v.Message.GetConversation()
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
		if v.Info.IsGroup && conversation == "!compbot" {
			message := `*Compbot online compiler bot*
_List command yang tersedia:_

!compbot@c				gcc compiler
!compbot@cpp			g++ compiler
			`
			log.Println(v.Message.GetConversation())
			c.WAClient.SendMessage(context.Background(), v.Info.Chat, v.Info.ID, &waProto.Message{
				ExtendedTextMessage: &waProto.ExtendedTextMessage{
					Text: proto.String(message),
					ContextInfo: &waProto.ContextInfo{
						StanzaId:      &v.Info.ID,
						Participant:   proto.String(v.Info.Sender.String()),
						QuotedMessage: v.Message,
					},
				},
			})
		} else if v.Info.IsGroup && conversation == "!compbot@c" {
			c.WAClient.SendMessage(context.Background(), v.Info.Chat, v.Info.ID, &waProto.Message{
				ExtendedTextMessage: &waProto.ExtendedTextMessage{
					Text: proto.String(replyPrepareLanguage("C")),
					ContextInfo: &waProto.ContextInfo{
						StanzaId:      &v.Info.ID,
						Participant:   proto.String(v.Info.Sender.String()),
						QuotedMessage: v.Message,
					},
				},
			})
		} else if v.Info.IsGroup && conversation == "!compbot@cpp" {
			c.WAClient.SendMessage(context.Background(), v.Info.Chat, v.Info.ID, &waProto.Message{
				ExtendedTextMessage: &waProto.ExtendedTextMessage{
					Text: proto.String(replyPrepareLanguage("C++")),
					ContextInfo: &waProto.ContextInfo{
						StanzaId:      &v.Info.ID,
						Participant:   proto.String(v.Info.Sender.String()),
						QuotedMessage: v.Message,
					},
				},
			})
		} else if v.Message.GetExtendedTextMessage() != nil && v.Info.IsGroup {
			replyBot := *v.Message.GetExtendedTextMessage().ContextInfo.QuotedMessage.Conversation
			receiveText := *v.Message.GetExtendedTextMessage().Text
			if replyBot == replyPrepareLanguage("C") {
				body := map[string]interface{}{
					"language":     "c",
					"versionIndex": "5",
					"clientId":     clientId,
					"clientSecret": clientSecret,
					"script":       receiveText,
				}
				res, err := onlineCompiler(body)
				if err != nil {
					log.Print(err)
				}
				res = strings.ReplaceAll(res, "jdoodle.c", "main.c")
				c.WAClient.SendMessage(context.Background(), v.Info.Chat, v.Info.ID, &waProto.Message{
					ExtendedTextMessage: &waProto.ExtendedTextMessage{
						Text: proto.String(fmt.Sprintf("*Output:*\n%s", res)),
						ContextInfo: &waProto.ContextInfo{
							StanzaId:      &v.Info.ID,
							Participant:   proto.String(v.Info.Sender.String()),
							QuotedMessage: v.Message,
						},
					},
				})

			} else if replyBot == replyPrepareLanguage("C++") {
				body := map[string]interface{}{
					"language":     "cpp17",
					"versionIndex": "1",
					"clientId":     clientId,
					"clientSecret": clientSecret,
					"script":       receiveText,
				}
				res, err := onlineCompiler(body)
				if err != nil {
					log.Print(err)
				}
				res = strings.ReplaceAll(res, "jdoodle.cpp", "main.cpp")
				c.WAClient.SendMessage(context.Background(), v.Info.Chat, v.Info.ID, &waProto.Message{
					ExtendedTextMessage: &waProto.ExtendedTextMessage{
						Text: proto.String(fmt.Sprintf("*Output:*\n%s", res)),
						ContextInfo: &waProto.ContextInfo{
							StanzaId:      &v.Info.ID,
							Participant:   proto.String(v.Info.Sender.String()),
							QuotedMessage: v.Message,
						},
					},
				})
			}
		}
	}
}

func replyPrepareLanguage(language string) string {
	return fmt.Sprintf("Ok, beri saya beberapa kode %s untuk dieksekusi", language)
}

func onlineCompiler(body map[string]interface{}) (string, error) {
	const url = "https://api.jdoodle.com/v1/execute"

	jsonValue, _ := json.Marshal(body)

	r, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}

	defer r.Body.Close()

	onlineCompiler := &OnlineCompiler{}
	err = json.NewDecoder(r.Body).Decode(onlineCompiler)

	if err != nil {
		return "", err
	}

	if r.StatusCode == http.StatusOK {
		return onlineCompiler.Output, nil
	}

	return "", errors.New("something went wrong")
}
