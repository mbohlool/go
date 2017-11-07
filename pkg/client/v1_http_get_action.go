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

// HTTPGetAction describes an action based on HTTP Get requests.
type V1HttpGetAction struct {

	// Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead.
	Host string `json:"host,omitempty"`

	// Custom headers to set in the request. HTTP allows repeated headers.
	HttpHeaders []V1HttpHeader `json:"httpHeaders,omitempty"`

	// Path to access on the HTTP server.
	Path string `json:"path,omitempty"`

	// Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	Port interface{} `json:"port"`

	// Scheme to use for connecting to the host. Defaults to HTTP.
	Scheme string `json:"scheme,omitempty"`
}
