# V1beta1SeLinuxStrategyOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | **string** | type is the strategy that will dictate the allowable labels that may be set. | [default to null]
**SeLinuxOptions** | [**V1SeLinuxOptions**](v1.SELinuxOptions.md) | seLinuxOptions required to run as; required for MustRunAs More info: https://git.k8s.io/community/contributors/design-proposals/security_context.md | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


