package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/eastwd/jctl/jenkins"
	"github.com/urfave/cli"
)

var lintCmd = func(c *cli.Context) error {
	jc := config.Client
	client := jenkins.NewClient(jc.Host, jc.TLSVerify)

	//JenkinsのCrumbを取得
	err := client.FetchCrumb()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	//Jenkinsfileの読み込み
	jenkinsfile, err := jenkins.ReadJenkinsfile(c.String("file"))
	if err != nil {
		log.Fatal(err)
		return nil
	}

	//バリデーションの結果を取得
	result, err := client.Validate(jenkinsfile)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	fmt.Println(result)
	return nil
}

var lintCmdFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "file, f",
		Value: "./Jenkinsfile",
		Usage: "Jenkinsfileのパス",
	},
}

var configCmd = func(c *cli.Context) error {
	config.Account.Password = strings.Repeat("*", len(config.Account.Password))
	fmt.Println(config.String())
	return nil
}
