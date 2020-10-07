package config

import (
	"errors"
	"strings"
)

// Config struct
type Config struct {
	Debug      bool
	ProjectDir string
	Name       string
	Version    string
}

// NewConfig boostrap a new config for the application
func NewConfig(name, projectDir, version string) (*Config, error) {
	appConfig := &Config{
		Name:       name,
		ProjectDir: projectDir,
		Version:    version,
	}

	return appConfig, nil
}

// ProjectName returns just the name of the current directory, without the rooted path
func (c *Config) ProjectName() (string, error) {
	if len(c.ProjectDir) == 0 {
		return "", errors.New("ProjectDir from Config is empty")
	}
	path := strings.Split(c.ProjectDir, "/")

	projectName := path[len(path)-1]

	return projectName, nil
}
