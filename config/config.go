package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	CMD              string   `json:"cmd"`
	ProjectDirectory string   `json:"project_directory"`
	ProjectDirs      []string `json:"-"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) LoadConfigFromFile() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	configFilePath := filepath.Join(homeDir, ".config", "project-selector", "config.json")
	fileContents, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	err = json.Unmarshal(fileContents, c)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	c.ProjectDirs = strings.Split(c.ProjectDirectory, "/")
	return nil
}

func (c *Config) SetDefaultConfig() {
	c.CMD = "zed"
	c.ProjectDirectory = "Development/work/bayonet"
	c.ProjectDirs = strings.Split(c.ProjectDirectory, "/")
}
