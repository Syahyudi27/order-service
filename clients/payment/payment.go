package clients

import (
	"context"
	"encoding/json"
	"order-service/clients/config"
	"order-service/common/util"
	config2 "order-service/config"
	"order-service/constants"
	"order-service/domain/dto"

	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type PaymentClient struct {
	client config.IClientConfig
}

type IPaymentClient interface {
	GetPaymentByUUID(context.Context, uuid.UUID) (*PaymentData, error)
	CreatePaymentLink(context.Context, *dto.PaymentRequest) (*PaymentData, error)
}

func NewPaymentClient(client config.IClientConfig) *PaymentClient {
	return &PaymentClient{
		client: client,
	}
}

func (p *PaymentClient) GetPaymentByUUID(ctx context.Context, uuid uuid.UUID) (*PaymentData, error) {
	unixTime := time.Now().Unix()
	generateAPIKey := fmt.Sprintf("%s:%s:%d",
		config2.Config.AppName,
		p.client.GetSignatureKey(),
		unixTime,
	)
	apiKey := util.GenerateSHA256(generateAPIKey)
	token := ctx.Value(constants.Token).(string)
	bearerToken := fmt.Sprintf("Bearer %s", token)

	var response PaymentResponse
	request := p.client.GetClient().Clone().
		Set(constants.Authorization, bearerToken).
		Set(constants.XServiceName, config2.Config.AppName).
		Set(constants.XApiKey, apiKey).
		Set(constants.XRequestAt, fmt.Sprintf("%d", unixTime)).
		Get(fmt.Sprintf("%s/api/v1/auth/user", p.client.GetBaseURL()))

	resp, _, errs := request.EndStruct(&response)
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("payment response: %s", response.Message)
	}

	return &response.Data, nil
}

func (p *PaymentClient) CreatePaymentLink(ctx context.Context, paymentRequest *dto.PaymentRequest) (*PaymentData, error) {
	unixTime := time.Now().Unix()
	generateAPIKey := fmt.Sprintf("%s:%s:%d",
		config2.Config.AppName,
		p.client.GetSignatureKey(),
		unixTime,
	)
	apiKey := util.GenerateSHA256(generateAPIKey)
	token := ctx.Value(constants.Token).(string)
	bearerToken := fmt.Sprintf("Bearer %s", token)

	body, err := json.Marshal(paymentRequest)
	if err != nil {
		return nil, err
	}

	resp, bodyResp, errs := p.client.GetClient().Clone().
		Post(fmt.Sprintf("%s/api/v1/auth/user", p.client.GetBaseURL())).
		Set(constants.Authorization, bearerToken).
		Set(constants.XServiceName, config2.Config.AppName).
		Set(constants.XApiKey, apiKey).
		Set(constants.XRequestAt, fmt.Sprintf("%d", unixTime)).
		Send(string(body)).
		End()
	
	if len(errs) > 0 {
		return nil, errs[0]
	}

	var response PaymentResponse
	if resp.StatusCode != http.StatusCreated {
		err = json.Unmarshal([]byte(bodyResp), &response)
		if err != nil {
			return nil, err
		}
		PaymentError := fmt.Errorf("payment response: %s", response.Message)
		return nil, PaymentError
	}

	err = json.Unmarshal([]byte(bodyResp), &response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}