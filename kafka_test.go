package main

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func Test_kafka(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := kafka(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("kafka() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
