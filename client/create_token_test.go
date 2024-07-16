package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"go.uber.org/zap"
)

func TestCreateToken(t *testing.T) {
	ctx := context.TODO()
	logger := zap.NewExample()
	httpClient := &http.Client{}
	client := NewClient(httpClient).
		WithDomain("https://t.ssn-571.com").
		WithAPIKey("<your-api-key>").
		WithTenantID("<your-tenant-id>")

	// Create token for tenant that will be used to authenticate the tenant in the system
	token, err := client.CreateToken(ctx)
	if err != nil {
		logger.Error("create token fail", zap.Error(err))
		return
	}
	
	fmt.Println(token)
}
