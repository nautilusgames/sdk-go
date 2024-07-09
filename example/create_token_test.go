package example

import (
	"fmt"
	"net/http"
	"testing"

	"go.uber.org/zap"

	"github.com/sdk-go/client"
)

func TestGetToken(t *testing.T) {
	log := zap.NewExample()
	sv := client.NewClient(&http.Client{}, log).
		WithDomain("your-domain").
		WithAPIKey("your-api-key").
		WithTenantID("your-tenant-id")
	token, err := sv.CreateToken()
	if err != nil {
		return
	}

	fmt.Println(token)
}
