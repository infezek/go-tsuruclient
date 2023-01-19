/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.12
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type ServiceList struct {
	Service string `json:"service,omitempty"`
	// [deprecated]
	Instances        []string          `json:"instances,omitempty"`
	Plans            []string          `json:"plans,omitempty"`
	ServiceInstances []ServiceInstance `json:"service_instances,omitempty"`
}
