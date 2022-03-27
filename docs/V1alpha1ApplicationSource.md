# V1alpha1ApplicationSource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chart** | **string** | Chart is a Helm chart name, and must be specified for applications sourced from a Helm repo. | [optional] [default to null]
**Directory** | [***V1alpha1ApplicationSourceDirectory**](v1alpha1ApplicationSourceDirectory.md) |  | [optional] [default to null]
**Helm** | [***V1alpha1ApplicationSourceHelm**](v1alpha1ApplicationSourceHelm.md) |  | [optional] [default to null]
**Ksonnet** | [***V1alpha1ApplicationSourceKsonnet**](v1alpha1ApplicationSourceKsonnet.md) |  | [optional] [default to null]
**Kustomize** | [***V1alpha1ApplicationSourceKustomize**](v1alpha1ApplicationSourceKustomize.md) |  | [optional] [default to null]
**Path** | **string** | Path is a directory path within the Git repository, and is only valid for applications sourced from Git. | [optional] [default to null]
**Plugin** | [***V1alpha1ApplicationSourcePlugin**](v1alpha1ApplicationSourcePlugin.md) |  | [optional] [default to null]
**RepoURL** | **string** |  | [optional] [default to null]
**TargetRevision** | **string** | TargetRevision defines the revision of the source to sync the application to. In case of Git, this can be commit, tag, or branch. If omitted, will equal to HEAD. In case of Helm, this is a semver tag for the Chart&#39;s version. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


