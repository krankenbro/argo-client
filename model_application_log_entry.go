/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApplicationLogEntry struct {
	Content string `json:"content,omitempty"`
	Last bool `json:"last,omitempty"`
	PodName string `json:"podName,omitempty"`
	TimeStamp *V1Time `json:"timeStamp,omitempty"`
	TimeStampStr string `json:"timeStampStr,omitempty"`
}
