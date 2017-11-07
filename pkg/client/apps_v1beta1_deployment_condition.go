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

// DeploymentCondition describes the state of a deployment at a certain point.
type AppsV1beta1DeploymentCondition struct {

	// Last time the condition transitioned from one status to another.
	LastTransitionTime time.Time `json:"lastTransitionTime,omitempty"`

	// The last time this condition was updated.
	LastUpdateTime time.Time `json:"lastUpdateTime,omitempty"`

	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty"`

	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// Status of the condition, one of True, False, Unknown.
	Status string `json:"status"`

	// Type of deployment condition.
	Type_ string `json:"type"`
}
