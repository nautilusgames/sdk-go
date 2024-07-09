package testing

import (
	"fmt"
	"net/http"
	"testing"

	"go.uber.org/zap"

	"github.com/sdk-go/client"
)

func TestGetToken(t *testing.T) {
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

	fmt.Println(token)
}
