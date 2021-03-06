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

// ResourceQuotaSpec defines the desired hard limits to enforce for Quota.
type V1ResourceQuotaSpec struct {

	// Hard is the set of desired hard limits for each named resource. More info: https://git.k8s.io/community/contributors/design-proposals/admission_control_resource_quota.md
	Hard map[string]string `json:"hard,omitempty"`

	// A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.
	Scopes []string `json:"scopes,omitempty"`
}
