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

// Initializers tracks the progress of initialization.
type V1Initializers struct {

	// Pending is a list of initializers that must execute in order before this object is visible. When the last pending initializer is removed, and no failing result is set, the initializers struct will be set to nil and the object is considered as initialized and visible to all clients.
	Pending []V1Initializer `json:"pending"`

	// If result is set with the Failure field, the object will be persisted to storage and then deleted, ensuring that other clients can observe the deletion.
	Result V1Status `json:"result,omitempty"`
}
