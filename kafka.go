package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/gofiber/fiber/v2"
)

func kafka(c *fiber.Ctx) error {

	defer timeTrack(time.Now(), "kafka")

	payload := struct {
		Topic     string `json:"topic"`
		Partition string `json:"partition"`
		Offset    string `json:"offset"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		fmt.Println(err.Error())
		return c.Status(500).SendString(err.Error())
	}

	c.Set("Content-type", "application/json; charset=utf-8")

	out, err := exec.Command("kafkacat", "-C", "-b", os.Getenv("KAFKA_HOST"), "-t",
		payload.Topic, "-p", payload.Partition, "-o", payload.Offset, "-c", "1", "-e", "-q").Output()
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString(err.Error())
	}

	fmt.Println("kafka:", payload.Topic, payload.Partition, payload.Offset)
	return c.Send(out)

}
