package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
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

	/*
		JSON, looks like this: bodyText {"active":true,"exp":1646656120,"scope":"openid"}
		This should verify the scope and return a boolean instead of the JSON string.
		Should be scope "kafkasearch"
	*/
	fmt.Println("bodyText", string(bodyText))

	return string(bodyText)

}

func getSsoContext(token string) string {

	client := &http.Client{}

	URL := os.Getenv("OIDC_SSO_CONTEXT")

	req, err := http.NewRequest("GET", URL, nil)

	req.Header.Add("Authorization", "Bearer "+token)

	if err != nil {
		fmt.Println(err)
	}

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

func verifySsoContext(token string, key string) bool {

	parts := strings.Split(token, ".")

	signingKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(key))

	if err != nil {
		fmt.Println(err)
	}

	err = jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], signingKey)

	if err != nil {
		fmt.Println(err)
		return false
	}

	ssoJson, _ := jwt.DecodeSegment(parts[1])

	var sso Sso

	json.Unmarshal(ssoJson, &sso)

	fmt.Println(string(ssoJson))

	// test if pv_entity_id == 3 == pvt_support1 / pv_entity_id == niels == 17899744)
	return sso.Organization.PvEntityId == 17899744

}
