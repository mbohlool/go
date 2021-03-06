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

// IngressStatus describe the current state of the Ingress.
type V1beta1IngressStatus struct {

	// LoadBalancer contains the current status of the load-balancer.
	LoadBalancer V1LoadBalancerStatus `json:"loadBalancer,omitempty"`
}
