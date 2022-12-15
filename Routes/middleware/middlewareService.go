package middleware

import (
	"EmployeeService/Constant"
	Response "EmployeeService/Library/Helper/Response"
	"EmployeeService/Library/Logging"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func (m Middleware) UserAuth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if len(tokenString) == 0 {
				Response.ResponseError(w, nil, Constant.StatusUnauthorizedTokenNotExists)
				return
			}

			if !strings.Contains(tokenString, "Bearer ") {
				Response.ResponseError(w, nil, Constant.StatusUnauthorizedBearerNotFound)
				return
			}
			tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

			claims, err := m.Jwt.Verify(tokenString)

			if err != nil {
				Response.ResponseError(w, err, Constant.StatusUnauthorizedErrorVerifying)
				return
			}
			ctx := context.WithValue(r.Context(), "claims_value", claims)
			r.Header.Set("claims_value", fmt.Sprintf("%v", claims))
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}

func (m Middleware) ApiKey() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			tokenHeader := r.Header.Get("Authorization")
			if tokenHeader == "" {
				Logging.LogError(map[string]interface{}{"error": Constant.StatusUnauthorizedTokenNotExists.Info().Description, "header": r.Header}, r)
				Response.ResponseError(w, nil, Constant.StatusUnauthorizedTokenNotExists)
				return
			}

			if m.TokenHeader.Token == tokenHeader {
				next.ServeHTTP(w, r)
				return
			}

			Logging.LogError(map[string]interface{}{"error": Constant.StatusUnauthorizedInvalidToken.Info().Description, "header": r.Header}, r)
			Response.ResponseError(w, nil, Constant.StatusUnauthorizedInvalidToken)
		})
	}
}

// func (m Middleware) BasicAuthSwagger() func(http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 			var username string
// 			var password string
// 			username, password, Ok := r.BasicAuth()
// 			if !Ok {
// 				w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic `+username+password))
// 				w.WriteHeader(http.StatusUnauthorized)
// 				w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
// 				return
// 			}

// 			if m.SwaggerSetting.Username == username && m.SwaggerSetting.Password == password {
// 				next.ServeHTTP(w, r)
// 				return
// 			}

// 			w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic `+username+password))
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
// 			return

// 		})
// 	}
// }

type limitBuffer struct {
	*bytes.Buffer
	limit int
}

func newLimitBuffer(size int) io.ReadWriter {
	return limitBuffer{
		Buffer: bytes.NewBuffer(make([]byte, 0, size)),
		limit:  size,
	}
}

func Logger(logger *zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			log := logger.With().Logger()

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			buf := newLimitBuffer(512)
			ww.Tee(buf)
			t1 := time.Now()
			defer func() {
				t2 := time.Now()
				var err error
				if rec := recover(); rec != nil {
					switch t := rec.(type) {
					case string:
						err = errors.New(t)
					case error:
						err = t
					default:
						err = errors.New("unknown error")
					}
					log.Error().Timestamp().Interface("recover_info", rec).Interface("httpRequest", r).Interface("httpRequestPath", r.URL.Path).Interface("httpRequestMethod", r.Method).Bytes("debug_stack", debug.Stack()).Msg("error_request")
					Response.ResponseError(ww, err, Constant.StatusInternalServerError)
				}

				// log end request
				log.Info().
					Str("type", "access").
					Timestamp().
					Fields(map[string]interface{}{
						"remoteIP":   r.RemoteAddr,
						"requestURL": r.RequestURI,
						"rawQuery":   r.URL.RawQuery,
						"proto":      r.Proto,
						"method":     r.Method,
						"userAgent":  r.Header.Get("User-Agent"),
						"status":     ww.Status(),
						"latencyMs":  float64(t2.Sub(t1).Nanoseconds()) / 1000000.0,
						"bytesIn":    r.Header.Get("Content-Length"),
						"bytesOut":   ww.BytesWritten(),
					}).
					Msg("incoming_request")

				switch status := ww.Status(); {
				case status >= 400 && status < 500:
					reqBody, _ := ioutil.ReadAll(r.Body)
					reqPayload := string(reqBody)
					respBody, _ := ioutil.ReadAll(buf)
					resPayload := string(respBody)
					log.Warn().
						Str("type", "warn").
						Timestamp().
						Interface("reqPayload", reqPayload).
						Interface("resPayload", resPayload).
						Interface("httpRequestPath", r.URL.Path).
						Interface("httpRequestMethod", r.Method).
						Msg("log system warn")
				case status >= 500:
					reqBody, _ := ioutil.ReadAll(r.Body)
					reqPayload := string(reqBody)
					respBody, _ := ioutil.ReadAll(buf)
					resPayload := string(respBody)
					log.Error().
						Str("type", "error").
						Timestamp().
						Interface("reqPayload", reqPayload).
						Interface("resPayload", resPayload).
						Interface("httpRequestPath", r.URL.Path).
						Interface("httpRequestMethod", r.Method).
						Msg("log system error")
				}

			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
