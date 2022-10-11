package message

import (
	"compbot/lib"
	"compbot/services"
	"fmt"

	"github.com/go-redis/redis/v8"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

const infoBot string = `*Compbot online compiler and interpreted bot*
_List command yang tersedia:_

!compbot@c					gcc compiler
!compbot@cpp				g++ compiler
!compbot@c#					.NET compiler
!compbot@dart				Dart compiler
!compbot@go					Golang compiler
!compbot@java				Java 11 compiler
!compbot@kotlin				Kotlin compiler
!compbot@pascal				Pascal compiler
!compbot@swift				Swift compiler
!compbot@python2			python2 interpreted
!compbpt@python3			python3 interpreted
!compbot@nodejs				node interpreted
!compbot@php				php interpreted
			`

func replyPrepareLanguage(language string) string {
	return fmt.Sprintf("Ok, beri saya beberapa kode %s untuk dieksekusi", language)
}

func Message(client *whatsmeow.Client, msg *events.Message, rdb *redis.Client) {
	l := lib.LiblImpl(client, msg)

	if msg.Info.IsGroup {
		// var messageExtended = make(chan proto.Message)
		comp := fmt.Sprintf("%s@compbot", msg.Info.Sender.User)
		if msg.Message.GetConversation() == "!compbot" {
			l.SendInfoBotMessage(infoBot, "Bot: compbot TPLE 09")
		} else if msg.Message.GetConversation() == "!compbot@c" {
			services.OnlineCompilerValidationService(rdb, comp, func() (whatsmeow.SendResponse, error) {
				return l.SendReplyMessage(replyPrepareLanguage("C"))
			})
		} else if msg.Message.GetConversation() == "!compbot@cpp" {
			services.OnlineCompilerValidationService(rdb, comp, func() (whatsmeow.SendResponse, error) {
				return l.SendReplyMessage(replyPrepareLanguage("C++"))
			})
		} else if msg.Message.GetConversation() == "!compbot@python2" {
			services.OnlineCompilerValidationService(rdb, comp, func() (whatsmeow.SendResponse, error) {
				return l.SendReplyMessage(replyPrepareLanguage("Python 2"))
			})
		} else if msg.Message.GetConversation() == "!compbot@python3" {
			services.OnlineCompilerValidationService(rdb, comp, func() (whatsmeow.SendResponse, error) {
				return l.SendReplyMessage(replyPrepareLanguage("Python 3"))
			})
		} else if msg.Message.GetConversation() == "!compbot@nodejs" {
			services.OnlineCompilerValidationService(rdb, comp, func() (whatsmeow.SendResponse, error) {
				return l.SendReplyMessage(replyPrepareLanguage("Nodejs"))
			})
		} else if msg.Message.ExtendedTextMessage != nil {
			// log.Println(msg.Info.Type)
			onlineCompilerConversation(msg, rdb, comp, l)
		}
	}
}
