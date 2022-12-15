package Config

import (
	"EmployeeService/Constant"
	"context"
	"database/sql"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"path"
	"runtime"
	"sync/atomic"
	"syscall"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/yaml.v2"
)

func GetEnvironment(env string) Config {
	_, filename, _, _ := runtime.Caller(1)
	envPath := path.Join(path.Dir(filename), Constant.ENVIRONMENT_PATH+env+".yml")
	_, err := os.Stat(envPath)
	if err != nil {
		log.Info().Msg(err.Error())
		panic(err)
	}
	content, err := ioutil.ReadFile(envPath)
	if err != nil {
		log.Info().Msg(err.Error())
		panic(err)
	}
	var config envFile = content
	return config
}

func (e envFile) LoadConfig() *ConfigSetting {

	var config Environment
	err := yaml.Unmarshal([]byte(string(e)), &config)
	if err != nil {
		log.Info().Msg(err.Error())
		panic(err)
	}
	if !config.App.Debug {
		log.Output(ioutil.Discard)
	}
	log.Info().Msg("Environment Config load successfully!")
	return &ConfigSetting{&config, nil, &config.App, &config}
}

func (e *Environment) BuildConnection() {
	var connectionPool connectionPool = &database{}
	for i := 0; i < len(e.Databases); i++ {
		connectionPool = &e.Databases[i]
		switch e.Databases[i].Engine {
		case Constant.MSSQL:
			con := sql.DB{}
			log.Info().Msg("ENGINE " + Constant.MSSQL + " start....")
			connectionPool.MssqlConnectionPool(&con)
			DbConSql[DbSqlConfigName(e.Databases[i].Connection)] = &con
		case Constant.POSTGRES:
			con := sql.DB{}
			log.Info().Msg("ENGINE " + Constant.POSTGRES + " start....")
			connectionPool.PostgresConnectionPool(&con)
			DbConSql[DbSqlConfigName(e.Databases[i].Connection)] = &con
		case Constant.POSTGRESGCP:
			con := sql.DB{}
			env := ""
			if env = e.App.Environment; env == "development" {
				env = "local"
			}
			log.Info().Msg("ENGINE " + Constant.POSTGRESGCP + " start....")
			connectionPool.PostgresGCPConnectionPool(&con, e.GoogleCP.ResourceID, env)
			DbConSql[DbSqlConfigName(e.Databases[i].Connection)] = &con
		case Constant.MONGO:
			con := mongo.Client{}
			log.Info().Msg("ENGINE " + Constant.MONGO + " start....")
			connectionPool.MonggoConnectionPool(&con)
			DbConMonggo[DbMongoConfigName(e.Databases[i].Connection)] = &con
		case Constant.FIRESTORE:
			if ENVIRONMENT_FIRESTORE = e.App.Environment; ENVIRONMENT_FIRESTORE == "local" {
				ENVIRONMENT_FIRESTORE = "development"
			}
			con := firestore.Client{}
			log.Info().Msg("ENGINE " + Constant.FIRESTORE + " start....")
			connectionPool.FirestoreConnectionPool(&con, e.GoogleCP.ResourceID, e.GoogleCP.ProjectID)
			DbConFirestore[DbFirestoreConfigName(e.Databases[i].Connection)] = &con
		}
	}
}

func (app *app) Run(route *chi.Mux) {
	//run with https or http
	if app.Service == "https" {
		app.runWithHttps(route)
	}
	app.runWithHttp(route)
}

func (app *app) runWithHttp(route *chi.Mux) {
	var healthy int32
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	address := app.Host + ":" + app.Port
	httpServer := &http.Server{
		Addr:    address,
		Handler: route,
	}
	go func() {
		<-quit
		log.Info().Msg("Server is shutting down...")
		atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		httpServer.SetKeepAlivesEnabled(false)
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatal().Msgf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	log.Info().Msgf("Http Service running on %s %s", address, " .....")
	atomic.StoreInt32(&healthy, 1)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Msgf("Could not listen on %s: %v\n", address, err)
	}

	<-done
	log.Info().Msg("Server stopped")

}

func (app *app) runWithHttps(route *chi.Mux) {
	_, filename, _, _ := runtime.Caller(1)
	filepathPem := path.Join(path.Dir(filename), "../Infrastructures/certificate/"+app.Pem_key)
	filepathKey := path.Join(path.Dir(filename), "../Infrastructures/certificate/"+app.Certificate)

	var healthy int32
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	address := app.Host + ":" + app.Port
	httpServer := &http.Server{
		Addr:    address,
		Handler: route,
	}
	go func() {
		<-quit
		log.Info().Msg("Server is shutting down...")
		atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		httpServer.SetKeepAlivesEnabled(false)
		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatal().Msgf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	log.Info().Msgf("Http Service running on %s %s", address, " .....")
	atomic.StoreInt32(&healthy, 1)
	if err := httpServer.ListenAndServeTLS(filepathKey, filepathPem); err != nil && err != http.ErrServerClosed {
		log.Fatal().Msgf("Could not listen on %s: %v\n", address, err)
	}

	<-done
	log.Info().Msg("Server stopped")
}

func (d DbSqlConfigName) Get() *sql.DB {
	return DbConSql[d]
}

func (d DbMongoConfigName) Get() *mongo.Client {
	return DbConMonggo[d]
}

func (d DbFirestoreConfigName) Get() *firestore.Client {
	return DbConFirestore[d]
}
