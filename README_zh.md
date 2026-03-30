# Cored IM OpenAPI SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/cored-im/openapi-sdk-go.svg)](https://pkg.go.dev/github.com/cored-im/openapi-sdk-go)
[![Go](https://github.com/cored-im/openapi-sdk-go/actions/workflows/go.yaml/badge.svg)](https://github.com/cored-im/openapi-sdk-go/actions/workflows/go.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/cored-im/openapi-sdk-go)](https://goreportcard.com/report/github.com/cored-im/openapi-sdk-go)
[![License](https://img.shields.io/github/license/cored-im/openapi-sdk-go)](LICENSE)

[English](README.md) | 中文

Cored 是一个安全、可自托管的团队协作平台，集成了即时通讯、组织架构、音视频会议和文件存储等功能。

本项目是Cored服务端的 Go SDK，用于通过 OpenAPI 与Cored服务端进行交互。使用前需要先自行部署Cored服务端，部署教程请参考[快速部署文档](https://coredim.com/docs/admin/install/quick-install)。

## 环境要求

- Go 1.12 及以上版本

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
    client := cosdk.NewClient("http://localhost:11000", "your-app-id", "your-app-secret")
    defer client.Close()

    // 可选：预热可提前获取访问凭证和同步服务端时间，减少首次调用的延迟
    _ = client.Preheat(context.Background())

    // 调用 API
    _, err := client.Im.Chat.CreateTyping(context.Background(), &coim.CreateTypingReq{})
    fmt.Println(err)
}
```

## 认证方式

本 SDK 使用应用级别认证。创建客户端时传入 App ID 和 App Secret，SDK 会自动管理访问凭证的获取与刷新。

## 运行示例

运行全部测试：

```bash
go test ./...
```

仅运行 IM 消息示例：

```bash
go test ./example -run TestImMessageSend
```

## 相关链接

- [官网](https://cored.im/)

## 许可证

[Apache-2.0 License](LICENSE)
