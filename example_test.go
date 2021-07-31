package modzy_test

import (
	"context"
	"fmt"
	"log"

	modzy "github.com/modzy/go-sdk"
)

func ExampleNewClient() {
	ctx := context.TODO()
	client := modzy.NewClient("https://your-base-url.example.com").WithAPIKey("your-api-key")
	details, err := client.Models().GetModelDetails(ctx, &modzy.GetModelDetailsInput{
		ModelID: "e3f73163d3",
	})
	if err != nil {
		log.Fatalf("Failed to get model details: %v", err)
	}
	fmt.Printf("Found latest version: %s\n", details.Details.LatestVersion)
}
