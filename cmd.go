package main

import (
	"fmt"
	"strings"

	"github.com/eastwd/jenkins-lint/jenkins"
	"github.com/urfave/cli"
)

var lintCmd = func(c *cli.Context) error {
	jc := config.Client
	ja := config.Account
	username := ja.Username
	if c.String("username") != "" {
		username = c.String("username")
	}
	apiToken := ja.APIToken
	if c.String("token") != "" {
		apiToken = c.String("token")
	}
	client := jenkins.NewClient(jc.Host, jc.Insecure, username, apiToken)

	//JenkinsのCrumbを取得
	err := client.FetchCrumb()
	if err != nil {
		return err
	}

	//Jenkinsfileの読み込み
	jenkinsfile, err := jenkins.ReadJenkinsfile(c.String("file"))
	if err != nil {
		return err
	}

	//バリデーションの結果を取得
	result, err := client.Validate(jenkinsfile)
	if err != nil {
		return err
	}

	if !strings.Contains(result, successMessage) {
		s := strings.Split(jenkinsfile, "\n")
		for i, row := range s {
			fmt.Printf("%3d : %s\n", i+1, row)
		}
	}
	fmt.Println(result)
	return nil
}

var lintCmdFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "file, f",
		Value: "Jenkinsfile",
		Usage: "Relative path of Jenkinsfile",
	},
	cli.StringFlag{
		Name:  "username, u",
		Value: "",
		Usage: "Relative path of Jenkinsfile",
	},
	cli.StringFlag{
		Name:  "token, t",
		Value: "",
		Usage: "Jenkins API token",
	},
}

var configCmd = func(c *cli.Context) error {
	config.Account.APIToken = strings.Repeat("*", len(config.Account.APIToken))
	fmt.Println(config.String())
	return nil
}
