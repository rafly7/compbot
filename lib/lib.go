package lib

import (
	"context"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func (simp *LibImpl) SendInfoBotMessage(content string, footer string) {
	simp.WAClient.SendMessage(context.Background(), simp.Message.Info.Chat, simp.Message.Info.ID, &waProto.Message{
		TemplateMessage: &waProto.TemplateMessage{
			HydratedTemplate: &waProto.TemplateMessage_HydratedFourRowTemplate{
				HydratedContentText: proto.String(content),
				HydratedFooterText:  proto.String(footer),
			},
		},
	})
}

func (simp *LibImpl) SendReplyMessage(text string) (whatsmeow.SendResponse, error) {
	return simp.WAClient.SendMessage(context.Background(), simp.Message.Info.Chat, simp.Message.Info.ID, &waProto.Message{
		ExtendedTextMessage: &waProto.ExtendedTextMessage{
			Text: proto.String(text),
			ContextInfo: &waProto.ContextInfo{
				StanzaId:      &simp.Message.Info.ID,
				Participant:   proto.String(simp.Message.Info.Sender.String()),
				QuotedMessage: simp.Message.Message,
			},
		},
	})
}
