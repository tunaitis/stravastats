package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"stravastats/internal/model"
)

func RefreshAccessToken(clientId, clientSecret, refreshToken string) (*model.Tokens, error) {
	u, err := GetTokenUrl()
	if err != nil {
		return nil, err
	}

	formData := url.Values{
		"client_id":     {clientId},
		"client_secret": {clientSecret},
		"refresh_token": {refreshToken},
		"grant_type":    {"refresh_token"},
	}

	resp, err := http.PostForm(u, formData)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	tokens := &model.Tokens{}
	err = json.Unmarshal(body, tokens)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func ExchangeCodeToAccessToken(clientId, clientSecret, code string) (*model.Tokens, error) {
	u, err := GetTokenUrl()
	if err != nil {
		return nil, err
	}

	formData := url.Values{
		"client_id":     {clientId},
		"client_secret": {clientSecret},
		"code":          {code},
		"grant_type":    {"authorization_code"},
	}

	resp, err := http.PostForm(u, formData)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	tokens := &model.Tokens{}
	err = json.Unmarshal(body, tokens)
	if err != nil {
		return nil, err
	}

	return tokens, nil
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
