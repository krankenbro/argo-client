/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type V1alpha1ConfigManagementPlugin struct {
	Generate *V1alpha1Command `json:"generate,omitempty"`
	Init *V1alpha1Command `json:"init,omitempty"`
	LockRepo bool `json:"lockRepo,omitempty"`
	Name string `json:"name,omitempty"`
}
