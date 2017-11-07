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

// SecurityContext holds security configuration that will be applied to a container. Some fields are present in both SecurityContext and PodSecurityContext.  When both are set, the values in SecurityContext take precedence.
type V1SecurityContext struct {

	// AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN
	AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation,omitempty"`

	// The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime.
	Capabilities V1Capabilities `json:"capabilities,omitempty"`

	// Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.
	Privileged bool `json:"privileged,omitempty"`

	// Whether this container has a read-only root filesystem. Default is false.
	ReadOnlyRootFilesystem bool `json:"readOnlyRootFilesystem,omitempty"`

	// Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	RunAsNonRoot bool `json:"runAsNonRoot,omitempty"`

	// The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	RunAsUser int64 `json:"runAsUser,omitempty"`

	// The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	SeLinuxOptions V1SeLinuxOptions `json:"seLinuxOptions,omitempty"`
}
