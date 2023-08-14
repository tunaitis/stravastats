package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"fmt"
	"context"
)

func request() {

}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func RefreshAccessToken(clientId, clientSecret, refreshToken string) (string, string, error) {
	u, err := GetTokenUrl()
	if err != nil {
		return "", "", err
	}

	formData := url.Values{
		"client_id":     {clientId},
		"client_secret": {clientSecret},
		"refresh_token":          {refreshToken},
		"grant_type":    {"refresh_token"},
	}

	resp, err := http.PostForm(u, formData)
	if err != nil {
		return "", "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	str := &TokenResponse{}
	err = json.Unmarshal(body, str)
	if err != nil {
		return "", "", err
	}

	return str.AccessToken, str.RefreshToken, nil
}

func ExchangeCodeToAccessToken(clientId, clientSecret, code string) (string, string, error) {
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

	str := &TokenResponse{}
	err = json.Unmarshal(body, str)
	if err != nil {
		return "", "", err
	}

	return str.AccessToken, str.RefreshToken, nil
}

func WaitForAuthorizationCode() (string, error) {

	m := http.NewServeMux()
	s := &http.Server{Addr: ":42001", Handler: m}

	var code = ""

	m.HandleFunc("/exchange_token", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, "The authorization code has been received. You can now close this window.")

		code = r.URL.Query().Get("code")

		go func() {
			s.Shutdown(context.Background())
		}()
	})

	s.ListenAndServe()

	return code, nil
}

