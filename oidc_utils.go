package main

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	// jwt "github.com/dgrijalva/jwt-go"
)

func getAccessToken(s string) string {

	data := strings.Split(s, " ")

	if len(data) > 1 {
		return data[1]
	}

	return ""

}

func validateAccessToken(token string) string {

	client := &http.Client{}
	URL := os.Getenv("OIDC_INTROSPECT")
	v := url.Values{}

	v.Set("token", token)

	req, err := http.NewRequest("POST", URL, strings.NewReader(v.Encode()))

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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

type Organization struct {
	PvEntityId int `json:"pv_entity_id"`
}

type Sso struct {
	Organization Organization
}
