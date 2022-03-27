/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type RepositoryRepoAppDetailsResponse struct {
	Directory *RepositoryDirectoryAppSpec `json:"directory,omitempty"`
	Helm *RepositoryHelmAppSpec `json:"helm,omitempty"`
	Ksonnet *RepositoryKsonnetAppSpec `json:"ksonnet,omitempty"`
	Kustomize *RepositoryKustomizeAppSpec `json:"kustomize,omitempty"`
	Type_ string `json:"type,omitempty"`
}