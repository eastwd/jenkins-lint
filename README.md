# jctl

Jenkinsfile linter

```
NAME:
   jctl - Jenkinsfileのlinter

USAGE:
   jctl [global options] command [command options] [arguments...]

VERSION:
   0.0.0

DESCRIPTION:
   JenkinsのバリデーションAPIにアクセスし、Jenkinsfileの構文チェックをする

COMMANDS:
     lint, validate, l, v  指定したJenkinsfileの構文チェックをします
     config, c             現在の設定を表示します
     help, h               Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## usage

.jctl.tomlをホームディレクトリに作成

```
$ jctl config
[Client]
  Host = "http://localhost:8080"
  TLSVerify = true

[Account]
  Username = ""
  Password = ""

```

Jenkinsfileのパスを指定して実行

```
$ jctl lint -f ./test/Jenkinsfile
Reading Jenkinsfile...
pipeline {
    agent {
      docker {
        image 'alpine'
        label 'linux'
      }
    }
    stages {
        stage('Build') {
            steps {
                echo 'hello-world'
                sh 'whoami'
                sh 'pwd'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}

Jenkinsfile successfully validated.

```
