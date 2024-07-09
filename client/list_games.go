package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nautilusgames/sdk-go/builder"
	"github.com/nautilusgames/sdk-go/model"
)

// ListGames API support get list game follow token from API GetToken
// params is pagination with page & page_size
func (s *Client) ListGames(ctx context.Context, request *model.ListGamesRequest, token string) (*model.GameResponse, error) {
	var resp *model.GameResponse
	url, err := builder.StructToURLValues(fmt.Sprintf("%s%s", s.domain, EndpointListGames), request)
	if err != nil {
		return nil, err
	}
	err = s.client.Send(ctx, url.String(), s.BuildHeader(token), "", http.MethodGet, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
