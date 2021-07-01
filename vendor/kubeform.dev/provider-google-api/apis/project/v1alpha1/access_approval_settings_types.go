/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type AccessApprovalSettings struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccessApprovalSettingsSpec   `json:"spec,omitempty"`
	Status            AccessApprovalSettingsStatus `json:"status,omitempty"`
}

type AccessApprovalSettingsSpecEnrolledServices struct {
	// The product for which Access Approval will be enrolled. Allowed values are listed (case-sensitive):
	//   all
	//   appengine.googleapis.com
	//   bigquery.googleapis.com
	//   bigtable.googleapis.com
	//   cloudkms.googleapis.com
	//   compute.googleapis.com
	//   dataflow.googleapis.com
	//   iam.googleapis.com
	//   pubsub.googleapis.com
	//   storage.googleapis.com
	CloudProduct *string `json:"cloudProduct" tf:"cloud_product"`
	// The enrollment level of the service. Default value: "BLOCK_ALL" Possible values: ["BLOCK_ALL"]
	// +optional
	EnrollmentLevel *string `json:"enrollmentLevel,omitempty" tf:"enrollment_level"`
}

type AccessApprovalSettingsSpec struct {
	KubeformOutput *AccessApprovalSettingsSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource AccessApprovalSettingsSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type AccessApprovalSettingsSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// If the field is true, that indicates that at least one service is enrolled for Access Approval in one or more ancestors of the Project.
	// +optional
	EnrolledAncestor *bool `json:"enrolledAncestor,omitempty" tf:"enrolled_ancestor"`
	// A list of Google Cloud Services for which the given resource has Access Approval enrolled.
	// Access requests for the resource given by name against any of these services contained here will be required
	// to have explicit approval. Enrollment can only be done on an all or nothing basis.
	//
	// A maximum of 10 enrolled services will be enforced, to be expanded as the set of supported services is expanded.
	EnrolledServices []AccessApprovalSettingsSpecEnrolledServices `json:"enrolledServices" tf:"enrolled_services"`
	// The resource name of the settings. Format is "projects/{project_id}/accessApprovalSettings"
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// A list of email addresses to which notifications relating to approval requests should be sent.
	// Notifications relating to a resource will be sent to all emails in the settings of ancestor
	// resources of that resource. A maximum of 50 email addresses are allowed.
	// +optional
	// +kubebuilder:validation:MaxItems=50
	NotificationEmails []string `json:"notificationEmails,omitempty" tf:"notification_emails"`
	// Deprecated in favor of 'project_id'
	// +optional
	// Deprecated
	Project *string `json:"project,omitempty" tf:"project"`
	// ID of the project of the access approval settings.
	ProjectID *string `json:"projectID" tf:"project_id"`
}

type AccessApprovalSettingsStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AccessApprovalSettingsList is a list of AccessApprovalSettingss
type AccessApprovalSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AccessApprovalSettings CRD objects
	Items []AccessApprovalSettings `json:"items,omitempty"`
}
