package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Logs struct {
		SSHLogPath   string `yaml:"ssh_log_path"`
		OtherLogPath string `yaml:"other_log_path"`
	} `yaml:"logs"`
	Prometheus struct {
		MetricsPort string `yaml:"metrics_port"`
		MetricsPath string `yaml:"metrics_path"`
	} `yaml:"prometheus"`
	Intervals struct {
		LogScanInterval string `yaml:"log_scan_interval"`
	} `yaml:"intervals"`
	Collectors struct {
		SSH    bool `yaml:"ssh"`
		System bool `yaml:"system"`
	} `yaml:"collectors"`
}

var AppConfig Config

// LoadConfig loads configuration from the provided YAML file
func LoadConfig(configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("could not open config file: %w", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		return fmt.Errorf("could not decode config file: %w", err)
	}
	return nil
}
