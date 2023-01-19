/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.12
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// Data sent to set up a node container
type NodeContainer struct {
	Name   string              `json:"name"`
	Pool   string              `json:"pool,omitempty"`
	Config NodeContainerConfig `json:"config,omitempty"`
}
