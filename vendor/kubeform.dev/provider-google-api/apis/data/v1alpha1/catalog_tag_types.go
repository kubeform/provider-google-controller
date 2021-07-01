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

type CatalogTag struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CatalogTagSpec   `json:"spec,omitempty"`
	Status            CatalogTagStatus `json:"status,omitempty"`
}

type CatalogTagSpecFields struct {
	// Holds the value for a tag field with boolean type.
	// +optional
	BoolValue *bool `json:"boolValue,omitempty" tf:"bool_value"`
	// The display name of this field
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// Holds the value for a tag field with double type.
	// +optional
	DoubleValue *float64 `json:"doubleValue,omitempty" tf:"double_value"`
	// The display name of the enum value.
	// +optional
	EnumValue *string `json:"enumValue,omitempty" tf:"enum_value"`
	FieldName *string `json:"fieldName" tf:"field_name"`
	// The order of this field with respect to other fields in this tag. For example, a higher value can indicate
	// a more important field. The value can be negative. Multiple fields can have the same order, and field orders
	// within a tag do not have to be sequential.
	// +optional
	Order *int64 `json:"order,omitempty" tf:"order"`
	// Holds the value for a tag field with string type.
	// +optional
	StringValue *string `json:"stringValue,omitempty" tf:"string_value"`
	// Holds the value for a tag field with timestamp type.
	// +optional
	TimestampValue *string `json:"timestampValue,omitempty" tf:"timestamp_value"`
}

type CatalogTagSpec struct {
	KubeformOutput *CatalogTagSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource CatalogTagSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type CatalogTagSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Resources like Entry can have schemas associated with them. This scope allows users to attach tags to an
	// individual column based on that schema.
	//
	// For attaching a tag to a nested column, use '.' to separate the column names. Example:
	// 'outer_column.inner_column'
	// +optional
	Column *string `json:"column,omitempty" tf:"column"`
	// This maps the ID of a tag field to the value of and additional information about that field.
	// Valid field IDs are defined by the tag's template. A tag must have at least 1 field and at most 500 fields.
	Fields []CatalogTagSpecFields `json:"fields" tf:"fields"`
	// The resource name of the tag in URL format. Example:
	// projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/entries/{entryId}/tags/{tag_id} or
	// projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/tags/{tag_id}
	// where tag_id is a system-generated identifier. Note that this Tag may not actually be stored in the location in this name.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The name of the parent this tag is attached to. This can be the name of an entry or an entry group. If an entry group, the tag will be attached to
	// all entries in that group.
	// +optional
	Parent *string `json:"parent,omitempty" tf:"parent"`
	// The resource name of the tag template that this tag uses. Example:
	// projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}
	// This field cannot be modified after creation.
	Template *string `json:"template" tf:"template"`
	// The display name of the tag template.
	// +optional
	TemplateDisplayname *string `json:"templateDisplayname,omitempty" tf:"template_displayname"`
}

type CatalogTagStatus struct {
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

// CatalogTagList is a list of CatalogTags
type CatalogTagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CatalogTag CRD objects
	Items []CatalogTag `json:"items,omitempty"`
}
