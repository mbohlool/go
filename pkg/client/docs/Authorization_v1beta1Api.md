# \Authorization_v1beta1Api

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespacedLocalSubjectAccessReview**](Authorization_v1beta1Api.md#CreateNamespacedLocalSubjectAccessReview) | **Post** /apis/authorization.k8s.io/v1beta1/namespaces/{namespace}/localsubjectaccessreviews | 
[**CreateSelfSubjectAccessReview**](Authorization_v1beta1Api.md#CreateSelfSubjectAccessReview) | **Post** /apis/authorization.k8s.io/v1beta1/selfsubjectaccessreviews | 
[**CreateSelfSubjectRulesReview**](Authorization_v1beta1Api.md#CreateSelfSubjectRulesReview) | **Post** /apis/authorization.k8s.io/v1beta1/selfsubjectrulesreviews | 
[**CreateSubjectAccessReview**](Authorization_v1beta1Api.md#CreateSubjectAccessReview) | **Post** /apis/authorization.k8s.io/v1beta1/subjectaccessreviews | 
[**GetAPIResources**](Authorization_v1beta1Api.md#GetAPIResources) | **Get** /apis/authorization.k8s.io/v1beta1/ | 


# **CreateNamespacedLocalSubjectAccessReview**
> V1beta1LocalSubjectAccessReview CreateNamespacedLocalSubjectAccessReview($namespace, $body, $pretty)



create a LocalSubjectAccessReview


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **body** | [**V1beta1LocalSubjectAccessReview**](V1beta1LocalSubjectAccessReview.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1beta1LocalSubjectAccessReview**](v1beta1.LocalSubjectAccessReview.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSelfSubjectAccessReview**
> V1beta1SelfSubjectAccessReview CreateSelfSubjectAccessReview($body, $pretty)



create a SelfSubjectAccessReview


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1beta1SelfSubjectAccessReview**](V1beta1SelfSubjectAccessReview.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1beta1SelfSubjectAccessReview**](v1beta1.SelfSubjectAccessReview.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSelfSubjectRulesReview**
> V1beta1SelfSubjectRulesReview CreateSelfSubjectRulesReview($body, $pretty)



create a SelfSubjectRulesReview


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1beta1SelfSubjectRulesReview**](V1beta1SelfSubjectRulesReview.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1beta1SelfSubjectRulesReview**](v1beta1.SelfSubjectRulesReview.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSubjectAccessReview**
> V1beta1SubjectAccessReview CreateSubjectAccessReview($body, $pretty)



create a SubjectAccessReview


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1beta1SubjectAccessReview**](V1beta1SubjectAccessReview.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1beta1SubjectAccessReview**](v1beta1.SubjectAccessReview.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAPIResources**
> V1ApiResourceList GetAPIResources()



get available resources


### Parameters
This endpoint does not need any parameter.

### Return type

[**V1ApiResourceList**](v1.APIResourceList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml, application/vnd.kubernetes.protobuf
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

