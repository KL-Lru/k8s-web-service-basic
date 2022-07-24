package mq

import (
	"context"

	"cloud.google.com/go/pubsub"
)

func (mq *PubSubMQ) Publish(topic, msg string) (string, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, mq.ProjectID)
	if err != nil {
		return "", err
	}
	defer client.Close()

	t := client.Topic(topic)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})

	return result.Get(ctx)
}
