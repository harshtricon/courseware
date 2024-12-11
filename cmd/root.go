package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-courseware",
	Short: "A CLI tool to fetch courseware data",
	Long:  `Go Courseware CLI fetches and displays courseware data using a REST API.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Go Courseware CLI. Use 'fetch' command to fetch courseware data.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
