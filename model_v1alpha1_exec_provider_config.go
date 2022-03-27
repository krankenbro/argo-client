/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type V1alpha1ExecProviderConfig struct {
	ApiVersion string `json:"apiVersion,omitempty"`
	Args []string `json:"args,omitempty"`
	Command string `json:"command,omitempty"`
	Env map[string]string `json:"env,omitempty"`
	InstallHint string `json:"installHint,omitempty"`
}