package main

import (
	"fmt"
	"io/ioutil"
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

/*
   data: `token=${token}`,
   method: 'post',
   baseURL: process.env.OID_BASE_URL,
   auth: {
     username: process.env.API_USERNAME,
     password: process.env.OPENID_PASSWORD
   }
*/

func validateAccessToken(token string) error {

	if err := call(os.Getenv("OIDC_INTROSPECT"), token); err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}

func call(u, token string) error {

	client := &http.Client{
		// Timeout: time.Second * 10,
	}

	params := url.Values{}
	params.Add("token", token)

	fmt.Println("Introspect token", token, u)

	req, err := http.NewRequest("PostForm", u, strings.NewReader(params.Encode()))

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("got error %s", err.Error())
	}

	req.SetBasicAuth("fhirstation-kafkasearch-backend", "HaQ3ew8aXN2kzZeS0Unwo3inE3chSB")

	response, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("got error %s", err.Error())
	}

	bodyText, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(bodyText), err)

	defer response.Body.Close()

	return nil

}
