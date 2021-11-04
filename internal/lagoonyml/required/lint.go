package required

import (
	"fmt"

	"github.com/uselagoon/lagoon-linter/internal/lagoonyml"
	"sigs.k8s.io/yaml"
)

// DefaultLinters is the list of default linters for this profile.
var DefaultLinters = []Linter{
	RouteAnnotation(),
}

// Linter validates the given Lagoon struct.
type Linter func(*Lagoon) error

// Lint takes a byte slice containing raw YAML and applies `.lagoon.yml` lint
// policy to it. Lint returns an error of type ErrLint if it finds problems
// with the file, a regular error if something else went wrong, and nil if the
// `.lagoon.yml` is valid.
func Lint(rawYAML []byte, linters []Linter) error {
	var l Lagoon
	if err := yaml.Unmarshal(rawYAML, &l); err != nil {
		return fmt.Errorf("couldn't unmarshal YAML: %v", err)
	}
	for _, linter := range linters {
		if err := linter(&l); err != nil {
			return &lagoonyml.ErrLint{
				Detail: err.Error(),
			}
		}
	}
	return nil
}
