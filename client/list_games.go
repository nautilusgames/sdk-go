package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nautilusgames/sdk-go/builder"
)

// ListGames API support get list game follow token from API GetToken
// params is pagination with page & page_size
func (s *Client) ListGames(ctx context.Context, request *ListGamesRequest) (*ListGamesReply, error) {
	var resp *ListGamesReply
	url, err := builder.StructToURLValues(fmt.Sprintf("%s%s", s.domain, EndpointListGames), request)
	if err != nil {
		return nil, err
	}
	err = s.client.Send(ctx, url.String(), s.BuildHeader(), "", http.MethodGet, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
