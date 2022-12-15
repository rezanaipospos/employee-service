package PubSub

import (
	"EmployeeService/Config"
	"EmployeeService/Library/SecretManager"
	"context"

	"github.com/rs/zerolog/log"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

func (c PubSubConfig) Initialize(env *Config.Environment) {
	ctx := context.Background()
	var err error
	secretKey, err := SecretManager.GetSecret(env.GoogleCP.ResourceID, "pubsub-secret")
	if err != nil {
		log.Fatal().Msgf("Failed to Get Secret Manager: %v", err)
	}
	optionClient := []option.ClientOption{
		option.WithCredentialsJSON(secretKey),
	}
	pubsubClient, err := pubsub.NewClient(ctx, env.GoogleCP.ProjectID, optionClient...)
	if err != nil {
		log.Fatal().Msgf("Failed to connect client: %v", err)
	}

	log.Info().Msg("CONNECTED with PubSub ....")
	//Collect All Publisher
	c.Publisher.SetClient(pubsubClient)
	c.Publisher.CollectPublisher()

	//Collect All Subcriber
	c.Subscriber.SetClient(pubsubClient)
	c.Subscriber.Subscribe()

	log.Info().Msg("PubSub Subcriber start ....")
}
