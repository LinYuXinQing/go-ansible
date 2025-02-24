package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/LinYuXinQing/gooo-ansible/pkg/execute"
	"github.com/LinYuXinQing/gooo-ansible/pkg/options"
	"github.com/LinYuXinQing/gooo-ansible/pkg/playbook"
	"github.com/LinYuXinQing/gooo-ansible/pkg/stdoutcallback/results"
)

func main() {

	var err error
	var res *results.AnsiblePlaybookJSONResults

	buff := new(bytes.Buffer)

	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
		User:       "apenella",
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
	}

	execute := execute.NewDefaultExecute(
		execute.WithWrite(io.Writer(buff)),
	)

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         []string{"site1.yml", "site2.yml"},
		Exec:              execute,
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
		StdoutCallback:    "json",
	}

	err = playbook.Run(context.TODO())
	if err != nil {
		fmt.Println(err.Error())
	}

	res, err = results.ParseJSONResultsStream(io.Reader(buff))
	if err != nil {
		panic(err)
	}

	msgOutput := struct {
		Host    string `json:"host"`
		Message string `json:"message"`
	}{}

	for _, play := range res.Plays {
		for _, task := range play.Tasks {
			if task.Task.Name == "ansibleplaybook-walk-through-json-output" {
				for _, content := range task.Hosts {

					err = json.Unmarshal([]byte(fmt.Sprint(content.Stdout)), &msgOutput)
					if err != nil {
						panic(err)
					}

					fmt.Printf("[%s] %s\n", msgOutput.Host, msgOutput.Message)
				}
			}
		}
	}
}
