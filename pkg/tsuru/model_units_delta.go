/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// number of units to add or remove from a process
type UnitsDelta struct {
	Units   string `json:"units,omitempty"`
	Process string `json:"process,omitempty"`
	Version string `json:"version,omitempty"`
}
