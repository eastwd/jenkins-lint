package jenkins

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func ReadJenkinsfile(filePath string) (string, error) {
	fmt.Println("Reading Jenkinsfile...")
	buf := bytes.NewBuffer(nil)
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return "", err
	}
	io.Copy(buf, f)
	s := string(buf.Bytes())
	return s, nil
}
