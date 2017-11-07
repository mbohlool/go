# V1SelfSubjectRulesReview

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] [default to null]
**Kind** | **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] [default to null]
**Metadata** | [**V1ObjectMeta**](v1.ObjectMeta.md) |  | [optional] [default to null]
**Spec** | [**V1SelfSubjectRulesReviewSpec**](v1.SelfSubjectRulesReviewSpec.md) | Spec holds information about the request being evaluated. | [default to null]
**Status** | [**V1SubjectRulesReviewStatus**](v1.SubjectRulesReviewStatus.md) | Status is filled in by the server and indicates the set of actions a user can perform. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


