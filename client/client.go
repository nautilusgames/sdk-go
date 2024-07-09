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

func NewClient(client *http.Client, log *zap.Logger) *Client {
	return &Client{
		Logger: log,
		client: sdkhttp.NewClientWrapper(client, log),
	}
}

func (c *Client) WithAPIKey(apiKey string) *Client {
	c.apiKey = apiKey
	return c
}

func (c *Client) WithTenantID(tenantID string) *Client {
	c.tenantID = tenantID
	return c
}

func (c *Client) WithDomain(domain string) *Client {
	c.domain = domain
	return c
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
