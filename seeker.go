package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func seeker(c *fiber.Ctx) error {
	payload := struct {
		QueryKafkaTopic      string `json:"queryKafkaTopic"`
		QueryIdentifierType  string `json:"queryIdentifierType"`
		QueryIdentifierValue string `json:"queryIdentifierValue"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}
	rows, err := db.Query(`SELECT /*greetz_from_golang*/
		* from func_identifier(in_identifier_value => $3,
		in_identifier_type => $2,
		in_kafka_topic => $1,
		in_kafka_offset => null,
		in_kafka_partition => null)
	`, payload.QueryKafkaTopic, payload.QueryIdentifierType, payload.QueryIdentifierValue)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := Messages{}
	for rows.Next() {
		message := Message{}
		if err := rows.Scan(
			&message.QueryKafkaTopic,
			&message.QueryKafkaPartition,
			&message.QueryKafkaOffset,
			&message.QueryIdentifierType,
			&message.QueryIdentifierValue); err != nil {
			fmt.Println(err)
			return err
		}
		result.Messages = append(result.Messages, message)
	}
	fmt.Printf("%d message%s received.", len(result.Messages),
		(map[bool]string{true: "", false: "s"})[len(result.Messages) == 1],
	)
	return c.JSON(result)
}
