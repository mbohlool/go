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

// TCPSocketAction describes an action based on opening a socket
type V1TcpSocketAction struct {

	// Optional: Host name to connect to, defaults to the pod IP.
	Host string `json:"host,omitempty"`

	// Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	Port interface{} `json:"port"`
}
