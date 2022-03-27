# \AccountServiceApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountServiceCanI**](AccountServiceApi.md#AccountServiceCanI) | **Get** /api/v1/account/can-i/{resource}/{action}/{subresource} | CanI checks if the current account has permission to perform an action
[**AccountServiceCreateToken**](AccountServiceApi.md#AccountServiceCreateToken) | **Post** /api/v1/account/{name}/token | CreateToken creates a token
[**AccountServiceDeleteToken**](AccountServiceApi.md#AccountServiceDeleteToken) | **Delete** /api/v1/account/{name}/token/{id} | DeleteToken deletes a token
[**AccountServiceGetAccount**](AccountServiceApi.md#AccountServiceGetAccount) | **Get** /api/v1/account/{name} | GetAccount returns an account
[**AccountServiceListAccounts**](AccountServiceApi.md#AccountServiceListAccounts) | **Get** /api/v1/account | ListAccounts returns the list of accounts
[**AccountServiceUpdatePassword**](AccountServiceApi.md#AccountServiceUpdatePassword) | **Put** /api/v1/account/password | UpdatePassword updates an account&#39;s password to a new value


# **AccountServiceCanI**
> AccountCanIResponse AccountServiceCanI(ctx, resource, action, subresource)
CanI checks if the current account has permission to perform an action

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resource** | **string**|  | 
  **action** | **string**|  | 
  **subresource** | **string**|  | 

### Return type

[**AccountCanIResponse**](accountCanIResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountServiceCreateToken**
> AccountCreateTokenResponse AccountServiceCreateToken(ctx, name, body)
CreateToken creates a token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **body** | [**AccountCreateTokenRequest**](AccountCreateTokenRequest.md)|  | 

### Return type

[**AccountCreateTokenResponse**](accountCreateTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountServiceDeleteToken**
> AccountEmptyResponse AccountServiceDeleteToken(ctx, name, id)
DeleteToken deletes a token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**AccountEmptyResponse**](accountEmptyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountServiceGetAccount**
> AccountAccount AccountServiceGetAccount(ctx, name)
GetAccount returns an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**AccountAccount**](accountAccount.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountServiceListAccounts**
> AccountAccountsList AccountServiceListAccounts(ctx, )
ListAccounts returns the list of accounts

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AccountAccountsList**](accountAccountsList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountServiceUpdatePassword**
> AccountUpdatePasswordResponse AccountServiceUpdatePassword(ctx, body)
UpdatePassword updates an account's password to a new value

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountUpdatePasswordRequest**](AccountUpdatePasswordRequest.md)|  | 

### Return type

[**AccountUpdatePasswordResponse**](accountUpdatePasswordResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

