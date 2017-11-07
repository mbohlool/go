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

// VolumeMount describes a mounting of a Volume within a container.
type V1VolumeMount struct {

	// Path within the container at which the volume should be mounted.  Must not contain ':'.
	MountPath string `json:"mountPath"`

	// mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationHostToContainer is used. This field is alpha in 1.8 and can be reworked or removed in a future release.
	MountPropagation string `json:"mountPropagation,omitempty"`

	// This must match the Name of a Volume.
	Name string `json:"name"`

	// Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.
	ReadOnly bool `json:"readOnly,omitempty"`

	// Path within the volume from which the container's volume should be mounted. Defaults to \"\" (volume's root).
	SubPath string `json:"subPath,omitempty"`
}
