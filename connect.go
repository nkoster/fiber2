package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

func Connect() error {

	var err error

	host := os.Getenv("PG_HOST")
	port, _ := strconv.Atoi(os.Getenv("PG_PORT"))
	sslcert := os.Getenv("PG_CERT_PATH")
	sslkey := os.Getenv("PG_KEY_PATH")
	sslrootcert := os.Getenv("PG_CA_PATH")
	sslmode := os.Getenv("PG_SSL_MODE")
	user := os.Getenv("PG_USER")
	database := os.Getenv("PG_DATABASE")

	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s sslmode=%s sslcert=%s sslkey=%s sslrootcert=%s database=%s",
		host, port, user, sslmode, sslcert, sslkey, sslrootcert, database))
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}
