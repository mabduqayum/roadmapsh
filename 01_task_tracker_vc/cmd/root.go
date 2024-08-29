package cmd

import (
	"fmt"
	"os"

	"github.com/mabduqayum/roadmapsh/01_task_tracker_vc/tracker"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	taskTracker *tracker.TaskTracker
)

var rootCmd = &cobra.Command{
	Use:   "task-cli",
	Short: "A simple CLI task tracker",
	Long:  `Task-cli is a simple command line interface to track and manage your tasks.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.db.yaml)")

	taskTracker = tracker.NewTaskTracker()
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".db")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
