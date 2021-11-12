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
	Id          string `json:"id"`
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

	app.Get("/topics", func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT * FROM dist_kafka_topic")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		defer rows.Close()
		result := Topics{}
		for rows.Next() {
			topic := Topic{}
			if err := rows.Scan(&topic.Id, &topic.Kafka_topic); err != nil {
				return err
			}
			fmt.Println(topic.Kafka_topic)
			result.Topics = append(result.Topics, topic)
		}
		fmt.Printf("%d topic%s received.", len(result.Topics),
			(map[bool]string{true: "", false: "s"})[len(result.Topics) == 1])
		return c.JSON(result)
	})

	log.Fatal(app.Listen(":3000"))

}
