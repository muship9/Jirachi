package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	user := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	requestUrl := os.Getenv("BASE_URL")

	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		log.Println(err)
	}

	// 認証情報を付与し、リクエストと一緒に送る
	req.SetBasicAuth(user, password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	// レスポンスボディをすべて読み出す
	body, _ := ioutil.ReadAll(resp.Body)
	// body は []byte
	fmt.Println(string(body))
}
