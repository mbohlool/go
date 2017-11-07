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

import (
	"time"
)

// ContainerStateRunning is a running state of a container.
type V1ContainerStateRunning struct {

	// Time at which the container was last (re-)started
	StartedAt time.Time `json:"startedAt,omitempty"`
}
