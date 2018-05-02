package jenkins

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	Host      string
	Crumb     string
	TlsVerify bool
	client    *http.Client
}

var (
	crumbUrlFormat    = "%s/crumbIssuer/api/xml?xpath=concat(//crumbRequestField,\":\",//crumb)"
	validateUrlFormat = "%s/pipeline-model-converter/validate"
)

func NewClient(host string, tlsVerify bool) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: tlsVerify,
		},
	}
	c := &http.Client{Transport: tr}
	client := &Client{
		Host:   host,
		client: c,
	}
	return client
}
func (c *Client) FetchCrumb() error {
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf(crumbUrlFormat, c.Host),
		nil,
	)
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	c.Crumb = strings.Split(string(body), ":")[0]
	return nil
}

func (c *Client) Validate(jenkinsfile string) (string, error) {
	form := url.Values{}
	form.Add("jenkinsfile", jenkinsfile)

	//バリデーションの結果を取得
	req, err := http.NewRequest("POST", fmt.Sprintf(validateUrlFormat, c.Host), strings.NewReader(form.Encode()))
	if err != nil {
		return "", nil
	}
	req.Header.Set("Jenkins-Crumb", c.Crumb)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
