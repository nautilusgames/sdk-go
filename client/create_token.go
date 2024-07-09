package client

import (
	"fmt"
	"net/http"

	"github.com/sdk-go/model"
)

// CreateToken Create token for tenant that will be used to authenticate the tenant in the system
func (s *Client) CreateToken() (*model.TokenResponse, error) {
	var resp *model.TokenResponse
	err := s.client.Send(fmt.Sprintf("%s%s", s.domain, EndpointGetToken), s.BuildHeader(""), "", http.MethodPost, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
