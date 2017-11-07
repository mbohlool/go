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

// SupplementalGroupsStrategyOptions defines the strategy type and options used to create the strategy.
type V1beta1SupplementalGroupsStrategyOptions struct {

	// Ranges are the allowed ranges of supplemental groups.  If you would like to force a single supplemental group then supply a single range with the same start and end.
	Ranges []V1beta1IdRange `json:"ranges,omitempty"`

	// Rule is the strategy that will dictate what supplemental groups is used in the SecurityContext.
	Rule string `json:"rule,omitempty"`
}