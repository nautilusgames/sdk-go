package client

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/sdk-go/builder"
	"github.com/sdk-go/model"
)

// API support get list game follow token from API GetToken
// params is pagination with page & page_size
func (s *Client) ListGames(request *model.ListGamesRequest, token string) (*model.GameResponse, error) {
	var resp *model.GameResponse
	url, err := builder.StructToURLValues(fmt.Sprintf("%s%s", s.domain, EndpointListGames), request)
	if err != nil {
		s.Logger.Error("build parameter query url fail", zap.Error(err))
		return nil, err
	}
	err = s.client.Send(url.String(), s.BuildHeader(token), "", http.MethodGet, &resp)
	if err != nil {
		s.Logger.Error("execute get list games error", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
