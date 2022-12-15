package Config

import (
	"EmployeeService/Constant"
	"EmployeeService/Library/SecretManager"
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"net"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/firestore"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/api/option"
)

type connectionPool interface {
	PostgresConnectionPool(Connection interface{})
	PostgresGCPConnectionPool(Connection interface{}, resourceId string, localConfig string)
	MssqlConnectionPool(Connection interface{})
	MonggoConnectionPool(Connection interface{})
	FirestoreConnectionPool(Connection interface{}, resourceId string, projectId string)
}

func (env *database) MssqlConnectionPool(Connection interface{}) {
	Con := Connection.(*sql.DB)
	connection_string := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;", env.Host, env.Username, env.Password, env.Port)
	Connection, err := sql.Open(Constant.MSSQL, connection_string)
	if err != nil {
		//panic err
		panic(err.Error())
	}
	*Con = *Connection.(*sql.DB)
	err = Con.Ping()
	if err != nil {
		log.Print(err.Error())
		panic(err.Error())
	}
}

func (env *database) PostgresConnectionPool(Connection interface{}) {
	Con := Connection.(*sql.DB)
	var buffer bytes.Buffer
	buffer.WriteString("postgres://")
	buffer.WriteString(env.Username + ":" + env.Password)
	buffer.WriteString("@")
	buffer.WriteString(env.Host + ":" + env.Port + "/")
	buffer.WriteString(env.Name)
	buffer.WriteString("?sslmode=disable")
	connection_string := buffer.String()
	Connection, err := sql.Open(Constant.POSTGRES, connection_string)
	if err != nil {
		//panic err
		panic(err.Error())
	}
	Connection.(*sql.DB).SetMaxOpenConns(env.Maximum_connection)
	*Con = *Connection.(*sql.DB)
	err = Con.Ping()
	if err != nil {
		log.Fatal().Msgf("Couldn't connect to the Postgres %v", err)
		panic(err.Error())
	} else {
		log.Info().Msg("Postgres Connected!")
	}
}

func (env *database) PostgresGCPConnectionPool(Connection interface{}, resourceId string, localConfig string) {
	Con := Connection.(*sql.DB)
	var (
		dbUser                      = env.Username
		dbIAMUser                   = env.DbIAMUser
		dbPwd                       = env.Password
		dbName                      = env.Name
		dbPort                      = env.Port
		instanceConnectionName      = env.Host
		PrivateIP                   = env.PrivateIP
		isPrivate              bool = env.IsPrivateIP //if true; use private ip cloud and set Host
	)
	if localConfig == "local" {
		var err error
		dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s",
			instanceConnectionName, dbUser, dbPwd, dbPort, dbName)

		Connection, err = sql.Open("pgx", dbURI)
		if err != nil {
			panic(err.Error())
		}
		Connection.(*sql.DB).SetMaxOpenConns(env.Maximum_connection)
		*Con = *Connection.(*sql.DB)
		err = Con.Ping()
		if err != nil {
			log.Fatal().Msgf("Couldn't connect to the Postgres %v", err)
			panic(err.Error())
		}
		log.Info().Msg("CONNECTED with Postgres CloudSql Local Proxy... ")
	} else {

		dsn := fmt.Sprintf("user=%s password=%s database=%s", dbUser, dbPwd, dbName)
		config, err := pgx.ParseConfig(dsn)
		if err != nil {
			panic(err.Error())
		}
		secret, err := SecretManager.GetSecret(resourceId, "cloudsql-secret")
		if err != nil {
			panic(err.Error())
		}
		option := []cloudsqlconn.Option{
			cloudsqlconn.WithCredentialsJSON(secret),
		}
		config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
			if isPrivate {
				log.Info().Msg("trying to connect with Postgres CloudSql private ...")
				os.Setenv("PRIVATE_IP", PrivateIP)
				println(os.Getenv("PRIVATE_IP"))
				d, err := cloudsqlconn.NewDialer(
					ctx,
					cloudsqlconn.WithDefaultDialOptions(cloudsqlconn.WithPrivateIP()),
					cloudsqlconn.WithCredentialsJSON(secret),
				)
				if err != nil {
					return nil, err
				}
				return d.Dial(ctx, instanceConnectionName)
			}
			if dbIAMUser != "" {
				log.Info().Msg("trying to connect with Postgres CloudSql iam_auth ...")
				// [START cloud_sql_postgres_databasesql_auto_iam_authn]

				d, err := cloudsqlconn.NewDialer(
					ctx,
					cloudsqlconn.WithIAMAuthN(),
					cloudsqlconn.WithCredentialsJSON(secret))
				if err != nil {
					return nil, err
				}
				return d.Dial(ctx, instanceConnectionName)
				// [END cloud_sql_postgres_databasesql_auto_iam_authn]
			}
			log.Info().Msg("trying to connect with Postgres CloudSql ...")
			// Use the Cloud SQL connector to handle connecting to the instance.
			// This approach does *NOT* require the Cloud SQL proxy.
			d, err := cloudsqlconn.NewDialer(ctx, option...)
			if err != nil {
				return nil, err
			}
			return d.Dial(ctx, instanceConnectionName)
		}
		dbURI := stdlib.RegisterConnConfig(config)
		Connection, err = sql.Open("pgx", dbURI)
		if err != nil {
			panic(err.Error())
		}
		Connection.(*sql.DB).SetMaxOpenConns(env.Maximum_connection)
		*Con = *Connection.(*sql.DB)
		err = Con.Ping()
		if err != nil {
			panic(err.Error())
		}
		log.Info().Msg("CONNECTED with Postgres CloudSql.. ")
	}
}

func (env *database) MonggoConnectionPool(Connection interface{}) {
	var buffer bytes.Buffer
	Con := Connection.(*mongo.Client)
	buffer.WriteString("mongodb://")
	if env.Username != "" || env.Password != "" {
		buffer.WriteString(env.Username + ":" + env.Password)
	}
	buffer.WriteString("@" + env.Host + ":" + env.Port)

	buffer.WriteString("/?tls=false&authSource=" + env.Name + "&authMechanism=")
	connectionString := buffer.String()

	clientOptions := options.Client().ApplyURI(connectionString)
	Connection, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic(err.Error())
	}
	*Con = *Connection.(*mongo.Client)
	err = Con.Connect(context.Background())
	if err != nil {
		log.Fatal().Msgf("Couldn't connect to the database", err)
	}
	err = Con.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal().Msgf("Couldn't PING to the database", err)
	} else {
		log.Info().Msg("MongoDb Connected!")
	}
}

func (env *database) FirestoreConnectionPool(Connection interface{}, resourceId string, projectId string) {
	Con := Connection.(*firestore.Client)
	secret, err := SecretManager.GetSecret(resourceId, "firestore-secret")
	if err != nil {
		panic(err.Error())
	}
	ctx := context.Background()

	Connection, err = firestore.NewClient(ctx, projectId, option.WithCredentialsJSON(secret))
	if err != nil {
		log.Fatal().Msgf("Couldn't connect to the firestore, error: ", err)
	}

	*Con = *Connection.(*firestore.Client)

	log.Info().Msg("CONNECTED with Firestore...")
}
