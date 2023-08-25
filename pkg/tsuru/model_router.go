/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.16
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type Router struct {
	Name         string            `json:"name,omitempty"`
	Opts         map[string]string `json:"opts,omitempty"`
	Address      string            `json:"address,omitempty"`
	Type         string            `json:"type,omitempty"`
	Status       string            `json:"status,omitempty"`
	StatusDetail string            `json:"status-detail,omitempty"`
}
