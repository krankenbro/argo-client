# \SessionServiceApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SessionServiceCreate**](SessionServiceApi.md#SessionServiceCreate) | **Post** /api/v1/session | Create a new JWT for authentication and set a cookie if using HTTP
[**SessionServiceDelete**](SessionServiceApi.md#SessionServiceDelete) | **Delete** /api/v1/session | Delete an existing JWT cookie if using HTTP
[**SessionServiceGetUserInfo**](SessionServiceApi.md#SessionServiceGetUserInfo) | **Get** /api/v1/session/userinfo | Get the current user&#39;s info


# **SessionServiceCreate**
> SessionSessionResponse SessionServiceCreate(ctx, body)
Create a new JWT for authentication and set a cookie if using HTTP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SessionSessionCreateRequest**](SessionSessionCreateRequest.md)|  | 

### Return type

[**SessionSessionResponse**](sessionSessionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionServiceDelete**
> SessionSessionResponse SessionServiceDelete(ctx, )
Delete an existing JWT cookie if using HTTP

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SessionSessionResponse**](sessionSessionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionServiceGetUserInfo**
> SessionGetUserInfoResponse SessionServiceGetUserInfo(ctx, )
Get the current user's info

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SessionGetUserInfoResponse**](sessionGetUserInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

