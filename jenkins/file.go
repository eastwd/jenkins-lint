package jenkins

import (
	"bytes"
	"io"
	"os"
)

func ReadJenkinsfile(filePath string) (string, error) {
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
