/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type V1alpha1RetryStrategy struct {
	Backoff *V1alpha1Backoff `json:"backoff,omitempty"`
	// Limit is the maximum number of attempts for retrying a failed sync. If set to 0, no retries will be performed.
	Limit string `json:"limit,omitempty"`
}