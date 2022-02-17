package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
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
		replaceStatus1 := strings.Replace(tasks.Issues[i].Fields.Status.Status, "{", "", -1)
		replaceStatus2 := strings.Replace(replaceStatus1, "}", "", -1)
		fmt.Printf("\n")
		color.Red(tasks.Issues[i].ID)
		color.Magenta(replaceStatus2)
		fmt.Println(tasks.Issues[i].Fields.Name)
		fmt.Println(tasks.Issues[i].Fields.Goal)
		fmt.Printf("\n")
	}
}
