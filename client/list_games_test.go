package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"go.uber.org/zap"
)

func TestListGames(t *testing.T) {
	ctx := context.TODO()
	logger := zap.NewExample()
	client := NewClient(&http.Client{}).
		WithDomain("your-domain").
		WithAPIKey("your-api-key").
		WithTenantID("your-tenant-id")

	// ListGames API support get list game follow token from API Key
	games, err := client.ListGames(ctx, &ListGamesRequest{
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		logger.Error("list games failed", zap.Error(err))
		return
	}

	fmt.Println(games)
}
