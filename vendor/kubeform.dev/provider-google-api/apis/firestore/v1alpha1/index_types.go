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

type Index struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IndexSpec   `json:"spec,omitempty"`
	Status            IndexStatus `json:"status,omitempty"`
}

type IndexSpecFields struct {
	// Indicates that this field supports operations on arrayValues. Only one of 'order' and 'arrayConfig' can
	// be specified. Possible values: ["CONTAINS"]
	// +optional
	ArrayConfig *string `json:"arrayConfig,omitempty" tf:"array_config"`
	// Name of the field.
	// +optional
	FieldPath *string `json:"fieldPath,omitempty" tf:"field_path"`
	// Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
	// Only one of 'order' and 'arrayConfig' can be specified. Possible values: ["ASCENDING", "DESCENDING"]
	// +optional
	Order *string `json:"order,omitempty" tf:"order"`
}

type IndexSpec struct {
	State *IndexSpecResource `json:"state,omitempty" tf:"-"`

	Resource IndexSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type IndexSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The collection being indexed.
	Collection *string `json:"collection" tf:"collection"`
	// The Firestore database id. Defaults to '"(default)"'.
	// +optional
	Database *string `json:"database,omitempty" tf:"database"`
	// The fields supported by this index. The last field entry is always for
	// the field path '__name__'. If, on creation, '__name__' was not
	// specified as the last field, it will be added automatically with the
	// same direction as that of the last field defined. If the final field
	// in a composite index is not directional, the '__name__' will be
	// ordered '"ASCENDING"' (unless explicitly specified otherwise).
	// +kubebuilder:validation:MinItems=2
	Fields []IndexSpecFields `json:"fields" tf:"fields"`
	// A server defined name for this index. Format:
	// 'projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}'
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The scope at which a query is run. Default value: "COLLECTION" Possible values: ["COLLECTION", "COLLECTION_GROUP"]
	// +optional
	QueryScope *string `json:"queryScope,omitempty" tf:"query_scope"`
}

type IndexStatus struct {
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

// IndexList is a list of Indexs
type IndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Index CRD objects
	Items []Index `json:"items,omitempty"`
}
