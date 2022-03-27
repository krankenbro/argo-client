# \GPGKeyServiceApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GPGKeyServiceCreate**](GPGKeyServiceApi.md#GPGKeyServiceCreate) | **Post** /api/v1/gpgkeys | Create one or more GPG public keys in the server&#39;s configuration
[**GPGKeyServiceDelete**](GPGKeyServiceApi.md#GPGKeyServiceDelete) | **Delete** /api/v1/gpgkeys | Delete specified GPG public key from the server&#39;s configuration
[**GPGKeyServiceGet**](GPGKeyServiceApi.md#GPGKeyServiceGet) | **Get** /api/v1/gpgkeys/{keyID} | Get information about specified GPG public key from the server
[**GPGKeyServiceList**](GPGKeyServiceApi.md#GPGKeyServiceList) | **Get** /api/v1/gpgkeys | List all available repository certificates


# **GPGKeyServiceCreate**
> GpgkeyGnuPgPublicKeyCreateResponse GPGKeyServiceCreate(ctx, body, optional)
Create one or more GPG public keys in the server's configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1alpha1GnuPgPublicKey**](V1alpha1GnuPgPublicKey.md)| Raw key data of the GPG key(s) to create | 
 **optional** | ***GPGKeyServiceApiGPGKeyServiceCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GPGKeyServiceApiGPGKeyServiceCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsert** | **optional.Bool**| Whether to upsert already existing public keys. | 

### Return type

[**GpgkeyGnuPgPublicKeyCreateResponse**](gpgkeyGnuPGPublicKeyCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GPGKeyServiceDelete**
> GpgkeyGnuPgPublicKeyResponse GPGKeyServiceDelete(ctx, optional)
Delete specified GPG public key from the server's configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GPGKeyServiceApiGPGKeyServiceDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GPGKeyServiceApiGPGKeyServiceDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyID** | **optional.String**| The GPG key ID to query for. | 

### Return type

[**GpgkeyGnuPgPublicKeyResponse**](gpgkeyGnuPGPublicKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GPGKeyServiceGet**
> V1alpha1GnuPgPublicKey GPGKeyServiceGet(ctx, keyID)
Get information about specified GPG public key from the server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyID** | **string**| The GPG key ID to query for | 

### Return type

[**V1alpha1GnuPgPublicKey**](v1alpha1GnuPGPublicKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GPGKeyServiceList**
> V1alpha1GnuPgPublicKeyList GPGKeyServiceList(ctx, optional)
List all available repository certificates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GPGKeyServiceApiGPGKeyServiceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GPGKeyServiceApiGPGKeyServiceListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyID** | **optional.String**| The GPG key ID to query for. | 

### Return type

[**V1alpha1GnuPgPublicKeyList**](v1alpha1GnuPGPublicKeyList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

