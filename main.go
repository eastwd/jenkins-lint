package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {

	//configの読み込み
	NewConfig()

	app := &cli.Command{
		Name:        "jenkins-lint",
		Usage:       "Linter of Jenkinsfile",
		Description: "Access the Jenkins validation API and lint Jenkinsfile",
		Version:     "0.3.4",
		Commands: []*cli.Command{
			{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Print current configuration",
				Action:  configCmd,
			},
		},
		Action: lintCmd,
		Flags:  lintCmdFlags,
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
