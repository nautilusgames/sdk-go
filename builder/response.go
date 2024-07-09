package builder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	_contentType    = "Content-Type"
	_contentTypeVal = "application/json; charset=UTF-8"
)

func SendResponse(writer http.ResponseWriter, res interface{}) {
	writer.Header().Set(_contentType, _contentTypeVal)
	writer.WriteHeader(http.StatusOK)

	bytes, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("json marshal response error %v \n", err)
		return
	}
	if _, err = writer.Write(bytes); err != nil {
		fmt.Printf("write message error %v \n", err)
	}
}

func ToResponse(resp *http.Response, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %s, body: %s", resp.Status, string(body))
	}

	return body, nil
}
