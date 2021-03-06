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

type BucketAccessControl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketAccessControlSpec   `json:"spec,omitempty"`
	Status            BucketAccessControlStatus `json:"status,omitempty"`
}

type BucketAccessControlSpec struct {
	State *BucketAccessControlSpecResource `json:"state,omitempty" tf:"-"`

	Resource BucketAccessControlSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type BucketAccessControlSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the bucket.
	Bucket *string `json:"bucket" tf:"bucket"`
	// The domain associated with the entity.
	// +optional
	Domain *string `json:"domain,omitempty" tf:"domain"`
	// The email address associated with the entity.
	// +optional
	Email *string `json:"email,omitempty" tf:"email"`
	// The entity holding the permission, in one of the following forms:
	//   user-userId
	//   user-email
	//   group-groupId
	//   group-email
	//   domain-domain
	//   project-team-projectId
	//   allUsers
	//   allAuthenticatedUsers
	// Examples:
	//   The user liz@example.com would be user-liz@example.com.
	//   The group example@googlegroups.com would be
	//   group-example@googlegroups.com.
	//   To refer to all members of the Google Apps for Business domain
	//   example.com, the entity would be domain-example.com.
	Entity *string `json:"entity" tf:"entity"`
	// The access permission for the entity. Possible values: ["OWNER", "READER", "WRITER"]
	// +optional
	Role *string `json:"role,omitempty" tf:"role"`
}

type BucketAccessControlStatus struct {
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

// BucketAccessControlList is a list of BucketAccessControls
type BucketAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BucketAccessControl CRD objects
	Items []BucketAccessControl `json:"items,omitempty"`
}
