package main

import (
	"bytes"
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
	homedir "github.com/mitchellh/go-homedir"
)

type Config struct {
	Client  ClientConfig
	Account AccountConfig
}

type ClientConfig struct {
	Host      string
	TLSVerify bool
}

type AccountConfig struct {
	Username string
	APIToken string
}

func NewConfig() {
	home, _ := homedir.Dir()
	if _, err := toml.DecodeFile(filepath.Join(home, configName), &config); err != nil {
		if _, err2 := toml.DecodeFile(defaultConfigPath, &config); err2 != nil {
			config = DefaultConfig()
		}
	}
}

func DefaultConfig() Config {
	return Config{
		Client: ClientConfig{
			Host:      defaultHost,
			TLSVerify: true,
		},
		Account: AccountConfig{
			Username: "",
			APIToken: "",
		},
	}
}

var config Config

func (c *Config) String() string {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(config); err != nil {
		log.Fatal(err)
		return ""
	}
	return buf.String()
}
