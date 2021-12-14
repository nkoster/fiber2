package main

import (
	"fmt"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func kafka(c *fiber.Ctx) error {
	c.Set("Content-type", "application/json; charset=utf-8")
	out, err := exec.Command("kafkacat", "-C", "-b", "localhost:9092", "-t", "fhir3.databus.portavita.pvt_medrie_zwolle.patient", "-p", "0", "-o", "989272", "-c", "1", "-e", "-q").Output()
	if err != nil {
		fmt.Println(err)
	}
	return c.Send(out)
}
