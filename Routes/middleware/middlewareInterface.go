package middleware

import (
	"EmployeeService/Config"
	"EmployeeService/Library/Helper/Jwt"
	"net/http"
)

type MiddlewareInterface interface {
	UserAuth() func(http.Handler) http.Handler
	BasicAuthSwagger() func(http.Handler) http.Handler
	ApiKey() func(http.Handler) http.Handler
}

type Middleware struct {
	MiddlewareInterface
	Jwt         Jwt.JwtServices
	TokenHeader Config.TokenHeader
}
