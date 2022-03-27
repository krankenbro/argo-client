# V1alpha1ApplicationSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | [***V1alpha1ApplicationDestination**](v1alpha1ApplicationDestination.md) |  | [optional] [default to null]
**IgnoreDifferences** | [**[]V1alpha1ResourceIgnoreDifferences**](v1alpha1ResourceIgnoreDifferences.md) |  | [optional] [default to null]
**Info** | [**[]V1alpha1Info**](v1alpha1Info.md) |  | [optional] [default to null]
**Project** | **string** | Project is a reference to the project this application belongs to. The empty string means that application belongs to the &#39;default&#39; project. | [optional] [default to null]
**RevisionHistoryLimit** | **string** | RevisionHistoryLimit limits the number of items kept in the application&#39;s revision history, which is used for informational purposes as well as for rollbacks to previous versions. This should only be changed in exceptional circumstances. Setting to zero will store no history. This will reduce storage used. Increasing will increase the space used to store the history, so we do not recommend increasing it. Default is 10. | [optional] [default to null]
**Source** | [***V1alpha1ApplicationSource**](v1alpha1ApplicationSource.md) |  | [optional] [default to null]
**SyncPolicy** | [***V1alpha1SyncPolicy**](v1alpha1SyncPolicy.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


