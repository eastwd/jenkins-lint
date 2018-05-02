# jctl

Jenkinsfile linter

## description

- JenkinsのAPIにアクセスし、Jenkinsfileの構文チェックをするツール
- 認証周りは作成中

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
Jenkinsfile successfully validated.

$ jctl lint -f ./test/Jenkinsfile.failed
  1 : pipeline {
  2 :     agent {
  3 :       docker {
  4 :         image 'alpine'
  5 :         label 'linux'
  6 :       }
  7 :     }
  8 :     stages {
  9 :             steps {
 10 :                 echo 'hello-world'
 11 :                 sh 'whoami'
 12 :                 sh 'pwd'
 13 :             }
 14 :         }
 15 :         stage('Test') {
 16 :             steps {
 17 :                 echo 'Testing..'
 18 :             }
 19 :         }
 20 :         stage('Deploy') {
 21 :             steps {
 22 :                 echo 'Deploying....'
 23 :             }
 24 :         }
 25 :     }
 26 : }
 27 : 
Errors encountered validating Jenkinsfile:
WorkflowScript: 26: unexpected token: } @ line 26, column 1.
   }
   ^

```
