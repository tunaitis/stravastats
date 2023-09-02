package api

import (
	"fmt"
	"log/slog"
	"net/url"
	"stravastats/internal/model"
	"strconv"
	"time"

	"github.com/schollz/progressbar/v3"
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
	pb := progressbar.Default(-1, "Downloading activities...")

	perPage := 200

	getPage := func(page int) ([]model.Activity, error) {
		query := url.Values{
			"per_page": {strconv.Itoa(perPage)},
			"after":    {strconv.FormatInt(from.Unix(), 10)},
			"page":     {strconv.Itoa(page)},
		}

		slog.Debug("Downloading activities", slog.Time("after", from), slog.Int("page", page))

		activities, err := Request[[]model.Activity]("athlete/activities", query)
		if err != nil {
			return nil, err
		}

		slog.Debug("Finished", slog.Int("activities", len(activities)))

		pb.Add(len(activities))

		return activities, nil
	}

	var activities []model.Activity

	page := 1

	result, err := getPage(page)
	if err != nil {
		return nil, err
	}

	activities = append(activities, result...)

	for len(result) == perPage || err != nil {
		page += 1
		result, err = getPage(page)
		if err != nil {
			return nil, err
		}

		activities = append(activities, result...)
	}

	// hide the progress bar and move the cursor one line up
	pb.Close()
	fmt.Printf("\033[%dA", 1)

	return activities, nil
}
