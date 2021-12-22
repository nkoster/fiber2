package main

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
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

	decN, err := base64.RawURLEncoding.DecodeString(pemData.Keys[0].N)

	if err != nil {
		fmt.Println(err)
	}

	n := big.NewInt(0)

	n.SetBytes(decN)

	decE, err := base64.RawURLEncoding.DecodeString(pemData.Keys[0].E)

	if err != nil {
		fmt.Println(err)
	}

	var eBytes []byte

	if len(decE) < 8 {
		eBytes = make([]byte, 8-len(decE), 8)
		eBytes = append(eBytes, decE...)
	} else {
		eBytes = decE
	}

	eReader := bytes.NewReader(eBytes)

	var e uint64

	err = binary.Read(eReader, binary.BigEndian, &e)

	if err != nil {
		fmt.Println(err)
	}

	return rsa.PublicKey{N: n, E: int(e)}, nil

}
