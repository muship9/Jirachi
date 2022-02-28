/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
			and usage of using your command. For example:
			
			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("move called")
	},
}

func init() {
	rootCmd.AddCommand(moveCmd)
}

func IssueStatusMove(id string, status string) {
	requestUrl := os.Getenv("MOVE_URL")
	log.Println(requestUrl)
}
