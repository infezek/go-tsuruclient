/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.14
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type Platform struct {
	Name     string `json:"name,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
}
