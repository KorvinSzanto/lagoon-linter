package deprecated

import (
	"fmt"
)

// validateEnvironment returns an error if the annotations on the environment
// are invalid, and nil otherwise.
func validateEnvironment(e *Environment) error {
	if e.MonitoringURLs != nil {
		return fmt.Errorf("deprecated monitoring_urls directive")
	}
	return nil
}

// MonitoringURLs checks for deprecated monitoring URLs on a Lagoon environment.
func MonitoringURLs() Linter {
	return func(l *Lagoon) error {
		for eName, e := range l.Environments {
			if err := validateEnvironment(&e); err != nil {
				return fmt.Errorf("environment %s: %v", eName, err)
			}
		}
		if l.ProductionRoutes != nil {
			if l.ProductionRoutes.Active != nil {
				if err := validateEnvironment(l.ProductionRoutes.Active); err != nil {
					return fmt.Errorf("active environment: %v", err)
				}
			}
			if l.ProductionRoutes.Standby != nil {
				if err := validateEnvironment(l.ProductionRoutes.Standby); err != nil {
					return fmt.Errorf("standby environment: %v", err)
				}
			}
		}
		return nil
	}
}
