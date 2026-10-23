package main

import (
	"bytes"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Client  ClientConfig
	Account AccountConfig
}

type ClientConfig struct {
	Host     string
	Insecure bool
}

type AccountConfig struct {
	Username string
	APIToken string
}

func homeDir() string {
	if home, err := os.UserHomeDir(); err == nil {
		return home
	}
	if u, err := user.Current(); err == nil {
		return u.HomeDir
	}
	return ""
}

func NewConfig() {
	home := homeDir()
	if _, err := toml.DecodeFile(defaultConfigPath, &config); err != nil {
		if _, err2 := toml.DecodeFile(filepath.Join(home, configName), &config); err2 != nil {
			config = DefaultConfig()
		}
	}
}

func DefaultConfig() Config {
	return Config{
		Client: ClientConfig{
			Host:     defaultHost,
			Insecure: false,
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
		fmt.Println(err)
		return ""
	}
	return buf.String()
}
