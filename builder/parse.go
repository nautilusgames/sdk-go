package builder

import "net/url"

func BuildParameterUrl(params map[string]string, baseURL string) (*url.URL, error) {
	// Create a URL object
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	// Add query parameters
	requestParameters := url.Values{}
	for k, v := range params {
		requestParameters.Add(k, v)
	}
	// Add the parameters to the URL
	u.RawQuery = requestParameters.Encode()
	return u, nil
}
