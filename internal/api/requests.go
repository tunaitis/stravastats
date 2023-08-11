package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func request() {

}

type ExchangeCodeToTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func ExchangeCodeToToken(clientId, clientSecret, code string) (string, string, error) {
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	str := &ExchangeCodeToTokenResponse{}
	err = json.Unmarshal(body, str)
	if err != nil {
		return "", "", err
	}

	return str.AccessToken, str.RefreshToken, nil
}
