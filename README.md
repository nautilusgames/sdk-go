# Nautilus Games Go SDK

The `sdk-go` package provides a simple and efficient way to interact with our API during the integration. Below is a basic example of how to use the SDK to perform a common task.

```go
package main

import (
    "fmt"

    sgc "github.com/nautilusgames/sdk-go/client"
)

func main() {
	ctx := context.TODO()
	logger := zap.NewExample()
	client := sgc.NewClient(&http.Client{}).
		WithDomain("https://t.ssn-571.com").
		WithAPIKey("your-api-key").
		WithTenantID("your-tenant-id")

	// Get list games from N11s with your API Key
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
```

To explore more features and functionalities of the `sdk-go` package, we highly recommend checking out our comprehensive documentation available at [https://docs.nautilusgames.com/](https://docs.nautilusgames.com/). The documentation provides detailed explanations, code examples, and usage guidelines to help you leverage the full potential of our API integration. Whether you are a beginner or an experienced developer, our documentation will serve as a valuable resource to enhance your understanding and streamline your development process. Happy coding!
