package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

// Config main configuration struct. Populated from yaml file
type Config struct {
	Telegram struct {
		Name  string `yaml:"name"`
		Token string `yaml:"token"`
		URL   string `yaml:"url"`
	} `yaml:"telegram"`
	Database struct {
		Driver   string `yaml:"driver"`
		User     string `yaml:"user"`
		Name     string `yaml:"name"`
		Password string `yaml:"password"`
		SSL      string `yaml:"ssl"`
	} `yaml:"database"`
}

// loadConfig loads yaml-configuration file to Config struct
func loadConfig(configPath string) (*Config, error) {
	config := &Config{}

	// Open file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init decoder and decode
	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, err
}

// validateConfigPath validates configuration file is not a directory
func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

// parseFlags create and parse the CLI flags
// and return the path to be used elsewhere
func parseFlags() (string, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./config.yaml", "path to config file")

	// Actually parse the flags
	flag.Parse()

	// Validate the path first
	if err := validateConfigPath(configPath); err != nil {
		return "", err
	}

	// Return the configuration path
	return configPath, nil
}
