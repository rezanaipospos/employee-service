package Dapr

import (
	"EmployeeService/Config"
	"time"
)

var Dapr ConfigDapr

type DaprService interface {
	Initialize(env *Config.Environment)
}

type ConfigDapr struct {
	DaprService
	ApiKey     string `json:"token"`
	StateStore string `json:"stateStore"`
	Host       string `json:"host"`
	Port       string `json:"port"`
}

type ConfigRequest struct {
	AppID   string `json:"app_id"`
	Payload []byte `json:"payload"`
	Method  string `json:"method"`
	Path    string `json:"path"`
}

type ConfigState struct {
	StateStore string   `json:"stateStore"`
	StateFmt   StateFmt `json:"stateFmt"`
}

type StateFmt struct {
	Key   string     `json:"key"`
	Value StateValue `json:"value"`
}

type StateValue struct {
	Services []StateService `json:"services"`
	State    interface{}    `json:"state"` //state is payload from client
}

// status failed,ok,progress
type StateService struct {
	Name      string     `json:"name"`
	Status    string     `json:"status"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
