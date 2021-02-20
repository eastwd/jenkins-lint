package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {

	//configの読み込み
	NewConfig()

	app := cli.NewApp()
	app.Name = "jenkins-lint"
	app.Usage = "Linter of Jenkinsfile"
	app.Description = "Access the Jenkins validation API and lint Jenkinsfile"
	app.Version = "0.3.4"
	app.Commands = []cli.Command{
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Print current configuration",
			Action:  configCmd,
		},
	}
	app.Action = lintCmd
	app.Flags = lintCmdFlags

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
