package clients

import (
	"order-service/clients/config"
	user "order-service/clients/user"
	field "order-service/clients/field"
	payment "order-service/clients/payment"
	config2 "order-service/config"
)

type ClientsRegistry struct {}

type IClientRegistry interface {
	GetUSer() user.IUserClient
	GetField() field.IFieldClient
	GetPayment() payment.IPaymentClient
}

func NewClientRegistry() IClientRegistry {
	return &ClientsRegistry{}
}

func (c *ClientsRegistry) GetUSer() user.IUserClient {
	return user.NewUserClients(
		config.NewClientConfig(
			config.WithBaseURL(config2.Config.InternalService.User.Host), 
			config.WithSignatureKey(config2.Config.InternalService.User.SignatureKey),
		),
	)
}

func (c *ClientsRegistry) GetField() field.IFieldClient {
	return field.NewFieldClient(
		config.NewClientConfig(
			config.WithBaseURL(config2.Config.InternalService.Field.Host), 
			config.WithSignatureKey(config2.Config.InternalService.Field.SignatureKey),
		),
	)
}

func (c *ClientsRegistry) GetPayment() payment.IPaymentClient {
	return payment.NewPaymentClient(
		config.NewClientConfig(
			config.WithBaseURL(config2.Config.InternalService.Payment.Host), 
			config.WithSignatureKey(config2.Config.InternalService.Payment.SignatureKey),
		),
	)
}