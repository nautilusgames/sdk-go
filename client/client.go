package client

import (
	"net/http"

	sdkhttp "github.com/sdk-go/http"
)

const (
	authorization     = "Authorization"
	xApiKey           = "x-api-key"
	xTenantId         = "x-tenant-id"
	tenantPlayerToken = "tenant-player-token"

	//Endpoint APIs
	EndpointGetToken  = "/api/hydra/v1/token"
	EndpointListGames = "/api/gnome/v1/games"
)

type Client struct {
	client   *sdkhttp.ClientWrapper
	domain   string
	tenantID string
	apiKey   string
}

func NewClient(client *http.Client) *Client {
	return &Client{
		client: sdkhttp.NewClientWrapper(client),
	}
}

// WithAPIKey is a function that helps add the API key value to the client struct for API calls.
func (c *Client) WithAPIKey(apiKey string) *Client {
	c.apiKey = apiKey
	return c
}

// WithTenantID is a function that helps add the TenantID value to the client struct for API calls.
func (c *Client) WithTenantID(tenantID string) *Client {
	c.tenantID = tenantID
	return c
}

// WithDomain is a function that helps add the Domain value to the client struct for API calls.
func (c *Client) WithDomain(domain string) *Client {
	c.domain = domain
	return c
}

// BuildHeader This function supports building the header to convey information in the API when making a call.
func (c *Client) BuildHeader(token string) map[string]string {
	mHeader := make(map[string]string)
	mHeader[xApiKey] = c.apiKey
	mHeader[xTenantId] = c.tenantID
	if len(token) > 0 {
		mHeader[authorization] = "Bearer " + token
	}
	return mHeader
}
