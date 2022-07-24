package mq

import (
	"context"

	"github.com/KL-Lru/sample-web-service/pkg/env"
)

type MQ interface {
	Publish(topic, msg string) (string, error)
	SubscribeSubscribe(subscribeID string, callback func(context.Context, string) error) error
}

type PubSubMQ struct {
	ProjectID string
}

func NewPubSubMQ() *PubSubMQ {
	projectId := env.GetEnvVal("PROJECT_ID", "")

	return &PubSubMQ{ProjectID: projectId}
}
