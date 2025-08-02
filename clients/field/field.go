package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"order-service/clients/config"
	"order-service/common/util"
	config2 "order-service/config"
	"order-service/constants"
	"order-service/domain/dto"
	"time"

	"github.com/google/uuid"
)

type FieldClient struct {
	client config.IClientConfig
}

type IFieldClient interface {
	GetFieldByUUID(context.Context, uuid.UUID) (*FieldData, error)
	UpdateStatus(request *dto.UpdateFieldScheduleStatusRequest) error
}

func NewFieldClient(client config.IClientConfig) *FieldClient {
	return &FieldClient{client: client}	
}

func (f *FieldClient) GetFieldByUUID(ctx context.Context, uuid uuid.UUID) (*FieldData, error) {
	unixTime := time.Now().Unix()
	generateAPIKey := fmt.Sprintf("%s:%s:%d",
		config2.Config.AppName,
		f.client.GetSignatureKey(),
		unixTime,
	)
	apiKey := util.GenerateSHA256(generateAPIKey)
	token := ctx.Value(constants.Token).(string)
	bearerToken := fmt.Sprintf("Bearer %s", token)

	var response FieldResponse
	request := f.client.GetClient().Clone().
		Set(constants.Authorization, bearerToken).
		Set(constants.XServiceName, config2.Config.AppName).
		Set(constants.XApiKey, apiKey).
		Set(constants.XRequestAt, fmt.Sprintf("%d", unixTime)).
		Get(fmt.Sprintf("%s/api/v1/auth/field/shecuhedule/%s", f.client.GetBaseURL(), uuid))

	resp, _, errs := request.EndStruct(&response)
	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("user response: %s", response.Message)
	}

	return &response.Data, nil
}

func (f *FieldClient) UpdateStatus(request *dto.UpdateFieldScheduleStatusRequest) error {
	unixTime := time.Now().Unix()
	generateAPIKey := fmt.Sprintf("%s:%s:%d",
		config2.Config.AppName,
		f.client.GetSignatureKey(),
		unixTime,
	)
	apiKey := util.GenerateSHA256(generateAPIKey)

	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	resp, bodyResp, errs := f.client.GetClient().Clone().
		Post(fmt.Sprintf("%s/api/v1/auth/payment", f.client.GetBaseURL())).
		Set(constants.XServiceName, config2.Config.AppName).
		Set(constants.XApiKey, apiKey).
		Set(constants.XRequestAt, fmt.Sprintf("%d", unixTime)).
		Send(string(body)).
		End()
	
	if len(errs) > 0 {
		return errs[0]
	}

	var response FieldResponse
	if resp.StatusCode != http.StatusCreated {
		err = json.Unmarshal([]byte(bodyResp), &response)
		if err != nil {
			return err
		}
		FieldError := fmt.Errorf("payment response: %s", response.Message)
		return FieldError
	}

	err = json.Unmarshal([]byte(bodyResp), &response)
	if err != nil {
		return err
	}

	return nil
}