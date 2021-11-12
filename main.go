package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host        = "db.fhirstation.net"
	port        = 50505
	sslcert     = "./client_postgres.crt"
	sslkey      = "./client_postgres.key"
	sslrootcert = "./root.crt"
	sslmode     = "verify-ca"
	user        = "postgres"
)

type Topic struct {
	Kafka_topic string `json:"kafka_topic"`
}

type Topics struct {
	Topics []Topic `json:"topics"`
}

func Connect() error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s sslmode=%s sslcert=%s sslkey=%s sslrootcert=%s", host, port, user, sslmode, sslcert, sslkey, sslrootcert))
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func main() {

	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/topics", topics)

	log.Fatal(app.Listen(":3000"))

}
