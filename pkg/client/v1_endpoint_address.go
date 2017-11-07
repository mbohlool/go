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

// EndpointAddress is a tuple that describes single IP address.
type V1EndpointAddress struct {

	// The Hostname of this endpoint
	Hostname string `json:"hostname,omitempty"`

	// The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24). IPv6 is also accepted but not fully supported on all platforms. Also, certain kubernetes components, like kube-proxy, are not IPv6 ready.
	Ip string `json:"ip"`

	// Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.
	NodeName string `json:"nodeName,omitempty"`

	// Reference to object providing the endpoint.
	TargetRef V1ObjectReference `json:"targetRef,omitempty"`
}
