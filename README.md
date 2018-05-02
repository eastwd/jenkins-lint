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
  1 : pipeline {
  2 :     agent {
  3 :       docker {
  4 :         image 'alpine'
  5 :         label 'linux'
  6 :       }
  7 :     }
  8 :     stages {
  9 :         stage('Build') {
 10 :             steps {
 11 :                 echo 'hello-world'
 12 :                 sh 'whoami'
 13 :                 sh 'pwd'
 14 :             }
 15 :         }
 16 :         stage('Test') {
 17 :             steps {
 18 :                 echo 'Testing..'
 19 :             }
 20 :         }
 21 :         stage('Deploy') {
 22 :             steps {
 23 :                 echo 'Deploying....'
 24 :             }
 25 :         }
 26 :     }
 27 : }
 28 :
Jenkinsfile successfully validated.
```
