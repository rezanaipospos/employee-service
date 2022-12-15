package SecretManager

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
)

var SecretManagerClient *secretmanager.Client

type SecretManagerServices interface {
	Initialize()
}

type SecretManagerConfig struct {
	SecretManagerServices
}
