package main

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestConnect(t *testing.T) {

	if err := godotenv.Load(); err != nil {
		t.Error("No .env file in current directory.")
	}

	if Connect() != nil {
		t.Error("DB connection failed.")
	}
}
