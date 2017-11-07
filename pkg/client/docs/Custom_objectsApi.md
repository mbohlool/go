# \Custom_objectsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClusterCustomObject**](Custom_objectsApi.md#CreateClusterCustomObject) | **Post** /apis/{group}/{version}/{plural} | 
[**CreateNamespacedCustomObject**](Custom_objectsApi.md#CreateNamespacedCustomObject) | **Post** /apis/{group}/{version}/namespaces/{namespace}/{plural} | 
[**DeleteClusterCustomObject**](Custom_objectsApi.md#DeleteClusterCustomObject) | **Delete** /apis/{group}/{version}/{plural}/{name} | 
[**DeleteNamespacedCustomObject**](Custom_objectsApi.md#DeleteNamespacedCustomObject) | **Delete** /apis/{group}/{version}/namespaces/{namespace}/{plural}/{name} | 
[**GetClusterCustomObject**](Custom_objectsApi.md#GetClusterCustomObject) | **Get** /apis/{group}/{version}/{plural}/{name} | 
[**GetNamespacedCustomObject**](Custom_objectsApi.md#GetNamespacedCustomObject) | **Get** /apis/{group}/{version}/namespaces/{namespace}/{plural}/{name} | 
[**ListClusterCustomObject**](Custom_objectsApi.md#ListClusterCustomObject) | **Get** /apis/{group}/{version}/{plural} | 
[**ListNamespacedCustomObject**](Custom_objectsApi.md#ListNamespacedCustomObject) | **Get** /apis/{group}/{version}/namespaces/{namespace}/{plural} | 
[**ReplaceClusterCustomObject**](Custom_objectsApi.md#ReplaceClusterCustomObject) | **Put** /apis/{group}/{version}/{plural}/{name} | 
[**ReplaceNamespacedCustomObject**](Custom_objectsApi.md#ReplaceNamespacedCustomObject) | **Put** /apis/{group}/{version}/namespaces/{namespace}/{plural}/{name} | 


# **CreateClusterCustomObject**
> interface{} CreateClusterCustomObject($group, $version, $plural, $body, $pretty)



Creates a cluster scoped Custom object


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string**| The custom resource&#39;s group name | 
 **version** | **string**| The custom resource&#39;s version | 
 **plural** | **string**| The custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **body** | [**interface{}**](interface{}.md)| The JSON schema of the Resource to create. | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNamespacedCustomObject**
> interface{} CreateNamespacedCustomObject($group, $version, $namespace, $plural, $body, $pretty)



Creates a namespace scoped Custom object


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string**| The custom resource&#39;s group name | 
 **version** | **string**| The custom resource&#39;s version | 
 **namespace** | **string**| The custom resource&#39;s namespace | 
 **plural** | **string**| The custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **body** | [**interface{}**](interface{}.md)| The JSON schema of the Resource to create. | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClusterCustomObject**
> interface{} DeleteClusterCustomObject($group, $version, $plural, $name, $body, $gracePeriodSeconds, $orphanDependents, $propagationPolicy)



Deletes the specified cluster scoped custom object


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string**| the custom resource&#39;s group | 
 **version** | **string**| the custom resource&#39;s version | 
 **plural** | **string**| the custom object&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **name** | **string**| the custom object&#39;s name | 
 **body** | [**V1DeleteOptions**](V1DeleteOptions.md)|  | 
 **gracePeriodSeconds** | **int32**| The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. | [optional] 
 **orphanDependents** | **bool**| Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \&quot;orphan\&quot; finalizer will be added to/removed from the object&#39;s finalizers list. Either this field or PropagationPolicy may be set, but not both. | [optional] 
 **propagationPolicy** | **string**| Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. | [optional] 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNamespacedCustomObject**
> interface{} DeleteNamespacedCustomObject($group, $version, $namespace, $plural, $name, $body, $gracePeriodSeconds, $orphanDependents, $propagationPolicy)



Deletes the specified namespace scoped custom object


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string**| the custom resource&#39;s group | 
 **version** | **string**| the custom resource&#39;s version | 
 **namespace** | **string**| The custom resource&#39;s namespace | 
 **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **name** | **string**| the custom object&#39;s name | 
 **body** | [**V1DeleteOptions**](V1DeleteOptions.md)|  | 
 **gracePeriodSeconds** | **int32**| The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. | [optional] 
 **orphanDependents** | **bool**| Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \&quot;orphan\&quot; finalizer will be added to/removed from the object&#39;s finalizers list. Either this field or PropagationPolicy may be set, but not both. | [optional] 
 **propagationPolicy** | **string**| Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. | [optional] 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterCustomObject**
> interface{} GetClusterCustomObject($group, $version, $plural, $name)



Returns a cluster scoped custom object


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string**| the custom resource&#39;s group | 
 **version** | **string**| the custom resource&#39;s version | 
 **plural** | **string**| the custom object&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **name** | **string**| the custom object&#39;s name | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNamespacedCustomObject**
> interface{} GetNamespacedCustomObject($group, $version, $namespace, $plural, $name)



Returns a namespace scoped custom object


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string**| the custom resource&#39;s group | 
 **version** | **string**| the custom resource&#39;s version | 
 **namespace** | **string**| The custom resource&#39;s namespace | 
 **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **name** | **string**| the custom object&#39;s name | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClusterCustomObject**
> interface{} ListClusterCustomObject($group, $version, $plural, $pretty, $labelSelector, $resourceVersion, $watch)



list or watch cluster scoped custom objects


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string**| The custom resource&#39;s group name | 
 **version** | **string**| The custom resource&#39;s version | 
 **plural** | **string**| The custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 
 **labelSelector** | **string**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | [optional] 
 **resourceVersion** | **string**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | [optional] 
 **watch** | **bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. | [optional] 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/json;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNamespacedCustomObject**
> interface{} ListNamespacedCustomObject($group, $version, $namespace, $plural, $pretty, $labelSelector, $resourceVersion, $watch)



list or watch namespace scoped custom objects


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string**| The custom resource&#39;s group name | 
 **version** | **string**| The custom resource&#39;s version | 
 **namespace** | **string**| The custom resource&#39;s namespace | 
 **plural** | **string**| The custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 
 **labelSelector** | **string**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | [optional] 
 **resourceVersion** | **string**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | [optional] 
 **watch** | **bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. | [optional] 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/json;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceClusterCustomObject**
> interface{} ReplaceClusterCustomObject($group, $version, $plural, $name, $body)



replace the specified cluster scoped custom object


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string**| the custom resource&#39;s group | 
 **version** | **string**| the custom resource&#39;s version | 
 **plural** | **string**| the custom object&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **name** | **string**| the custom object&#39;s name | 
 **body** | [**interface{}**](interface{}.md)| The JSON schema of the Resource to replace. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceNamespacedCustomObject**
> interface{} ReplaceNamespacedCustomObject($group, $version, $namespace, $plural, $name, $body)



replace the specified namespace scoped custom object


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string**| the custom resource&#39;s group | 
 **version** | **string**| the custom resource&#39;s version | 
 **namespace** | **string**| The custom resource&#39;s namespace | 
 **plural** | **string**| the custom resource&#39;s plural name. For TPRs this would be lowercase plural kind. | 
 **name** | **string**| the custom object&#39;s name | 
 **body** | [**interface{}**](interface{}.md)| The JSON schema of the Resource to replace. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

