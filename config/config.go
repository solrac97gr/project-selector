package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	CMD              string   `json:"cmd"`
	ProjectDirs      []string `json:"project_dirs"`
	DirsToIgnore     []string `json:"dirs_to_ignore"`
	NumberOfProjects int      `json:"number_of_projects"`
	Style            struct {
		Title struct {
			Template string `json:"template"`
			Icon     string `json:"icon"`
		} `json:"title"`
		Active struct {
			Template string `json:"template"`
			Icon     string `json:"icon"`
		} `json:"active"`
		Inactive struct {
			Template string `json:"template"`
			Icon     string `json:"icon"`
		} `json:"inactive"`
	} `json:"style"`
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
	fileContents, err := os.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	err = json.Unmarshal(fileContents, c)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return nil
}

func (c *Config) SetDefaultConfig() {
	c.CMD = "zed"
	c.ProjectDirs = []string{
		"Development/work",
		"Development/personal",
	}
	c.DirsToIgnore = []string{
		"node_modules",
		".git",
	}
	c.NumberOfProjects = 5

	c.Style.Title.Template = "{{ . | blue | bold }}"
	c.Style.Title.Icon = "🔎"

	c.Style.Active.Template = "{{ . | blue | underline | bold}}"
	c.Style.Active.Icon = "🚀"

	c.Style.Inactive.Template = "{{ . | cyan }}"
	c.Style.Inactive.Icon = "📁"

}
