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
	app.Name = "jctl"
	app.Usage = "Jenkinsfileのlinter"
	app.Description = "JenkinsのバリデーションAPIにアクセスし、Jenkinsfileの構文チェックをする"
	app.Commands = []cli.Command{
		{
			Name:    "lint",
			Aliases: []string{"validate", "l", "v"},
			Usage:   "指定したJenkinsfileの構文チェックをします",
			Action:  lintCmd,
			Flags:   lintCmdFlags,
		},
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "現在の設定を表示します",
			Action:  configCmd,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
