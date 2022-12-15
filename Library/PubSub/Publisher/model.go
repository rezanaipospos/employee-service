package Publisher

import (
	"EmployeeService/Controller/Dto"

	"cloud.google.com/go/pubsub"
)

var pubsubClient *pubsub.Client
var Topics []PubSubTopic
var TopicPublisherController = map[PubSubPublisherTopic]PubSubTopic{}

type PubSubPublisherTopic string
type PubSubTopic struct {
	TopicID   string
	Publisher *pubsub.Topic
}

type PubSubPublisherConfig struct {
}

type Publisher struct {
	Key               string                  `json:"key"`
	StateStore        string                  `json:"stateStore"`
	NotificationData  Dto.NotificationInfoDTO `json:"notificationData"`
	JWTDocodedPayload map[string]interface{}  `json:"jwtDocodedPayload"`
	Data              interface{}             `json:"data"`
}
