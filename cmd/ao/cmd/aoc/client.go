package aoc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL       *url.URL
	sessionCookie string
	httpClient    *http.Client
}

type Member struct {
	Name               string `json:"name"`
	LocalScore         int    `json:"local_score"`
	Stars              int    `json:"stars"`
	GlobalScore        int    `json:"global_score"`
	Id                 int    `json:"id"`
	LastStarTs64       int    `json:"last_star_ts"`
	CompletionDayLevel map[string]map[string]struct {
		StarIndex int   `json:"star_index"`
		GetStarTs int64 `json:"get_star_ts"`
	} `json:"completion_day_level"`
}

type Leaderboard struct {
	OwnerId int               `json:"owner_id"`
	Event   string            `json:"event"`
	Members map[string]Member `json:"members"`
}

func NewClient(sessionCookie string) *Client {
	baseUrl, _ := url.Parse("https://adventofcode.com")
	return &Client{
		BaseURL:       baseUrl,
		sessionCookie: sessionCookie,
		httpClient:    http.DefaultClient,
	}
}

func (c *Client) GetPrivateLeaderboard(year int, leaderboardId string) (*Leaderboard, error) {
	req, err := c.newRequest(
		"GET",
		fmt.Sprintf("/%d/leaderboard/private/view/%s.json", year, leaderboardId),
		nil,
	)
	if err != nil {
		return nil, err
	}

	leaderboard := &Leaderboard{}
	_, err = c.do(req, leaderboard)
	if err != nil {
		return nil, err
	}

	return leaderboard, err
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: c.sessionCookie,
	})
	return req, nil
}

func (c *Client) do(req *http.Request, responseBody interface{}) (*http.Response, error) {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(responseBody)
	return res, err
}
