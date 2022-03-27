# \RepoCredsServiceApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepoCredsServiceCreateRepositoryCredentials**](RepoCredsServiceApi.md#RepoCredsServiceCreateRepositoryCredentials) | **Post** /api/v1/repocreds | CreateRepositoryCredentials creates a new repository credential set
[**RepoCredsServiceDeleteRepositoryCredentials**](RepoCredsServiceApi.md#RepoCredsServiceDeleteRepositoryCredentials) | **Delete** /api/v1/repocreds/{url} | DeleteRepositoryCredentials deletes a repository credential set from the configuration
[**RepoCredsServiceListRepositoryCredentials**](RepoCredsServiceApi.md#RepoCredsServiceListRepositoryCredentials) | **Get** /api/v1/repocreds | ListRepositoryCredentials gets a list of all configured repository credential sets
[**RepoCredsServiceUpdateRepositoryCredentials**](RepoCredsServiceApi.md#RepoCredsServiceUpdateRepositoryCredentials) | **Put** /api/v1/repocreds/{creds.url} | UpdateRepositoryCredentials updates a repository credential set


# **RepoCredsServiceCreateRepositoryCredentials**
> V1alpha1RepoCreds RepoCredsServiceCreateRepositoryCredentials(ctx, body, optional)
CreateRepositoryCredentials creates a new repository credential set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1alpha1RepoCreds**](V1alpha1RepoCreds.md)| Repository definition | 
 **optional** | ***RepoCredsServiceApiRepoCredsServiceCreateRepositoryCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepoCredsServiceApiRepoCredsServiceCreateRepositoryCredentialsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsert** | **optional.Bool**| Whether to create in upsert mode. | 

### Return type

[**V1alpha1RepoCreds**](v1alpha1RepoCreds.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepoCredsServiceDeleteRepositoryCredentials**
> RepocredsRepoCredsResponse RepoCredsServiceDeleteRepositoryCredentials(ctx, url)
DeleteRepositoryCredentials deletes a repository credential set from the configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **url** | **string**|  | 

### Return type

[**RepocredsRepoCredsResponse**](repocredsRepoCredsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepoCredsServiceListRepositoryCredentials**
> V1alpha1RepoCredsList RepoCredsServiceListRepositoryCredentials(ctx, optional)
ListRepositoryCredentials gets a list of all configured repository credential sets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepoCredsServiceApiRepoCredsServiceListRepositoryCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepoCredsServiceApiRepoCredsServiceListRepositoryCredentialsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **optional.String**| Repo URL for query. | 

### Return type

[**V1alpha1RepoCredsList**](v1alpha1RepoCredsList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepoCredsServiceUpdateRepositoryCredentials**
> V1alpha1RepoCreds RepoCredsServiceUpdateRepositoryCredentials(ctx, credsUrl, body)
UpdateRepositoryCredentials updates a repository credential set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **credsUrl** | **string**| URL is the URL that this credentials matches to | 
  **body** | [**V1alpha1RepoCreds**](V1alpha1RepoCreds.md)|  | 

### Return type

[**V1alpha1RepoCreds**](v1alpha1RepoCreds.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

