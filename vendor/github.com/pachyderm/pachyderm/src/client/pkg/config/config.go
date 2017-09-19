package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pachyderm/pachyderm/src/client/pkg/uuid"
)

var configDirPath = filepath.Join(os.Getenv("HOME"), ".pachyderm")
var configPath = filepath.Join(configDirPath, "config.json")

// Read loads the Pachyderm config on this machine.
// If an existing configuration cannot be found, it sets up the defaults. Read
// returns a nil Config if and only if it returns a non-nil error.
func Read() (*Config, error) {
	var c *Config

	// Read json file
	if raw, err := ioutil.ReadFile(configPath); err == nil {
		err = json.Unmarshal(raw, &c)
		if err != nil {
			return nil, err
		}
	} else {
		// File doesn't exist, so create a new config
		fmt.Println("No config detected. Generating new config...")
		c = &Config{}
	}
	if c.UserID == "" {
		fmt.Printf("No UserID present in config. Generating new UserID and "+
			"updating config at %s\n", configPath)
		c.UserID = uuid.NewWithoutDashes()
		if err := c.Write(); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// Write writes the configuration in 'c' to this machine's Pachyderm config
// file.
func (c *Config) Write() error {
	rawConfig, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	err = os.MkdirAll(configDirPath, 0755)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(configPath, rawConfig, 0644)
}
