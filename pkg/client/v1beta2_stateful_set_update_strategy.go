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

// StatefulSetUpdateStrategy indicates the strategy that the StatefulSet controller will use to perform updates. It includes any additional parameters necessary to perform the update for the indicated strategy.
type V1beta2StatefulSetUpdateStrategy struct {

	// RollingUpdate is used to communicate parameters when Type is RollingUpdateStatefulSetStrategyType.
	RollingUpdate V1beta2RollingUpdateStatefulSetStrategy `json:"rollingUpdate,omitempty"`

	// Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.
	Type_ string `json:"type,omitempty"`
}
