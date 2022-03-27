/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type V1alpha1ApplicationSource struct {
	// Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo.
	Chart string `json:"chart,omitempty"`
	Directory *V1alpha1ApplicationSourceDirectory `json:"directory,omitempty"`
	Helm *V1alpha1ApplicationSourceHelm `json:"helm,omitempty"`
	Ksonnet *V1alpha1ApplicationSourceKsonnet `json:"ksonnet,omitempty"`
	Kustomize *V1alpha1ApplicationSourceKustomize `json:"kustomize,omitempty"`
	// Path is a directory path within the Git repository, and is only valid for applications sourced from Git.
	Path string `json:"path,omitempty"`
	Plugin *V1alpha1ApplicationSourcePlugin `json:"plugin,omitempty"`
	RepoURL string `json:"repoURL,omitempty"`
	// TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart's version.
	TargetRevision string `json:"targetRevision,omitempty"`
}
