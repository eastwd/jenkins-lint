package jenkins_test

import (
	"testing"

	"github.com/eastwd/jenkins-lint/jenkins"
)

func TestNewClient(t *testing.T) {
	client := jenkins.NewClient("test_host", true, "hoge", "fuga")
	expected := "test_host"
	if client.Host != expected {
		t.Fatalf("hostname Must be %s", expected)
	}
}
