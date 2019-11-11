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
	Username  string
	APIToken  string
	Crumb     string
	TlsVerify bool
	client    *http.Client
}

var (
	crumbUrlFormat    = "%s/crumbIssuer/api/xml?xpath=concat(//crumbRequestField,\":\",//crumb)"
	validateUrlFormat = "%s/pipeline-model-converter/validate"
)

func NewClient(host string, tlsVerify bool, username string, apiToken string) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: tlsVerify,
		},
	}
	c := &http.Client{Transport: tr}
	client := &Client{
		Host:     host,
		Username: username,
		APIToken: apiToken,
		client:   c,
	}
	return client
}
func (c *Client) FetchCrumb() error {
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf(crumbUrlFormat, c.Host),
		nil,
	)
	if c.Username != "" && c.APIToken != "" {
		req.SetBasicAuth(c.Username, c.APIToken)
	}
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
	if c.Username != "" && c.APIToken != "" {
		req.SetBasicAuth(c.Username, c.APIToken)
	}

	req.Header.Set("Jenkins-Crumb", c.Crumb)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
