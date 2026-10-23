package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not get home directory: %w", err)
	}
	loc := filepath.Join(home, ".gatorconfig.json")
	data, err := os.ReadFile(loc)
	if err != nil {
		return nil, fmt.Errorf("could not read config file: %v", err)
	}
	var config *Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("could not parse config file: %v", err)
	}
	return config, nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("could not marshal config: %v", err)
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not get home directory: %w", err)
	}
	loc := filepath.Join(home, ".gatorconfig.json")
	err = os.WriteFile(loc, data, 0644)
	if err != nil {
		return fmt.Errorf("could not write config file: %v", err)
	}
	return nil
}
