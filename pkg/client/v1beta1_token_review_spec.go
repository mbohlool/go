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

// TokenReviewSpec is a description of the token authentication request.
type V1beta1TokenReviewSpec struct {

	// Token is the opaque bearer token.
	Token string `json:"token,omitempty"`
}
