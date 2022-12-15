package Logging

import (
	"EmployeeService/Config"

	logging "cloud.google.com/go/logging"
)

var Logger *logging.Logger
var Environment string
var ServiceGroup string

type LoggingServices interface {
	Initialize(env *Config.Environment)
}

type LoggingConfig struct {
	LoggingServices
	Config *Config.GCPConfig
}
