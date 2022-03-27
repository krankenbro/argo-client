/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type V1alpha1OperationState struct {
	FinishedAt *V1Time `json:"finishedAt,omitempty"`
	// Message holds any pertinent messages when attempting to perform operation (typically errors).
	Message string `json:"message,omitempty"`
	Operation *V1alpha1Operation `json:"operation,omitempty"`
	Phase string `json:"phase,omitempty"`
	RetryCount string `json:"retryCount,omitempty"`
	StartedAt *V1Time `json:"startedAt,omitempty"`
	SyncResult *V1alpha1SyncOperationResult `json:"syncResult,omitempty"`
}
