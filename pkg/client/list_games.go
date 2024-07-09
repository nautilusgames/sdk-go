package client

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/sdk-go/pkg/helper"
	"github.com/sdk-go/pkg/model"
)

// API support get list game follow token from API GetToken
// Header :
// Authorization : Your Token
func (s *Client) ListGames(params map[string]string, header map[string]string) (*model.GameResponse, error) {
	var resp *model.GameResponse
	url, err := helper.BuildParameterUrl(params, fmt.Sprintf("%s%s", s.domain, EndpointListGames))
	if err != nil {
		return nil, err
	}
	err = s.client.Send(url.String(), header, "", http.MethodGet, &resp)
	if err != nil {
		s.Logger.Error("call get list games error", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
