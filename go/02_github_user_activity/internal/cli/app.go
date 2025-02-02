package cli

import (
	"fmt"

	"github_user_activity/config"
	"github_user_activity/internal/github"

	"github.com/urfave/cli/v2"
)

func NewApp(cfg *config.Config) *cli.App {
	client := github.NewClient(cfg)

	return &cli.App{
		Name:  "github-activity",
		Usage: "Fetch and display GitHub user activity",
		Action: func(c *cli.Context) error {
			if c.NArg() < 1 {
				return cli.Exit("Please provide a GitHub username", 1)
			}
			username := c.Args().Get(0)
			return fetchAndDisplayActivity(username, client)
		},
	}
}

func fetchAndDisplayActivity(username string, client *github.Client) error {
	events, err := client.FetchUserActivity(username)
	if err != nil {
		return err
	}

	for _, event := range events {
		displayEvent(event)
	}

	return nil
}

func displayEvent(event github.Event) {
	switch event.Type {
	case "PushEvent":
		fmt.Printf("- Pushed %d commits to %s\n", len(event.Payload.Commits), event.Repo.Name)
	case "IssuesEvent":
		fmt.Printf("- %s an issue in %s\n", event.Payload.Action, event.Repo.Name)
	case "WatchEvent":
		fmt.Printf("- Starred %s\n", event.Repo.Name)
	default:
		fmt.Printf("- %s on %s\n", event.Type, event.Repo.Name)
	}
}
