package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"os"

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

type Key struct {
	KTY string `json:"kty"`
	E   string `json:"e"`
	USE string `json:"use"`
	KID string `json:"kid"`
	ALG string `json:"alg"`
	N   string `json:"n"`
}

type Keys struct {
	Keys []Key `json:"keys"`
}

type TokenState struct {
	Active bool    `json:"active"`
	Exp    big.Int `json:"exp"`
	Scope  string  `json:"scope"`
}

var pemFile Keys

func main() {

	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err = Connect(); err != nil {
		log.Fatal(err)
	}

	pemFile = getPem()

	// if pem, err = getPem(); err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println(pemFile)

	app := fiber.New()

	app.Static("/", "./ui")

	if os.Getenv("DEV_MODE") == "true" {
		app.Use(cors.New())
		fmt.Print("MODE=DEV ")
	} else {
		fmt.Print("MODE=PROD ")
	}

	if os.Getenv("USE_AUTH") == "false" {
		fmt.Println("AUTH=false")
	} else {
		app.Use(oidc)
		fmt.Println("AUTH=true")
	}

	app.Post("/kafka", kafka)
	app.Post("/seeker", seeker)
	app.Post("/topics", topics)

	log.Fatal(app.Listen(":3030"))

}
