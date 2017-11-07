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

// ResourceRule is the list of actions the subject is allowed to perform on resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.
type V1ResourceRule struct {

	// APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are specified, any action requested against one of the enumerated resources in any API group will be allowed.  \"*\" means all.
	ApiGroups []string `json:"apiGroups,omitempty"`

	// ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.  \"*\" means all.
	ResourceNames []string `json:"resourceNames,omitempty"`

	// Resources is a list of resources this rule applies to.  ResourceAll represents all resources.  \"*\" means all.
	Resources []string `json:"resources,omitempty"`

	// Verb is a list of kubernetes resource API verbs, like: get, list, watch, create, update, delete, proxy.  \"*\" means all.
	Verbs []string `json:"verbs"`
}
