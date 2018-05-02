jenkins-lint:
	go build -o jenkins-lint
	go install

run:
	go run *.go
