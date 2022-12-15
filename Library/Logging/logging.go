package Logging

import (
	"EmployeeService/Config"
	"EmployeeService/Library/SecretManager"
	"context"
	"net/http"

	"cloud.google.com/go/logging"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
)

func (c LoggingConfig) Initialize(env *Config.Environment) {
	Environment = env.App.Environment
	ServiceGroup = env.App.ServiceGroup
	secretKey, err := SecretManager.GetSecret(env.GoogleCP.ResourceID, "logging-secret")
	if err != nil {
		log.Fatal().Msgf("Failed to Get Secret Manager: %v", err)
	}
	optionClient := []option.ClientOption{
		option.WithCredentialsJSON(secretKey),
	}
	client, err := logging.NewClient(context.Background(), env.GoogleCP.ProjectID, optionClient...)
	if err != nil {
		log.Fatal().Msgf("Google Logging initialization failed: %v", err)
	}
	Logger = client.Logger(env.App.Appname)

	log.Info().Msg("Google Logging start and Connected... ")

}

func LogError(payload interface{}, r *http.Request) {
	if r != nil {
		Logger.Log(logging.Entry{
			Severity: logging.Error,
			Payload:  payload,
			HTTPRequest: &logging.HTTPRequest{
				Request: r,
			},
			Labels: map[string]string{
				"name":          "Microservices",
				"service-group": ServiceGroup,
				"environment":   Environment,
			},
		})
	} else {
		Logger.Log(logging.Entry{
			Severity: logging.Error,
			Payload:  payload,
			Labels: map[string]string{
				"name":          "Microservices",
				"service-group": ServiceGroup,
				"environment":   Environment,
			},
		})
	}
}

func LogInfo(payload interface{}) {
	Logger.Log(logging.Entry{
		Severity: logging.Info,
		Payload:  payload,
		Labels: map[string]string{
			"name":          "Microservices",
			"service-group": ServiceGroup,
			"environment":   Environment,
		},
	})
}

func LogWarning(payload interface{}) {
	Logger.Log(logging.Entry{
		Severity: logging.Warning,
		Payload:  payload,
		Labels: map[string]string{
			"name":          "Microservices",
			"service-group": ServiceGroup,
			"environment":   Environment,
		},
	})
}
