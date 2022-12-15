package PubSub

import (
	"EmployeeService/Config"
	"EmployeeService/Library/PubSub/Publisher"
	"EmployeeService/Library/PubSub/Subscriber"
)

type PubSubServices interface {
	Initialize(env *Config.Environment)
}

type PubSubConfig struct {
	PubSubServices
	Config     *Config.GCPConfig
	Publisher  Publisher.PubSubPublisherConfig
	Subscriber Subscriber.PubSubSubscriberConfig
}
