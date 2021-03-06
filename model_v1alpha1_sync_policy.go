/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type V1alpha1SyncPolicy struct {
	Automated *V1alpha1SyncPolicyAutomated `json:"automated,omitempty"`
	Retry *V1alpha1RetryStrategy `json:"retry,omitempty"`
	SyncOptions []string `json:"syncOptions,omitempty"`
}
