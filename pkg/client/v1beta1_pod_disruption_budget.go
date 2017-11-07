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

// PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods
type V1beta1PodDisruptionBudget struct {

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	ApiVersion string `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	Metadata V1ObjectMeta `json:"metadata,omitempty"`

	// Specification of the desired behavior of the PodDisruptionBudget.
	Spec V1beta1PodDisruptionBudgetSpec `json:"spec,omitempty"`

	// Most recently observed status of the PodDisruptionBudget.
	Status V1beta1PodDisruptionBudgetStatus `json:"status,omitempty"`
}
