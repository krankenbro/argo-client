/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// ApplicationSpec represents desired application state. Contains link to repository with application definition and additional parameters link definition revision.
type V1alpha1ApplicationSpec struct {
	Destination *V1alpha1ApplicationDestination `json:"destination,omitempty"`
	IgnoreDifferences []V1alpha1ResourceIgnoreDifferences `json:"ignoreDifferences,omitempty"`
	Info []V1alpha1Info `json:"info,omitempty"`
	// Project is a reference to the project this application belongs to. The empty string means that application belongs to the 'default' project.
	Project string `json:"project,omitempty"`
	// RevisionHistoryLimit limits the number of items kept in the application's revision history, which is used for informational purposes as well as for rollbacks to previous versions. This should only be changed in exceptional circumstances. Setting to zero will store no history. This will reduce storage used. Increasing will increase the space used to store the history, so we do not recommend increasing it. Default is 10.
	RevisionHistoryLimit string `json:"revisionHistoryLimit,omitempty"`
	Source *V1alpha1ApplicationSource `json:"source,omitempty"`
	SyncPolicy *V1alpha1SyncPolicy `json:"syncPolicy,omitempty"`
}
