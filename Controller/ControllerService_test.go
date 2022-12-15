package Controller

import (
	"EmployeeService/Config"
	"EmployeeService/Constant"
	"EmployeeService/Controller/ControllerSubscriber"
	"EmployeeService/Library/Helper/Jwt"
	"EmployeeService/Library/Logging"
	"EmployeeService/Library/PubSub"
	"EmployeeService/Library/PubSub/Subscriber"
	"EmployeeService/Library/SecretManager"
	"flag"
	"os"
	"testing"
)

var AppEnv = flag.String("env", "", "define environment")

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	flag.Parse()
	if *AppEnv == "" {
		*AppEnv = Constant.Local
	}
	newConfig := Config.GetEnvironment(*AppEnv).LoadConfig()
	service := ServiceInit(newConfig.Environment)
	service.SecretManager.Initialize()

	newConfig.Database.BuildConnection()
	service.Pubsub.Initialize(newConfig.Environment)
	service.Logging.Initialize(newConfig.Environment)

}

type service struct {
	Jwt           Jwt.JwtServices
	Pubsub        PubSub.PubSubServices
	Logging       Logging.LoggingServices
	SecretManager SecretManager.SecretManagerServices
}

func ServiceInit(Env *Config.Environment) service {
	svc := service{
		Jwt:           Jwt.JwtStruct{Config: &Env.Jwt},
		Pubsub:        PubSub.PubSubConfig{Config: &Env.GoogleCP, Subscriber: Subscriber.PubSubSubscriberConfig{SubscriberController: &ControllerSubscriber.Subscriber{}}},
		Logging:       Logging.LoggingConfig{Config: &Env.GoogleCP},
		SecretManager: SecretManager.SecretManagerConfig{},
	}
	return svc
}
