package main

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"
)

func getPem() (rsa.PublicKey, error) {

	defer timeTrack(time.Now(), "getPem")

	var err error

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

	if err != nil {
		log.Fatalln(err)
	}

	N := big.NewInt(0)
	N.SetBytes([]byte(pemData.Keys[0].N))

	E, err := strconv.Atoi(pemData.Keys[0].E)
	if err != nil {
		fmt.Println(err)
	}

	pem := rsa.PublicKey{N: N, E: E}

	fmt.Println("pemmie", pem)

	return pem, nil

}
