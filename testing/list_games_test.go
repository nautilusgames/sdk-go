package testing

import (
	"fmt"
	"github.com/sdk-go/client"
	"go.uber.org/zap"
	"net/http"
	"testing"
)

func TestListGames(t *testing.T) {
	log := &zap.Logger{}
	sv, err := client.NewClient(&http.Client{}, "https://p.ssn-571.com", log)
	if err != nil {
		return
	}

	mHeader := make(map[string]string)
	mHeader["x-api-key"] = "1KAjw4y2Y5h6lOTb28OZojnBoBi+Rf73eIA/DoKID8w="
	mHeader["x-tenant-id"] = "6"

	token, err := sv.GetToken(mHeader)
	if err != nil {
		fmt.Println(err)
		return
	}

	mHeaderG := make(map[string]string)
	mHeaderG["Authorization"] = "Bearer " + token.Token

	mParams := make(map[string]string)
	mParams["page"] = "1"
	mParams["page_size"] = "20"
	games, err := sv.ListGames(mParams, mHeaderG)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(games)
}
