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

type InstanceIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceIamBindingSpec   `json:"spec,omitempty"`
	Status            InstanceIamBindingStatus `json:"status,omitempty"`
}

type InstanceIamBindingSpecCondition struct {
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	Expression  *string `json:"expression" tf:"expression"`
	Title       *string `json:"title" tf:"title"`
}

type InstanceIamBindingSpec struct {
	State *InstanceIamBindingSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstanceIamBindingSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type InstanceIamBindingSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Condition *InstanceIamBindingSpecCondition `json:"condition,omitempty" tf:"condition"`
	// +optional
	Etag         *string `json:"etag,omitempty" tf:"etag"`
	InstanceName *string `json:"instanceName" tf:"instance_name"`
	// +optional
	Location *string  `json:"location,omitempty" tf:"location"`
	Members  []string `json:"members" tf:"members"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	Role    *string `json:"role" tf:"role"`
}

type InstanceIamBindingStatus struct {
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

// InstanceIamBindingList is a list of InstanceIamBindings
type InstanceIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of InstanceIamBinding CRD objects
	Items []InstanceIamBinding `json:"items,omitempty"`
}
