/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type V1alpha1AppProject struct {
	Metadata *V1ObjectMeta `json:"metadata,omitempty"`
	Spec *V1alpha1AppProjectSpec `json:"spec,omitempty"`
	Status *V1alpha1AppProjectStatus `json:"status,omitempty"`
}
