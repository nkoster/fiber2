package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func oidc(c *fiber.Ctx) error {

	fmt.Println("oidc:", c.Get("Authorization"))

	c.Next()

	return nil

}
