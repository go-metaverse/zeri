package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// NewConfig reads a configuration file in either JSON or YAML format and populates the provided struct model.
// Returns the populated model or an error if the path is invalid, the file cannot be read, or if unmarshalling fails.
//
// The file format is determined by the file extension:
// - .json for JSON files
// - .yml or .yaml for YAML files
//
// Parameters:
// - configPath: Path to the configuration file (JSON or YAML).
// - configModel: Pointer to a struct that will be populated with the configuration data.
//
// Returns:
// - *T: A pointer to the populated struct or nil if an error occurs.
// - error: An error message if reading the file or unmarshalling fails.
//
// Example usage:
//
//	cfg, err := NewConfig("config.yaml", &MyConfig{})
//	if err != nil {
//	    log.Fatal(err)
//	}
func NewConfig[T any](configPath string, configModel *T) (*T, error) {
	if configPath == "" {
		return nil, fmt.Errorf("config path is required")
	}

	// Read the file content.
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Determine the file format based on the file extension.
	switch ext := filepath.Ext(configPath); ext {
	case ".json":
		// Unmarshal JSON content
		if err := json.Unmarshal(data, configModel); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
		}
	case ".yml", ".yaml":
		// Unmarshal YAML content
		if err := yaml.Unmarshal(data, configModel); err != nil {
			return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported file extension: %s", ext)
	}

	return configModel, nil
}

// GetConfigPath returns the configuration file path based on the environment and file extension.
// It supports both YAML (.yml, .yaml) and JSON (.json) formats.
//
// Parameters:
// - env: The environment name (e.g., "qc", "staging", "prod").
// - ext: The file extension to return (either "json", "yml", or "yaml").
//
// Returns:
// - string: The path to the appropriate configuration file.
//
// Example usage:
//
//	GetConfigPath("qc", "json")  -> returns "./env/env.qc.json"
//	GetConfigPath("prod", "yml") -> returns "./env/env.prod.yml"
func GetConfigPath(env, ext string) string {
	if ext != "json" && ext != "yml" && ext != "yaml" {
		ext = "yml" // Default to YAML if the extension is invalid
	}

	configPaths := map[string]string{
		"qc":      "./env/env.qc." + ext,
		"staging": "./env/env.staging." + ext,
		"prod":    "./env/env.prod." + ext,
	}

	// Return the config path for the environment or the default local file.
	if path, exists := configPaths[env]; exists {
		return path
	}

	// Default to a local configuration file based on the provided extension.
	return "./env/env.local." + ext
}
