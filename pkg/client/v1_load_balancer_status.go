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

// LoadBalancerStatus represents the status of a load-balancer.
type V1LoadBalancerStatus struct {

	// Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points.
	Ingress []V1LoadBalancerIngress `json:"ingress,omitempty"`
}
