package example

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"go.uber.org/zap"

	"github.com/sdk-go/client"
	"github.com/sdk-go/model"
)

func TestListGames(t *testing.T) {
	var (
		ctx = context.TODO()
		log = zap.NewExample()
		sv  = client.NewClient(&http.Client{}).
			WithDomain("your-domain").
			WithAPIKey("your-api-key").
			WithTenantID("your-tenant-id")
	)

	// Create Token support call API ListGames
	token, err := sv.CreateToken(ctx)
	if err != nil {
		return
	}

	games, err := sv.ListGames(ctx, &model.ListGamesRequest{
		Page:     1,
		PageSize: 10,
	}, token.Token)

	if err != nil {
		log.Error("get list gamesfail", zap.Error(err))
		return
	}

	fmt.Println(games)
}
