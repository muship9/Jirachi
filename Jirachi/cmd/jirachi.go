package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type User struct {
	name       string
	password   string
	requestUrl string
}

var user User

// jirachiCmd represents the jirachi command
var jirachiCmd = &cobra.Command{
	Use:   "jirachi",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		process := args[0]
		switch process {
		case "task":
			GetTask(user)
		case "show":
			fmt.Println("Sorry , I have no tasks to show")
		default:
			// TODO : cmd一覧を表示するようにしたい
			fmt.Println("Sorry ,Invalid command")
		}
	}}

func init() {
	rootCmd.AddCommand(jirachiCmd)

	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".envの読み込みに失敗しました。")
	}
	user = User{
		name:       os.Getenv("USER_NAME"),
		password:   os.Getenv("PASSWORD"),
		requestUrl: os.Getenv("BASE_URL"),
	}
}
