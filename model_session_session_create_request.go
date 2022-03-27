/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// SessionCreateRequest is for logging in.
type SessionSessionCreateRequest struct {
	Password string `json:"password,omitempty"`
	Token string `json:"token,omitempty"`
	Username string `json:"username,omitempty"`
}