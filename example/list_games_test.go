package example

import (
	"fmt"
	"net/http"
	"testing"

	"go.uber.org/zap"

	"github.com/sdk-go/client"
	"github.com/sdk-go/model"
)

func TestListGames(t *testing.T) {
	log := zap.NewExample()
	sv := client.NewClient(&http.Client{}, log).
		WithDomain("your-domain").
		WithAPIKey("your-api-key").
		WithTenantID("your-tenant-id")
	// Call API Get Token support API call ListGames
	token, err := sv.CreateToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	games, err := sv.ListGames(&model.ListGamesRequest{
		Page:     1,
		PageSize: 10,
	}, token.Token)
	if err != nil {
		return
	}

	fmt.Println(games)
}
