package main

import (
	"fmt"
	"os"

	"github.com/mabduqayum/roadmapsh/01_task_tracker_vc/cmd"
	"github.com/mabduqayum/roadmapsh/01_task_tracker_vc/config"
)

func main() {
	cfg := config.Load()
	command := cmd.NewRootCmd(cfg.Storage.File)

	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
