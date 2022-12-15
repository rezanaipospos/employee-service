package Jwt

import (
	"EmployeeService/Config"
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

func (j JwtStruct) authKey() *jwtauth.JWTAuth {
	return jwtauth.New(j.Config.Encrypt, []byte(j.Config.Key), nil)
}

func (j JwtStruct) Encode(ParamKeys jwt.MapClaims) (token string, err error) {
	_, token, err = j.authKey().Encode(ParamKeys)
	return
}

func (j JwtStruct) Decode(token string) (res map[string]interface{}, err error) {
	jwtToken, err := j.authKey().Decode(token)
	if err != nil {
		return
	}
	res = jwtToken.PrivateClaims()
	return
}

func (j JwtStruct) Verify(token string) (res map[string]interface{}, err error) {
	jwtToken, err := jwtauth.VerifyToken(j.authKey(), token)
	if jwtToken != nil {
		res = jwtToken.PrivateClaims()
	}
	if err != nil {
		return
	}
	if res == nil {
		res = jwtToken.PrivateClaims()
	}
	exp := jwtToken.Expiration().UnixMilli()
	if exp == 0 {
		err = errors.New("token not valid, claims exp not found")
	}
	return
}

func (j JwtStruct) GetConfig() *Config.JwtSetting {
	config := Config.JwtSetting{
		Key:     j.Config.Key,
		Encrypt: j.Config.Encrypt,
	}

	return &config
}
