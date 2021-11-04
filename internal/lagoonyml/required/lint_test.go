package required_test

import (
	"os"
	"testing"

	"github.com/uselagoon/lagoon-linter/internal/lagoonyml/required"
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
		"valid.1.lagoon.yml": {
			input: "testdata/valid.1.lagoon.yml",
			valid: true,
		},
		"tls-acme boolean": {
			input: "testdata/valid.2.lagoon.yml",
			valid: true,
		},
		"complete .lagoon.yml": {
			input: "testdata/valid.3.lagoon.yml",
			valid: true,
		},
		"single quoted arguments": {
			input: "testdata/valid.4.lagoon.yml",
			valid: true,
		},
		"add_header ... always": {
			input: "testdata/valid.5.lagoon.yml",
			valid: true,
		},
		"standby route valid annotation": {
			input: "testdata/valid.6.lagoon.yml",
			valid: true,
		},
		"invalid.0.lagoon.yml": {
			input: "testdata/invalid.0.lagoon.yml",
			valid: false,
		},
		"invalid.1.lagoon.yml": {
			input: "testdata/invalid.1.lagoon.yml",
			valid: false,
		},
		"standby route invalid annotation": {
			input: "testdata/invalid.2.lagoon.yml",
			valid: false,
		},
		"active route invalid annotation": {
			input: "testdata/invalid.3.lagoon.yml",
			valid: false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			rawYAML, err := os.ReadFile(tc.input)
			if err != nil {
				tt.Fatalf("couldn't read %v: %v", tc.input, err)
			}
			err = required.Lint(rawYAML, required.DefaultLinters)
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
