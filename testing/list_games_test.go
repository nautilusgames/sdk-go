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
	sv, err := client.NewClient(&http.Client{}, "https://p.ssn-571.com", "6", "1KAjw4y2Y5h6lOTb28OZojnBoBi+Rf73eIA/DoKID8w=", log)
	if err != nil {
		return
	}

	mHeader := make(map[string]string)
	mHeader["x-api-key"] = "1KAjw4y2Y5h6lOTb28OZojnBoBi+Rf73eIA/DoKID8w="
	mHeader["x-tenant-id"] = "6"

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
