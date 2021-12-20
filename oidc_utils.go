package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func getAccessToken(s string) string {

	data := strings.Split(s, " ")

	if len(data) > 0 {
		return data[1]
	}

	return ""

}

func validateAccessToken(token string) string {

	client := &http.Client{}
	URL := os.Getenv("OIDC_INTRSPECT")
	v := url.Values{}

	v.Set("token", token)

	req, err := http.NewRequest("POST", URL, strings.NewReader(v.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		fmt.Println(err)
	}

	req.SetBasicAuth(os.Getenv("OIDC_API_USER"), os.Getenv("OIDC_API_PASSWORD"))

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	return string(bodyText)

}
