/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.12
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// App plan.
type Plan struct {
	Name     string       `json:"name,omitempty"`
	Memory   int64        `json:"memory,omitempty"`
	Swap     int64        `json:"swap,omitempty"`
	Cpushare int32        `json:"cpushare,omitempty"`
	Cpumilli int32        `json:"cpumilli,omitempty"`
	Default  bool         `json:"default,omitempty"`
	Router   string       `json:"router,omitempty"`
	Override PlanOverride `json:"override,omitempty"`
}
