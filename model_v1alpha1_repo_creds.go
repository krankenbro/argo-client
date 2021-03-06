/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type V1alpha1RepoCreds struct {
	EnableOCI bool `json:"enableOCI,omitempty"`
	GithubAppEnterpriseBaseUrl string `json:"githubAppEnterpriseBaseUrl,omitempty"`
	GithubAppID string `json:"githubAppID,omitempty"`
	GithubAppInstallationID string `json:"githubAppInstallationID,omitempty"`
	GithubAppPrivateKey string `json:"githubAppPrivateKey,omitempty"`
	Password string `json:"password,omitempty"`
	SshPrivateKey string `json:"sshPrivateKey,omitempty"`
	TlsClientCertData string `json:"tlsClientCertData,omitempty"`
	TlsClientCertKey string `json:"tlsClientCertKey,omitempty"`
	// Type specifies the type of the repoCreds. Can be either \"git\" or \"helm. \"git\" is assumed if empty or absent.
	Type_ string `json:"type,omitempty"`
	Url string `json:"url,omitempty"`
	Username string `json:"username,omitempty"`
}
