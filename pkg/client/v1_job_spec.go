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

// JobSpec describes how the job execution will look like.
type V1JobSpec struct {

	// Specifies the duration in seconds relative to the startTime that the job may be active before the system tries to terminate it; value must be positive integer
	ActiveDeadlineSeconds int64 `json:"activeDeadlineSeconds,omitempty"`

	// Specifies the number of retries before marking this job failed. Defaults to 6
	BackoffLimit int32 `json:"backoffLimit,omitempty"`

	// Specifies the desired number of successfully finished pods the job should be run with.  Setting to nil means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value.  Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Completions int32 `json:"completions,omitempty"`

	// manualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template.  When true, the user is responsible for picking unique labels and specifying the selector.  Failure to pick a unique label may cause this and other jobs to not function correctly.  However, You may see `manualSelector=true` in jobs that were created with the old `extensions/v1beta1` API. More info: https://git.k8s.io/community/contributors/design-proposals/selector-generation.md
	ManualSelector bool `json:"manualSelector,omitempty"`

	// Specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Parallelism int32 `json:"parallelism,omitempty"`

	// A label query over pods that should match the pod count. Normally, the system sets this field for you. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
	Selector V1LabelSelector `json:"selector,omitempty"`

	// Describes the pod that will be created when executing a job. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	Template V1PodTemplateSpec `json:"template"`
}