// Copyright (c) 2026 Cored Limited
// SPDX-License-Identifier: Apache-2.0

package coim

import (
	"context"

	cocore "github.com/cored-im/openapi-sdk-go/core"
)

// Send a message (Request)
type SendMessageReq struct {
	MessageType    *string         `json:"message_type,omitempty"`     // Message type
	MessageContent *MessageContent `json:"message_content,omitempty"`  // Message content
	ChatId         *string         `json:"chat_id,omitempty"`          // Chat ID
	ReplyMessageId *string         `json:"reply_message_id,omitempty"` // ID of the message being replied to
}

// Send a message (Response)
type SendMessageResp struct {
	MessageId *string `json:"message_id,omitempty"` // Message ID
}

// Send a message
func (impl *v1Message) SendMessage(ctx context.Context, req *SendMessageReq) (*SendMessageResp, error) {
	apiResp, err := impl.config.ApiClient.Request(ctx, &cocore.ApiRequest{
		Method:             "POST",
		Path:               "/oapi/im/v1/messages",
		Body:               req,
		WithAppAccessToken: true,
		WithWebSocket:      true,
	})
	if err != nil {
		return nil, err
	}
	var resp SendMessageResp
	if err = apiResp.JSON(&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Get a message (Request)
type GetMessageReq struct {
	MessageId *string `json:"message_id,omitempty"` // Message ID
}

// Get a message (Response)
type GetMessageResp struct {
	Message *Message `json:"message,omitempty"` // Message
}

// Get a message
func (impl *v1Message) GetMessage(ctx context.Context, req *GetMessageReq) (*GetMessageResp, error) {
	apiResp, err := impl.config.ApiClient.Request(ctx, &cocore.ApiRequest{
		Method: "GET",
		Path:   "/oapi/im/v1/messages/:message_id",
		Body:   req,
		PathParams: map[string]string{
			"message_id": stringOrEmpty(req.MessageId),
		},
		WithAppAccessToken: true,
	})
	if err != nil {
		return nil, err
	}
	var resp GetMessageResp
	if err = apiResp.JSON(&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Recall a message (Request)
type RecallMessageReq struct {
	MessageId *string `json:"message_id,omitempty"` // Message ID
}

// Recall a message (Response)
type RecallMessageResp struct {
}

// Recall a message
func (impl *v1Message) RecallMessage(ctx context.Context, req *RecallMessageReq) (*RecallMessageResp, error) {
	apiResp, err := impl.config.ApiClient.Request(ctx, &cocore.ApiRequest{
		Method: "POST",
		Path:   "/oapi/im/v1/messages/:message_id/recall",
		Body:   req,
		PathParams: map[string]string{
			"message_id": stringOrEmpty(req.MessageId),
		},
		WithAppAccessToken: true,
	})
	if err != nil {
		return nil, err
	}
	var resp RecallMessageResp
	if err = apiResp.JSON(&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Mark message as read (Request)
type ReadMessageReq struct {
	MessageId *string `json:"message_id,omitempty"` // Message ID
}

// Mark message as read (Response)
type ReadMessageResp struct {
}

// Mark message as read
func (impl *v1Message) ReadMessage(ctx context.Context, req *ReadMessageReq) (*ReadMessageResp, error) {
	apiResp, err := impl.config.ApiClient.Request(ctx, &cocore.ApiRequest{
		Method: "POST",
		Path:   "/oapi/im/v1/messages/:message_id/read",
		Body:   req,
		PathParams: map[string]string{
			"message_id": stringOrEmpty(req.MessageId),
		},
		WithAppAccessToken: true,
	})
	if err != nil {
		return nil, err
	}
	var resp ReadMessageResp
	if err = apiResp.JSON(&resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
