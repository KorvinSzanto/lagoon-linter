package deprecated_test

import (
	"os"
	"testing"

	"github.com/uselagoon/lagoon-linter/internal/lagoonyml/deprecated"
)

func TestLint(t *testing.T) {
	var testCases = map[string]struct {
		input string
		valid bool
	}{
		"valid.0.lagoon.yml": {
			input: "testdata/valid.0.lagoon.yml",
			valid: true,
		},
		"invalid standard monitoring_urls": {
			input: "testdata/invalid.0.lagoon.yml",
			valid: false,
		},
		"invalid weird monitoring_urls": {
			input: "testdata/invalid.1.lagoon.yml",
			valid: false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			rawYAML, err := os.ReadFile(tc.input)
			if err != nil {
				tt.Fatalf("couldn't read %v: %v", tc.input, err)
			}
			err = deprecated.Lint(rawYAML, deprecated.DefaultLinters)
			if tc.valid {
				if err != nil {
					tt.Fatalf("unexpected error %v", err)
				}
			} else {
				tt.Log(err)
				if err == nil {
					tt.Fatalf("expected error, but got nil")
				}
			}
		})
	}
}
