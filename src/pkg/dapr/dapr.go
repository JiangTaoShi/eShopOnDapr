package dapr

import (
	dapr "github.com/dapr/go-sdk/client"
)

var DaprClient dapr.Client

// Setup initializes the dapr instance
func Setup() {
	var err error
	DaprClient, err = dapr.NewClient()
	if err != nil {
		panic(err)
	}
}

// Close
func Close() {
	defer DaprClient.Close()
}
