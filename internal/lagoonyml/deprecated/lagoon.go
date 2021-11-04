package deprecated

// Environment represents a Lagoon environment.
type Environment struct {
	MonitoringURLs interface{} `json:"monitoring_urls"`
}

// ProductionRoutes represents an active/standby configuration.
type ProductionRoutes struct {
	Active  *Environment `json:"active"`
	Standby *Environment `json:"standby"`
}

// Lagoon represents the .lagoon.yml file.
type Lagoon struct {
	Environments     map[string]Environment `json:"environments"`
	ProductionRoutes *ProductionRoutes      `json:"production_routes"`
}
