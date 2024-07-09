package testing

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
	sv, err := client.NewClient(&http.Client{},
		log,
		"https://your-domain.com",
		"your-tenant-id",
		"your-api-key",
	)
	if err != nil {
		return
	}
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
