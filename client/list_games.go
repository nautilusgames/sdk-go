package client

import (
	"fmt"
	"net/http"

	"github.com/sdk-go/builder"
	"github.com/sdk-go/model"
)

// API support get list game follow token from API GetToken
// params is pagination with page & page_size
func (s *Client) ListGames(params map[string]string, token string) (*model.GameResponse, error) {
	var resp *model.GameResponse
	url, err := builder.BuildParameterUrl(params, fmt.Sprintf("%s%s", s.domain, EndpointListGames))
	if err != nil {

		return nil, err
	}
	fmt.Println(url.String())
	err = s.client.Send(url.String(), s.BuildHeader(token), "", http.MethodGet, &resp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resp, nil
}
