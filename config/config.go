package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type DB struct {
	Host     string `json:"host"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// Config represents the configuration information.
type Config struct {
	Db DB `json:"db"`
}

func (c *Config) String() string {
	return fmt.Sprintf("%s:%s@(%s)/%s", c.Db.User, c.Db.Password, c.Db.Host, c.Db.Database)
}

var Conf Config

func init() {
	// default to development
	configFile := "./config-dev.json"

	// Get the config file name if passed in. For prd, pass it in
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}

	config_file, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
	}
	json.Unmarshal(config_file, &Conf)
}
