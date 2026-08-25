package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	AppName       string
	TasksFilePath string
}

func New() *Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}

	configDir := filepath.Join(homeDir, ".todosher")
	_ = os.MkdirAll(configDir, 0755)

	return &Config{
		AppName:       "todosher",
		TasksFilePath: filepath.Join(configDir, "tasks.json"),
	}
}
