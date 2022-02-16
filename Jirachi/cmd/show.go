package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type IssueList struct {
	Issues []Issue `json:"issues"`
}

type Issue struct {
	ID     string `json:"key"`
	Fields Fields `json:"fields"`
}

type Fields struct {
	Goal   string     `json:"customfield_10031"`
	Name   string     `json:"summary"`
	Status StatusName `json:"status"`
}

type StatusName struct {
	Status string `json:"name"`
}

func ConsistencyTasks(resp *http.Response, err error) {
	// レスポンスボディをすべて読み出す
	body, err := ioutil.ReadAll(resp.Body)
	// 配列の作成
	tasks := IssueList{}
	//tasks := make([]*Issues , 0)
	// jsonから構造体に変える
	err = json.Unmarshal(body, &tasks)
	if err != nil {
		log.Println("unmarshall error")
	}
	fmt.Printf("\n\n", tasks)
}
