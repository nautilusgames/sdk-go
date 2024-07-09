package testing

import (
	"fmt"
	"net/http"
	"testing"

	"go.uber.org/zap"

	"github.com/sdk-go/client"
)

func TestListGames(t *testing.T) {
	log := &zap.Logger{}
	sv, err := client.NewClient(&http.Client{},
		log,
		"https://test.com",
		"your-tenant-id",
		"your-api-key",
	)
	if err != nil {
		return
	}
	// Call API Get Token support API call ListGames
	token, err := sv.GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	mParams := make(map[string]string)
	mParams["page"] = "1"
	mParams["page_size"] = "20"
	games, err := sv.ListGames(mParams, token.Token)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(games)
}
