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

// Selects a key from a ConfigMap.
type V1ConfigMapKeySelector struct {

	// The key to select.
	Key string `json:"key"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Specify whether the ConfigMap or it's key must be defined
	Optional bool `json:"optional,omitempty"`
}