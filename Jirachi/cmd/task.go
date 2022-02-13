package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("task called")
	},
}

func init() {
	rootCmd.AddCommand(taskCmd)
}

func GetTask(user User) {
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
