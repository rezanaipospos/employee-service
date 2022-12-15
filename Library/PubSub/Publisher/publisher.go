package Publisher

import (
	"context"
	"encoding/json"
	"errors"

	"cloud.google.com/go/pubsub"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/iterator"
)

const (
	// Employee
	TOPIC_EMPLOYEEADDED               PubSubPublisherTopic = "EmployeeAdded"
	TOPIC_EMPLOYEEPERSONALINFOUPDATED PubSubPublisherTopic = "EmployeePersonalInfoUpdated"
	TOPIC_EMPLOYEEDELETED             PubSubPublisherTopic = "EmployeeDeleted"
	TOPIC_EMPLOYEEWORKSTATUSUPDATED   PubSubPublisherTopic = "EmployeeWorkStatusUpdated"
	TOPIC_EMPLOYEERESIGNED            PubSubPublisherTopic = "EmployeeResigned"
	TOPIC_EMPLOYEEFINGERUPDATED       PubSubPublisherTopic = "EmployeeFingerUpdated"
	TOPIC_EMPLOYEEMACHINEIDUPDATED    PubSubPublisherTopic = "EmployeeMachineIdUpdated"
	TOPIC_EMPLOYEEFACEIDUPDATED       PubSubPublisherTopic = "EmployeeFaceIdUpdated"

	// Employee Transfer
	TOPIC_EMPLOYEETRANSFERED PubSubPublisherTopic = "EmployeeTransfered"
)

func (c PubSubPublisherConfig) CollectPublisher() {
	// Employee
	c.CreateTopicClient(TOPIC_EMPLOYEEADDED)
	c.CreateTopicClient(TOPIC_EMPLOYEEPERSONALINFOUPDATED)
	c.CreateTopicClient(TOPIC_EMPLOYEEDELETED)
	c.CreateTopicClient(TOPIC_EMPLOYEEFINGERUPDATED)
	c.CreateTopicClient(TOPIC_EMPLOYEEWORKSTATUSUPDATED)
	c.CreateTopicClient(TOPIC_EMPLOYEERESIGNED)
	c.CreateTopicClient(TOPIC_EMPLOYEEMACHINEIDUPDATED)
	c.CreateTopicClient(TOPIC_EMPLOYEEFACEIDUPDATED)

	// Employee Transfer
	c.CreateTopicClient(TOPIC_EMPLOYEETRANSFERED)
}
func (d PubSubPublisherConfig) SetClient(client *pubsub.Client) {
	pubsubClient = client
}
func (c PubSubPublisherConfig) CreateTopicClient(topic PubSubPublisherTopic) {
	if _, ok := TopicPublisherController[topic]; !ok {
		clientTopic := pubsubClient.Topic(string(topic))
		exists, err := clientTopic.Exists(context.Background())
		if err != nil {
			log.Fatal().Msgf("Failed to connect client Topic: %s, error: %v", topic, err)
		}
		if !exists {
			log.Fatal().Msgf("PubSub Topic: %s, Not Found", topic)
		}
		TopicPublisherController[topic] = PubSubTopic{TopicID: string(topic), Publisher: clientTopic}
	}
}

func (d PubSubPublisherTopic) Get() (PubSubTopic, error) {
	if _, ok := TopicPublisherController[d]; !ok {
		return PubSubTopic{}, errors.New("pubsub topic not found")
	}
	return TopicPublisherController[d], nil
}

func (p PubSubTopic) Publish(data interface{}) (serverID string, err error) {
	ctx := context.Background()
	byteData, _ := json.Marshal(data)
	res := p.Publisher.Publish(ctx, &pubsub.Message{Data: []byte(string(byteData))})
	return res.Get(context.Background())
}

func (p PubSubTopic) Subscriptions() (subs []string, err error) {
	ctx := context.Background()
	it := p.Publisher.Subscriptions(ctx)
	for {
		sub, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		subs = append(subs, sub.String())
	}
	return subs, nil
}
