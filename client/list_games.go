package client

import (
	"fmt"
	"net/http"

	"github.com/sdk-go/builder"
	"github.com/sdk-go/model"
)

// API support get list game follow token from API GetToken
// Header :
// Authorization : Your Token
func (s *Client) ListGames(params map[string]string, header map[string]string) (*model.GameResponse, error) {
	var resp *model.GameResponse
	url, err := builder.BuildParameterUrl(params, fmt.Sprintf("%s%s", s.domain, EndpointListGames))
	if err != nil {

		return nil, err
	}
	fmt.Println(url.String())
	err = s.client.Send(url.String(), header, "", http.MethodGet, &resp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resp, nil
}
