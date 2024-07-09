package example

import (
	"fmt"
	"net/http"
	"testing"

	"go.uber.org/zap"

	"github.com/sdk-go/client"
)

func TestCreateToken(t *testing.T) {
	log := zap.NewExample()
	sv := client.NewClient(&http.Client{}).
		WithDomain("your-domain").
		WithAPIKey("your-api-key").
		WithTenantID("your-tenant-id")
	token, err := sv.CreateToken()
	if err != nil {
		log.Error("create token fail", zap.Error(err))
		return
	}

	fmt.Println(token)
}
