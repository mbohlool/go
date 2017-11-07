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

// ID Range provides a min/max of an allowed range of IDs.
type V1beta1IdRange struct {

	// Max is the end of the range, inclusive.
	Max int64 `json:"max"`

	// Min is the start of the range, inclusive.
	Min int64 `json:"min"`
}