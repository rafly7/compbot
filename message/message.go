package message

import (
	"compbot/lib"
	"compbot/services"
	"compbot/utils"
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

_Bot: compbot TPLE 09_`

func replyPrepareLanguage(language string) string {
	return fmt.Sprintf("Ok, beri saya beberapa kode %s untuk dieksekusi", language)
}

func Message(client *whatsmeow.Client, msg *events.Message, rdb *redis.Client) {
	l := lib.LiblImpl(client, msg)

	if msg.Info.IsGroup {
		comp := fmt.Sprintf("%s@compbot", msg.Info.Sender.User)
		msgConversation := msg.Message.GetConversation()
		if msg.Message.ExtendedTextMessage != nil {
			onlineCompilerConversation(msg, rdb, comp, l)
			return
		}
		switch msgConversation {
		case utils.OCBanner:
			l.SendInfoBotMessage(infoBot, "Bot: compbot TPLE 09")
		case utils.OCLangC:
			services.OnlineCompilerValidationService(rdb, comp, func() (whatsmeow.SendResponse, error) {
				return l.SendReplyMessage(replyPrepareLanguage(utils.OCRunC))
			})
		case utils.OCLangCPP:
			services.OnlineCompilerValidationService(rdb, comp, func() (whatsmeow.SendResponse, error) {
				return l.SendReplyMessage(replyPrepareLanguage(utils.OCRunCPP))
			})
		case utils.OCLangCSharp:
			services.OnlineCompilerValidationService(rdb, comp, func() (whatsmeow.SendResponse, error) {
				return l.SendReplyMessage(replyPrepareLanguage(utils.OCRunCSharp))
			})
		case utils.OCLangDart:
			services.OnlineCompilerValidationService(rdb, comp, func() (whatsmeow.SendResponse, error) {
				return l.SendReplyMessage(replyPrepareLanguage(utils.OCRunDart))
			})
		case utils.OCLangGo:
			services.OnlineCompilerValidationService(rdb, comp, func() (whatsmeow.SendResponse, error) {
				return l.SendReplyMessage(replyPrepareLanguage(utils.OCRunGo))
			})
		case utils.OCLangJava:
			services.OnlineCompilerValidationService(rdb, comp, func() (whatsmeow.SendResponse, error) {
				return l.SendReplyMessage(replyPrepareLanguage(utils.OCRunJava))
			})
		case utils.OCLangKotlin:
			services.OnlineCompilerValidationService(rdb, comp, func() (whatsmeow.SendResponse, error) {
				return l.SendReplyMessage(replyPrepareLanguage(utils.OCRunKotlin))
			})
		case utils.OCLangPascal:
			services.OnlineCompilerValidationService(rdb, comp, func() (whatsmeow.SendResponse, error) {
				return l.SendReplyMessage(replyPrepareLanguage(utils.OCRunPascal))
			})
		case utils.OCLangSwift:
			services.OnlineCompilerValidationService(rdb, comp, func() (whatsmeow.SendResponse, error) {
				return l.SendReplyMessage(replyPrepareLanguage(utils.OCRunSwift))
			})
		case utils.OCLangPython2:
			services.OnlineCompilerValidationService(rdb, comp, func() (whatsmeow.SendResponse, error) {
				return l.SendReplyMessage(replyPrepareLanguage(utils.OCRunPython2))
			})
		case utils.OCLangPython3:
			services.OnlineCompilerValidationService(rdb, comp, func() (whatsmeow.SendResponse, error) {
				return l.SendReplyMessage(replyPrepareLanguage(utils.OCRunPython3))
			})
		case utils.OCLangNodejs:
			services.OnlineCompilerValidationService(rdb, comp, func() (whatsmeow.SendResponse, error) {
				return l.SendReplyMessage(replyPrepareLanguage(utils.OCRunNodejs))
			})
		case utils.OCLangPhp:
			services.OnlineCompilerValidationService(rdb, comp, func() (whatsmeow.SendResponse, error) {
				return l.SendReplyMessage(replyPrepareLanguage(utils.OCRunPhp))
			})
		}
	}
}
