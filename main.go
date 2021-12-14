package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
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
	QueryKafkaTopic      string `json:"queryKafkaTopic"`
	QueryKafkaPartition  string `json:"queryKafkaPartition"`
	QueryKafkaOffset     string `json:"queryKafkaOffset"`
	QueryIdentifierType  string `json:"queryIdentifierType"`
	QueryIdentifierValue string `json:"queryIdentifierValue"`
}

type Messages struct {
	Messages []Message `json:"messages"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Post("/kafka", kafka)
	app.Post("/seeker", seeker)
	app.Post("/topics", topics)

	log.Fatal(app.Listen(":3030"))

}
