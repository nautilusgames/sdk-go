package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"github.com/sdk-go/builder"
)

type ClientWrapper struct {
	logger     *zap.Logger
	HTTPClient *http.Client
}

func NewClientWrapper(client *http.Client, logger *zap.Logger) *ClientWrapper {
	t := http.DefaultTransport.(*http.Transport).Clone()
	//proxyUrl, _ := url.Parse("http://127.0.0.1:7951")
	//t := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100
	client.Transport = t

	return &ClientWrapper{
		logger:     logger,
		HTTPClient: client,
	}
}

func (c *ClientWrapper) Send(endpoint string, header map[string]string, body string, method string, resultObject interface{}) error {
	req, err := http.NewRequest(method, endpoint, strings.NewReader(body))
	if err != nil {
		return err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	response, err := builder.ToResponse(c.HTTPClient.Do(req))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(response, &resultObject); err != nil {
		c.logger.Error(fmt.Sprintf("cannot unmarshal response due to %v", err))
		return err
	}
	return nil
}
