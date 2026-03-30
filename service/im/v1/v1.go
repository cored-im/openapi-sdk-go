// Copyright (c) 2026 Cored Limited
// SPDX-License-Identifier: Apache-2.0

package coim

import (
	cocore "github.com/cored-im/openapi-sdk-go/core"
)

type V1 struct {
	Chat    *v1Chat
	Message *v1Message
}

func New(config *cocore.Config) *V1 {
	return &V1{
		Chat:    &v1Chat{config: config},
		Message: &v1Message{config: config, Event: &v1MessageEvent{config: config}},
	}
}

func stringOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
