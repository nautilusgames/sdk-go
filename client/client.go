package client

import (
	"github.com/sdk-go/constant"
	"net/http"

	"go.uber.org/zap"

	sdkhttp "github.com/sdk-go/http"
)

const (
	//Endpoint APIs
	EndpointGetToken  = "/api/hydra/v1/token"
	EndpointListGames = "/api/gnome/v1/games"
)

type Client struct {
	Logger   *zap.Logger
	client   *sdkhttp.ClientWrapper
	domain   string
	tenantID string
	apiKey   string
}

func NewClient(client *http.Client, domain, tenantID, apiKey string, log *zap.Logger) (*Client, error) {
	return &Client{
		Logger:   log,
		client:   sdkhttp.NewClientWrapper(client, log),
		domain:   domain,
		tenantID: tenantID,
		apiKey:   apiKey,
	}, nil
}

func (c *Client) BuildHeader(token string) map[string]string {
	mHeader := make(map[string]string)
	mHeader[constant.XAPIkey] = c.apiKey
	mHeader[constant.XTenantId] = c.tenantID
	if len(token) > 0 {
		mHeader[constant.Authorization] = "Bearer " + token
	}
	return mHeader
}
