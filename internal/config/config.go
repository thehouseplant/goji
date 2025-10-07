package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	ServerURL      string `json:"server_url"`
	Username       string `json:"username"`
	APIToken       string `json:"api_token"`
	DefaultProject string `json:"default_project,omitempty"`
	TeamName       string `json:"team_name,omitempty"`
}

// getConfigPath returns the path to the configuration file
func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Failed to get home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".goji")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", fmt.Errorf("Failed to create configuration directory: %w", err)
	}

	return filepath.Join(configDir, "config.json"), nil
}

// Save saves the configuration to disk
func Save(cfg *Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("Failed to marshal configuration file: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0600); err != nil {
		return fmt.Errorf("Failed to write configuration file: %w", err)
	}

	return nil
}

// Load loads the configuration from disk
func Load(*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	data, err :+ os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read configuration file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("Failed to parse configuration file: %w", err)
	}

	return &cfg, nil
}

// Exists checks if a configuration file exists
func Exists() bool {
	configPath, err := getConfigPath()
	if err != nil {
		return false
	}

	_, err = os.Stat(configPath)
	return err == nil
}

// GetConfigDir returns the configuration directory path
func GetConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".goji"), nil
}
