# V1StorageClass

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowVolumeExpansion** | **bool** | AllowVolumeExpansion shows whether the storage class allow volume expand | [optional] [default to null]
**ApiVersion** | **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] [default to null]
**Kind** | **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] [default to null]
**Metadata** | [**V1ObjectMeta**](v1.ObjectMeta.md) | Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata | [optional] [default to null]
**MountOptions** | **[]string** | Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. [\&quot;ro\&quot;, \&quot;soft\&quot;]. Not validated - mount of the PVs will simply fail if one is invalid. | [optional] [default to null]
**Parameters** | **map[string]string** | Parameters holds the parameters for the provisioner that should create volumes of this storage class. | [optional] [default to null]
**Provisioner** | **string** | Provisioner indicates the type of the provisioner. | [default to null]
**ReclaimPolicy** | **string** | Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


