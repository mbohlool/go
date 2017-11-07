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

// The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)
type V1WeightedPodAffinityTerm struct {

	// Required. A pod affinity term, associated with the corresponding weight.
	PodAffinityTerm V1PodAffinityTerm `json:"podAffinityTerm"`

	// weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
	Weight int32 `json:"weight"`
}
