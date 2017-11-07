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

// A null or empty node selector term matches no objects.
type V1NodeSelectorTerm struct {

	// Required. A list of node selector requirements. The requirements are ANDed.
	MatchExpressions []V1NodeSelectorRequirement `json:"matchExpressions"`
}
