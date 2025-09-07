package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitgo",
	Short: "Git.go - A simple and small VCS CLI build in Go.",
	Long:  `Git.go is a simple lightweighted version control system CLI build in go.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Git.go!")
		fmt.Println("- Commit your sins and past mistakes. We’ll remember them all like your toxic girlfriend.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil { // if command runs without err then make the err = nil
		fmt.Println(err)
		os.Exit(1)
	}
}
