/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type V1alpha1TlsClientConfig struct {
	CaData string `json:"caData,omitempty"`
	CertData string `json:"certData,omitempty"`
	// Insecure specifies that the server should be accessed without verifying the TLS certificate. For testing only.
	Insecure bool `json:"insecure,omitempty"`
	KeyData string `json:"keyData,omitempty"`
	// ServerName is passed to the server for SNI and is used in the client to check server certificates against. If ServerName is empty, the hostname used to contact the server is used.
	ServerName string `json:"serverName,omitempty"`
}
