# V1alpha1AdmissionHookClientConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaBundle** | **string** | CABundle is a PEM encoded CA bundle which will be used to validate webhook&#39;s server certificate. Required | [default to null]
**Service** | [**V1alpha1ServiceReference**](v1alpha1.ServiceReference.md) | Service is a reference to the service for this webhook. If there is only one port open for the service, that port will be used. If there are multiple ports open, port 443 will be used if it is open, otherwise it is an error. Required | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


