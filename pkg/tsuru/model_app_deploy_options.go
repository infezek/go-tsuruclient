/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type AppDeployOptions struct {
	Image            string `json:"image"`
	Message          string `json:"message,omitempty"`
	NewVersion       bool   `json:"new-version,omitempty"`
	OverrideVersions bool   `json:"override-versions,omitempty"`
}
