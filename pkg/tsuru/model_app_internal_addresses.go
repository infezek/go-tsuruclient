/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.14
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type AppInternalAddresses struct {
	Version  string `json:"Version,omitempty"`
	Port     int64  `json:"Port,omitempty"`
	Process  string `json:"Process,omitempty"`
	Domain   string `json:"Domain,omitempty"`
	Protocol string `json:"Protocol,omitempty"`
}
