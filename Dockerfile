FROM golang:1.27

WORKDIR /app

COPY . .
RUN go build

FROM ubuntu:latest
COPY --from=0 /app/jenkins-lint /jenkins-lint
ENTRYPOINT ["/jenkins-lint"]
