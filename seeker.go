package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func seeker(c *fiber.Ctx) error {

	var err error

	payload := struct {
		QueryKafkaTopic      string `json:"queryKafkaTopic"`
		QueryIdentifierType  string `json:"queryIdentifierType"`
		QueryIdentifierValue string `json:"queryIdentifierValue"`
		QueryId              string `json:"queryId"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	sqlCancel := `WITH pids AS (
		/*notthisone*/
		SELECT pid
		FROM   pg_stat_activity
		WHERE  query LIKE '%%/*%s*/%%'
		AND    query NOT LIKE '%%/*notthisone*/%%'
		AND    state='active'
	)
	SELECT pg_cancel_backend(pid) FROM pids
	`
	sqlCancel = fmt.Sprintf(sqlCancel, payload.QueryId)

	cancel, err := db.Query(sqlCancel)
	if err != nil {
		fmt.Println(err)
	}

	defer cancel.Close()

	counter := 0
	for cancel.Next() {
		counter++
	}

	if counter > 0 {
		fmt.Println("Previous query was cancelled.")
	}

	sqlSelect := `SELECT /*%s*/
	* from func_identifier(in_identifier_value => $3,
	in_identifier_type => $2,
	in_kafka_topic => $1,
	in_kafka_offset => null,
	in_kafka_partition => null)
	`
	sqlSelect = fmt.Sprintf(sqlSelect, payload.QueryId)

	rows, err := db.Query(sqlSelect, payload.QueryKafkaTopic, payload.QueryIdentifierType, payload.QueryIdentifierValue)
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

	fmt.Printf("pg: %d message%s received.\n", len(result.Messages),
		(map[bool]string{true: "", false: "s"})[len(result.Messages) == 1],
	)

	return c.JSON(result)

}
