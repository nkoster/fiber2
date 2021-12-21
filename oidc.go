package main

import (
	"encoding/json"
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
		var token_state TokenState
		json.Unmarshal([]byte(body), &token_state)
		fmt.Println("STATE", token_state)
		if token_state.Active {
			fmt.Println("ACTIVE")
			// get SSO context
			ssoContext := getSsoContext(access_token)
			fmt.Println(ssoContext)
			// Verify ssoContext JWT
			if len(ssoContext) > 1 {
				if verifySsoContext(ssoContext) {
					return c.Next()
				} else {
					fmt.Println("DENIED")
					return c.Status(401).SendString("Please login first.")
				}
			}
			c.Next()
		} else {
			fmt.Println("DENIED")
			return c.Status(401).SendString("Please login first.")
		}
		// c.Next()
	}

	return nil

}
