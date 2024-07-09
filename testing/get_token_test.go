package testing

import (
	"fmt"
	"net/http"
	"testing"

	"go.uber.org/zap"

	"github.com/sdk-go/client"
)

func TestGetToken(t *testing.T) {
	log := &zap.Logger{}
	sv, err := client.NewClient(&http.Client{},
		log,
		"https://test.com",
		"1",
		"api-test",
	)
	if err != nil {
		return
	}
	token, err := sv.GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(token)
}
