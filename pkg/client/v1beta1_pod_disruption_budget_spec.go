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

// PodDisruptionBudgetSpec is a description of a PodDisruptionBudget.
type V1beta1PodDisruptionBudgetSpec struct {

	// An eviction is allowed if at most \"maxUnavailable\" pods selected by \"selector\" are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with \"minAvailable\".
	MaxUnavailable interface{} `json:"maxUnavailable,omitempty"`

	// An eviction is allowed if at least \"minAvailable\" pods selected by \"selector\" will still be available after the eviction, i.e. even in the absence of the evicted pod.  So for example you can prevent all voluntary evictions by specifying \"100%\".
	MinAvailable interface{} `json:"minAvailable,omitempty"`

	// Label query over pods whose evictions are managed by the disruption budget.
	Selector V1LabelSelector `json:"selector,omitempty"`
}
