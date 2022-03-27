# \ClusterServiceApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterServiceCreate**](ClusterServiceApi.md#ClusterServiceCreate) | **Post** /api/v1/clusters | Create creates a cluster
[**ClusterServiceDelete**](ClusterServiceApi.md#ClusterServiceDelete) | **Delete** /api/v1/clusters/{server} | Delete deletes a cluster
[**ClusterServiceGet**](ClusterServiceApi.md#ClusterServiceGet) | **Get** /api/v1/clusters/{server} | Get returns a cluster by server address
[**ClusterServiceInvalidateCache**](ClusterServiceApi.md#ClusterServiceInvalidateCache) | **Post** /api/v1/clusters/{server}/invalidate-cache | InvalidateCache invalidates cluster cache
[**ClusterServiceList**](ClusterServiceApi.md#ClusterServiceList) | **Get** /api/v1/clusters | List returns list of clusters
[**ClusterServiceRotateAuth**](ClusterServiceApi.md#ClusterServiceRotateAuth) | **Post** /api/v1/clusters/{server}/rotate-auth | RotateAuth rotates the bearer token used for a cluster
[**ClusterServiceUpdate**](ClusterServiceApi.md#ClusterServiceUpdate) | **Put** /api/v1/clusters/{cluster.server} | Update updates a cluster


# **ClusterServiceCreate**
> V1alpha1Cluster ClusterServiceCreate(ctx, body, optional)
Create creates a cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1alpha1Cluster**](V1alpha1Cluster.md)|  | 
 **optional** | ***ClusterServiceApiClusterServiceCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterServiceApiClusterServiceCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsert** | **optional.Bool**|  | 

### Return type

[**V1alpha1Cluster**](v1alpha1Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterServiceDelete**
> ClusterClusterResponse ClusterServiceDelete(ctx, server, optional)
Delete deletes a cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**|  | 
 **optional** | ***ClusterServiceApiClusterServiceDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterServiceApiClusterServiceDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**|  | 

### Return type

[**ClusterClusterResponse**](clusterClusterResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterServiceGet**
> V1alpha1Cluster ClusterServiceGet(ctx, server, optional)
Get returns a cluster by server address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**|  | 
 **optional** | ***ClusterServiceApiClusterServiceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterServiceApiClusterServiceGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**|  | 

### Return type

[**V1alpha1Cluster**](v1alpha1Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterServiceInvalidateCache**
> V1alpha1Cluster ClusterServiceInvalidateCache(ctx, server)
InvalidateCache invalidates cluster cache

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**|  | 

### Return type

[**V1alpha1Cluster**](v1alpha1Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterServiceList**
> V1alpha1ClusterList ClusterServiceList(ctx, optional)
List returns list of clusters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClusterServiceApiClusterServiceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterServiceApiClusterServiceListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **server** | **optional.String**|  | 
 **name** | **optional.String**|  | 

### Return type

[**V1alpha1ClusterList**](v1alpha1ClusterList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterServiceRotateAuth**
> ClusterClusterResponse ClusterServiceRotateAuth(ctx, server)
RotateAuth rotates the bearer token used for a cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **server** | **string**|  | 

### Return type

[**ClusterClusterResponse**](clusterClusterResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClusterServiceUpdate**
> V1alpha1Cluster ClusterServiceUpdate(ctx, clusterServer, body, optional)
Update updates a cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterServer** | **string**| Server is the API server URL of the Kubernetes cluster | 
  **body** | [**V1alpha1Cluster**](V1alpha1Cluster.md)|  | 
 **optional** | ***ClusterServiceApiClusterServiceUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClusterServiceApiClusterServiceUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updatedFields** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**V1alpha1Cluster**](v1alpha1Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

