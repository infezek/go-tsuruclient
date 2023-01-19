/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.12
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// Plan Router
type PlanRouter struct {
	Name           string                 `json:"name,omitempty"`
	Type           string                 `json:"type,omitempty"`
	Config         map[string]interface{} `json:"config,omitempty"`
	Info           map[string]string      `json:"info,omitempty"`
	Default        bool                   `json:"default,omitempty"`
	Dynamic        bool                   `json:"dynamic,omitempty"`
	ReadinessGates []string               `json:"readinessGates,omitempty"`
}
