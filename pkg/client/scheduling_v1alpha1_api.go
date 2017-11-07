/* 
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: v1.8.3
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package client

import (
	"net/url"
	"strings"
	"encoding/json"
	"fmt"
)

type Scheduling_v1alpha1Api struct {
	Configuration *Configuration
}

func NewScheduling_v1alpha1Api() *Scheduling_v1alpha1Api {
	configuration := NewConfiguration()
	return &Scheduling_v1alpha1Api{
		Configuration: configuration,
	}
}

func NewScheduling_v1alpha1ApiWithBasePath(basePath string) *Scheduling_v1alpha1Api {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &Scheduling_v1alpha1Api{
		Configuration: configuration,
	}
}

/**
 * 
 * create a PriorityClass
 *
 * @param body 
 * @param pretty If &#39;true&#39;, then the output is pretty printed.
 * @return *V1alpha1PriorityClass
 */
func (a Scheduling_v1alpha1Api) CreatePriorityClass(body V1alpha1PriorityClass, pretty string) (*V1alpha1PriorityClass, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/apis/scheduling.k8s.io/v1alpha1/priorityclasses"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(BearerToken)' required
	// set key with prefix in header
	localVarHeaderParams["authorization"] = a.Configuration.GetAPIKeyWithPrefix("authorization")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("pretty", a.Configuration.APIClient.ParameterToString(pretty, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "*/*",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/yaml",
		"application/vnd.kubernetes.protobuf",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(V1alpha1PriorityClass)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "CreatePriorityClass", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 * delete collection of PriorityClass
 *
 * @param pretty If &#39;true&#39;, then the output is pretty printed.
 * @param continue_ The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.
 * @param fieldSelector A selector to restrict the list of returned objects by their fields. Defaults to everything.
 * @param includeUninitialized If true, partially initialized resources are included in the response.
 * @param labelSelector A selector to restrict the list of returned objects by their labels. Defaults to everything.
 * @param limit limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.
 * @param resourceVersion When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.
 * @param timeoutSeconds Timeout for the list/watch call.
 * @param watch Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.
 * @return *V1Status
 */
func (a Scheduling_v1alpha1Api) DeleteCollectionPriorityClass(pretty string, continue_ string, fieldSelector string, includeUninitialized bool, labelSelector string, limit int32, resourceVersion string, timeoutSeconds int32, watch bool) (*V1Status, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/apis/scheduling.k8s.io/v1alpha1/priorityclasses"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(BearerToken)' required
	// set key with prefix in header
	localVarHeaderParams["authorization"] = a.Configuration.GetAPIKeyWithPrefix("authorization")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("pretty", a.Configuration.APIClient.ParameterToString(pretty, ""))
	localVarQueryParams.Add("continue", a.Configuration.APIClient.ParameterToString(continue_, ""))
	localVarQueryParams.Add("fieldSelector", a.Configuration.APIClient.ParameterToString(fieldSelector, ""))
	localVarQueryParams.Add("includeUninitialized", a.Configuration.APIClient.ParameterToString(includeUninitialized, ""))
	localVarQueryParams.Add("labelSelector", a.Configuration.APIClient.ParameterToString(labelSelector, ""))
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("resourceVersion", a.Configuration.APIClient.ParameterToString(resourceVersion, ""))
	localVarQueryParams.Add("timeoutSeconds", a.Configuration.APIClient.ParameterToString(timeoutSeconds, ""))
	localVarQueryParams.Add("watch", a.Configuration.APIClient.ParameterToString(watch, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "*/*",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/yaml",
		"application/vnd.kubernetes.protobuf",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(V1Status)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeleteCollectionPriorityClass", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 * delete a PriorityClass
 *
 * @param name name of the PriorityClass
 * @param body 
 * @param pretty If &#39;true&#39;, then the output is pretty printed.
 * @param gracePeriodSeconds The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.
 * @param orphanDependents Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \&quot;orphan\&quot; finalizer will be added to/removed from the object&#39;s finalizers list. Either this field or PropagationPolicy may be set, but not both.
 * @param propagationPolicy Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy.
 * @return *V1Status
 */
func (a Scheduling_v1alpha1Api) DeletePriorityClass(name string, body V1DeleteOptions, pretty string, gracePeriodSeconds int32, orphanDependents bool, propagationPolicy string) (*V1Status, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/apis/scheduling.k8s.io/v1alpha1/priorityclasses/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(BearerToken)' required
	// set key with prefix in header
	localVarHeaderParams["authorization"] = a.Configuration.GetAPIKeyWithPrefix("authorization")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("pretty", a.Configuration.APIClient.ParameterToString(pretty, ""))
	localVarQueryParams.Add("gracePeriodSeconds", a.Configuration.APIClient.ParameterToString(gracePeriodSeconds, ""))
	localVarQueryParams.Add("orphanDependents", a.Configuration.APIClient.ParameterToString(orphanDependents, ""))
	localVarQueryParams.Add("propagationPolicy", a.Configuration.APIClient.ParameterToString(propagationPolicy, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "*/*",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/yaml",
		"application/vnd.kubernetes.protobuf",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(V1Status)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "DeletePriorityClass", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 * get available resources
 *
 * @return *V1ApiResourceList
 */
func (a Scheduling_v1alpha1Api) GetAPIResources() (*V1ApiResourceList, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/apis/scheduling.k8s.io/v1alpha1/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(BearerToken)' required
	// set key with prefix in header
	localVarHeaderParams["authorization"] = a.Configuration.GetAPIKeyWithPrefix("authorization")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json", "application/yaml", "application/vnd.kubernetes.protobuf",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/yaml",
		"application/vnd.kubernetes.protobuf",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(V1ApiResourceList)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetAPIResources", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 * list or watch objects of kind PriorityClass
 *
 * @param pretty If &#39;true&#39;, then the output is pretty printed.
 * @param continue_ The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.
 * @param fieldSelector A selector to restrict the list of returned objects by their fields. Defaults to everything.
 * @param includeUninitialized If true, partially initialized resources are included in the response.
 * @param labelSelector A selector to restrict the list of returned objects by their labels. Defaults to everything.
 * @param limit limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.
 * @param resourceVersion When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv.
 * @param timeoutSeconds Timeout for the list/watch call.
 * @param watch Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.
 * @return *V1alpha1PriorityClassList
 */
func (a Scheduling_v1alpha1Api) ListPriorityClass(pretty string, continue_ string, fieldSelector string, includeUninitialized bool, labelSelector string, limit int32, resourceVersion string, timeoutSeconds int32, watch bool) (*V1alpha1PriorityClassList, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/apis/scheduling.k8s.io/v1alpha1/priorityclasses"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(BearerToken)' required
	// set key with prefix in header
	localVarHeaderParams["authorization"] = a.Configuration.GetAPIKeyWithPrefix("authorization")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("pretty", a.Configuration.APIClient.ParameterToString(pretty, ""))
	localVarQueryParams.Add("continue", a.Configuration.APIClient.ParameterToString(continue_, ""))
	localVarQueryParams.Add("fieldSelector", a.Configuration.APIClient.ParameterToString(fieldSelector, ""))
	localVarQueryParams.Add("includeUninitialized", a.Configuration.APIClient.ParameterToString(includeUninitialized, ""))
	localVarQueryParams.Add("labelSelector", a.Configuration.APIClient.ParameterToString(labelSelector, ""))
	localVarQueryParams.Add("limit", a.Configuration.APIClient.ParameterToString(limit, ""))
	localVarQueryParams.Add("resourceVersion", a.Configuration.APIClient.ParameterToString(resourceVersion, ""))
	localVarQueryParams.Add("timeoutSeconds", a.Configuration.APIClient.ParameterToString(timeoutSeconds, ""))
	localVarQueryParams.Add("watch", a.Configuration.APIClient.ParameterToString(watch, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "*/*",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/yaml",
		"application/vnd.kubernetes.protobuf",
		"application/json;stream=watch",
		"application/vnd.kubernetes.protobuf;stream=watch",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(V1alpha1PriorityClassList)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ListPriorityClass", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 * partially update the specified PriorityClass
 *
 * @param name name of the PriorityClass
 * @param body 
 * @param pretty If &#39;true&#39;, then the output is pretty printed.
 * @return *V1alpha1PriorityClass
 */
func (a Scheduling_v1alpha1Api) PatchPriorityClass(name string, body interface{}, pretty string) (*V1alpha1PriorityClass, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Patch")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/apis/scheduling.k8s.io/v1alpha1/priorityclasses/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(BearerToken)' required
	// set key with prefix in header
	localVarHeaderParams["authorization"] = a.Configuration.GetAPIKeyWithPrefix("authorization")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("pretty", a.Configuration.APIClient.ParameterToString(pretty, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json-patch+json", "application/merge-patch+json", "application/strategic-merge-patch+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/yaml",
		"application/vnd.kubernetes.protobuf",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(V1alpha1PriorityClass)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "PatchPriorityClass", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 * read the specified PriorityClass
 *
 * @param name name of the PriorityClass
 * @param pretty If &#39;true&#39;, then the output is pretty printed.
 * @param exact Should the export be exact.  Exact export maintains cluster-specific fields like &#39;Namespace&#39;.
 * @param export Should this value be exported.  Export strips fields that a user can not specify.
 * @return *V1alpha1PriorityClass
 */
func (a Scheduling_v1alpha1Api) ReadPriorityClass(name string, pretty string, exact bool, export bool) (*V1alpha1PriorityClass, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/apis/scheduling.k8s.io/v1alpha1/priorityclasses/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(BearerToken)' required
	// set key with prefix in header
	localVarHeaderParams["authorization"] = a.Configuration.GetAPIKeyWithPrefix("authorization")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("pretty", a.Configuration.APIClient.ParameterToString(pretty, ""))
	localVarQueryParams.Add("exact", a.Configuration.APIClient.ParameterToString(exact, ""))
	localVarQueryParams.Add("export", a.Configuration.APIClient.ParameterToString(export, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "*/*",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/yaml",
		"application/vnd.kubernetes.protobuf",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(V1alpha1PriorityClass)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReadPriorityClass", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 * replace the specified PriorityClass
 *
 * @param name name of the PriorityClass
 * @param body 
 * @param pretty If &#39;true&#39;, then the output is pretty printed.
 * @return *V1alpha1PriorityClass
 */
func (a Scheduling_v1alpha1Api) ReplacePriorityClass(name string, body V1alpha1PriorityClass, pretty string) (*V1alpha1PriorityClass, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/apis/scheduling.k8s.io/v1alpha1/priorityclasses/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(BearerToken)' required
	// set key with prefix in header
	localVarHeaderParams["authorization"] = a.Configuration.GetAPIKeyWithPrefix("authorization")
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("pretty", a.Configuration.APIClient.ParameterToString(pretty, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "*/*",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/yaml",
		"application/vnd.kubernetes.protobuf",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(V1alpha1PriorityClass)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "ReplacePriorityClass", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}
