package main

import "testing"

func TestConnect(t *testing.T) {
	if Connect() != nil {
		t.Error("DB connection failed")
	}
}
