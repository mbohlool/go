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

type V1beta1ApiServiceCondition struct {

	// Last time the condition transitioned from one status to another.
	LastTransitionTime time.Time `json:"lastTransitionTime,omitempty"`

	// Human-readable message indicating details about last transition.
	Message string `json:"message,omitempty"`

	// Unique, one-word, CamelCase reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// Status is the status of the condition. Can be True, False, Unknown.
	Status string `json:"status"`

	// Type is the type of the condition.
	Type_ string `json:"type"`
}
