package main

import (
	"context"
	"encoding/json"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("success")
	})
	s := daprd.NewServiceWithMux(":9091", r)

	sub := &common.Subscription{
		PubsubName: "rabbitmq",
		Topic:      "topic1",
		Route:      "/topic1",
	}
	s.AddTopicEventHandler(sub, eventHandler)
	log.Println("start success")
	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error: %v", err)
	}
}
func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName:%s, Topic:%s, ID:%s, Data: %v", e.PubsubName, e.Topic, e.ID, e.Data)
	// do something with the event
	return true, nil
}
