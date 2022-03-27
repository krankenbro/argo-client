# \ApplicationServiceApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationServiceCreate**](ApplicationServiceApi.md#ApplicationServiceCreate) | **Post** /api/v1/applications | Create creates an application
[**ApplicationServiceDelete**](ApplicationServiceApi.md#ApplicationServiceDelete) | **Delete** /api/v1/applications/{name} | Delete deletes an application
[**ApplicationServiceDeleteResource**](ApplicationServiceApi.md#ApplicationServiceDeleteResource) | **Delete** /api/v1/applications/{name}/resource | DeleteResource deletes a single application resource
[**ApplicationServiceGet**](ApplicationServiceApi.md#ApplicationServiceGet) | **Get** /api/v1/applications/{name} | Get returns an application by name
[**ApplicationServiceGetApplicationSyncWindows**](ApplicationServiceApi.md#ApplicationServiceGetApplicationSyncWindows) | **Get** /api/v1/applications/{name}/syncwindows | Get returns sync windows of the application
[**ApplicationServiceGetManifests**](ApplicationServiceApi.md#ApplicationServiceGetManifests) | **Get** /api/v1/applications/{name}/manifests | GetManifests returns application manifests
[**ApplicationServiceGetResource**](ApplicationServiceApi.md#ApplicationServiceGetResource) | **Get** /api/v1/applications/{name}/resource | GetResource returns single application resource
[**ApplicationServiceList**](ApplicationServiceApi.md#ApplicationServiceList) | **Get** /api/v1/applications | List returns list of applications
[**ApplicationServiceListResourceActions**](ApplicationServiceApi.md#ApplicationServiceListResourceActions) | **Get** /api/v1/applications/{name}/resource/actions | ListResourceActions returns list of resource actions
[**ApplicationServiceListResourceEvents**](ApplicationServiceApi.md#ApplicationServiceListResourceEvents) | **Get** /api/v1/applications/{name}/events | ListResourceEvents returns a list of event resources
[**ApplicationServiceManagedResources**](ApplicationServiceApi.md#ApplicationServiceManagedResources) | **Get** /api/v1/applications/{applicationName}/managed-resources | ManagedResources returns list of managed resources
[**ApplicationServicePatch**](ApplicationServiceApi.md#ApplicationServicePatch) | **Patch** /api/v1/applications/{name} | Patch patch an application
[**ApplicationServicePatchResource**](ApplicationServiceApi.md#ApplicationServicePatchResource) | **Post** /api/v1/applications/{name}/resource | PatchResource patch single application resource
[**ApplicationServicePodLogs**](ApplicationServiceApi.md#ApplicationServicePodLogs) | **Get** /api/v1/applications/{name}/pods/{podName}/logs | PodLogs returns stream of log entries for the specified pod. Pod
[**ApplicationServicePodLogs2**](ApplicationServiceApi.md#ApplicationServicePodLogs2) | **Get** /api/v1/applications/{name}/logs | PodLogs returns stream of log entries for the specified pod. Pod
[**ApplicationServiceResourceTree**](ApplicationServiceApi.md#ApplicationServiceResourceTree) | **Get** /api/v1/applications/{applicationName}/resource-tree | ResourceTree returns resource tree
[**ApplicationServiceRevisionMetadata**](ApplicationServiceApi.md#ApplicationServiceRevisionMetadata) | **Get** /api/v1/applications/{name}/revisions/{revision}/metadata | Get the meta-data (author, date, tags, message) for a specific revision of the application
[**ApplicationServiceRollback**](ApplicationServiceApi.md#ApplicationServiceRollback) | **Post** /api/v1/applications/{name}/rollback | Rollback syncs an application to its target state
[**ApplicationServiceRunResourceAction**](ApplicationServiceApi.md#ApplicationServiceRunResourceAction) | **Post** /api/v1/applications/{name}/resource/actions | RunResourceAction run resource action
[**ApplicationServiceSync**](ApplicationServiceApi.md#ApplicationServiceSync) | **Post** /api/v1/applications/{name}/sync | Sync syncs an application to its target state
[**ApplicationServiceTerminateOperation**](ApplicationServiceApi.md#ApplicationServiceTerminateOperation) | **Delete** /api/v1/applications/{name}/operation | TerminateOperation terminates the currently running operation
[**ApplicationServiceUpdate**](ApplicationServiceApi.md#ApplicationServiceUpdate) | **Put** /api/v1/applications/{application.metadata.name} | Update updates an application
[**ApplicationServiceUpdateSpec**](ApplicationServiceApi.md#ApplicationServiceUpdateSpec) | **Put** /api/v1/applications/{name}/spec | UpdateSpec updates an application spec
[**ApplicationServiceWatch**](ApplicationServiceApi.md#ApplicationServiceWatch) | **Get** /api/v1/stream/applications | Watch returns stream of application change events
[**ApplicationServiceWatchResourceTree**](ApplicationServiceApi.md#ApplicationServiceWatchResourceTree) | **Get** /api/v1/stream/applications/{applicationName}/resource-tree | Watch returns stream of application resource tree


# **ApplicationServiceCreate**
> V1alpha1Application ApplicationServiceCreate(ctx, body, optional)
Create creates an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1alpha1Application**](V1alpha1Application.md)|  | 
 **optional** | ***ApplicationServiceApiApplicationServiceCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServiceCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsert** | **optional.Bool**|  | 
 **validate** | **optional.Bool**|  | 

### Return type

[**V1alpha1Application**](v1alpha1Application.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceDelete**
> ApplicationApplicationResponse ApplicationServiceDelete(ctx, name, optional)
Delete deletes an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
 **optional** | ***ApplicationServiceApiApplicationServiceDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServiceDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascade** | **optional.Bool**|  | 
 **propagationPolicy** | **optional.String**|  | 

### Return type

[**ApplicationApplicationResponse**](applicationApplicationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceDeleteResource**
> ApplicationApplicationResponse ApplicationServiceDeleteResource(ctx, name, optional)
DeleteResource deletes a single application resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
 **optional** | ***ApplicationServiceApiApplicationServiceDeleteResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServiceDeleteResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **optional.String**|  | 
 **resourceName** | **optional.String**|  | 
 **version** | **optional.String**|  | 
 **group** | **optional.String**|  | 
 **kind** | **optional.String**|  | 
 **force** | **optional.Bool**|  | 
 **orphan** | **optional.Bool**|  | 

### Return type

[**ApplicationApplicationResponse**](applicationApplicationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceGet**
> V1alpha1Application ApplicationServiceGet(ctx, name, optional)
Get returns an application by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the application&#39;s name | 
 **optional** | ***ApplicationServiceApiApplicationServiceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServiceGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refresh** | **optional.String**| forces application reconciliation if set to true. | 
 **project** | [**optional.Interface of []string**](string.md)| the project names to restrict returned list applications. | 
 **resourceVersion** | **optional.String**| when specified with a watch call, shows changes that occur after that particular version of a resource. | 
 **selector** | **optional.String**| the selector to to restrict returned list to applications only with matched labels. | 
 **repo** | **optional.String**| the repoURL to restrict returned list applications. | 

### Return type

[**V1alpha1Application**](v1alpha1Application.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceGetApplicationSyncWindows**
> ApplicationApplicationSyncWindowsResponse ApplicationServiceGetApplicationSyncWindows(ctx, name)
Get returns sync windows of the application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**ApplicationApplicationSyncWindowsResponse**](applicationApplicationSyncWindowsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceGetManifests**
> RepositoryManifestResponse ApplicationServiceGetManifests(ctx, name, optional)
GetManifests returns application manifests

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
 **optional** | ***ApplicationServiceApiApplicationServiceGetManifestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServiceGetManifestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **optional.String**|  | 

### Return type

[**RepositoryManifestResponse**](repositoryManifestResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceGetResource**
> ApplicationApplicationResourceResponse ApplicationServiceGetResource(ctx, name, optional)
GetResource returns single application resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
 **optional** | ***ApplicationServiceApiApplicationServiceGetResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServiceGetResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **optional.String**|  | 
 **resourceName** | **optional.String**|  | 
 **version** | **optional.String**|  | 
 **group** | **optional.String**|  | 
 **kind** | **optional.String**|  | 

### Return type

[**ApplicationApplicationResourceResponse**](applicationApplicationResourceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceList**
> V1alpha1ApplicationList ApplicationServiceList(ctx, optional)
List returns list of applications

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationServiceApiApplicationServiceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServiceListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| the application&#39;s name. | 
 **refresh** | **optional.String**| forces application reconciliation if set to true. | 
 **project** | [**optional.Interface of []string**](string.md)| the project names to restrict returned list applications. | 
 **resourceVersion** | **optional.String**| when specified with a watch call, shows changes that occur after that particular version of a resource. | 
 **selector** | **optional.String**| the selector to to restrict returned list to applications only with matched labels. | 
 **repo** | **optional.String**| the repoURL to restrict returned list applications. | 

### Return type

[**V1alpha1ApplicationList**](v1alpha1ApplicationList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceListResourceActions**
> ApplicationResourceActionsListResponse ApplicationServiceListResourceActions(ctx, name, optional)
ListResourceActions returns list of resource actions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
 **optional** | ***ApplicationServiceApiApplicationServiceListResourceActionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServiceListResourceActionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **optional.String**|  | 
 **resourceName** | **optional.String**|  | 
 **version** | **optional.String**|  | 
 **group** | **optional.String**|  | 
 **kind** | **optional.String**|  | 

### Return type

[**ApplicationResourceActionsListResponse**](applicationResourceActionsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceListResourceEvents**
> V1EventList ApplicationServiceListResourceEvents(ctx, name, optional)
ListResourceEvents returns a list of event resources

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
 **optional** | ***ApplicationServiceApiApplicationServiceListResourceEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServiceListResourceEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceNamespace** | **optional.String**|  | 
 **resourceName** | **optional.String**|  | 
 **resourceUID** | **optional.String**|  | 

### Return type

[**V1EventList**](v1EventList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceManagedResources**
> ApplicationManagedResourcesResponse ApplicationServiceManagedResources(ctx, applicationName, optional)
ManagedResources returns list of managed resources

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationName** | **string**|  | 
 **optional** | ***ApplicationServiceApiApplicationServiceManagedResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServiceManagedResourcesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **version** | **optional.String**|  | 
 **group** | **optional.String**|  | 
 **kind** | **optional.String**|  | 

### Return type

[**ApplicationManagedResourcesResponse**](applicationManagedResourcesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServicePatch**
> V1alpha1Application ApplicationServicePatch(ctx, name, body)
Patch patch an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **body** | [**ApplicationApplicationPatchRequest**](ApplicationApplicationPatchRequest.md)|  | 

### Return type

[**V1alpha1Application**](v1alpha1Application.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServicePatchResource**
> ApplicationApplicationResourceResponse ApplicationServicePatchResource(ctx, name, body, optional)
PatchResource patch single application resource

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **body** | **string**|  | 
 **optional** | ***ApplicationServiceApiApplicationServicePatchResourceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServicePatchResourceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **namespace** | **optional.String**|  | 
 **resourceName** | **optional.String**|  | 
 **version** | **optional.String**|  | 
 **group** | **optional.String**|  | 
 **kind** | **optional.String**|  | 
 **patchType** | **optional.String**|  | 

### Return type

[**ApplicationApplicationResourceResponse**](applicationApplicationResourceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServicePodLogs**
> StreamResultOfApplicationLogEntry ApplicationServicePodLogs(ctx, name, podName, optional)
PodLogs returns stream of log entries for the specified pod. Pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **podName** | **string**|  | 
 **optional** | ***ApplicationServiceApiApplicationServicePodLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServicePodLogsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **namespace** | **optional.String**|  | 
 **container** | **optional.String**|  | 
 **sinceSeconds** | **optional.String**|  | 
 **sinceTimeSeconds** | **optional.String**| Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive. | 
 **sinceTimeNanos** | **optional.Int32**| Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context. | 
 **tailLines** | **optional.String**|  | 
 **follow** | **optional.Bool**|  | 
 **untilTime** | **optional.String**|  | 
 **filter** | **optional.String**|  | 
 **kind** | **optional.String**|  | 
 **group** | **optional.String**|  | 
 **resourceName** | **optional.String**|  | 
 **previous** | **optional.Bool**|  | 

### Return type

[**StreamResultOfApplicationLogEntry**](Stream result of applicationLogEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServicePodLogs2**
> StreamResultOfApplicationLogEntry ApplicationServicePodLogs2(ctx, name, optional)
PodLogs returns stream of log entries for the specified pod. Pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
 **optional** | ***ApplicationServiceApiApplicationServicePodLogs2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServicePodLogs2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **optional.String**|  | 
 **podName** | **optional.String**|  | 
 **container** | **optional.String**|  | 
 **sinceSeconds** | **optional.String**|  | 
 **sinceTimeSeconds** | **optional.String**| Represents seconds of UTC time since Unix epoch 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59Z inclusive. | 
 **sinceTimeNanos** | **optional.Int32**| Non-negative fractions of a second at nanosecond resolution. Negative second values with fractions must still have non-negative nanos values that count forward in time. Must be from 0 to 999,999,999 inclusive. This field may be limited in precision depending on context. | 
 **tailLines** | **optional.String**|  | 
 **follow** | **optional.Bool**|  | 
 **untilTime** | **optional.String**|  | 
 **filter** | **optional.String**|  | 
 **kind** | **optional.String**|  | 
 **group** | **optional.String**|  | 
 **resourceName** | **optional.String**|  | 
 **previous** | **optional.Bool**|  | 

### Return type

[**StreamResultOfApplicationLogEntry**](Stream result of applicationLogEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceResourceTree**
> V1alpha1ApplicationTree ApplicationServiceResourceTree(ctx, applicationName, optional)
ResourceTree returns resource tree

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationName** | **string**|  | 
 **optional** | ***ApplicationServiceApiApplicationServiceResourceTreeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServiceResourceTreeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **version** | **optional.String**|  | 
 **group** | **optional.String**|  | 
 **kind** | **optional.String**|  | 

### Return type

[**V1alpha1ApplicationTree**](v1alpha1ApplicationTree.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceRevisionMetadata**
> V1alpha1RevisionMetadata ApplicationServiceRevisionMetadata(ctx, name, revision)
Get the meta-data (author, date, tags, message) for a specific revision of the application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the application&#39;s name | 
  **revision** | **string**| the revision of the app | 

### Return type

[**V1alpha1RevisionMetadata**](v1alpha1RevisionMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceRollback**
> V1alpha1Application ApplicationServiceRollback(ctx, name, body)
Rollback syncs an application to its target state

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **body** | [**ApplicationApplicationRollbackRequest**](ApplicationApplicationRollbackRequest.md)|  | 

### Return type

[**V1alpha1Application**](v1alpha1Application.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceRunResourceAction**
> ApplicationApplicationResponse ApplicationServiceRunResourceAction(ctx, name, body, optional)
RunResourceAction run resource action

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **body** | **string**|  | 
 **optional** | ***ApplicationServiceApiApplicationServiceRunResourceActionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServiceRunResourceActionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **namespace** | **optional.String**|  | 
 **resourceName** | **optional.String**|  | 
 **version** | **optional.String**|  | 
 **group** | **optional.String**|  | 
 **kind** | **optional.String**|  | 

### Return type

[**ApplicationApplicationResponse**](applicationApplicationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceSync**
> V1alpha1Application ApplicationServiceSync(ctx, name, body)
Sync syncs an application to its target state

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **body** | [**ApplicationApplicationSyncRequest**](ApplicationApplicationSyncRequest.md)|  | 

### Return type

[**V1alpha1Application**](v1alpha1Application.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceTerminateOperation**
> ApplicationOperationTerminateResponse ApplicationServiceTerminateOperation(ctx, name)
TerminateOperation terminates the currently running operation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**ApplicationOperationTerminateResponse**](applicationOperationTerminateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceUpdate**
> V1alpha1Application ApplicationServiceUpdate(ctx, applicationMetadataName, body, optional)
Update updates an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationMetadataName** | **string**| Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names +optional | 
  **body** | [**V1alpha1Application**](V1alpha1Application.md)|  | 
 **optional** | ***ApplicationServiceApiApplicationServiceUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServiceUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validate** | **optional.Bool**|  | 

### Return type

[**V1alpha1Application**](v1alpha1Application.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceUpdateSpec**
> V1alpha1ApplicationSpec ApplicationServiceUpdateSpec(ctx, name, body, optional)
UpdateSpec updates an application spec

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **body** | [**V1alpha1ApplicationSpec**](V1alpha1ApplicationSpec.md)|  | 
 **optional** | ***ApplicationServiceApiApplicationServiceUpdateSpecOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServiceUpdateSpecOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validate** | **optional.Bool**|  | 

### Return type

[**V1alpha1ApplicationSpec**](v1alpha1ApplicationSpec.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceWatch**
> StreamResultOfV1alpha1ApplicationWatchEvent ApplicationServiceWatch(ctx, optional)
Watch returns stream of application change events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ApplicationServiceApiApplicationServiceWatchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServiceWatchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| the application&#39;s name. | 
 **refresh** | **optional.String**| forces application reconciliation if set to true. | 
 **project** | [**optional.Interface of []string**](string.md)| the project names to restrict returned list applications. | 
 **resourceVersion** | **optional.String**| when specified with a watch call, shows changes that occur after that particular version of a resource. | 
 **selector** | **optional.String**| the selector to to restrict returned list to applications only with matched labels. | 
 **repo** | **optional.String**| the repoURL to restrict returned list applications. | 

### Return type

[**StreamResultOfV1alpha1ApplicationWatchEvent**](Stream result of v1alpha1ApplicationWatchEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplicationServiceWatchResourceTree**
> StreamResultOfV1alpha1ApplicationTree ApplicationServiceWatchResourceTree(ctx, applicationName, optional)
Watch returns stream of application resource tree

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationName** | **string**|  | 
 **optional** | ***ApplicationServiceApiApplicationServiceWatchResourceTreeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ApplicationServiceApiApplicationServiceWatchResourceTreeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **version** | **optional.String**|  | 
 **group** | **optional.String**|  | 
 **kind** | **optional.String**|  | 

### Return type

[**StreamResultOfV1alpha1ApplicationTree**](Stream result of v1alpha1ApplicationTree.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

