package cmd

import (
	"fmt"
)

// []byteを見やすいように整える
func ConsistencyTasks(res []byte) {
	fmt.Println(string(res))
}
