/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type V1alpha1ClusterInfo struct {
	ApiVersions []string `json:"apiVersions,omitempty"`
	ApplicationsCount string `json:"applicationsCount,omitempty"`
	CacheInfo *V1alpha1ClusterCacheInfo `json:"cacheInfo,omitempty"`
	ConnectionState *V1alpha1ConnectionState `json:"connectionState,omitempty"`
	ServerVersion string `json:"serverVersion,omitempty"`
}
