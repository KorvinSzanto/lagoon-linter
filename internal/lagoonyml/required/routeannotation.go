package required

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	authSnippet          = "nginx.ingress.kubernetes.io/auth-snippet"
	configurationSnippet = "nginx.ingress.kubernetes.io/configuration-snippet"
	modsecuritySnippet   = "nginx.ingress.kubernetes.io/modsecurity-snippet"
	serverSnippet        = "nginx.ingress.kubernetes.io/server-snippet"
)

// validSnippets is the allow-list of snippets that Lagoon will accept.
// Currently these are only valid in server-snippet and configuration-snippet
// annotations.
var validSnippets = regexp.MustCompile(
	`^(rewrite +[^; ]+ +[^; ]+( (last|break|redirect|permanent))?;|` +
		`add_header +([^; ]+|"[^"]+"|'[^']+') +([^; ]+|"[^"]+"|'[^']+')( always)?;|` +
		`set_real_ip_from +[^; ]+;|` +
		`more_set_headers +(-s +("[^"]+"|'[^']+')|-t +("[^"]+"|'[^']+')|("[^"]+"|'[^']+'))+;|` +
		` )+$`)

// validate returns true if the annotations are valid, and false otherwise.
func validate(annotations map[string]string, r *regexp.Regexp,
	annotation string) (string, bool) {
	if ss, ok := annotations[annotation]; ok {
		for _, line := range strings.Split(ss, "\n") {
			line = strings.TrimSpace(line)
			if len(line) > 0 && !r.MatchString(line) {
				return line, false
			}
		}
	}
	return "", true
}

// validateEnvironment returns an error if the annotations on the environment
// are invalid, and nil otherwise.
func validateEnvironment(e *Environment) error {
	for _, routeMap := range e.Routes {
		for rName, lagoonRoutes := range routeMap {
			for _, lagoonRoute := range lagoonRoutes {
				for iName, ingress := range lagoonRoute.Ingresses {
					// auth-snippet
					if _, ok := ingress.Annotations[authSnippet]; ok {
						return fmt.Errorf(
							"invalid %s annotation on route %s, ingress %s: %s",
							authSnippet, rName, iName,
							"this annotation is restricted")
					}
					// configuration-snippet
					if annotation, ok := validate(ingress.Annotations, validSnippets,
						configurationSnippet); !ok {
						return fmt.Errorf(
							"invalid %s annotation on route %s, ingress %s: %s",
							configurationSnippet, rName, iName, annotation)
					}
					// modsecurity-snippet
					if _, ok := ingress.Annotations[modsecuritySnippet]; ok {
						return fmt.Errorf(
							"invalid %s annotation on route %s, ingress %s: %s",
							modsecuritySnippet, rName, iName,
							"this annotation is restricted")
					}
					// server-snippet
					if annotation, ok := validate(ingress.Annotations, validSnippets,
						serverSnippet); !ok {
						return fmt.Errorf(
							"invalid %s annotation on route %s, ingress %s: %s",
							serverSnippet, rName, iName, annotation)
					}
				}
			}
		}
	}
	return nil
}

// RouteAnnotation checks for valid annotations on defined routes.
func RouteAnnotation() Linter {
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
