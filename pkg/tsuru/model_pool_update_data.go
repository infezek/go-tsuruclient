/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.14
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type PoolUpdateData struct {
	Default bool              `json:"default,omitempty"`
	Public  bool              `json:"public,omitempty"`
	Force   bool              `json:"force,omitempty"`
	Labels  map[string]string `json:"labels,omitempty"`
}
