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

// JSONSchemaPropsOrStringArray represents a JSONSchemaProps or a string array.
type V1beta1JsonSchemaPropsOrStringArray struct {

	Property []string `json:"Property"`

	Schema V1beta1JsonSchemaProps `json:"Schema"`
}
