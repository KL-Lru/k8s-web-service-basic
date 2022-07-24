package main

import (
	"context"
	"fmt"

	"github.com/KL-Lru/sample-web-service/pkg/env"
	"github.com/KL-Lru/sample-web-service/pkg/heartbeat"
	"github.com/KL-Lru/sample-web-service/pkg/mq"
)

func main() {
	go heartbeat.Pulse()
	subscribeID := env.GetEnvVal("SUBSCRIBE_ID", "")
	client := mq.NewPubSubMQ()

	err := client.Subscribe(subscribeID, act)
	if err != nil {
		fmt.Println(err)
	}
}

func act(_ context.Context, message string) error {
	fmt.Printf("Receive Message: %s\n", message)
	return nil
}
