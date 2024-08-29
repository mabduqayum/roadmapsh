package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mabduqayum/roadmapsh/02_github_user_activity/config"
)

type Client struct {
	config *config.Config
}

type Event struct {
	Type      string    `json:"type"`
	Repo      RepoInfo  `json:"repo"`
	Payload   Payload   `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}

type RepoInfo struct {
	Name string `json:"name"`
}

type Payload struct {
	Action  string   `json:"action"`
	Commits []Commit `json:"commits"`
}

type Commit struct {
	Message string `json:"message"`
}

func NewClient(cfg *config.Config) *Client {
	return &Client{config: cfg}
}

func (c *Client) FetchUserActivity(username string) ([]Event, error) {
	url := fmt.Sprintf("%s/users/%s/events", c.config.GitHubAPIURL, username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching GitHub activity: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned non-OK status: %s", resp.Status)
	}

	var events []Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, fmt.Errorf("error decoding GitHub response: %v", err)
	}

	return events, nil
}
