package persistence

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	dbUser                 = mustGetenv("DB_USER")                  // e.g. 'my-db-user'
	dbPwd                  = mustGetenv("DB_PASS")                  // e.g. 'my-db-password'
	instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
	dbName                 = mustGetenv("DB_NAME")                  // e.g. 'my-database'
)

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Panic("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func ConnectDB() (*gorm.DB, error) {

	/*socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}*/

	//dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)
	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=localhost port=5432 sslmode=disable", dbUser, dbPwd, dbName)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}

	return db, nil
}
