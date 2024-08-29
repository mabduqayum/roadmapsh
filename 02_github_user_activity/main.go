package main

import (
	"log"
	"os"

	"github.com/mabduqayum/roadmapsh/02_github_user_activity/config"
	"github.com/mabduqayum/roadmapsh/02_github_user_activity/internal/cli"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	app := cli.NewApp(cfg)
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
