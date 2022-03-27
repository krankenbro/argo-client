/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ClusterOidcConfig struct {
	CliClientID string `json:"cliClientID,omitempty"`
	ClientID string `json:"clientID,omitempty"`
	IdTokenClaims map[string]OidcClaim `json:"idTokenClaims,omitempty"`
	Issuer string `json:"issuer,omitempty"`
	Name string `json:"name,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
}
