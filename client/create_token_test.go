package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"go.uber.org/zap"
)

func TestCreateToken(t *testing.T) {
	var (
		ctx = context.TODO()
		log = zap.NewExample()
		sv  = NewClient(&http.Client{}).
			WithDomain("your-domain").
			WithAPIKey("your-api-key").
			WithTenantID("your-tenant-id")
	)

	token, err := sv.CreateToken(ctx)
	if err != nil {
		log.Error("create token fail", zap.Error(err))
		return
	}

	fmt.Println(token)
}
