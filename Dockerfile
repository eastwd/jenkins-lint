FROM golang:1.16-alpine as builder

ENV APP_DIR=/go/src/github.com/eastwd/jenkins-lint

RUN apk update
RUN apk upgrade
RUN apk --no-cache add git musl-dev

COPY . $APP_DIR
WORKDIR $APP_DIR
RUN go build

FROM alpine:latest
COPY --from=builder /go/src/github.com/eastwd/jenkins-lint/jenkins-lint /jenkins-lint
ENTRYPOINT ["/jenkins-lint"]
