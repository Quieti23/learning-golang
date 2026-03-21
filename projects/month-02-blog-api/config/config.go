package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type AppConfig struct {
	ServerPort string `json:"server_port"`
}

func Load(path string) (AppConfig, error) {
	cfg := AppConfig{
		ServerPort: "8081",
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return AppConfig{}, fmt.Errorf("read config file failed: %w", err)
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return AppConfig{}, fmt.Errorf("parse config file failed: %w", err)
	}

	cfg.ServerPort = strings.TrimSpace(cfg.ServerPort)
	if cfg.ServerPort == "" {
		return AppConfig{}, fmt.Errorf("server_port is required")
	}

	return cfg, nil
}

func (c AppConfig) Address() string {
	if strings.HasPrefix(c.ServerPort, ":") {
		return c.ServerPort
	}

	return ":" + c.ServerPort
}
