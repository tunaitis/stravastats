package api

import "net/url"

const baseAuthUrl = "https://www.strava.com/oauth/"

func GetAuthUrl(clientId string) (string, error) {

	u, err := url.Parse(baseAuthUrl + "authorize")
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Add("client_id", clientId)
	q.Add("response_type", "code")
	q.Add("redirect_uri", "http://localhost:42001/exchange_token")
	q.Add("approval_prompt", "force")
	q.Add("scope", "read,activity:read_all")

	u.RawQuery = q.Encode()

	return u.String(), nil
}

func GetTokenUrl() (string, error) {
	u, err := url.Parse(baseAuthUrl + "token")
	if err != nil {
		return "", err
	}

	q := u.Query()

	u.RawQuery = q.Encode()

	return u.String(), nil
}
