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

// DownwardAPIVolumeSource represents a volume containing downward API info. Downward API volumes support ownership management and SELinux relabeling.
type V1DownwardApiVolumeSource struct {

	// Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	DefaultMode int32 `json:"defaultMode,omitempty"`

	// Items is a list of downward API volume file
	Items []V1DownwardApiVolumeFile `json:"items,omitempty"`
}
