package config

import (
	"os"
	"strings"
)

type Config struct {
	ClickHouseIP       string
	ClickHouseUsername string
	ClickHousePassword string
	ClickHouseDatabase string
	ClickHouseTable    string
}

func LoadClickHouseConfig() *Config {
	return &Config{
		ClickHouseIP:       getEnvOrDefault("TEST_DB_IP", " "),
		ClickHouseUsername: getEnvOrDefault("TEST_DB_USERNAME", " "),
		ClickHousePassword: getEnvOrDefault("TEST_DB_PASSWORD", " "),
		ClickHouseDatabase: "test_db",
		ClickHouseTable:    "kv7",
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return strings.Trim(value, "\n")
	}
	return defaultValue
}
