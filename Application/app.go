package Application

import (
	"EmployeeService/Constant"
	"EmployeeService/Controller/ControllerSubscriber"
	"EmployeeService/Library/Dapr"
	"EmployeeService/Library/Helper/Jwt"
	"EmployeeService/Library/Logging"
	"EmployeeService/Library/PubSub"
	"EmployeeService/Library/PubSub/Subscriber"
	"EmployeeService/Library/Scheduler"
	"EmployeeService/Library/SecretManager"
	"EmployeeService/Library/Storage"
	"flag"

	"EmployeeService/Routes"

	"EmployeeService/Routes/middleware"

	"EmployeeService/Controller"

	"EmployeeService/Config"

	"github.com/go-chi/chi/v5"
)

var AppEnv = flag.String("env", "", "define environment")

func init() {
	flag.Parse()
	if *AppEnv == "" {
		*AppEnv = Constant.Local

	}
}

func AppInitialization() {
	newConfig := Config.GetEnvironment(*AppEnv).LoadConfig()
	service := ServiceInit(newConfig.Environment)
	service.SecretManager.Initialize()

	newConfig.Database.BuildConnection()
	service.Logging.Initialize(newConfig.Environment)
	service.Pubsub.Initialize(newConfig.Environment)
	service.Storage.Initialize(newConfig.Environment)
	service.Dapr.Initialize(newConfig.Environment)
	newConfig.Routes = &Routes.Routes{
		Chi:        chi.NewRouter(),
		Controller: &Controller.Controller{},
		Middleware: &middleware.Middleware{Jwt: service.Jwt, TokenHeader: newConfig.Environment.ApiToken},
	}
	service.Scheduler.SchedulerStart()
	newConfig.Routes.SetCors()
	routes := newConfig.Routes.CollectRoutes()
	newConfig.HttpEngine.Run(routes)
}

type service struct {
	Jwt           Jwt.JwtServices
	Pubsub        PubSub.PubSubServices
	Logging       Logging.LoggingServices
	SecretManager SecretManager.SecretManagerServices
	Scheduler     Scheduler.SchedulerServices
	Storage       Storage.StorageConfigServices
	Dapr          Dapr.DaprService
}

func ServiceInit(Env *Config.Environment) service {
	svc := service{
		Jwt:           Jwt.JwtStruct{Config: &Env.Jwt},
		Pubsub:        PubSub.PubSubConfig{Config: &Env.GoogleCP, Subscriber: Subscriber.PubSubSubscriberConfig{SubscriberController: &ControllerSubscriber.Subscriber{}}},
		Logging:       Logging.LoggingConfig{Config: &Env.GoogleCP},
		SecretManager: SecretManager.SecretManagerConfig{},
		Scheduler:     Scheduler.Scheduler{},
		Storage:       Storage.StorageConfig{},
		Dapr:          Dapr.ConfigDapr{},
	}
	return svc
}
