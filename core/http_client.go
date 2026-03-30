// Copyright (c) 2026 Cored Limited
// SPDX-License-Identifier: Apache-2.0

package cocore

import (
	"net/http"
	"time"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewDefaultHttpClient(requestTimeout time.Duration) HttpClient {
	if requestTimeout == 0 {
		return http.DefaultClient
	} else {
		return &http.Client{Timeout: requestTimeout}
	}
}
