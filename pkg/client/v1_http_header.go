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

// HTTPHeader describes a custom header to be used in HTTP probes
type V1HttpHeader struct {

	// The header field name
	Name string `json:"name"`

	// The header field value
	Value string `json:"value"`
}
