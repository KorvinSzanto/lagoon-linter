package lagoonyml

import (
	"fmt"
	"os"

	"sigs.k8s.io/yaml"
)

// Linter validates the given Lagoon struct.
type Linter func(*Lagoon) error

// LintFile takes a file path, reads it, and applies `.lagoon.yml` lint policy to
// it. Lint returns an error of type ErrLint if it finds problems with the
// file, a regular error if something else went wrong, and nil if the
// `.lagoon.yml` is valid.
func LintFile(path string, linters ...Linter) error {
	var l Lagoon
	rawYAML, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("couldn't read %v: %v", path, err)
	}
	err = yaml.Unmarshal(rawYAML, &l)
	if err != nil {
		return fmt.Errorf("couldn't unmarshal %v: %v", path, err)
	}
	for _, linter := range linters {
		if err := linter(&l); err != nil {
			return &ErrLint{
				Detail: err.Error(),
			}
		}
	}
	return nil
}

// LintYAML takes a byte slice containing raw YAML and applies `.lagoon.yml` lint policy to
// it. Lint returns an error of type ErrLint if it finds problems with the
// file, a regular error if something else went wrong, and nil if the
// `.lagoon.yml` is valid.
func LintYAML(rawYAML []byte, linters ...Linter) error {
	var l Lagoon
	if err := yaml.Unmarshal(rawYAML, &l); err != nil {
		return fmt.Errorf("couldn't unmarshal YAML: %v", err)
	}
	for _, linter := range linters {
		if err := linter(&l); err != nil {
			return &ErrLint{
				Detail: err.Error(),
			}
		}
	}
	return nil
}
