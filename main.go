package main

import (
	"context"
	"fmt"

	client "github.com/y16ra/gmo-coin-go-sdk/client"
)

func main() {
	// GMO Coin client
	cfg := client.NewConfiguration()
	apiClient := client.NewAPIClient(cfg)
	ctx := context.Background()
	req := apiClient.PublicApi.PublicV1StatusGet(ctx)
	res, _, _ := req.Execute()
	fmt.Println(res.Data)
}
