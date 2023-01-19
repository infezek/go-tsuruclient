/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.12
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package tsuru

type ClusterKubeConfigCluster struct {
	Server                   string `json:"server"`
	CertificateAuthorityData string `json:"certificate-authority-data,omitempty"`
	TlsServerName            string `json:"tls-server-name,omitempty"`
	InsecureSkipTlsVerify    bool   `json:"insecure-skip-tls-verify,omitempty"`
}
