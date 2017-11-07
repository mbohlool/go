# V1ResourceRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiGroups** | **[]string** | APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.  \&quot;*\&quot; means all. | [optional] [default to null]
**ResourceNames** | **[]string** | ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.  \&quot;*\&quot; means all. | [optional] [default to null]
**Resources** | **[]string** | Resources is a list of resources this rule applies to.  ResourceAll represents all resources.  \&quot;*\&quot; means all. | [optional] [default to null]
**Verbs** | **[]string** | Verb is a list of kubernetes resource API verbs, like: get, list, watch, create, update, delete, proxy.  \&quot;*\&quot; means all. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


