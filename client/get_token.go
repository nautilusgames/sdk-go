package client

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/sdk-go/model"
)

// Create token for tenant that will be used to authenticate the tenant in the system
// Header :
// x-api-key : Your API Key created on N11s Back Office
// x-tenant-id: Your Tenant ID
func (s *Client) GetToken(header map[string]string) (*model.TokenResponse, error) {
	var resp *model.TokenResponse
	err := s.client.Send(fmt.Sprintf("%s%s", s.domain, EndpointGetToken), header, "", http.MethodPost, &resp)
	if err != nil {
		s.Logger.Error("call get token error", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
