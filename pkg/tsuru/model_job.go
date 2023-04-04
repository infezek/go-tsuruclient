/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.14
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type Job struct {
	// Job name.
	Name string `json:"name,omitempty"`
	// job pool name.
	Pool string `json:"pool,omitempty"`
	// job description.
	Description string `json:"description,omitempty"`
	// who created this job.
	Owner string `json:"owner,omitempty"`
	// Team that owns the job.
	TeamOwner string `json:"teamOwner,omitempty"`
	// Teams that have access to this job
	Teams    []string `json:"teams,omitempty"`
	Plan     Plan     `json:"plan,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
	Spec     JobSpec  `json:"spec,omitempty"`
}
