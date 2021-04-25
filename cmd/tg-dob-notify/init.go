package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func initDatabaseConnection(driver string, user string, password string, dbName string, sslMode string) (*sqlx.DB, error) {
	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	connOptions := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		user,
		password,
		dbName,
		sslMode,
	)
	db, err := sqlx.Connect(driver, connOptions)
	if err != nil {
		return nil, err
	}

	return db, nil
}
