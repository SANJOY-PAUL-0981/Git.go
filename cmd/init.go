package cmd

import (
	"fmt"
	"gitgo/internal/repo"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initilize a new Git.go repository",
	Run: func(cmd *cobra.Command, args []string) {
		err := repo.InitRepo()

		if err != nil {
			fmt.Println("Error initilizing Git.go repository")
			return
		}

		fmt.Println("Initilized empty Git.go repository in .gitgo/")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
