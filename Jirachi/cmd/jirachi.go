package cmd

import (
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
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
		GetSplintTask()
	},
}

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

func GetSplintTask() {
	req, err := http.NewRequest(http.MethodGet, user.requestUrl, nil)
	if err != nil {
		log.Println(err)
	}

	// 認証情報を付与し、リクエストと一緒に送る
	req.SetBasicAuth(user.name, user.password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	// レスポンスボディをすべて読み出す
	body, _ := ioutil.ReadAll(resp.Body)
	// body は []byte
	fmt.Println(string(body))
}
