package Config

import (
	"github.com/go-chi/chi/v5"
)

type app struct {
	Appname      string `yaml:"name"`
	ServiceGroup string `yaml:"service_group"`
	Debug        bool   `yaml:"debug"`
	Port         string `yaml:"port"`
	Service      string `yaml:"service"`
	Certificate  string `yaml:"certificate"`
	Pem_key      string `yaml:"pem_key"`
	Host         string `yaml:"host"`
	Environment  string `yaml:"environment"`
}

type database struct {
	Name               string `yaml:"name"`
	Username           string `yaml:"username"`
	Password           string `yaml:"password"`
	Port               string `yaml:"port"`
	Engine             string `yaml:"engine"`
	Host               string `yaml:"host"`
	Maximum_connection int    `yaml:"maximum_connection"`
	Usage              string `yaml:"usage"`
	Connection         string `yaml:"connection"`
	IsPrivateIP        bool   `yaml:"is_private_ip"`
	PrivateIP          string `yaml:"private_ip"`
	DbIAMUser          string `yaml:"db_iam_user"`
}

type dbConfig []database

type Environment struct {
	App       app            `yaml:"app"`
	Databases dbConfig       `yaml:"databases"`
	Jwt       JwtSetting     `yaml:"jwt"`
	Kafka     KafkaConfig    `yaml:"kafka"`
	GoogleCP  GCPConfig      `yaml:"google_cp"`
	Dapr      DaprConfig     `yaml:"dapr"`
	ApiToken  TokenHeader    `yaml:"auth"`
	Swagger   SwaggerSetting `yaml:"swagger"`
}

type DaprConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type JwtSetting struct {
	Key     string `yaml:"secretkey"`
	Encrypt string `yaml:"encrypt"`
}

type SwaggerSetting struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
	// Topic   string   `yaml:"topic"`
}

type GCPConfig struct {
	PubSub     PubSubConfig       `yaml:"pubsub"`
	Logging    LoggingConfig      `yaml:"logging"`
	ProjectID  string             `yaml:"project_id"`
	ResourceID string             `yaml:"resource_id"`
	GACPath    string             `yaml:"gac_path"`
	Storage    CloudStorageConfig `yaml:"storage"`
}
type LoggingConfig struct{}
type PubSubConfig struct {
	ProjectID string `yaml:"projectId"`
	GACPath   string `yaml:"gacPath"`
}
type CloudStorageConfig struct {
	Bucket string `yaml:"bucket"`
}
type envFile []byte

type Config interface {
	LoadConfig() *ConfigSetting
}
type Db interface {
	BuildConnection()
}

type ConfigSetting struct {
	Database    Db
	Routes      RouteInterface
	HttpEngine  HttpEngine
	Environment *Environment
}

type TokenHeader struct {
	Token string `yaml:"token"`
}

type HttpEngine interface {
	runWithHttp(route *chi.Mux)
	runWithHttps(route *chi.Mux)
	Run(route *chi.Mux)
}

type RouteInterface interface {
	SetCors()
	CollectRoutes() *chi.Mux
}
