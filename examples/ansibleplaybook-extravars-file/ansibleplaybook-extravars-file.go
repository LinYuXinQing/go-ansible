package main

import (
	"context"

	"github.com/LinYuXinQing/gooo-ansible/pkg/options"
	"github.com/LinYuXinQing/gooo-ansible/pkg/playbook"
)

func main() {

	ansiblePlaybookConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "local",
		User:       "aleix",
	}

	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		Inventory: "127.0.0.1,",
		ExtraVarsFile: []string{
			"@vars-file1.yml",
			"@vars-file2.yml",
		},
	}

	playbook := &playbook.AnsiblePlaybookCmd{
		Playbooks:         []string{"site.yml"},
		ConnectionOptions: ansiblePlaybookConnectionOptions,
		Options:           ansiblePlaybookOptions,
	}

	err := playbook.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}
