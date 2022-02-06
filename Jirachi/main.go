package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type User struct {
	name       string
	password   string
	requestUrl string
}

var user User

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("envの読み込みに失敗しました。")
	}
	user = User{
		name:       os.Getenv("USER_NAME"),
		password:   os.Getenv("PASSWORD"),
		requestUrl: os.Getenv("BASE_URL"),
	}

}

func main() {
	getSplintTask()
}

func getSplintTask() {
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
