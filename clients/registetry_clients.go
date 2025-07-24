package clients

import (
	"order-service/clients/config"
	clients "order-service/clients/user"
	config2 "order-service/config"
)

type ClientsRegistry struct {}

type IClientRegistry interface {
	GetUSer() clients.IUserClient
}

func NewClientRegistry() IClientRegistry {
	return &ClientsRegistry{}
}

func (c *ClientsRegistry) GetUSer() clients.IUserClient {
	return clients.NewUserClients(
		config.NewClientConfig(
			config.WithBaseURL(config2.Config.InternalService.User.Host), 
			config.WithSignatureKey(config2.Config.InternalService.User.SignatureKey),
		),
	)
}