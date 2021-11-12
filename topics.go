package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func topics(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT * FROM dist_kafka_topic")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := Topics{}
	id := ""
	// this is basically a dummy because we don't need the id,
	// but we need to pass something to rows.Scan(), like &id
	for rows.Next() {
		topic := Topic{}
		if err := rows.Scan(&id, &topic.Kafka_topic); err != nil {
			return err
		}
		fmt.Println(id, topic.Kafka_topic)
		result.Topics = append(result.Topics, topic)
	}
	fmt.Printf("%d topic%s received.", len(result.Topics),
		(map[bool]string{true: "", false: "s"})[len(result.Topics) == 1],
	)
	return c.JSON(result)
}
