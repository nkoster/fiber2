package main

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
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

	N, err := base64.StdEncoding.DecodeString(pemData.Keys[0].N)

	if err != nil {
		fmt.Println(N, err)
	}

	NN := new(big.Int)
	NN.SetBytes(N)

	E, err := base64.StdEncoding.DecodeString(pemData.Keys[0].E)
	if err != nil {
		fmt.Println(E, err)
	}

	// EE := new(big.Int)
	// EEE := EE.SetBytes(E).SetInt64()

	// EEE = EE.Int64()
	// fmt.Println("pemData ALG", pemData.Keys[0].ALG)
	// fmt.Println("pemData E", pemData.Keys[0].E)
	// fmt.Println("pemData KID", pemData.Keys[0].KID)
	// fmt.Println("pemData KTY", pemData.Keys[0].KTY)
	// fmt.Println("pemData N", pemData.Keys[0].N)
	// fmt.Println("pemData USE", pemData.Keys[0].USE)

	// pem := rsa.PublicKey{N: NN, E: int(EEE)}

	fmt.Println("pemmie", pem)

	return pem, nil

}
