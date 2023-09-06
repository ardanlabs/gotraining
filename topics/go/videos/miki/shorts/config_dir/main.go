package main

import (
	"fmt"
	"os"
	"path"
)

var (
	appName = "httpd"
)

func configFile() (string, error) {
	fileName := os.Getenv("HTTPD_CONFIG_FILE")
	if fileName != "" {
		return fileName, nil
	}

	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("can't get user config dir - %w", err)
	}

	fileName = path.Join(cfgDir, appName, "config.toml")
	return fileName, nil
}

func main() {
	fmt.Println(configFile())
}
