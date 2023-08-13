package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var baseApiUrl = "https://www.strava.com/api/v3/"

type Client struct {
	client 			*http.Client
	accessToken 	string
	refreshToken 	string
}

func NewClient(accessToken string, refreshToken string) *Client {
	httpClient := &http.Client{

	}

	return &Client{
		client: httpClient,
		accessToken: accessToken,
		refreshToken: refreshToken,
	}
}

// https://www.strava.com/api/v3/athlete/activities

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

	req.Header.Add("Authorization", "Bearer " + c.accessToken)

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
