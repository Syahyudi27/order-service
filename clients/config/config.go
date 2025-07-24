package config

import "github.com/parnurzeal/gorequest"

type ClientConfig struct {
	client *gorequest.SuperAgent
	baseURL string
	singnatureKey string
}

type IClientConfig interface {
	GetClient() *gorequest.SuperAgent
	GetBaseURL() string
	GetSignatureKey() string
}

type Option func(*ClientConfig)

func NewClientConfig(options ...Option) *ClientConfig {
	clientConfig := &ClientConfig{
		client: gorequest.New().
		Set("Content-Type", "application/json").
		Set("Accept", "application/json"),
	}

	for _, option := range options {
		option(clientConfig)
	}

	return clientConfig
}

func (c *ClientConfig) GetClient() *gorequest.SuperAgent {
	return c.client
}

func(c *ClientConfig) GetBaseURL() string {
	return c.baseURL
}

func (c *ClientConfig) GetSignatureKey() string {
	return c.singnatureKey
}

func WithBaseURL(baseURL string) Option {
	return func(c *ClientConfig) {
		c.baseURL = baseURL
	}
}

func WithSignatureKey(signatureKey string) Option {
	return func(c *ClientConfig) {
		c.singnatureKey = signatureKey
	}
}