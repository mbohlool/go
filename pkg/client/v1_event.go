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

// Event is a report of an event somewhere in the cluster.
type V1Event struct {

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	ApiVersion string `json:"apiVersion,omitempty"`

	// The number of times this event has occurred.
	Count int32 `json:"count,omitempty"`

	// The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
	FirstTimestamp time.Time `json:"firstTimestamp,omitempty"`

	// The object that this event is about.
	InvolvedObject V1ObjectReference `json:"involvedObject"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// The time at which the most recent occurrence of this event was recorded.
	LastTimestamp time.Time `json:"lastTimestamp,omitempty"`

	// A human-readable description of the status of this operation.
	Message string `json:"message,omitempty"`

	// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	Metadata V1ObjectMeta `json:"metadata"`

	// This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
	Reason string `json:"reason,omitempty"`

	// The component reporting this event. Should be a short machine understandable string.
	Source V1EventSource `json:"source,omitempty"`

	// Type of this event (Normal, Warning), new types could be added in the future
	Type_ string `json:"type,omitempty"`
}
