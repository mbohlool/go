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

// CustomResourceValidation is a list of validation methods for CustomResources.
type V1beta1CustomResourceValidation struct {

	// OpenAPIV3Schema is the OpenAPI v3 schema to be validated against.
	OpenAPIV3Schema V1beta1JsonSchemaProps `json:"openAPIV3Schema,omitempty"`
}