package main

import (
	"fmt"
	"os"

	"github.com/uselagoon/lagoon-linter/internal/lagoonyml/deprecated"
	"github.com/uselagoon/lagoon-linter/internal/lagoonyml/required"
)

// ValidateCmd represents the validate command.
type ValidateCmd struct {
	LagoonYAML string `kong:"default='.lagoon.yml',help='Specify the Lagoon YAML file.'"`
	Profile    string `kong:"default='required',enum='required,deprecated',help='Set the linting profile (required,deprecated)'"`
}

// Run the validation of the Lagoon YAML.
func (cmd *ValidateCmd) Run() error {
	rawYAML, err := os.ReadFile(cmd.LagoonYAML)
	if err != nil {
		return fmt.Errorf("couldn't read %v: %v", cmd.LagoonYAML, err)
	}
	switch cmd.Profile {
	case "required":
		return required.Lint(rawYAML, required.DefaultLinters)
	case "deprecated":
		return deprecated.Lint(rawYAML, deprecated.DefaultLinters)
	default:
		return fmt.Errorf("invalid profile: %v", cmd.Profile)
	}
}
