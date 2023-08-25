/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.16
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type WebhookEventFilter struct {
	TargetTypes  []string `json:"target_types,omitempty"`
	TargetValues []string `json:"target_values,omitempty"`
	KindTypes    []string `json:"kind_types,omitempty"`
	KindNames    []string `json:"kind_names,omitempty"`
	ErrorOnly    bool     `json:"error_only,omitempty"`
	SuccessOnly  bool     `json:"success_only,omitempty"`
}
