package main

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"
)

func getPem() string {

	defer timeTrack(time.Now(), "getPem")

	oidc := os.Getenv("OIDC_CERTS")
	resp, err := http.Get(oidc)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	sb := string(body)

	var pemData Keys

	json.Unmarshal([]byte(sb), &pemData)

	nb, err := base64.RawURLEncoding.DecodeString(pemData.Keys[0].N)

	if err != nil {
		fmt.Println(err)
	}

	e := 65537

	pk := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nb),
		E: e,
	}

	der, err := x509.MarshalPKIXPublicKey(pk)
	if err != nil {
		log.Fatal(err)
	}

	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: der,
	}

	var out bytes.Buffer
	pem.Encode(&out, block)
	return out.String()

}
