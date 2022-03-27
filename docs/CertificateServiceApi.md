# \CertificateServiceApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateServiceCreateCertificate**](CertificateServiceApi.md#CertificateServiceCreateCertificate) | **Post** /api/v1/certificates | Creates repository certificates on the server
[**CertificateServiceDeleteCertificate**](CertificateServiceApi.md#CertificateServiceDeleteCertificate) | **Delete** /api/v1/certificates | Delete the certificates that match the RepositoryCertificateQuery
[**CertificateServiceListCertificates**](CertificateServiceApi.md#CertificateServiceListCertificates) | **Get** /api/v1/certificates | List all available repository certificates


# **CertificateServiceCreateCertificate**
> V1alpha1RepositoryCertificateList CertificateServiceCreateCertificate(ctx, body, optional)
Creates repository certificates on the server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1alpha1RepositoryCertificateList**](V1alpha1RepositoryCertificateList.md)| List of certificates to be created | 
 **optional** | ***CertificateServiceApiCertificateServiceCreateCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertificateServiceApiCertificateServiceCreateCertificateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsert** | **optional.Bool**| Whether to upsert already existing certificates. | 

### Return type

[**V1alpha1RepositoryCertificateList**](v1alpha1RepositoryCertificateList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CertificateServiceDeleteCertificate**
> V1alpha1RepositoryCertificateList CertificateServiceDeleteCertificate(ctx, optional)
Delete the certificates that match the RepositoryCertificateQuery

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CertificateServiceApiCertificateServiceDeleteCertificateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertificateServiceApiCertificateServiceDeleteCertificateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostNamePattern** | **optional.String**| A file-glob pattern (not regular expression) the host name has to match. | 
 **certType** | **optional.String**| The type of the certificate to match (ssh or https). | 
 **certSubType** | **optional.String**| The sub type of the certificate to match (protocol dependent, usually only used for ssh certs). | 

### Return type

[**V1alpha1RepositoryCertificateList**](v1alpha1RepositoryCertificateList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CertificateServiceListCertificates**
> V1alpha1RepositoryCertificateList CertificateServiceListCertificates(ctx, optional)
List all available repository certificates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CertificateServiceApiCertificateServiceListCertificatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CertificateServiceApiCertificateServiceListCertificatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hostNamePattern** | **optional.String**| A file-glob pattern (not regular expression) the host name has to match. | 
 **certType** | **optional.String**| The type of the certificate to match (ssh or https). | 
 **certSubType** | **optional.String**| The sub type of the certificate to match (protocol dependent, usually only used for ssh certs). | 

### Return type

[**V1alpha1RepositoryCertificateList**](v1alpha1RepositoryCertificateList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

