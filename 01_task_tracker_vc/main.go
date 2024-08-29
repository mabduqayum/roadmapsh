package main

import (
	"fmt"
	"os"

	"github.com/mabduqayum/roadmapsh/01_task_tracker_vc/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
