# \ProjectServiceApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectServiceCreate**](ProjectServiceApi.md#ProjectServiceCreate) | **Post** /api/v1/projects | Create a new project
[**ProjectServiceCreateToken**](ProjectServiceApi.md#ProjectServiceCreateToken) | **Post** /api/v1/projects/{project}/roles/{role}/token | Create a new project token
[**ProjectServiceDelete**](ProjectServiceApi.md#ProjectServiceDelete) | **Delete** /api/v1/projects/{name} | Delete deletes a project
[**ProjectServiceDeleteToken**](ProjectServiceApi.md#ProjectServiceDeleteToken) | **Delete** /api/v1/projects/{project}/roles/{role}/token/{iat} | Delete a new project token
[**ProjectServiceGet**](ProjectServiceApi.md#ProjectServiceGet) | **Get** /api/v1/projects/{name} | Get returns a project by name
[**ProjectServiceGetDetailedProject**](ProjectServiceApi.md#ProjectServiceGetDetailedProject) | **Get** /api/v1/projects/{name}/detailed | GetDetailedProject returns a project that include project, global project and scoped resources by name
[**ProjectServiceGetGlobalProjects**](ProjectServiceApi.md#ProjectServiceGetGlobalProjects) | **Get** /api/v1/projects/{name}/globalprojects | Get returns a virtual project by name
[**ProjectServiceGetSyncWindowsState**](ProjectServiceApi.md#ProjectServiceGetSyncWindowsState) | **Get** /api/v1/projects/{name}/syncwindows | GetSchedulesState returns true if there are any active sync syncWindows
[**ProjectServiceList**](ProjectServiceApi.md#ProjectServiceList) | **Get** /api/v1/projects | List returns list of projects
[**ProjectServiceListEvents**](ProjectServiceApi.md#ProjectServiceListEvents) | **Get** /api/v1/projects/{name}/events | ListEvents returns a list of project events
[**ProjectServiceUpdate**](ProjectServiceApi.md#ProjectServiceUpdate) | **Put** /api/v1/projects/{project.metadata.name} | Update updates a project


# **ProjectServiceCreate**
> V1alpha1AppProject ProjectServiceCreate(ctx, body)
Create a new project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectProjectCreateRequest**](ProjectProjectCreateRequest.md)|  | 

### Return type

[**V1alpha1AppProject**](v1alpha1AppProject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectServiceCreateToken**
> ProjectProjectTokenResponse ProjectServiceCreateToken(ctx, project, role, body)
Create a new project token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | **string**|  | 
  **role** | **string**|  | 
  **body** | [**ProjectProjectTokenCreateRequest**](ProjectProjectTokenCreateRequest.md)|  | 

### Return type

[**ProjectProjectTokenResponse**](projectProjectTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectServiceDelete**
> ProjectEmptyResponse ProjectServiceDelete(ctx, name)
Delete deletes a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**ProjectEmptyResponse**](projectEmptyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectServiceDeleteToken**
> ProjectEmptyResponse ProjectServiceDeleteToken(ctx, project, role, iat, optional)
Delete a new project token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | **string**|  | 
  **role** | **string**|  | 
  **iat** | **string**|  | 
 **optional** | ***ProjectServiceApiProjectServiceDeleteTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectServiceApiProjectServiceDeleteTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **id** | **optional.String**|  | 

### Return type

[**ProjectEmptyResponse**](projectEmptyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectServiceGet**
> V1alpha1AppProject ProjectServiceGet(ctx, name)
Get returns a project by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**V1alpha1AppProject**](v1alpha1AppProject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectServiceGetDetailedProject**
> ProjectDetailedProjectsResponse ProjectServiceGetDetailedProject(ctx, name)
GetDetailedProject returns a project that include project, global project and scoped resources by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**ProjectDetailedProjectsResponse**](projectDetailedProjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectServiceGetGlobalProjects**
> ProjectGlobalProjectsResponse ProjectServiceGetGlobalProjects(ctx, name)
Get returns a virtual project by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**ProjectGlobalProjectsResponse**](projectGlobalProjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectServiceGetSyncWindowsState**
> ProjectSyncWindowsResponse ProjectServiceGetSyncWindowsState(ctx, name)
GetSchedulesState returns true if there are any active sync syncWindows

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**ProjectSyncWindowsResponse**](projectSyncWindowsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectServiceList**
> V1alpha1AppProjectList ProjectServiceList(ctx, optional)
List returns list of projects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectServiceApiProjectServiceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectServiceApiProjectServiceListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**|  | 

### Return type

[**V1alpha1AppProjectList**](v1alpha1AppProjectList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectServiceListEvents**
> V1EventList ProjectServiceListEvents(ctx, name)
ListEvents returns a list of project events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**V1EventList**](v1EventList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectServiceUpdate**
> V1alpha1AppProject ProjectServiceUpdate(ctx, projectMetadataName, body)
Update updates a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectMetadataName** | **string**| Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names +optional | 
  **body** | [**ProjectProjectUpdateRequest**](ProjectProjectUpdateRequest.md)|  | 

### Return type

[**V1alpha1AppProject**](v1alpha1AppProject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

