package api

import (
	"net/url"
	"stravastats/internal/model"
	"strconv"
	"time"
)

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

func GetActivities(from time.Time) ([]model.Activity, error) {

	perPage := 200

	getPage := func(page int) ([]model.Activity, error) {
		query := url.Values{
			"per_page": {strconv.Itoa(perPage)},
			"after":    {strconv.FormatInt(from.Unix(), 10)},
			"page":     {strconv.Itoa(page)},
		}

		activities, err := Request[[]model.Activity]("athlete/activities", query)
		if err != nil {
			return nil, err
		}

		return activities, nil
	}

	var activities []model.Activity

	page := 1

	result, err := getPage(page)
	if err != nil {
		return nil, err
	}

	activities = append(activities, result...)

	for len(result) == 200 || err != nil {
		page += 1
		result, err = getPage(page)
		if err != nil {
			return nil, err
		}

		activities = append(activities, result...)
	}

	return activities, nil
}
