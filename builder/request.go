package builder

import (
	"encoding/json"
	"fmt"
	"io"
)

// ToRequest This function supports reading values from the body into the desired request/struct based on the v interface
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
