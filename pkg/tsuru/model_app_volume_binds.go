/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.20
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type AppVolumeBinds struct {
	ReadOnly bool  `json:"ReadOnly,omitempty"`
	ID       AppId `json:"ID,omitempty"`
}
