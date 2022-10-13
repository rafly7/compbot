package message

import (
	"compbot/configs"
	"compbot/lib"
	"compbot/services"
	"compbot/utils"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/go-redis/redis/v8"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types/events"
)

func onlineCompilerConversation(msg *events.Message, rdb *redis.Client, infoSender string, l *lib.LibImpl) {
	defer func() {
		if r := recover(); r != nil {
			log.Print(r)
		}
	}()
	extendedMsg := msg.Message.GetExtendedTextMessage()
	if extendedMsg != nil {
		if msg.Message.ExtendedTextMessage.ContextInfo != nil && msg.Message.ExtendedTextMessage.Text != nil {
			// log.Print(msg.Message.ExtendedTextMessage)
			val, err := rdb.Get(context.Background(), infoSender).Result()
			if err != redis.Nil {
				if err != nil {
					utils.Recover(err)
					return
				}
				if val == *msg.Message.ExtendedTextMessage.ContextInfo.StanzaId {
					replyBot := *msg.Message.ExtendedTextMessage.ContextInfo.QuotedMessage.Conversation
					receiveText := *msg.Message.ExtendedTextMessage.Text
					switch replyBot {
					case replyPrepareLanguage(utils.OCRunC):
						{
							body := *payloadCodeCompiler("c", 5, receiveText)
							res, err := services.RunCode(body)
							if err != nil {
								utils.Recover(err)
								return
							}
							res = strings.ReplaceAll(res, "jdoodle.c", "main.c")
							output := fmt.Sprintf("*Output:*\n%s", res)
							deleteInfoUserOnlineCompiler(rdb, infoSender, func() (whatsmeow.SendResponse, error) {
								return l.SendReplyMessage(output)
							})
						}
					case replyPrepareLanguage(utils.OCRunCPP):
						{
							body := *payloadCodeCompiler("cpp17", 1, receiveText)
							res, err := services.RunCode(body)
							if err != nil {
								utils.Recover(err)
								return
							}
							res = strings.ReplaceAll(res, "jdoodle.cpp", "main.cpp")
							output := fmt.Sprintf("*Output:*\n%s", res)
							deleteInfoUserOnlineCompiler(rdb, infoSender, func() (whatsmeow.SendResponse, error) {
								return l.SendReplyMessage(output)
							})
						}
					case replyPrepareLanguage(utils.OCRunCSharp):
						{
							body := *payloadCodeCompiler("csharp", 4, receiveText)
							res, err := services.RunCode(body)
							if err != nil {
								utils.Recover(err)
								return
							}
							res = strings.ReplaceAll(res, "jdoodle.cs", "main.cs")
							output := fmt.Sprintf("*Output:*\n%s", res)
							deleteInfoUserOnlineCompiler(rdb, infoSender, func() (whatsmeow.SendResponse, error) {
								return l.SendReplyMessage(output)
							})
						}
					case replyPrepareLanguage(utils.OCRunDart):
						{
							body := *payloadCodeCompiler("dart", 4, receiveText)
							res, err := services.RunCode(body)
							if err != nil {
								utils.Recover(err)
								return
							}
							res = strings.ReplaceAll(res, "jdoodle.dart", "main.dart")
							output := fmt.Sprintf("*Output:*\n%s", res)
							deleteInfoUserOnlineCompiler(rdb, infoSender, func() (whatsmeow.SendResponse, error) {
								return l.SendReplyMessage(output)
							})
						}
					case replyPrepareLanguage(utils.OCRunGo):
						{
							body := *payloadCodeCompiler("go", 4, receiveText)
							res, err := services.RunCode(body)
							if err != nil {
								utils.Recover(err)
								return
							}
							res = strings.ReplaceAll(res, "jdoodle.go", "main.go")
							output := fmt.Sprintf("*Output:*\n%s", res)
							deleteInfoUserOnlineCompiler(rdb, infoSender, func() (whatsmeow.SendResponse, error) {
								return l.SendReplyMessage(output)
							})
						}
					case replyPrepareLanguage(utils.OCRunJava):
						{
							body := *payloadCodeCompiler("java", 4, receiveText)
							res, err := services.RunCode(body)
							if err != nil {
								utils.Recover(err)
								return
							}
							res = strings.ReplaceAll(res, "jdoodle.java", "main.java")
							output := fmt.Sprintf("*Output:*\n%s", res)
							deleteInfoUserOnlineCompiler(rdb, infoSender, func() (whatsmeow.SendResponse, error) {
								return l.SendReplyMessage(output)
							})
						}
					case replyPrepareLanguage(utils.OCRunKotlin):
						{
							body := *payloadCodeCompiler("kotlin", 3, receiveText)
							res, err := services.RunCode(body)
							if err != nil {
								utils.Recover(err)
								return
							}
							res = strings.ReplaceAll(res, "jdoodle.kt", "main.kt")
							output := fmt.Sprintf("*Output:*\n%s", res)
							deleteInfoUserOnlineCompiler(rdb, infoSender, func() (whatsmeow.SendResponse, error) {
								return l.SendReplyMessage(output)
							})
						}
					case replyPrepareLanguage(utils.OCRunPascal):
						{
							body := *payloadCodeCompiler("pascal", 3, receiveText)
							res, err := services.RunCode(body)
							if err != nil {
								utils.Recover(err)
								return
							}
							res = strings.ReplaceAll(res, "jdoodle.pas", "main.pas")
							output := fmt.Sprintf("*Output:*\n%s", res)
							deleteInfoUserOnlineCompiler(rdb, infoSender, func() (whatsmeow.SendResponse, error) {
								return l.SendReplyMessage(output)
							})
						}
					case replyPrepareLanguage(utils.OCRunSwift):
						{
							body := *payloadCodeCompiler("swift", 4, receiveText)
							res, err := services.RunCode(body)
							if err != nil {
								utils.Recover(err)
								return
							}
							res = strings.ReplaceAll(res, "jdoodle.swift", "main.swift")
							output := fmt.Sprintf("*Output:*\n%s", res)
							deleteInfoUserOnlineCompiler(rdb, infoSender, func() (whatsmeow.SendResponse, error) {
								return l.SendReplyMessage(output)
							})
						}
					case replyPrepareLanguage(utils.OCRunPython2):
						{
							body := *payloadCodeCompiler("python2", 3, receiveText)
							res, err := services.RunCode(body)
							if err != nil {
								utils.Recover(err)
								return
							}
							res = strings.ReplaceAll(res, "jdoodle.py", "main.py")
							output := fmt.Sprintf("*Output:*\n%s", res)
							deleteInfoUserOnlineCompiler(rdb, infoSender, func() (whatsmeow.SendResponse, error) {
								return l.SendReplyMessage(output)
							})
						}
					case replyPrepareLanguage(utils.OCRunPython3):
						{
							body := *payloadCodeCompiler("python3", 4, receiveText)
							res, err := services.RunCode(body)
							if err != nil {
								utils.Recover(err)
								return
							}
							res = strings.ReplaceAll(res, "jdoodle.py", "main.py")
							output := fmt.Sprintf("*Output:*\n%s", res)
							deleteInfoUserOnlineCompiler(rdb, infoSender, func() (whatsmeow.SendResponse, error) {
								return l.SendReplyMessage(output)
							})
						}
					case replyPrepareLanguage(utils.OCRunNodejs):
						{
							body := *payloadCodeCompiler("nodejs", 4, receiveText)
							res, err := services.RunCode(body)
							if err != nil {
								utils.Recover(err)
								return
							}
							res = strings.ReplaceAll(res, "jdoodle.js", "main.js")
							output := fmt.Sprintf("*Output:*\n%s", res)
							deleteInfoUserOnlineCompiler(rdb, infoSender, func() (whatsmeow.SendResponse, error) {
								return l.SendReplyMessage(output)
							})
						}
					case replyPrepareLanguage(utils.OCRunPhp):
						{
							body := *payloadCodeCompiler("php", 4, receiveText)
							res, err := services.RunCode(body)
							if err != nil {
								utils.Recover(err)
								return
							}
							res = strings.ReplaceAll(res, "jdoodle.php", "main.php")
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
		utils.Recover(errCB)
		return
	}
	_, errRDB := rdb.Del(context.Background(), infoSender).Result()
	if errRDB != nil {
		utils.Recover(errRDB)
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
