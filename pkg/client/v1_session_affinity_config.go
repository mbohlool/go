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

// SessionAffinityConfig represents the configurations of session affinity.
type V1SessionAffinityConfig struct {

	// clientIP contains the configurations of Client IP based session affinity.
	ClientIP V1ClientIpConfig `json:"clientIP,omitempty"`
}
