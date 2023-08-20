package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"stravastats/internal/config"
	"time"
)

var baseApiUrl = "https://www.strava.com/api/v3/"
var ErrUnauthorized = &http.ProtocolError{"Unauthorized"}

func Request[T interface{}](resource string, query url.Values) (T, error) {
	var result T
	tokens, err := config.ReadTokens()
	if err != nil {
		return result, err
	}

	cfg, err := config.ReadConfig()
	if err != nil {
		return result, err
	}

	accessToken := tokens.AccessToken
	expires := time.Unix(int64(tokens.ExpiresAt), 0).UTC()

	if expires.Before(time.Now().UTC()) {
		refreshed, err := RefreshAccessToken(cfg.Api.ClientId, cfg.Api.ClientSecret, tokens.RefreshToken)
		if err != nil {
			return result, err
		}

		accessToken = refreshed.AccessToken

		err = config.SaveTokens(refreshed)
		if err != nil {
			return result, err
		}
	}

	resp, err := httpGet(accessToken, resource, query)
	if err != nil {
		return result, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func httpGet(accessToken string, resource string, query url.Values) (*http.Response, error) {
	u, err := url.Parse(baseApiUrl + resource)
	if err != nil {
		return nil, err
	}


	if query != nil {
		u.RawQuery = query.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer " +accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, ErrUnauthorized
	}

	return resp, nil
}
