/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.16
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type Machine struct {
	Id             string            `json:"id,omitempty"`
	Iaas           string            `json:"iaas,omitempty"`
	Address        string            `json:"address,omitempty"`
	Port           int32             `json:"port,omitempty"`
	Protocol       string            `json:"protocol,omitempty"`
	Creationparams map[string]string `json:"creationparams,omitempty"`
}
