# V1alpha1PriorityClass

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] [default to null]
**Description** | **string** | description is an arbitrary string that usually provides guidelines on when this priority class should be used. | [optional] [default to null]
**GlobalDefault** | **bool** | globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. | [optional] [default to null]
**Kind** | **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] [default to null]
**Metadata** | [**V1ObjectMeta**](v1.ObjectMeta.md) | Standard object&#39;s metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata | [optional] [default to null]
**Value** | **int32** | The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


