package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Topic struct {
	Kafka_topic string `json:"kafka_topic"`
}

type Topics struct {
	Topics []Topic `json:"topics"`
}

type Message struct {
	QueryKafkaTopic      string `json:"kafka_topic"`
	QueryKafkaPartition  string `json:"kafka_partition"`
	QueryKafkaOffset     string `json:"kafka_offset"`
	QueryIdentifierType  string `json:"identifier_type"`
	QueryIdentifierValue string `json:"identifier_value"`
}

type Messages struct {
	Messages []Message `json:"messages"`
}

func main() {

	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err = Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Post("/kafka", kafka)
	app.Post("/seeker", seeker)
	app.Post("/topics", topics)

	log.Fatal(app.Listen(":3030"))

}
