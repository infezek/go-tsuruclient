/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package tsuru

type ServiceInstance struct {
	Name string `json:"name,omitempty"`

	Tags []string `json:"tags,omitempty"`

	Id int32 `json:"id,omitempty"`

	ServiceName string `json:"service_name,omitempty"`

	PlanName string `json:"plan_name,omitempty"`

	Apps []string `json:"apps,omitempty"`

	BoundUnits []ServiceInstanceBoundUnit `json:"bound_units,omitempty"`

	Teams []string `json:"teams,omitempty"`

	TeamOwner string `json:"team_owner,omitempty"`

	Description string `json:"description,omitempty"`
}
