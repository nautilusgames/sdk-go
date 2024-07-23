package client

import (
	"fmt"
	"net/http"

	sdkhttp "github.com/nautilusgames/sdk-go/http"
	"github.com/nautilusgames/sdk-go/webhook"
)

const (
	// Endpoint APIs
	EndpointGetToken  = "/api/hydra/v1/token"
	EndpointListGames = "/api/gnome/v1/games"
)

type Client struct {
	client   *sdkhttp.ClientWrapper
	domain   string
	tenantID string
	apiKey   string
	// tenantToken is optional, you can call API with token or apiKey
	tenantToken string
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

// WithDomain is a function that helps add the Domain value to the client struct for API calls.
func (c *Client) WithTenantToken(tenantToken string) *Client {
	c.tenantToken = tenantToken
	return c
}

// BuildHeader This function supports building the header to convey information in the API when making a call.
func (c *Client) BuildHeader() map[string]string {
	mHeader := make(map[string]string)
	mHeader[webhook.XTenantId] = c.tenantID
	if len(c.tenantToken) > 0 {
		mHeader[webhook.Authorization] = fmt.Sprintf("Bearer %s", c.tenantToken)
	} else if len(c.apiKey) > 0 {
		mHeader[webhook.XApiKey] = c.apiKey
	}
	return mHeader
}
