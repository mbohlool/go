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

// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
type V1AzureFilePersistentVolumeSource struct {

	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
	ReadOnly bool `json:"readOnly,omitempty"`

	// the name of secret that contains Azure Storage Account Name and Key
	SecretName string `json:"secretName"`

	// the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod
	SecretNamespace string `json:"secretNamespace,omitempty"`

	// Share Name
	ShareName string `json:"shareName"`
}
