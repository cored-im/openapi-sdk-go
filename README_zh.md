# Cored IM OpenAPI SDK - Golang

[![Go Reference](https://pkg.go.dev/badge/github.com/cored-im/openapi-sdk-go.svg)](https://pkg.go.dev/github.com/cored-im/openapi-sdk-go)
[![Go Version](https://img.shields.io/github/v/tag/cored-im/openapi-sdk-go)](https://github.com/cored-im/openapi-sdk-go/tags)
[![Go](https://github.com/cored-im/openapi-sdk-go/actions/workflows/go.yaml/badge.svg)](https://github.com/cored-im/openapi-sdk-go/actions/workflows/go.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/cored-im/openapi-sdk-go)](https://goreportcard.com/report/github.com/cored-im/openapi-sdk-go)
[![License](https://img.shields.io/github/license/cored-im/openapi-sdk-go)](LICENSE)

[English](README.md) | 中文

Cored 是一个安全、可自托管的团队协作平台，集成了即时通讯、组织架构、音视频会议和文件存储等功能。

本项目是Cored服务端的 Go SDK，用于通过 OpenAPI 与Cored服务端进行交互。使用前需要先自行部署Cored服务端，部署教程请参考[快速部署文档](https://coredim.com/docs/admin/install/quick-install)。

## 安装

```bash
go get github.com/cored-im/openapi-sdk-go
```

## 快速开始

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

    // 可选：预热可提前获取访问凭证和同步服务端时间，减少首次调用的延迟
    _ = client.Preheat(context.Background())

    // 调用 API
    resp, err := client.Im.Message.SendMessage(context.Background(), &coim.SendMessageReq{
        ChatId:      cosdk.String("chat-id"),
        MessageType: cosdk.String(coim.MessageType_TEXT),
        MessageContent: &coim.MessageContent{
            Text: &coim.MessageText{
                Content: cosdk.String("Cored 新版本发布！"),
            },
        },
    })
    fmt.Println(resp, err)
}
```

## 客户端配置

`NewClient()` 支持通过可选参数配置客户端行为：

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
    cosdk.WithLogLevel(cocore.LoggerLevelDebug),            // 日志级别（默认: Info）
    cosdk.WithRequestTimeout(30 * time.Second),             // 请求超时（默认: 60s）
    cosdk.WithEnableEncryption(false),                      // 启用请求加密（默认: true）
)
```

## 事件订阅

通过 WebSocket 接收实时事件推送：

```go
// 注册事件处理函数
client.Im.Message.Event.OnMessageReceive(func(ctx context.Context, event *coim.EventMessageReceive) {
    fmt.Println("收到消息:", event.Body.Message.MessageId)
})
```

## 错误处理

API 调用失败时返回的 `error` 可以断言为 `*cocore.ApiError`，包含 `Code`、`Msg` 和 `LogId` 字段：

```go
import cocore "github.com/cored-im/openapi-sdk-go/core"

resp, err := client.Im.Message.SendMessage(ctx, req)
if err != nil {
    if apiErr, ok := err.(*cocore.ApiError); ok {
        fmt.Printf("API 错误: code=%d, msg=%s, logId=%s\n", apiErr.Code, apiErr.Msg, apiErr.LogId)
    } else {
        fmt.Println("请求错误:", err)
    }
}
```

## 环境要求

- Go 1.12 及以上版本

## 相关链接

- [官网](https://cored.im/)

## 许可证

[Apache-2.0 License](LICENSE)
