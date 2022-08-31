package main

import (
	"context"
	dapr "github.com/dapr/go-sdk/client"
)

func main() {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	data := []byte("hello")
	if err := client.PublishEvent(ctx, "component-name", "topic-name", data); err != nil {
		panic(err)
	}
	defer client.Close()
}
