package testing

import (
	"fmt"
	"net/http"
	"testing"

	"go.uber.org/zap"

	"github.com/sdk-go/client"
)

func TestGetToken(t *testing.T) {
	log := zap.NewExample()
	sv, err := client.NewClient(&http.Client{},
		log,
		"https://your-domain.com",
		"your-tenant-id",
		"your-api-key",
	)
	if err != nil {
		return
	}
	token, err := sv.CreateToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(token)
}
