/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.14
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type VolumeBind struct {
	Id VolumeBindId `json:"id,omitempty"`
	// Volume is read-only.
	Readonly bool `json:"readonly,omitempty"`
}
