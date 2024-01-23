package database

import (
	"github.com/dlyycious/gopos/config"
)

func Run() {

	database := dbConfig(config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASS, config.DB_NAME, config.DB_PROV)
	db, err := database.connect()

	if err != nil {
		panic(err)
	}

	DB = db
	autoMigrate(db)
	database.connectionPool(config.DB_CONN_POOL_IDLE, config.DB_CONN_POOL_OPEN, config.DB_CONN_POOL_TIME)
}
