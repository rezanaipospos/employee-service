package Jwt

import (
	"EmployeeService/Config"

	"github.com/dgrijalva/jwt-go"
)

type JwtStruct struct {
	Config *Config.JwtSetting
}

type JwtServices interface {
	Encode(ParamKeys jwt.MapClaims) (token string, err error)
	Decode(token string) (res map[string]interface{}, err error)
	Verify(token string) (res map[string]interface{}, err error)
	GetConfig() *Config.JwtSetting
}
