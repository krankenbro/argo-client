/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SessionGetUserInfoResponse struct {
	Groups []string `json:"groups,omitempty"`
	Iss string `json:"iss,omitempty"`
	LoggedIn bool `json:"loggedIn,omitempty"`
	Username string `json:"username,omitempty"`
}
