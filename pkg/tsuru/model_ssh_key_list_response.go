/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.14
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

// Response returned by User SSH Key List.
type SshKeyListResponse struct {
	// keyvalue
	Keyname string `json:"keyname,omitempty"`
}
