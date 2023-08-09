package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func request() {

}

func ExchangeAuthCodeToToken(clientId, clientSecret, code string) (string, string, error) {
	u, err := GetTokenUrl()
	if err != nil {
		return "", "", err
	}

	formData := url.Values{
		"client_id":     {clientId},
		"client_secret": {clientSecret},
		"code":          {code},
		"grant_type":    {"authorization_code"},
	}

	resp, err := http.PostForm(u, formData)
	if err != nil {
		return "", "", err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	fmt.Println(string(respBody))

	return "", "", nil
}
