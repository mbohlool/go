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

// APIServiceStatus contains derived information about an API server
type V1beta1ApiServiceStatus struct {

	// Current service state of apiService.
	Conditions []V1beta1ApiServiceCondition `json:"conditions,omitempty"`
}