package Storage

import (
	"EmployeeService/Config"
	"EmployeeService/Library/Helper/Convert"
	"EmployeeService/Library/SecretManager"
	"context"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/rs/zerolog/log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func (c StorageConfig) Initialize(env *Config.Environment) {
	ctx := context.Background()
	var err error
	secretKey, err := SecretManager.GetSecret(env.GoogleCP.ResourceID, "cloud-storage")
	if err != nil {
		log.Fatal().Msgf("Failed to Get Secret Manager: %v", err)
	}
	optionClient := []option.ClientOption{
		option.WithCredentialsJSON(secretKey),
	}
	clientConfig = StorageConfig{
		bucketName: env.GoogleCP.Storage.Bucket,
	}
	clientConfig.client, err = storage.NewClient(ctx, optionClient...)
	if err != nil {
		log.Fatal().Msgf("Failed to connect cloud storage client: %v", err)
	}
	clientConfig.bucket = clientConfig.client.Bucket(clientConfig.bucketName)
	log.Info().Msg(" Cloud Storage start and Connected...")
}

func UploadFile(fileName string, file io.Reader) error {
	obj := clientConfig.client.Bucket(clientConfig.bucketName).Object(Convert.GetMD5Hash(fileName))
	wc := obj.NewWriter(context.Background())
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		fmt.Println(err)
		return fmt.Errorf("Writer.Close: %v", err)
	}
	return nil
}

func ReadFile(fileName string) ([]byte, error) {
	obj := clientConfig.client.Bucket(clientConfig.bucketName).Object(Convert.GetMD5Hash(fileName))
	rc, err := obj.NewReader(context.Background())
	if err != nil {
		return nil, fmt.Errorf("NewReader error: %v", err)
	}
	defer rc.Close()

	file, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, fmt.Errorf("readFile: unable to read data from bucket %q, file %q: %v", clientConfig.bucketName, fileName, err)
	}
	return file, nil
}
