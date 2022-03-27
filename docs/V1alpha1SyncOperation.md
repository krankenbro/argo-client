# V1alpha1SyncOperation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | **bool** |  | [optional] [default to null]
**Manifests** | **[]string** |  | [optional] [default to null]
**Prune** | **bool** |  | [optional] [default to null]
**Resources** | [**[]V1alpha1SyncOperationResource**](v1alpha1SyncOperationResource.md) |  | [optional] [default to null]
**Revision** | **string** | Revision is the revision (Git) or chart version (Helm) which to sync the application to If omitted, will use the revision specified in app spec. | [optional] [default to null]
**Source** | [***V1alpha1ApplicationSource**](v1alpha1ApplicationSource.md) |  | [optional] [default to null]
**SyncOptions** | **[]string** |  | [optional] [default to null]
**SyncStrategy** | [***V1alpha1SyncStrategy**](v1alpha1SyncStrategy.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


