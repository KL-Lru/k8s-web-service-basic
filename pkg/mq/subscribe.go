package mq

import (
	"context"

	"cloud.google.com/go/pubsub"
)

func (mq *PubSubMQ) Subscribe(subscribeID string, callback func(context.Context, string) error) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, mq.ProjectID)
	if err != nil {
		return err
	}
	defer client.Close()

	sub := client.Subscription(subscribeID)
	// ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	// defer cancel()

	return sub.Receive(ctx, func(c context.Context, msg *pubsub.Message) {
		message := string(msg.Data)
		err := callback(c, message)
		if err != nil {
			msg.Nack()
		} else {
			msg.Ack()
		}
	})
}
