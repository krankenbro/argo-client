/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type V1alpha1ApplicationCondition struct {
	LastTransitionTime *V1Time `json:"lastTransitionTime,omitempty"`
	Message string `json:"message,omitempty"`
	Type_ string `json:"type,omitempty"`
}
