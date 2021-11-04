package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/uselagoon/lagoon-linter/internal/lagoonyml/deprecated"
	"github.com/uselagoon/lagoon-linter/internal/lagoonyml/required"
)

// ValidateConfigMapJSONCmd represents the validate command.
type ValidateConfigMapJSONCmd struct {
	ConfigMapJSON string `kong:"default='configmap.json',help='Specify the configmap JSON file dump.'"`
	Profile       string `kong:"default='required',enum='required,deprecated',help='Set the linting profile (required,deprecated)'"`
}

// ConfigMap represents an individual configmap.
type ConfigMap struct {
	Data     map[string]string      `json:"data"`
	Metadata map[string]interface{} `json:"metadata"`
}

// ConfigMapList represents a list of configmaps.
type ConfigMapList struct {
	ConfigMaps []ConfigMap `json:"items"`
}

// Run the validation of the Lagoon YAML dumps.
func (cmd *ValidateConfigMapJSONCmd) Run() error {
	// open the file
	rawJSON, err := os.ReadFile(cmd.ConfigMapJSON)
	if err != nil {
		return fmt.Errorf("couldn't read file: %v", err)
	}
	var cml ConfigMapList
	// unmarshal ConfigMapList
	if err := json.Unmarshal(rawJSON, &cml); err != nil {
		return fmt.Errorf("couldn't unmarshal JSON: %v", err)
	}
	// lint it
	for _, cm := range cml.ConfigMaps {
		if lagoonYAML, ok := cm.Data[".lagoon.yml"]; ok {
			switch cmd.Profile {
			case "required":
				err = required.Lint([]byte(lagoonYAML), required.DefaultLinters)
			case "deprecated":
				err = deprecated.Lint([]byte(lagoonYAML), deprecated.DefaultLinters)
			default:
				return fmt.Errorf("invalid profile: %v", cmd.Profile)
			}
			if err != nil {
				fmt.Printf("bad .lagoon.yml: %s: %v\n", cm.Metadata["namespace"], err)
			}
		}
	}
	return nil
}
