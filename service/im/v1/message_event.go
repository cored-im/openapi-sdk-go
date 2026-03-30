// Copyright (c) 2026 Cored Limited
// SPDX-License-Identifier: Apache-2.0

package coim

import (
	"context"
	"reflect"

	cocore "github.com/cored-im/openapi-sdk-go/core"
)

// Message received
type EventMessageReceive struct {
	Header *cocore.EventHeader      `json:"header,omitempty"`
	Body   *EventMessageReceiveBody `json:"body,omitempty"`
}

// Message received
type EventMessageReceiveBody struct {
	Message *Message `json:"message,omitempty"` // Message
}

// Message received
func (impl *v1MessageEvent) OnMessageReceive(handler func(ctx context.Context, event *EventMessageReceive)) {
	var eventHandler cocore.EventHandler = func(ctx context.Context, header *cocore.EventHeader, body []byte) error {
		event := &EventMessageReceive{Header: header, Body: &EventMessageReceiveBody{}}
		if err := impl.config.JsonUnmarshal(body, event.Body); err != nil {
			return err
		}
		handler(ctx, event)
		return nil
	}
	impl.handlerMap.Store(reflect.ValueOf(handler).Pointer(), eventHandler)
	impl.config.ApiClient.OnEvent("im.v1.message.receive", eventHandler)
}

func (impl *v1MessageEvent) OffMessageReceive(handler func(ctx context.Context, event *EventMessageReceive)) {
	key := reflect.ValueOf(handler).Pointer()
	eventHandler, ok := impl.handlerMap.Load(key)
	if ok {
		impl.config.ApiClient.OffEvent("im.v1.message.receive", eventHandler.(cocore.EventHandler))
		impl.handlerMap.Delete(key)
	}
}
