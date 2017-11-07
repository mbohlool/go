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

// ContainerState holds a possible state of container. Only one of its members may be specified. If none of them is specified, the default one is ContainerStateWaiting.
type V1ContainerState struct {

	// Details about a running container
	Running V1ContainerStateRunning `json:"running,omitempty"`

	// Details about a terminated container
	Terminated V1ContainerStateTerminated `json:"terminated,omitempty"`

	// Details about a waiting container
	Waiting V1ContainerStateWaiting `json:"waiting,omitempty"`
}