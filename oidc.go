package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func oidc(c *fiber.Ctx) error {

	accessToken := getAccessToken(c.Get("Authorization"))

	if len(accessToken) == 0 {
		fmt.Println("oidc: No token supplied.")
		return c.Status(401).SendString("Please login first.")
	}

	// check if accessToken is valid via /introspect.
	if body := validateAccessToken(accessToken); body == "" {
		fmt.Println("oidc: Access token introspect failed.")
		return c.Status(401).SendString("Please login first.")
	} else {
		// Test if token is active.
		var tokenState TokenState
		if err := json.Unmarshal([]byte(body), &tokenState); err != nil {
			fmt.Println("oidc: Cannot unmarshal JSON.")
			return c.Status(500).SendString("Cannot unmarshal JSON.")
		}
		// Verify if scope is correct, and if the token is active.
		if strings.Contains(tokenState.Scope, os.Getenv("OIDC_SCOPE")) {
			if tokenState.Active {
				fmt.Printf("oidc: Allow %s, %s\n", accessToken, tokenState.Scope)
				return c.Next()
			} else {
				fmt.Println("oidc: Token was not active.")
				return c.Status(401).SendString("Please login first.")
			}
		} else {
			fmt.Printf("oidc: Scope %s not allowed\n", tokenState.Scope)
			return c.Status(401).SendString("Please login first.")
		}
	}

}
