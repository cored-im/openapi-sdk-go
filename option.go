package cosdk

import (
	"time"

	cocore "github.com/cored-im/openapi-sdk-go/core"
)

type clientOption struct {
	httpClient       cocore.HttpClient
	requestTimeout   time.Duration
	enableEncryption bool

	logLevel    cocore.LoggerLevel
	logger      cocore.Logger
	timeManager cocore.TimeManager

	jsonMarshaller   cocore.Marshaller
	jsonUnmarshaller cocore.Unmarshaller
}

type clientOptionFunc func(option *clientOption)

func WithHttpClient(httpClient cocore.HttpClient) clientOptionFunc {
	return func(option *clientOption) {
		option.httpClient = httpClient
	}
}

func WithRequestTimeout(requestTimeout time.Duration) clientOptionFunc {
	return func(option *clientOption) {
		if requestTimeout > 0 {
			option.requestTimeout = requestTimeout
		}
	}
}

func WithEnableEncryption(enableEncryption bool) clientOptionFunc {
	return func(option *clientOption) {
		option.enableEncryption = enableEncryption
	}
}

func WithLogLevel(logLevel cocore.LoggerLevel) clientOptionFunc {
	return func(option *clientOption) {
		option.logLevel = logLevel
	}
}

func WithLogger(logger cocore.Logger) clientOptionFunc {
	return func(option *clientOption) {
		option.logger = logger
	}
}

func WithTimeManager(timeManager cocore.TimeManager) clientOptionFunc {
	return func(option *clientOption) {
		option.timeManager = timeManager
	}
}

func WithJsonMarshaller(jsonMarshaller cocore.Marshaller) clientOptionFunc {
	return func(option *clientOption) {
		option.jsonMarshaller = jsonMarshaller
	}
}

func WithJsonUnmarshaller(jsonUnmarshaller cocore.Unmarshaller) clientOptionFunc {
	return func(option *clientOption) {
		option.jsonUnmarshaller = jsonUnmarshaller
	}
}
