/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type StreamResultOfApplicationLogEntry struct {
	Error_ *RuntimeStreamError `json:"error,omitempty"`
	Result *ApplicationLogEntry `json:"result,omitempty"`
}