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

// SecretKeySelector selects a key of a Secret.
type V1SecretKeySelector struct {

	// The key of the secret to select from.  Must be a valid secret key.
	Key string `json:"key"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Specify whether the Secret or it's key must be defined
	Optional bool `json:"optional,omitempty"`
}
