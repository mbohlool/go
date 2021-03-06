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

// Info contains versioning information. how we'll want to distribute that information.
type VersionInfo struct {

	BuildDate string `json:"buildDate"`

	Compiler string `json:"compiler"`

	GitCommit string `json:"gitCommit"`

	GitTreeState string `json:"gitTreeState"`

	GitVersion string `json:"gitVersion"`

	GoVersion string `json:"goVersion"`

	Major string `json:"major"`

	Minor string `json:"minor"`

	Platform string `json:"platform"`
}
