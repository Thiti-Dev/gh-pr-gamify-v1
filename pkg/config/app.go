package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Repository defines the structure for a single repository configuration
type Repository struct {
	Name         string `yaml:"name"`
	Organization string `yaml:"organization"`
}

// GithubConfig defines the structure for Github-related configuration
type GithubConfig struct {
	Token        string       `yaml:"token"`
	Repositories []Repository `yaml:"repositories"`
}

// Config defines the overall configuration structure
type GithubConfigRoot struct {
	Github GithubConfig `yaml:"github"`
}

// LoadConfig loads configuration from a YAML file
func LoadApplicationConfig() (*GithubConfigRoot, error) {
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		return nil, err
	}

	var cfg GithubConfigRoot
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
