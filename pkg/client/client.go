package client

import (
	"net/http"

	"go.uber.org/zap"

	sdkhttp "github.com/sdk-go/pkg/http"
)

const (
	//Endpoint APIs for report
	EndpointGetToken  = "/api/hydra/v1/token"
	EndpointListGames = "/api/gnome/v1/games"
)

type Client struct {
	Logger *zap.Logger
	client *sdkhttp.ClientWrapper
	domain string
}

func NewClient(client *http.Client, domain string, log *zap.Logger) (*Client, error) {
	return &Client{
		Logger: log,
		client: sdkhttp.NewClientWrapper(client, log),
		domain: domain,
	}, nil
}
