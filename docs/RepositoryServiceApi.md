# \RepositoryServiceApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoryServiceCreateRepository**](RepositoryServiceApi.md#RepositoryServiceCreateRepository) | **Post** /api/v1/repositories | CreateRepository creates a new repository configuration
[**RepositoryServiceDeleteRepository**](RepositoryServiceApi.md#RepositoryServiceDeleteRepository) | **Delete** /api/v1/repositories/{repo} | DeleteRepository deletes a repository from the configuration
[**RepositoryServiceGet**](RepositoryServiceApi.md#RepositoryServiceGet) | **Get** /api/v1/repositories/{repo} | Get returns a repository or its credentials
[**RepositoryServiceGetAppDetails**](RepositoryServiceApi.md#RepositoryServiceGetAppDetails) | **Post** /api/v1/repositories/{source.repoURL}/appdetails | GetAppDetails returns application details by given path
[**RepositoryServiceGetHelmCharts**](RepositoryServiceApi.md#RepositoryServiceGetHelmCharts) | **Get** /api/v1/repositories/{repo}/helmcharts | GetHelmCharts returns list of helm charts in the specified repository
[**RepositoryServiceListApps**](RepositoryServiceApi.md#RepositoryServiceListApps) | **Get** /api/v1/repositories/{repo}/apps | ListApps returns list of apps in the repo
[**RepositoryServiceListRefs**](RepositoryServiceApi.md#RepositoryServiceListRefs) | **Get** /api/v1/repositories/{repo}/refs | 
[**RepositoryServiceListRepositories**](RepositoryServiceApi.md#RepositoryServiceListRepositories) | **Get** /api/v1/repositories | ListRepositories gets a list of all configured repositories
[**RepositoryServiceUpdateRepository**](RepositoryServiceApi.md#RepositoryServiceUpdateRepository) | **Put** /api/v1/repositories/{repo.repo} | UpdateRepository updates a repository configuration
[**RepositoryServiceValidateAccess**](RepositoryServiceApi.md#RepositoryServiceValidateAccess) | **Post** /api/v1/repositories/{repo}/validate | ValidateAccess validates access to a repository with given parameters


# **RepositoryServiceCreateRepository**
> V1alpha1Repository RepositoryServiceCreateRepository(ctx, body, optional)
CreateRepository creates a new repository configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1alpha1Repository**](V1alpha1Repository.md)| Repository definition | 
 **optional** | ***RepositoryServiceApiRepositoryServiceCreateRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryServiceApiRepositoryServiceCreateRepositoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsert** | **optional.Bool**| Whether to create in upsert mode. | 
 **credsOnly** | **optional.Bool**| Whether to operate on credential set instead of repository. | 

### Return type

[**V1alpha1Repository**](v1alpha1Repository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoryServiceDeleteRepository**
> RepositoryRepoResponse RepositoryServiceDeleteRepository(ctx, repo, optional)
DeleteRepository deletes a repository from the configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| Repo URL for query | 
 **optional** | ***RepositoryServiceApiRepositoryServiceDeleteRepositoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryServiceApiRepositoryServiceDeleteRepositoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceRefresh** | **optional.Bool**| Whether to force a cache refresh on repo&#39;s connection state. | 

### Return type

[**RepositoryRepoResponse**](repositoryRepoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoryServiceGet**
> V1alpha1Repository RepositoryServiceGet(ctx, repo, optional)
Get returns a repository or its credentials

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| Repo URL for query | 
 **optional** | ***RepositoryServiceApiRepositoryServiceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryServiceApiRepositoryServiceGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceRefresh** | **optional.Bool**| Whether to force a cache refresh on repo&#39;s connection state. | 

### Return type

[**V1alpha1Repository**](v1alpha1Repository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoryServiceGetAppDetails**
> RepositoryRepoAppDetailsResponse RepositoryServiceGetAppDetails(ctx, sourceRepoURL, body)
GetAppDetails returns application details by given path

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceRepoURL** | **string**| RepoURL is the URL to the repository (Git or Helm) that contains the application manifests | 
  **body** | [**RepositoryRepoAppDetailsQuery**](RepositoryRepoAppDetailsQuery.md)|  | 

### Return type

[**RepositoryRepoAppDetailsResponse**](repositoryRepoAppDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoryServiceGetHelmCharts**
> RepositoryHelmChartsResponse RepositoryServiceGetHelmCharts(ctx, repo, optional)
GetHelmCharts returns list of helm charts in the specified repository

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| Repo URL for query | 
 **optional** | ***RepositoryServiceApiRepositoryServiceGetHelmChartsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryServiceApiRepositoryServiceGetHelmChartsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceRefresh** | **optional.Bool**| Whether to force a cache refresh on repo&#39;s connection state. | 

### Return type

[**RepositoryHelmChartsResponse**](repositoryHelmChartsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoryServiceListApps**
> RepositoryRepoAppsResponse RepositoryServiceListApps(ctx, repo, optional)
ListApps returns list of apps in the repo

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**|  | 
 **optional** | ***RepositoryServiceApiRepositoryServiceListAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryServiceApiRepositoryServiceListAppsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revision** | **optional.String**|  | 

### Return type

[**RepositoryRepoAppsResponse**](repositoryRepoAppsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoryServiceListRefs**
> RepositoryRefs RepositoryServiceListRefs(ctx, repo, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| Repo URL for query | 
 **optional** | ***RepositoryServiceApiRepositoryServiceListRefsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryServiceApiRepositoryServiceListRefsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceRefresh** | **optional.Bool**| Whether to force a cache refresh on repo&#39;s connection state. | 

### Return type

[**RepositoryRefs**](repositoryRefs.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoryServiceListRepositories**
> V1alpha1RepositoryList RepositoryServiceListRepositories(ctx, optional)
ListRepositories gets a list of all configured repositories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryServiceApiRepositoryServiceListRepositoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryServiceApiRepositoryServiceListRepositoriesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repo** | **optional.String**| Repo URL for query. | 
 **forceRefresh** | **optional.Bool**| Whether to force a cache refresh on repo&#39;s connection state. | 

### Return type

[**V1alpha1RepositoryList**](v1alpha1RepositoryList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoryServiceUpdateRepository**
> V1alpha1Repository RepositoryServiceUpdateRepository(ctx, repoRepo, body)
UpdateRepository updates a repository configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoRepo** | **string**| Repo contains the URL to the remote repository | 
  **body** | [**V1alpha1Repository**](V1alpha1Repository.md)|  | 

### Return type

[**V1alpha1Repository**](v1alpha1Repository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoryServiceValidateAccess**
> RepositoryRepoResponse RepositoryServiceValidateAccess(ctx, repo, body, optional)
ValidateAccess validates access to a repository with given parameters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The URL to the repo | 
  **body** | **string**| The URL to the repo | 
 **optional** | ***RepositoryServiceApiRepositoryServiceValidateAccessOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryServiceApiRepositoryServiceValidateAccessOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **username** | **optional.String**| Username for accessing repo. | 
 **password** | **optional.String**| Password for accessing repo. | 
 **sshPrivateKey** | **optional.String**| Private key data for accessing SSH repository. | 
 **insecure** | **optional.Bool**| Whether to skip certificate or host key validation. | 
 **tlsClientCertData** | **optional.String**| TLS client cert data for accessing HTTPS repository. | 
 **tlsClientCertKey** | **optional.String**| TLS client cert key for accessing HTTPS repository. | 
 **type_** | **optional.String**| The type of the repo. | 
 **name** | **optional.String**| The name of the repo. | 
 **enableOci** | **optional.Bool**| Whether helm-oci support should be enabled for this repo. | 
 **githubAppPrivateKey** | **optional.String**| Github App Private Key PEM data. | 
 **githubAppID** | **optional.String**| Github App ID of the app used to access the repo. | 
 **githubAppInstallationID** | **optional.String**| Github App Installation ID of the installed GitHub App. | 
 **githubAppEnterpriseBaseUrl** | **optional.String**| Github App Enterprise base url if empty will default to https://api.github.com. | 
 **proxy** | **optional.String**| HTTP/HTTPS proxy to access the repository. | 
 **project** | **optional.String**| Reference between project and repository that allow you automatically to be added as item inside SourceRepos project entity. | 

### Return type

[**RepositoryRepoResponse**](repositoryRepoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

