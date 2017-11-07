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

// ReplicaSetStatus represents the current status of a ReplicaSet.
type V1beta2ReplicaSetStatus struct {

	// The number of available replicas (ready for at least minReadySeconds) for this replica set.
	AvailableReplicas int32 `json:"availableReplicas,omitempty"`

	// Represents the latest available observations of a replica set's current state.
	Conditions []V1beta2ReplicaSetCondition `json:"conditions,omitempty"`

	// The number of pods that have labels matching the labels of the pod template of the replicaset.
	FullyLabeledReplicas int32 `json:"fullyLabeledReplicas,omitempty"`

	// ObservedGeneration reflects the generation of the most recently observed ReplicaSet.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// The number of ready replicas for this replica set.
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`

	// Replicas is the most recently oberved number of replicas. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller/#what-is-a-replicationcontroller
	Replicas int32 `json:"replicas"`
}
