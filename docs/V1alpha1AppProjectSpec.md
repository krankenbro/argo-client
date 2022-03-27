# V1alpha1AppProjectSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterResourceBlacklist** | [**[]V1GroupKind**](v1GroupKind.md) |  | [optional] [default to null]
**ClusterResourceWhitelist** | [**[]V1GroupKind**](v1GroupKind.md) |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Destinations** | [**[]V1alpha1ApplicationDestination**](v1alpha1ApplicationDestination.md) |  | [optional] [default to null]
**NamespaceResourceBlacklist** | [**[]V1GroupKind**](v1GroupKind.md) |  | [optional] [default to null]
**NamespaceResourceWhitelist** | [**[]V1GroupKind**](v1GroupKind.md) |  | [optional] [default to null]
**OrphanedResources** | [***V1alpha1OrphanedResourcesMonitorSettings**](v1alpha1OrphanedResourcesMonitorSettings.md) |  | [optional] [default to null]
**Roles** | [**[]V1alpha1ProjectRole**](v1alpha1ProjectRole.md) |  | [optional] [default to null]
**SignatureKeys** | [**[]V1alpha1SignatureKey**](v1alpha1SignatureKey.md) |  | [optional] [default to null]
**SourceRepos** | **[]string** |  | [optional] [default to null]
**SyncWindows** | [**[]V1alpha1SyncWindow**](v1alpha1SyncWindow.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


