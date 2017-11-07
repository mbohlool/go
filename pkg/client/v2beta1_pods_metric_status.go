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

// PodsMetricStatus indicates the current value of a metric describing each pod in the current scale target (for example, transactions-processed-per-second).
type V2beta1PodsMetricStatus struct {

	// currentAverageValue is the current value of the average of the metric across all relevant pods (as a quantity)
	CurrentAverageValue string `json:"currentAverageValue"`

	// metricName is the name of the metric in question
	MetricName string `json:"metricName"`
}
