// Copyright (c) 2026 Cored Limited
// SPDX-License-Identifier: Apache-2.0

package cocore

import "time"

type Config struct {
	AppId      string
	AppSecret  string
	BackendUrl string

	HttpClient       HttpClient
	ApiClient        ApiClient
	EnableEncryption bool
	RequestTimeout   time.Duration
	TimeManager      TimeManager

	Logger Logger

	JsonMarshal   Marshaller
	JsonUnmarshal Unmarshaller
}
