package Subscriber

import (
	"EmployeeService/Controller/ControllerSubscriber"
	"context"

	"cloud.google.com/go/pubsub"
)

var pubsubClient *pubsub.Client
var TopicSubscriptionController = map[PubSubSubscription]PubSubTopicSubscription{}

type PubSubSubscription string

type PubSubTopicSubscription struct {
	SubscriptionID     PubSubSubscription
	Subscription       *pubsub.Subscription
	SubcriberContoller func(ctx context.Context, msg *pubsub.Message)
}

type PubSubSubscriberConfig struct {
	SubscriberController ControllerSubscriber.SubscriberInterface
}
