package cosdk

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	cocore "github.com/cored-im/openapi-sdk-go/core"
	coserviceim "github.com/cored-im/openapi-sdk-go/service/im"
)

type Client struct {
	config    *cocore.Config
	ApiClient cocore.ApiClient
	Im        *coserviceim.Service
}

// Creates a client
func NewClient(backendUrl string, appId string, appSecret string, options ...clientOptionFunc) *Client {
	// init option
	option := &clientOption{
		logLevel:         cocore.LoggerLevelInfo,
		requestTimeout:   1 * time.Minute,
		enableEncryption: true,
		jsonMarshaller:   json.Marshal,
		jsonUnmarshaller: json.Unmarshal,
	}
	for _, fn := range options {
		fn(option)
	}

	// init config
	config := &cocore.Config{
		AppId:            appId,
		AppSecret:        appSecret,
		BackendUrl:       strings.TrimSpace(strings.TrimSuffix(backendUrl, "/")),
		HttpClient:       option.httpClient,
		EnableEncryption: option.enableEncryption,
		RequestTimeout:   option.requestTimeout,
		TimeManager:      option.timeManager,
		Logger:           option.logger,
		JsonMarshal:      option.jsonMarshaller,
		JsonUnmarshal:    option.jsonUnmarshaller,
	}
	if config.TimeManager == nil {
		config.TimeManager = cocore.NewDefaultTimeManager()
	}
	if config.Logger == nil {
		config.Logger = cocore.NewDefaultLogger(option.logLevel)
	}
	if config.HttpClient == nil {
		config.HttpClient = cocore.NewDefaultHttpClient(option.requestTimeout)
	}
	config.ApiClient = cocore.NewDefaultApiClient(config)

	// init client
	client := &Client{
		config:    config,
		ApiClient: config.ApiClient,
		Im:        coserviceim.New(config),
	}

	return client
}

// Preheating to prevent delay in the first request
func (c *Client) Preheat(ctx context.Context) error {
	return c.ApiClient.Preheat(ctx)
}

// Close client
func (c *Client) Close() error {
	return c.ApiClient.Close()
}
