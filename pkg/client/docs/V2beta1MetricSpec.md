# V2beta1MetricSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | [**V2beta1ObjectMetricSource**](v2beta1.ObjectMetricSource.md) | object refers to a metric describing a single kubernetes object (for example, hits-per-second on an Ingress object). | [optional] [default to null]
**Pods** | [**V2beta1PodsMetricSource**](v2beta1.PodsMetricSource.md) | pods refers to a metric describing each pod in the current scale target (for example, transactions-processed-per-second).  The values will be averaged together before being compared to the target value. | [optional] [default to null]
**Resource** | [**V2beta1ResourceMetricSource**](v2beta1.ResourceMetricSource.md) | resource refers to a resource metric (such as those specified in requests and limits) known to Kubernetes describing each pod in the current scale target (e.g. CPU or memory). Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the \&quot;pods\&quot; source. | [optional] [default to null]
**Type_** | **string** | type is the type of metric source.  It should match one of the fields below. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


