package main

import (
	"fmt"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func kafka(c *fiber.Ctx) error {

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
	}
	return c.Send(out)
}
