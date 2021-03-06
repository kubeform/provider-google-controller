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

type CatalogEntryGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CatalogEntryGroupSpec   `json:"spec,omitempty"`
	Status            CatalogEntryGroupStatus `json:"status,omitempty"`
}

type CatalogEntryGroupSpec struct {
	State *CatalogEntryGroupSpecResource `json:"state,omitempty" tf:"-"`

	Resource CatalogEntryGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type CatalogEntryGroupSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// A short name to identify the entry group, for example, "analytics data - jan 2011".
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// The id of the entry group to create. The id must begin with a letter or underscore,
	// contain only English letters, numbers and underscores, and be at most 64 characters.
	EntryGroupID *string `json:"entryGroupID" tf:"entry_group_id"`
	// The resource name of the entry group in URL format. Example: projects/{project}/locations/{location}/entryGroups/{entryGroupId}
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// EntryGroup location region.
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
}

type CatalogEntryGroupStatus struct {
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

// CatalogEntryGroupList is a list of CatalogEntryGroups
type CatalogEntryGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CatalogEntryGroup CRD objects
	Items []CatalogEntryGroup `json:"items,omitempty"`
}
