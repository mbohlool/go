# \RbacAuthorization_v1Api

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClusterRole**](RbacAuthorization_v1Api.md#CreateClusterRole) | **Post** /apis/rbac.authorization.k8s.io/v1/clusterroles | 
[**CreateClusterRoleBinding**](RbacAuthorization_v1Api.md#CreateClusterRoleBinding) | **Post** /apis/rbac.authorization.k8s.io/v1/clusterrolebindings | 
[**CreateNamespacedRole**](RbacAuthorization_v1Api.md#CreateNamespacedRole) | **Post** /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/roles | 
[**CreateNamespacedRoleBinding**](RbacAuthorization_v1Api.md#CreateNamespacedRoleBinding) | **Post** /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/rolebindings | 
[**DeleteClusterRole**](RbacAuthorization_v1Api.md#DeleteClusterRole) | **Delete** /apis/rbac.authorization.k8s.io/v1/clusterroles/{name} | 
[**DeleteClusterRoleBinding**](RbacAuthorization_v1Api.md#DeleteClusterRoleBinding) | **Delete** /apis/rbac.authorization.k8s.io/v1/clusterrolebindings/{name} | 
[**DeleteCollectionClusterRole**](RbacAuthorization_v1Api.md#DeleteCollectionClusterRole) | **Delete** /apis/rbac.authorization.k8s.io/v1/clusterroles | 
[**DeleteCollectionClusterRoleBinding**](RbacAuthorization_v1Api.md#DeleteCollectionClusterRoleBinding) | **Delete** /apis/rbac.authorization.k8s.io/v1/clusterrolebindings | 
[**DeleteCollectionNamespacedRole**](RbacAuthorization_v1Api.md#DeleteCollectionNamespacedRole) | **Delete** /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/roles | 
[**DeleteCollectionNamespacedRoleBinding**](RbacAuthorization_v1Api.md#DeleteCollectionNamespacedRoleBinding) | **Delete** /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/rolebindings | 
[**DeleteNamespacedRole**](RbacAuthorization_v1Api.md#DeleteNamespacedRole) | **Delete** /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/roles/{name} | 
[**DeleteNamespacedRoleBinding**](RbacAuthorization_v1Api.md#DeleteNamespacedRoleBinding) | **Delete** /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/rolebindings/{name} | 
[**GetAPIResources**](RbacAuthorization_v1Api.md#GetAPIResources) | **Get** /apis/rbac.authorization.k8s.io/v1/ | 
[**ListClusterRole**](RbacAuthorization_v1Api.md#ListClusterRole) | **Get** /apis/rbac.authorization.k8s.io/v1/clusterroles | 
[**ListClusterRoleBinding**](RbacAuthorization_v1Api.md#ListClusterRoleBinding) | **Get** /apis/rbac.authorization.k8s.io/v1/clusterrolebindings | 
[**ListNamespacedRole**](RbacAuthorization_v1Api.md#ListNamespacedRole) | **Get** /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/roles | 
[**ListNamespacedRoleBinding**](RbacAuthorization_v1Api.md#ListNamespacedRoleBinding) | **Get** /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/rolebindings | 
[**ListRoleBindingForAllNamespaces**](RbacAuthorization_v1Api.md#ListRoleBindingForAllNamespaces) | **Get** /apis/rbac.authorization.k8s.io/v1/rolebindings | 
[**ListRoleForAllNamespaces**](RbacAuthorization_v1Api.md#ListRoleForAllNamespaces) | **Get** /apis/rbac.authorization.k8s.io/v1/roles | 
[**PatchClusterRole**](RbacAuthorization_v1Api.md#PatchClusterRole) | **Patch** /apis/rbac.authorization.k8s.io/v1/clusterroles/{name} | 
[**PatchClusterRoleBinding**](RbacAuthorization_v1Api.md#PatchClusterRoleBinding) | **Patch** /apis/rbac.authorization.k8s.io/v1/clusterrolebindings/{name} | 
[**PatchNamespacedRole**](RbacAuthorization_v1Api.md#PatchNamespacedRole) | **Patch** /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/roles/{name} | 
[**PatchNamespacedRoleBinding**](RbacAuthorization_v1Api.md#PatchNamespacedRoleBinding) | **Patch** /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/rolebindings/{name} | 
[**ReadClusterRole**](RbacAuthorization_v1Api.md#ReadClusterRole) | **Get** /apis/rbac.authorization.k8s.io/v1/clusterroles/{name} | 
[**ReadClusterRoleBinding**](RbacAuthorization_v1Api.md#ReadClusterRoleBinding) | **Get** /apis/rbac.authorization.k8s.io/v1/clusterrolebindings/{name} | 
[**ReadNamespacedRole**](RbacAuthorization_v1Api.md#ReadNamespacedRole) | **Get** /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/roles/{name} | 
[**ReadNamespacedRoleBinding**](RbacAuthorization_v1Api.md#ReadNamespacedRoleBinding) | **Get** /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/rolebindings/{name} | 
[**ReplaceClusterRole**](RbacAuthorization_v1Api.md#ReplaceClusterRole) | **Put** /apis/rbac.authorization.k8s.io/v1/clusterroles/{name} | 
[**ReplaceClusterRoleBinding**](RbacAuthorization_v1Api.md#ReplaceClusterRoleBinding) | **Put** /apis/rbac.authorization.k8s.io/v1/clusterrolebindings/{name} | 
[**ReplaceNamespacedRole**](RbacAuthorization_v1Api.md#ReplaceNamespacedRole) | **Put** /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/roles/{name} | 
[**ReplaceNamespacedRoleBinding**](RbacAuthorization_v1Api.md#ReplaceNamespacedRoleBinding) | **Put** /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/rolebindings/{name} | 


# **CreateClusterRole**
> V1ClusterRole CreateClusterRole($body, $pretty)



create a ClusterRole


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1ClusterRole**](V1ClusterRole.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1ClusterRole**](v1.ClusterRole.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateClusterRoleBinding**
> V1ClusterRoleBinding CreateClusterRoleBinding($body, $pretty)



create a ClusterRoleBinding


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1ClusterRoleBinding**](V1ClusterRoleBinding.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1ClusterRoleBinding**](v1.ClusterRoleBinding.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNamespacedRole**
> V1Role CreateNamespacedRole($namespace, $body, $pretty)



create a Role


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **body** | [**V1Role**](V1Role.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1Role**](v1.Role.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNamespacedRoleBinding**
> V1RoleBinding CreateNamespacedRoleBinding($namespace, $body, $pretty)



create a RoleBinding


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **body** | [**V1RoleBinding**](V1RoleBinding.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1RoleBinding**](v1.RoleBinding.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClusterRole**
> V1Status DeleteClusterRole($name, $body, $pretty, $gracePeriodSeconds, $orphanDependents, $propagationPolicy)



delete a ClusterRole


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the ClusterRole | 
 **body** | [**V1DeleteOptions**](V1DeleteOptions.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 
 **gracePeriodSeconds** | **int32**| The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. | [optional] 
 **orphanDependents** | **bool**| Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \&quot;orphan\&quot; finalizer will be added to/removed from the object&#39;s finalizers list. Either this field or PropagationPolicy may be set, but not both. | [optional] 
 **propagationPolicy** | **string**| Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. | [optional] 

### Return type

[**V1Status**](v1.Status.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClusterRoleBinding**
> V1Status DeleteClusterRoleBinding($name, $body, $pretty, $gracePeriodSeconds, $orphanDependents, $propagationPolicy)



delete a ClusterRoleBinding


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the ClusterRoleBinding | 
 **body** | [**V1DeleteOptions**](V1DeleteOptions.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 
 **gracePeriodSeconds** | **int32**| The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. | [optional] 
 **orphanDependents** | **bool**| Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \&quot;orphan\&quot; finalizer will be added to/removed from the object&#39;s finalizers list. Either this field or PropagationPolicy may be set, but not both. | [optional] 
 **propagationPolicy** | **string**| Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. | [optional] 

### Return type

[**V1Status**](v1.Status.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCollectionClusterRole**
> V1Status DeleteCollectionClusterRole($pretty, $continue_, $fieldSelector, $includeUninitialized, $labelSelector, $limit, $resourceVersion, $timeoutSeconds, $watch)



delete collection of ClusterRole


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 
 **continue_** | **string**| The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | [optional] 
 **fieldSelector** | **string**| A selector to restrict the list of returned objects by their fields. Defaults to everything. | [optional] 
 **includeUninitialized** | **bool**| If true, partially initialized resources are included in the response. | [optional] 
 **labelSelector** | **string**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | [optional] 
 **limit** | **int32**| limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | [optional] 
 **resourceVersion** | **string**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | [optional] 
 **timeoutSeconds** | **int32**| Timeout for the list/watch call. | [optional] 
 **watch** | **bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | [optional] 

### Return type

[**V1Status**](v1.Status.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCollectionClusterRoleBinding**
> V1Status DeleteCollectionClusterRoleBinding($pretty, $continue_, $fieldSelector, $includeUninitialized, $labelSelector, $limit, $resourceVersion, $timeoutSeconds, $watch)



delete collection of ClusterRoleBinding


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 
 **continue_** | **string**| The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | [optional] 
 **fieldSelector** | **string**| A selector to restrict the list of returned objects by their fields. Defaults to everything. | [optional] 
 **includeUninitialized** | **bool**| If true, partially initialized resources are included in the response. | [optional] 
 **labelSelector** | **string**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | [optional] 
 **limit** | **int32**| limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | [optional] 
 **resourceVersion** | **string**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | [optional] 
 **timeoutSeconds** | **int32**| Timeout for the list/watch call. | [optional] 
 **watch** | **bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | [optional] 

### Return type

[**V1Status**](v1.Status.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCollectionNamespacedRole**
> V1Status DeleteCollectionNamespacedRole($namespace, $pretty, $continue_, $fieldSelector, $includeUninitialized, $labelSelector, $limit, $resourceVersion, $timeoutSeconds, $watch)



delete collection of Role


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 
 **continue_** | **string**| The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | [optional] 
 **fieldSelector** | **string**| A selector to restrict the list of returned objects by their fields. Defaults to everything. | [optional] 
 **includeUninitialized** | **bool**| If true, partially initialized resources are included in the response. | [optional] 
 **labelSelector** | **string**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | [optional] 
 **limit** | **int32**| limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | [optional] 
 **resourceVersion** | **string**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | [optional] 
 **timeoutSeconds** | **int32**| Timeout for the list/watch call. | [optional] 
 **watch** | **bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | [optional] 

### Return type

[**V1Status**](v1.Status.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCollectionNamespacedRoleBinding**
> V1Status DeleteCollectionNamespacedRoleBinding($namespace, $pretty, $continue_, $fieldSelector, $includeUninitialized, $labelSelector, $limit, $resourceVersion, $timeoutSeconds, $watch)



delete collection of RoleBinding


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 
 **continue_** | **string**| The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | [optional] 
 **fieldSelector** | **string**| A selector to restrict the list of returned objects by their fields. Defaults to everything. | [optional] 
 **includeUninitialized** | **bool**| If true, partially initialized resources are included in the response. | [optional] 
 **labelSelector** | **string**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | [optional] 
 **limit** | **int32**| limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | [optional] 
 **resourceVersion** | **string**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | [optional] 
 **timeoutSeconds** | **int32**| Timeout for the list/watch call. | [optional] 
 **watch** | **bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | [optional] 

### Return type

[**V1Status**](v1.Status.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNamespacedRole**
> V1Status DeleteNamespacedRole($name, $namespace, $body, $pretty, $gracePeriodSeconds, $orphanDependents, $propagationPolicy)



delete a Role


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the Role | 
 **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **body** | [**V1DeleteOptions**](V1DeleteOptions.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 
 **gracePeriodSeconds** | **int32**| The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. | [optional] 
 **orphanDependents** | **bool**| Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \&quot;orphan\&quot; finalizer will be added to/removed from the object&#39;s finalizers list. Either this field or PropagationPolicy may be set, but not both. | [optional] 
 **propagationPolicy** | **string**| Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. | [optional] 

### Return type

[**V1Status**](v1.Status.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNamespacedRoleBinding**
> V1Status DeleteNamespacedRoleBinding($name, $namespace, $body, $pretty, $gracePeriodSeconds, $orphanDependents, $propagationPolicy)



delete a RoleBinding


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the RoleBinding | 
 **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **body** | [**V1DeleteOptions**](V1DeleteOptions.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 
 **gracePeriodSeconds** | **int32**| The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. | [optional] 
 **orphanDependents** | **bool**| Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \&quot;orphan\&quot; finalizer will be added to/removed from the object&#39;s finalizers list. Either this field or PropagationPolicy may be set, but not both. | [optional] 
 **propagationPolicy** | **string**| Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. | [optional] 

### Return type

[**V1Status**](v1.Status.md)

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

# **ListClusterRole**
> V1ClusterRoleList ListClusterRole($pretty, $continue_, $fieldSelector, $includeUninitialized, $labelSelector, $limit, $resourceVersion, $timeoutSeconds, $watch)



list or watch objects of kind ClusterRole


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 
 **continue_** | **string**| The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | [optional] 
 **fieldSelector** | **string**| A selector to restrict the list of returned objects by their fields. Defaults to everything. | [optional] 
 **includeUninitialized** | **bool**| If true, partially initialized resources are included in the response. | [optional] 
 **labelSelector** | **string**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | [optional] 
 **limit** | **int32**| limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | [optional] 
 **resourceVersion** | **string**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | [optional] 
 **timeoutSeconds** | **int32**| Timeout for the list/watch call. | [optional] 
 **watch** | **bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | [optional] 

### Return type

[**V1ClusterRoleList**](v1.ClusterRoleList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClusterRoleBinding**
> V1ClusterRoleBindingList ListClusterRoleBinding($pretty, $continue_, $fieldSelector, $includeUninitialized, $labelSelector, $limit, $resourceVersion, $timeoutSeconds, $watch)



list or watch objects of kind ClusterRoleBinding


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 
 **continue_** | **string**| The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | [optional] 
 **fieldSelector** | **string**| A selector to restrict the list of returned objects by their fields. Defaults to everything. | [optional] 
 **includeUninitialized** | **bool**| If true, partially initialized resources are included in the response. | [optional] 
 **labelSelector** | **string**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | [optional] 
 **limit** | **int32**| limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | [optional] 
 **resourceVersion** | **string**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | [optional] 
 **timeoutSeconds** | **int32**| Timeout for the list/watch call. | [optional] 
 **watch** | **bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | [optional] 

### Return type

[**V1ClusterRoleBindingList**](v1.ClusterRoleBindingList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNamespacedRole**
> V1RoleList ListNamespacedRole($namespace, $pretty, $continue_, $fieldSelector, $includeUninitialized, $labelSelector, $limit, $resourceVersion, $timeoutSeconds, $watch)



list or watch objects of kind Role


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 
 **continue_** | **string**| The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | [optional] 
 **fieldSelector** | **string**| A selector to restrict the list of returned objects by their fields. Defaults to everything. | [optional] 
 **includeUninitialized** | **bool**| If true, partially initialized resources are included in the response. | [optional] 
 **labelSelector** | **string**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | [optional] 
 **limit** | **int32**| limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | [optional] 
 **resourceVersion** | **string**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | [optional] 
 **timeoutSeconds** | **int32**| Timeout for the list/watch call. | [optional] 
 **watch** | **bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | [optional] 

### Return type

[**V1RoleList**](v1.RoleList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNamespacedRoleBinding**
> V1RoleBindingList ListNamespacedRoleBinding($namespace, $pretty, $continue_, $fieldSelector, $includeUninitialized, $labelSelector, $limit, $resourceVersion, $timeoutSeconds, $watch)



list or watch objects of kind RoleBinding


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 
 **continue_** | **string**| The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | [optional] 
 **fieldSelector** | **string**| A selector to restrict the list of returned objects by their fields. Defaults to everything. | [optional] 
 **includeUninitialized** | **bool**| If true, partially initialized resources are included in the response. | [optional] 
 **labelSelector** | **string**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | [optional] 
 **limit** | **int32**| limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | [optional] 
 **resourceVersion** | **string**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | [optional] 
 **timeoutSeconds** | **int32**| Timeout for the list/watch call. | [optional] 
 **watch** | **bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | [optional] 

### Return type

[**V1RoleBindingList**](v1.RoleBindingList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRoleBindingForAllNamespaces**
> V1RoleBindingList ListRoleBindingForAllNamespaces($continue_, $fieldSelector, $includeUninitialized, $labelSelector, $limit, $pretty, $resourceVersion, $timeoutSeconds, $watch)



list or watch objects of kind RoleBinding


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **continue_** | **string**| The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | [optional] 
 **fieldSelector** | **string**| A selector to restrict the list of returned objects by their fields. Defaults to everything. | [optional] 
 **includeUninitialized** | **bool**| If true, partially initialized resources are included in the response. | [optional] 
 **labelSelector** | **string**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | [optional] 
 **limit** | **int32**| limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | [optional] 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 
 **resourceVersion** | **string**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | [optional] 
 **timeoutSeconds** | **int32**| Timeout for the list/watch call. | [optional] 
 **watch** | **bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | [optional] 

### Return type

[**V1RoleBindingList**](v1.RoleBindingList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRoleForAllNamespaces**
> V1RoleList ListRoleForAllNamespaces($continue_, $fieldSelector, $includeUninitialized, $labelSelector, $limit, $pretty, $resourceVersion, $timeoutSeconds, $watch)



list or watch objects of kind Role


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **continue_** | **string**| The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | [optional] 
 **fieldSelector** | **string**| A selector to restrict the list of returned objects by their fields. Defaults to everything. | [optional] 
 **includeUninitialized** | **bool**| If true, partially initialized resources are included in the response. | [optional] 
 **labelSelector** | **string**| A selector to restrict the list of returned objects by their labels. Defaults to everything. | [optional] 
 **limit** | **int32**| limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | [optional] 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 
 **resourceVersion** | **string**| When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | [optional] 
 **timeoutSeconds** | **int32**| Timeout for the list/watch call. | [optional] 
 **watch** | **bool**| Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | [optional] 

### Return type

[**V1RoleList**](v1.RoleList.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchClusterRole**
> V1ClusterRole PatchClusterRole($name, $body, $pretty)



partially update the specified ClusterRole


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the ClusterRole | 
 **body** | [**interface{}**](interface{}.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1ClusterRole**](v1.ClusterRole.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchClusterRoleBinding**
> V1ClusterRoleBinding PatchClusterRoleBinding($name, $body, $pretty)



partially update the specified ClusterRoleBinding


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the ClusterRoleBinding | 
 **body** | [**interface{}**](interface{}.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1ClusterRoleBinding**](v1.ClusterRoleBinding.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchNamespacedRole**
> V1Role PatchNamespacedRole($name, $namespace, $body, $pretty)



partially update the specified Role


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the Role | 
 **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **body** | [**interface{}**](interface{}.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1Role**](v1.Role.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchNamespacedRoleBinding**
> V1RoleBinding PatchNamespacedRoleBinding($name, $namespace, $body, $pretty)



partially update the specified RoleBinding


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the RoleBinding | 
 **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **body** | [**interface{}**](interface{}.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1RoleBinding**](v1.RoleBinding.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadClusterRole**
> V1ClusterRole ReadClusterRole($name, $pretty)



read the specified ClusterRole


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the ClusterRole | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1ClusterRole**](v1.ClusterRole.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadClusterRoleBinding**
> V1ClusterRoleBinding ReadClusterRoleBinding($name, $pretty)



read the specified ClusterRoleBinding


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the ClusterRoleBinding | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1ClusterRoleBinding**](v1.ClusterRoleBinding.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNamespacedRole**
> V1Role ReadNamespacedRole($name, $namespace, $pretty)



read the specified Role


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the Role | 
 **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1Role**](v1.Role.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNamespacedRoleBinding**
> V1RoleBinding ReadNamespacedRoleBinding($name, $namespace, $pretty)



read the specified RoleBinding


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the RoleBinding | 
 **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1RoleBinding**](v1.RoleBinding.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceClusterRole**
> V1ClusterRole ReplaceClusterRole($name, $body, $pretty)



replace the specified ClusterRole


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the ClusterRole | 
 **body** | [**V1ClusterRole**](V1ClusterRole.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1ClusterRole**](v1.ClusterRole.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceClusterRoleBinding**
> V1ClusterRoleBinding ReplaceClusterRoleBinding($name, $body, $pretty)



replace the specified ClusterRoleBinding


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the ClusterRoleBinding | 
 **body** | [**V1ClusterRoleBinding**](V1ClusterRoleBinding.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1ClusterRoleBinding**](v1.ClusterRoleBinding.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceNamespacedRole**
> V1Role ReplaceNamespacedRole($name, $namespace, $body, $pretty)



replace the specified Role


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the Role | 
 **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **body** | [**V1Role**](V1Role.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1Role**](v1.Role.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceNamespacedRoleBinding**
> V1RoleBinding ReplaceNamespacedRoleBinding($name, $namespace, $body, $pretty)



replace the specified RoleBinding


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name of the RoleBinding | 
 **namespace** | **string**| object name and auth scope, such as for teams and projects | 
 **body** | [**V1RoleBinding**](V1RoleBinding.md)|  | 
 **pretty** | **string**| If &#39;true&#39;, then the output is pretty printed. | [optional] 

### Return type

[**V1RoleBinding**](v1.RoleBinding.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

