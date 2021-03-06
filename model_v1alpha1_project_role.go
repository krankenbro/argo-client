/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type V1alpha1ProjectRole struct {
	Description string `json:"description,omitempty"`
	Groups []string `json:"groups,omitempty"`
	JwtTokens []V1alpha1JwtToken `json:"jwtTokens,omitempty"`
	Name string `json:"name,omitempty"`
	Policies []string `json:"policies,omitempty"`
}
