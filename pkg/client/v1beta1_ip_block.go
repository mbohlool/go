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

// IPBlock describes a particular CIDR (Ex. \"192.168.1.1/24\") that is allowed to the pods matched by a NetworkPolicySpec's podSelector. The except entry describes CIDRs that should not be included within this rule.
type V1beta1IpBlock struct {

	// CIDR is a string representing the IP Block Valid examples are \"192.168.1.1/24\"
	Cidr string `json:"cidr"`

	// Except is a slice of CIDRs that should not be included within an IP Block Valid examples are \"192.168.1.1/24\" Except values will be rejected if they are outside the CIDR range
	Except []string `json:"except,omitempty"`
}
