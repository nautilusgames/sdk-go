package client

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/sdk-go/model"
)

// Create token for tenant that will be used to authenticate the tenant in the system
func (s *Client) GetToken() (*model.TokenResponse, error) {
	var resp *model.TokenResponse
	err := s.client.Send(fmt.Sprintf("%s%s", s.domain, EndpointGetToken), s.BuildHeader(""), "", http.MethodPost, &resp)
	if err != nil {
		s.Logger.Error("call get token error", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
