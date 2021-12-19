package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func getPem() (Keys, error) {

	defer timeTrack(time.Now(), "getPem")

	var err error

	oidc := os.Getenv("OIDC_ENDPOINT")
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

	return pemData, nil

}
