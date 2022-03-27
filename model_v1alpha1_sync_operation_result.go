/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type V1alpha1SyncOperationResult struct {
	Resources []V1alpha1ResourceResult `json:"resources,omitempty"`
	Revision string `json:"revision,omitempty"`
	Source *V1alpha1ApplicationSource `json:"source,omitempty"`
}
