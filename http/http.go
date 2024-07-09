package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sdk-go/builder"
)

type ClientWrapper struct {
	HTTPClient *http.Client
}

func NewClientWrapper(client *http.Client) *ClientWrapper {
	return &ClientWrapper{
		HTTPClient: client,
	}
}

// Send This function executes the API call based on the provided information, including method, body, header, and URL.
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
		return err
	}
	return nil
}
