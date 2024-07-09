package helper

import (
	"encoding/json"
	"fmt"
	"io"
)

func ToRequest(b io.ReadCloser, v interface{}) error {
	// Read the request body
	bytes, err := io.ReadAll(b)
	if err != nil {
		fmt.Printf("read body error %v \n", err)
		return err
	}
	if err := json.Unmarshal(bytes, v); err != nil {
		fmt.Printf("unmarshal request body error %v \n", err)
		return err
	}
	defer func() {
		if b != nil {
			err := b.Close()
			if err != nil {
				return
			}
		}
	}()

	return nil
}
