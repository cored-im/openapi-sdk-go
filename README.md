# Cored IM OpenAPI SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/cored-im/openapi-sdk-go.svg)](https://pkg.go.dev/github.com/cored-im/openapi-sdk-go)
[![Go](https://github.com/cored-im/openapi-sdk-go/actions/workflows/go.yaml/badge.svg)](https://github.com/cored-im/openapi-sdk-go/actions/workflows/go.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/cored-im/openapi-sdk-go)](https://goreportcard.com/report/github.com/cored-im/openapi-sdk-go)
[![License](https://img.shields.io/github/license/cored-im/openapi-sdk-go)](LICENSE)

English | [中文](README_zh.md)

Cored is a secure, self-hosted productivity platform for teams, integrating instant messaging, organizational structures, video conferencing, and file storage.

This is the official Go SDK for Cored server, used to interact with the Cored server via OpenAPI. You need to deploy the Cored server before using this SDK. See the [Quick Deploy Guide](https://coredim.com/docs/admin/install/quick-install) for setup instructions.

## Requirements

- Go 1.12 or later

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
    client := cosdk.NewClient("http://localhost:11000", "your-app-id", "your-app-secret")
    defer client.Close()

    // Optional: preheat fetches access token and syncs server time upfront,
    // reducing latency on the first API call
    _ = client.Preheat(context.Background())

    // Call API
    _, err := client.Im.Chat.CreateTyping(context.Background(), &coim.CreateTypingReq{})
    fmt.Println(err)
}
```

## Authentication

This SDK uses app-level authentication. Pass your App ID and App Secret when creating the client. The SDK automatically manages access token retrieval and refresh.

## Examples

Run all tests:

```bash
go test ./...
```

Run the IM message example only:

```bash
go test ./example -run TestImMessageSend
```

## Links

- [Website](https://cored.im/)

## License

[Apache-2.0 License](LICENSE)
