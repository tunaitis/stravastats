package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"stravastats/internal/config"
	"time"
)

var baseApiUrl = "https://www.strava.com/api/v3/"
var ErrUnauthorized = &http.ProtocolError{"Unauthorized"}

type Client struct {
	client       *http.Client
	accessToken  string
	refreshToken string
}

func NewClient(accessToken string, refreshToken string) *Client {
	httpClient := &http.Client{}

	return &Client{
		client:       httpClient,
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}
}

// https://www.strava.com/api/v3/athlete/activities
func Request(resource string) error {
	tokens, err := config.ReadTokens()
	if err != nil {
		return err
	}

	cfg, err := config.ReadConfig()
	if err != nil {
		return err
	}

	accessToken := tokens.AccessToken
	expires := time.Unix(int64(tokens.ExpiresAt), 0).UTC()

	if expires.Before(time.Now().UTC()) {
		refreshed, err := RefreshAccessToken(cfg.Api.ClientId, cfg.Api.ClientSecret, tokens.RefreshToken)
		if err != nil {
			return err
		}

		err = config.SaveTokens(refreshed.AccessToken, refreshed.RefreshToken, refreshed.ExpiresAt)
		if err != nil {
			return err
		}
	}

	resp, err := httpGet(accessToken, resource, "")
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return nil
}

func httpGet(accessToken string, resource string, query string) (*http.Response, error) {
	u, err := url.Parse(baseApiUrl + resource)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, ErrUnauthorized
	}

	return resp, nil
}

func (c *Client) Do(resource string) (*http.Response, error) {
	u, err := url.Parse(baseApiUrl + resource)
	if err != nil {
		return nil, err
	}

	fmt.Println(u.String())

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.accessToken)

	fmt.Println(c.accessToken)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) ListActivities() error {

	resp, err := c.Do("athlete/activities")
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	return nil
}
