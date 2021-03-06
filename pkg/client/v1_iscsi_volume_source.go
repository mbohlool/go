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

// Represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.
type V1IscsiVolumeSource struct {

	// whether support iSCSI Discovery CHAP authentication
	ChapAuthDiscovery bool `json:"chapAuthDiscovery,omitempty"`

	// whether support iSCSI Session CHAP authentication
	ChapAuthSession bool `json:"chapAuthSession,omitempty"`

	// Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: \"ext4\", \"xfs\", \"ntfs\". Implicitly inferred to be \"ext4\" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
	FsType string `json:"fsType,omitempty"`

	// Custom iSCSI initiator name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.
	InitiatorName string `json:"initiatorName,omitempty"`

	// Target iSCSI Qualified Name.
	Iqn string `json:"iqn"`

	// Optional: Defaults to 'default' (tcp). iSCSI interface name that uses an iSCSI transport.
	IscsiInterface string `json:"iscsiInterface,omitempty"`

	// iSCSI target lun number.
	Lun int32 `json:"lun"`

	// iSCSI target portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	Portals []string `json:"portals,omitempty"`

	// ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.
	ReadOnly bool `json:"readOnly,omitempty"`

	// CHAP secret for iSCSI target and initiator authentication
	SecretRef V1LocalObjectReference `json:"secretRef,omitempty"`

	// iSCSI target portal. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
	TargetPortal string `json:"targetPortal"`
}
