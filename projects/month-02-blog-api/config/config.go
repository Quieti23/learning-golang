package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type AppConfig struct {
	ServerPort          string `json:"server_port"`
	MySQLDSN            string `json:"mysql_dsn"`
	MaxOpenConns        int    `json:"max_open_conns"`
	MaxIdleConns        int    `json:"max_idle_conns"`
	ConnMaxLifetimeMins int    `json:"conn_max_lifetime_minutes"`
	RequestTimeoutMs    int    `json:"request_timeout_ms"`
}

func Load(path string) (AppConfig, error) {
	cfg := AppConfig{
		ServerPort:          "8081",
		MaxOpenConns:        10,
		MaxIdleConns:        5,
		ConnMaxLifetimeMins: 5,
		RequestTimeoutMs:    3000,
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
	if strings.TrimSpace(cfg.MySQLDSN) == "" {
		return AppConfig{}, fmt.Errorf("mysql_dsn is required")
	}
	if cfg.MaxOpenConns <= 0 {
		return AppConfig{}, fmt.Errorf("max_open_conns must be greater than 0")
	}
	if cfg.MaxIdleConns < 0 {
		return AppConfig{}, fmt.Errorf("max_idle_conns must be 0 or greater")
	}
	if cfg.ConnMaxLifetimeMins <= 0 {
		return AppConfig{}, fmt.Errorf("conn_max_lifetime_minutes must be greater than 0")
	}
	if cfg.RequestTimeoutMs <= 0 {
		return AppConfig{}, fmt.Errorf("request_timeout_ms must be greater than 0")
	}

	return cfg, nil
}

func (c AppConfig) Address() string {
	if strings.HasPrefix(c.ServerPort, ":") {
		return c.ServerPort
	}

	return ":" + c.ServerPort
}

func (c AppConfig) ConnMaxLifetime() time.Duration {
	return time.Duration(c.ConnMaxLifetimeMins) * time.Minute
}

func (c AppConfig) RequestTimeout() time.Duration {
	return time.Duration(c.RequestTimeoutMs) * time.Millisecond
}

func (c AppConfig) PoolSummary() string {
	return "max_open=" + strconv.Itoa(c.MaxOpenConns) + ", max_idle=" + strconv.Itoa(c.MaxIdleConns) + ", conn_lifetime_minutes=" + strconv.Itoa(c.ConnMaxLifetimeMins)
}
