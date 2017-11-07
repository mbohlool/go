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

// JSONSchemaPropsOrArray represents a value that can either be a JSONSchemaProps or an array of JSONSchemaProps. Mainly here for serialization purposes.
type V1beta1JsonSchemaPropsOrArray struct {

	JSONSchemas []V1beta1JsonSchemaProps `json:"JSONSchemas"`

	Schema V1beta1JsonSchemaProps `json:"Schema"`
}