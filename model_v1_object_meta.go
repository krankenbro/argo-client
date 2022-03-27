/*
 * Consolidate Services
 *
 * Description of all APIs
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// ObjectMeta is metadata that all persisted resources must have, which includes all objects users must create.
type V1ObjectMeta struct {
	Annotations map[string]string `json:"annotations,omitempty"`
	ClusterName string `json:"clusterName,omitempty"`
	CreationTimestamp *V1Time `json:"creationTimestamp,omitempty"`
	DeletionGracePeriodSeconds string `json:"deletionGracePeriodSeconds,omitempty"`
	DeletionTimestamp *V1Time `json:"deletionTimestamp,omitempty"`
	Finalizers []string `json:"finalizers,omitempty"`
	// GenerateName is an optional prefix, used by the server, to generate a unique name ONLY IF the Name field has not been provided. If this field is used, the name returned to the client will be different than the name passed. This value will also be combined with a unique suffix. The provided value has the same validation rules as the Name field, and may be truncated by the length of the suffix required to make the value unique on the server.  If this field is specified and the generated name exists, the server will NOT return a 409 - instead, it will either return 201 Created or 500 with Reason ServerTimeout indicating a unique name could not be found in the time allotted, and the client should retry (optionally after the time indicated in the Retry-After header).  Applied only if Name is not specified. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#idempotency +optional
	GenerateName string `json:"generateName,omitempty"`
	Generation string `json:"generation,omitempty"`
	Labels map[string]string `json:"labels,omitempty"`
	// ManagedFields maps workflow-id and version to the set of fields that are managed by that workflow. This is mostly for internal housekeeping, and users typically shouldn't need to set or understand this field. A workflow can be the user's name, a controller's name, or the name of a specific apply path like \"ci-cd\". The set of fields is always in the version that the workflow used when modifying the object.  +optional
	ManagedFields []V1ManagedFieldsEntry `json:"managedFields,omitempty"`
	Name string `json:"name,omitempty"`
	// Namespace defines the space within which each name must be unique. An empty namespace is equivalent to the \"default\" namespace, but \"default\" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.  Must be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces +optional
	Namespace string `json:"namespace,omitempty"`
	OwnerReferences []V1OwnerReference `json:"ownerReferences,omitempty"`
	// An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. Clients must treat these values as opaque and passed unmodified back to the server. They may only be valid for a particular resource or set of resources.  Populated by the system. Read-only. Value must be treated as opaque by clients and . More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency +optional
	ResourceVersion string `json:"resourceVersion,omitempty"`
	// SelfLink is a URL representing this object. Populated by the system. Read-only.  DEPRECATED Kubernetes will stop propagating this field in 1.20 release and the field is planned to be removed in 1.21 release. +optional
	SelfLink string `json:"selfLink,omitempty"`
	// UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.  Populated by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids +optional
	Uid string `json:"uid,omitempty"`
}