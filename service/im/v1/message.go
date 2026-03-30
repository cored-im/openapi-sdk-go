// Copyright (c) 2026 Cored Limited
// SPDX-License-Identifier: Apache-2.0

package coim

import (
	"sync"

	cocore "github.com/cored-im/openapi-sdk-go/core"
)

type v1Message struct {
	config *cocore.Config
	Event  *v1MessageEvent
}

type v1MessageEvent struct {
	config     *cocore.Config
	handlerMap sync.Map
}
