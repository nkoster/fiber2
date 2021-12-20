package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func oidc(c *fiber.Ctx) error {

	fmt.Println("oidc:", c.Get("Authorization"))

	access_token := getAccessToken(c.Get("Authorization"))

	if len(access_token) == 0 {
		fmt.Println("No token supplied.")
		// no token supplied logic, WIP
	}

	fmt.Println(access_token)
	// check if access_token is valid via /introspect
	if body := validateAccessToken(access_token); body == "" {
		fmt.Println("Invalid token")
	} else {
		fmt.Println(body)
		// Parse json in body
		// Test if token is active
		c.Next()
	}

	return nil

}
