package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func oidc(c *fiber.Ctx) error {

	access_token := getAccessToken(c.Get("Authorization"))

	if len(access_token) == 0 {
		// no token supplied
		fmt.Println("oidc: No token supplied.")
		return c.Status(401).SendString("Please login first.")
	}

	// check if access_token is valid via /introspect
	if body := validateAccessToken(access_token); body == "" {
		fmt.Println("oidc: Access token introspect failed.")
		return c.Status(401).SendString("Please login first.")
	} else {
		// Parse json in body
		// Test if token is active
		var token_state TokenState
		json.Unmarshal([]byte(body), &token_state)
		// Verify if token is active and if the scope is correct
		if token_state.Scope == "openid" {
			fmt.Println("oidc: Allowed:", token_state.Scope, access_token)
			if token_state.Active {
				return c.Next()
			} else {
				fmt.Println("oidc: Token was not active.")
				return c.Status(401).SendString("Please login first.")
			}
		} else {
			fmt.Printf("oidc: Scope %s not allowed\n", token_state.Scope)
			return c.Status(401).SendString("Please login first.")
		}
	}

}
