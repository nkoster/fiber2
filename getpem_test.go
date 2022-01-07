package main

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestGetPem(t *testing.T) {

	PEM := `-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuSAabsm0HbxycbCspKaB
e33QAe5kgmLylTQJUSIYhVI+X8NA6HtxMjiMKuLA3G3GEDEXiWqL/v1rtMCJoWt4
a4F9f29XqgjH3vAVtP4qSTYl8xHCshEjwnB6HPw/LtNgkiad3aI8HRHu/w6FM5Ve
M76DeowAk7nO/g2znljhpUXkJhehWseCQbe17teGL+Q9TW39tSh7rNSx9rPhycdW
U3UWio2gTydfQS+i6tvq0Tgg3ruFO0ECTS/AC6EiII4eUTL+DDKb1b+CkbtOE3Pk
WITkdBiKT5eb9AQZjf3lvNyrc2oLQDijOX1T7J/f1FPDmPzRgOeYL7hLU4OxtO9+
hwIDAQAB
-----END RSA PUBLIC KEY-----
`

	if err := godotenv.Load(); err != nil {
		t.Error("No .env file in current directory.")
	}

	result := getPem()

	if result != PEM {
		t.Error("PEM file has unexpected layout.")
	}
}
