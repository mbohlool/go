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

// specification of a horizontal pod autoscaler.
type V1HorizontalPodAutoscalerSpec struct {

	// upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
	MaxReplicas int32 `json:"maxReplicas"`

	// lower limit for the number of pods that can be set by the autoscaler, default 1.
	MinReplicas int32 `json:"minReplicas,omitempty"`

	// reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.
	ScaleTargetRef V1CrossVersionObjectReference `json:"scaleTargetRef"`

	// target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
	TargetCPUUtilizationPercentage int32 `json:"targetCPUUtilizationPercentage,omitempty"`
}
