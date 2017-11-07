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

type V1beta1NetworkPolicyPort struct {

	// If specified, the port on the given protocol.  This can either be a numerical or named port on a pod.  If this field is not provided, this matches all port names and numbers. If present, only traffic on the specified protocol AND port will be matched.
	Port interface{} `json:"port,omitempty"`

	// Optional.  The protocol (TCP or UDP) which traffic must match. If not specified, this field defaults to TCP.
	Protocol string `json:"protocol,omitempty"`
}