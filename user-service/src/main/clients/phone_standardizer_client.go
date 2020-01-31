package clients

import (
	"bytes"
	"encoding/json"
	"github.com/vtomar01/user-service/src/main/context"
	"github.com/vtomar01/user-service/src/main/dtos"
	"net/http"
)

type PhoneStandardizerClient struct {
	basePath string
	client   *http.Client
}

func NewPhoneStandardizerClient(basePath string, client *http.Client) *PhoneStandardizerClient {
	return &PhoneStandardizerClient{basePath: basePath, client: client}
}

func (p *PhoneStandardizerClient) Standardize(ctx *context.Context,
	request *dtos.PhoneStandardizerRequest) (*dtos.PhoneStandardizerResponse, error) {

	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest("POST",
		p.basePath+"/api/v1/standardize/", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	var response dtos.PhoneStandardizerResponse
	err = Execute(ctx, p.client, httpReq, &response, http.StatusOK)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
