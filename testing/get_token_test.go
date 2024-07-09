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
	sv, err := client.NewClient(&http.Client{}, "https://p.ssn-571.com", "6", "1KAjw4y2Y5h6lOTb28OZojnBoBi+Rf73eIA/DoKID8w=", log)
	if err != nil {
		return
	}
	token, err := sv.GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(token)
}
