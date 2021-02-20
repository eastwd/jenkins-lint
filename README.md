# jenkins-lint

Jenkinsfile linter

## description

- Access the Jenkins validation API and lint Jenkinsfile

## install

```
go get -u  github.com/eastwd/jenkins-lint
```

## setting

- Generate your API token
  - https://[your jenkins url]/user/[your name]/configure

- Create .jlint.toml in home directory

```
[Client]
  Host = "http://localhost:8080"
  Insecure = false

[Account]
  Username = "your name"
  APIToken = "your API token"
```

## usage

```
$ jenkins-lint -f ./test/Jenkinsfile
Jenkinsfile successfully validated.

$ jenkins-lint -f ./test/Jenkinsfile.failed
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
