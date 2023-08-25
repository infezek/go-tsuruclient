/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.16
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type Lock struct {
	Locked      bool   `json:"locked,omitempty"`
	Reason      string `json:"reason,omitempty"`
	Owner       string `json:"owner,omitempty"`
	AcquireDate string `json:"acquireDate,omitempty"`
}
