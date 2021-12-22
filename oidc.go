package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func oidc(c *fiber.Ctx) error {

	access_token := getAccessToken(c.Get("Authorization"))

	if len(access_token) == 0 {
		// no token supplied
		return c.Status(401).SendString("Please login first.")
	}

	// check if access_token is valid via /introspect
	if body := validateAccessToken(access_token); body == "" {
		return c.Status(401).SendString("Please login first.")
	} else {
		// Parse json in body
		// Test if token is active
		var token_state TokenState
		json.Unmarshal([]byte(body), &token_state)
		if token_state.Active {
			// get SSO context
			ssoContext := getSsoContext(access_token)
			// Verify ssoContext JWT
			if len(ssoContext) > 1 {
				if verifySsoContext(ssoContext, pemFile) {
					// All good, allow access
					// To do: verify scope
					return c.Next()
				} else {
					return c.Status(401).SendString("Please login first.")
				}
			}
			c.Next()
		} else {
			return c.Status(401).SendString("Please login first.")
		}
	}

	return nil
}
