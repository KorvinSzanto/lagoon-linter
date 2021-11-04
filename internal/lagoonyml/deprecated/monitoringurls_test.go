package deprecated

import (
	"testing"
)

func TestMonitoringURLS(t *testing.T) {
	var testCases = map[string]struct {
		input interface{}
		valid bool
	}{
		"valid monitoring_urls absent": {
			input: nil,
			valid: true,
		},
		"invalid monitoring_urls present": {
			input: []string{"example.com", "www.example.com"},
			valid: false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			l := Lagoon{
				Environments: map[string]Environment{
					"testenv": {
						MonitoringURLs: tc.input,
					},
				},
			}
			err := MonitoringURLs()(&l)
			if tc.valid {
				if err != nil {
					tt.Fatalf("unexpected error %v", err)
				}
			} else {
				if err == nil {
					tt.Fatalf("expected error, but got nil")
				}
			}
		})
	}
}
