/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type V1alpha1ConnectionState struct {
	AttemptedAt *V1Time `json:"attemptedAt,omitempty"`
	Message string `json:"message,omitempty"`
	Status string `json:"status,omitempty"`
}
