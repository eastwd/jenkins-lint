package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {

	//configの読み込み
	NewConfig()

	app := cli.NewApp()
	app.Name = "jenkins-lint"
	app.Usage = "Jenkinsfileのlinter"
	app.Description = "JenkinsのバリデーションAPIにアクセスし、Jenkinsfileの構文チェックをする"
	app.Commands = []cli.Command{
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "現在の設定を表示します",
			Action:  configCmd,
		},
	}
	app.Action = lintCmd
	app.Flags = lintCmdFlags

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
