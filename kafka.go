package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/gofiber/fiber/v2"
)

func kafka(c *fiber.Ctx) error {

	defer timeTrack(time.Now(), "seeker")

	payload := struct {
		Topic     string `json:"topic"`
		Partition string `json:"partition"`
		Offset    string `json:"offset"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	c.Set("Content-type", "application/json; charset=utf-8")

	out, err := exec.Command("kafkacat", "-C", "-b", "localhost:9092", "-t",
		payload.Topic, "-p", payload.Partition, "-o", payload.Offset, "-c", "1", "-e", "-q").Output()
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString(err.Error())
	}

	fmt.Println("kafka:", payload.Topic, payload.Partition, payload.Offset)
	return c.Send(out)

}
