package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func getPem() error {

	var err error

	resp, err := http.Get("https://oic.dev.portavita.nl/certs")

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	sb := string(body)
	log.Println(sb)

	if err != nil {
		log.Fatalln(err)
	}

	return nil

}
