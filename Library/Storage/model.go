package Storage

import (
	"EmployeeService/Config"

	"cloud.google.com/go/storage"
)

var clientConfig StorageConfig

type StorageConfigServices interface {
	Initialize(env *Config.Environment)
}

type StorageConfig struct {
	StorageConfigServices
	client     *storage.Client
	bucketName string
	bucket     *storage.BucketHandle
}
