/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.14
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type ClusterKubeConfigUserAuthprovider struct {
	Name   string            `json:"name,omitempty"`
	Config map[string]string `json:"config,omitempty"`
}
