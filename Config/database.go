package Config

import (
	"database/sql"

	"cloud.google.com/go/firestore"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
)

type DbSqlConfigName string
type DbMongoConfigName string
type DbFirestoreConfigName string

const (
	// Database Connection Constant name
	DATABASE_MAIN      DbSqlConfigName       = "mainDB"
	DATABASE_MONGO     DbMongoConfigName     = "mainDBMonggo"
	DATABASE_FIRESTORE DbFirestoreConfigName = "mainDBFirestore"
)

//mapping all sql connection
var DbConSql = map[DbSqlConfigName]*sql.DB{
	DATABASE_MAIN: nil,
}

var DbConMonggo = map[DbMongoConfigName]*mongo.Client{
	DATABASE_MONGO: nil,
}
var DbConFirestore = map[DbFirestoreConfigName]*firestore.Client{
	DATABASE_FIRESTORE: nil,
}

var ENVIRONMENT_FIRESTORE = ""
