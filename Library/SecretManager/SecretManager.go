package SecretManager

import (
	"context"
	"errors"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/rs/zerolog/log"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func (c SecretManagerConfig) Initialize() {

	var err error
	// Create the client.
	ctx := context.Background()
	SecretManagerClient, err = secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatal().Msgf("failed to setup client: %v", err)
	}
	log.Info().Msg("CONNECTED with SecretManager.. ")
}

func GetSecret(resourceId string, secretName string) ([]byte, error) {
	if SecretManagerClient != nil {
		// Call the API.
		result, err := SecretManagerClient.AccessSecretVersion(context.Background(), &secretmanagerpb.AccessSecretVersionRequest{
			Name: fmt.Sprintf("projects/%s/secrets/%s/versions/1", resourceId, secretName),
		})
		if err != nil {
			return nil, err
		}
		// Print the secret payload.
		//
		// WARNING: Do not print the secret in a production environment - this
		// snippet is showing how to access the secret material.
		// log.Printf("Plaintext: %s", result.Payload.Data)
		return result.Payload.Data, nil
	}
	return nil, errors.New("no client secret manager found")
}
