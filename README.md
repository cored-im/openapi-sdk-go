# Cored IM OpenAPI SDK - Golang

[![Go Reference](https://pkg.go.dev/badge/github.com/cored-im/openapi-sdk-go.svg)](https://pkg.go.dev/github.com/cored-im/openapi-sdk-go)
[![Go Version](https://img.shields.io/github/v/tag/cored-im/openapi-sdk-go)](https://github.com/cored-im/openapi-sdk-go/tags)
[![Go](https://github.com/cored-im/openapi-sdk-go/actions/workflows/go.yaml/badge.svg)](https://github.com/cored-im/openapi-sdk-go/actions/workflows/go.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/cored-im/openapi-sdk-go)](https://goreportcard.com/report/github.com/cored-im/openapi-sdk-go)
[![License](https://img.shields.io/github/license/cored-im/openapi-sdk-go)](LICENSE)

English | [中文](README_zh.md)

Cored is a secure, self-hosted productivity platform for teams, integrating instant messaging, organizational structures, video conferencing, and file storage.

This is the official Go SDK for Cored server, used to interact with the Cored server via OpenAPI. You need to deploy the Cored server before using this SDK. See the [Quick Deploy Guide](https://coredim.com/docs/admin/install/quick-install) for setup instructions.

## Installation

```bash
go get github.com/cored-im/openapi-sdk-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"

    cosdk "github.com/cored-im/openapi-sdk-go"
    coim "github.com/cored-im/openapi-sdk-go/service/im/v1"
)

func main() {
    client := cosdk.NewClient("https://your-backend-url.com", "your-app-id", "your-app-secret")
    defer client.Close()

    // Optional: preheat fetches access token and syncs server time upfront,
    // reducing latency on the first API call
    _ = client.Preheat(context.Background())

    // Call API
    resp, err := client.Im.Message.SendMessage(context.Background(), &coim.SendMessageReq{
        ChatId:      cosdk.String("chat-id"),
        MessageType: cosdk.String(coim.MessageType_TEXT),
        MessageContent: &coim.MessageContent{
            Text: &coim.MessageText{
                Content: cosdk.String("Cored new version released!"),
            },
        },
    })
    fmt.Println(resp, err)
}
```

## Configuration

`NewClient()` accepts optional functional options to configure client behavior:

```go
import (
    "time"

    cosdk "github.com/cored-im/openapi-sdk-go"
    cocore "github.com/cored-im/openapi-sdk-go/core"
)

client := cosdk.NewClient(
    "https://your-backend-url.com",
    "your-app-id",
    "your-app-secret",
    cosdk.WithLogLevel(cocore.LoggerLevelDebug),            // Log level (default: Info)
    cosdk.WithRequestTimeout(30 * time.Second),             // Request timeout (default: 60s)
    cosdk.WithEnableEncryption(false),                      // Enable request encryption (default: true)
)
```

## Event Subscription

Receive real-time events via WebSocket:

```go
// Register event handler
client.Im.Message.Event.OnMessageReceive(func(ctx context.Context, event *coim.EventMessageReceive) {
    fmt.Println("Message received:", event.Body.Message.MessageId)
})
```

## Error Handling

When an API call fails, the returned `error` can be type-asserted to `*cocore.ApiError`, which contains `Code`, `Msg`, and `LogId` fields:

```go
import cocore "github.com/cored-im/openapi-sdk-go/core"

resp, err := client.Im.Message.SendMessage(ctx, req)
if err != nil {
    if apiErr, ok := err.(*cocore.ApiError); ok {
        fmt.Printf("API error: code=%d, msg=%s, logId=%s\n", apiErr.Code, apiErr.Msg, apiErr.LogId)
    } else {
        fmt.Println("Request error:", err)
    }
}
```

## Requirements

- Go 1.12 or later

## Links

- [Website](https://cored.im/)

## License

[Apache-2.0 License](LICENSE)
