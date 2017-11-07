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

// ResourceFieldSelector represents container resources (cpu, memory) and their output format
type V1ResourceFieldSelector struct {

	// Container name: required for volumes, optional for env vars
	ContainerName string `json:"containerName,omitempty"`

	// Specifies the output format of the exposed resources, defaults to \"1\"
	Divisor string `json:"divisor,omitempty"`

	// Required: resource to select
	Resource string `json:"resource"`
}
