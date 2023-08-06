package api

import "net/url"

const baseUrl = "https://www.strava.com"

func GetAuthUrl(clientId string) (string, error) {

	u, err := url.Parse(baseUrl + "/oauth/authorize")
	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Add("client_id", clientId)
	q.Add("response_type", "code")
	q.Add("redirect_uri", "http://localhost:42001/exchange_token")
	q.Add("approval_prompt", "force")
	q.Add("scope", "read")

	u.RawQuery = q.Encode()

	return u.String(), nil
}
