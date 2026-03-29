package example

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	cosdk "github.com/cored-im/openapi-sdk-go"
	coim "github.com/cored-im/openapi-sdk-go/service/im/v1"
)

func TestImMessageSend(t *testing.T) {
	client := cosdk.NewClient(
		"http://localhost:11000",
		"c-TestAppId2",
		"TestAppSecret2",
	)
	client.Im.Message.Event.OnMessageReceive(func(ctx context.Context, event *coim.EventMessageReceive) {
		_, _ = client.Im.Chat.CreateTyping(ctx, &coim.CreateTypingReq{
			ChatId: event.Body.Message.ChatId,
		})
		defer func() {
			_, _ = client.Im.Chat.DeleteTyping(ctx, &coim.DeleteTypingReq{
				ChatId: event.Body.Message.ChatId,
			})
		}()
		_, _ = json.MarshalIndent(event, "", "  ")
		_, _ = client.Im.Message.ReadMessage(ctx, &coim.ReadMessageReq{
			MessageId: event.Body.Message.MessageId,
		})
		_, err := client.Im.Message.SendMessage(ctx, &coim.SendMessageReq{
			MessageType: cosdk.String(coim.MessageType_CARD),
			MessageContent: &coim.MessageContent{
				Card: &coim.MessageCard{
					Schema: cosdk.String("1.0"),
					V1: &coim.MessageCardV1{
						Header: &coim.MessageCardV1Header{
							Title: cosdk.String("Cored new version released!"),
							TitleI18n: map[string]string{
								"en": "Cored new version released!",
							},
							Template: cosdk.String("green"),
						},
						Body: &coim.MessageCardV1Body{
							MessageText: &coim.MessageText{
								Content: cosdk.String("New version features:\n- Added a Night Mode theme\n- Added multilingual support\n- Fixed the iOS video playback crash issue"),
							},
							MessageTextI18n: map[string]*coim.MessageText{
								"en": {
									Content: cosdk.String("New version features:\n- Added a Night Mode theme\n- Added multilingual support\n- Fixed the iOS video playback crash issue"),
								},
							},
						},
						Footer: &coim.MessageCardV1Footer{
							ButtonList: []*coim.MessageCardV1Button{{
								ButtonText: cosdk.String("Open website"),
								ButtonTextI18n: map[string]string{
									"en": "Jump to official website",
								},
								Link: &coim.MessageCardV1ButtonLink{
									Url: cosdk.String("https://cored.im/"),
								},
								Template: cosdk.String("default"),
							}, {
								ButtonText: cosdk.String("Open website"),
								ButtonTextI18n: map[string]string{
									"en": "Jump to official website",
								},
								Link: &coim.MessageCardV1ButtonLink{
									Url: cosdk.String("https://cored.im/"),
								},
								Template: cosdk.String("primary_filled"),
							}, {
								ButtonText: cosdk.String("Open website"),
								ButtonTextI18n: map[string]string{
									"en": "Jump to official website",
								},
								Link: &coim.MessageCardV1ButtonLink{
									Url: cosdk.String("https://cored.im/"),
								},
								Template: cosdk.String("primary"),
							}, {
								ButtonText: cosdk.String("Open website"),
								ButtonTextI18n: map[string]string{
									"en": "Jump to official website",
								},
								Link: &coim.MessageCardV1ButtonLink{
									Url: cosdk.String("https://cored.im/"),
								},
								Template: cosdk.String("danger"),
							}, {
								ButtonText: cosdk.String("Open website"),
								ButtonTextI18n: map[string]string{
									"en": "Jump to official website",
								},
								Link: &coim.MessageCardV1ButtonLink{
									Url: cosdk.String("https://cored.im/"),
								},
								Template: cosdk.String("danger_filled"),
							}, {
								ButtonText: cosdk.String("Open website"),
								ButtonTextI18n: map[string]string{
									"en": "Jump to official website",
								},
								Link: &coim.MessageCardV1ButtonLink{
									Url: cosdk.String("https://cored.im/"),
								},
								Template: cosdk.String("danger_text"),
							}, {
								ButtonText: cosdk.String("Open website"),
								ButtonTextI18n: map[string]string{
									"en": "Jump to official website",
								},
								Link: &coim.MessageCardV1ButtonLink{
									Url: cosdk.String("https://cored.im/"),
								},
								Template: cosdk.String("primary_text"),
							}},
							ButtonAlign: cosdk.String("start"),
						},
					},
				},
			},
			ChatId: event.Body.Message.ChatId,
			// ReplyMessageId: event.Body.Message.MessageId,
		})
		if err != nil {
			return
		}
		time.Sleep(2 * time.Second)
		// _, _ = client.Im.Message.SendMessage(ctx, &coim.SendMessageReq{
		// 	MessageType:    event.Body.Message.MessageType,
		// 	MessageContent: event.Body.Message.MessageContent,
		// 	ChatId:         event.Body.Message.ChatId,
		// 	ReplyMessageId: event.Body.Message.MessageId,
		// })
		// _, _ = client.Im.Message.RecallMessage(ctx, &coim.RecallMessageReq{
		// 	MessageId: resp.MessageId,
		// })
	})
	time.Sleep(10 * time.Second)
}
