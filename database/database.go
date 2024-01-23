package database

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB holds the connection to the database.
var DB *gorm.DB

// databaseType is a struct representing the configuration for database connection.
type databaseType struct {
	Hostname string
	Port     uint
	Username string
	Password string
	Database string
	Provider string
}

// dsn returns the Data Source Name (DSN) based on the configured provider.
func (db *databaseType) dsn() string {
	if db.Provider == "mysql" {
		return db.mysqlDSN()
	}
	return db.postgresDSN()
}

// mysqlDSN returns the MySQL-specific DSN.
func (db *databaseType) mysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.Username, db.Password, db.Hostname, db.Port, db.Database)
}

// postgresDSN returns the PostgreSQL-specific DSN.
func (db *databaseType) postgresDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", db.Hostname, db.Username, db.Password, db.Database, db.Port)
}

// connect establishes a connection to the database based on the configured provider.
func (d *databaseType) connect() (*gorm.DB, error) {
	var database *gorm.DB
	if d.Provider == "mysql" {
		var db, err = gorm.Open(mysql.Open(d.dsn()), &gorm.Config{
			SkipDefaultTransaction: true,
		})

		if err != nil {
			return &gorm.DB{}, err
		}
		database = db
	} else {
		var db, err = gorm.Open(postgres.Open(d.dsn()), &gorm.Config{
			SkipDefaultTransaction: true,
		})

		if err != nil {
			return &gorm.DB{}, err
		}
		database = db
	}

	return database, nil
}

// connectionPool sets up and returns the connection pool configurations.
func (d *databaseType) connectionPool(idle int, open int, time time.Duration) *sql.DB {
	conn, _ := DB.DB()
	conn.SetMaxIdleConns(idle)
	conn.SetMaxOpenConns(open)
	conn.SetConnMaxLifetime(time)
	return conn
}

// dbConfig creates and returns a databaseType configuration.
func dbConfig(hostname string, port uint, username string, password string, database string, provider string) databaseType {
	return databaseType{
		Hostname: hostname,
		Port:     port,
		Username: username,
		Password: password,
		Database: database,
		Provider: provider,
	}
}
