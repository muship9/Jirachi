package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	for i := 0; i < len(tasks.Issues); i++ {
		fmt.Printf("\n")
		replaceStatus1 := strings.Replace(tasks.Issues[i].Fields.Status.Status, "{", "", -1)
		replaceStatus2 := strings.Replace(replaceStatus1, "}", "", -1)
		fmt.Println(tasks.Issues[i].ID, tasks.Issues[i].Fields.Name, tasks.Issues[i].Fields.Goal, replaceStatus2)
		fmt.Printf("\n")
	}
}
