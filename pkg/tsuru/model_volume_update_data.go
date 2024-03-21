/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.20
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// Volume
type VolumeUpdateData struct {
	// Volume name.
	Name string `json:"name,omitempty"`
	// Volume pool.
	Pool string `json:"pool,omitempty"`
	// Team that owns the volume.
	TeamOwner string `json:"teamOwner,omitempty"`
	// Volume status.
	Status string     `json:"status,omitempty"`
	Plan   VolumePlan `json:"plan,omitempty"`
	// Custom volume options.
	Opts map[string]string `json:"opts,omitempty"`
}
