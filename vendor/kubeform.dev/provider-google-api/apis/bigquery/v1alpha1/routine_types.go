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

type Routine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoutineSpec   `json:"spec,omitempty"`
	Status            RoutineStatus `json:"status,omitempty"`
}

type RoutineSpecArguments struct {
	// Defaults to FIXED_TYPE. Default value: "FIXED_TYPE" Possible values: ["FIXED_TYPE", "ANY_TYPE"]
	// +optional
	ArgumentKind *string `json:"argumentKind,omitempty" tf:"argument_kind"`
	// A JSON schema for the data type. Required unless argumentKind = ANY_TYPE.
	// ~>**NOTE**: Because this field expects a JSON string, any changes to the string
	// will create a diff, even if the JSON itself hasn't changed. If the API returns
	// a different value for the same schema, e.g. it switched the order of values
	// or replaced STRUCT field type with RECORD field type, we currently cannot
	// suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	// +optional
	DataType *string `json:"dataType,omitempty" tf:"data_type"`
	// Specifies whether the argument is input or output. Can be set for procedures only. Possible values: ["IN", "OUT", "INOUT"]
	// +optional
	Mode *string `json:"mode,omitempty" tf:"mode"`
	// The name of this argument. Can be absent for function return argument.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type RoutineSpec struct {
	State *RoutineSpecResource `json:"state,omitempty" tf:"-"`

	Resource RoutineSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type RoutineSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Input/output argument of a function or a stored procedure.
	// +optional
	Arguments []RoutineSpecArguments `json:"arguments,omitempty" tf:"arguments"`
	// The time when this routine was created, in milliseconds since the
	// epoch.
	// +optional
	CreationTime *int64 `json:"creationTime,omitempty" tf:"creation_time"`
	// The ID of the dataset containing this routine
	DatasetID *string `json:"datasetID" tf:"dataset_id"`
	// The body of the routine. For functions, this is the expression in the AS clause.
	// If language=SQL, it is the substring inside (but excluding) the parentheses.
	DefinitionBody *string `json:"definitionBody" tf:"definition_body"`
	// The description of the routine if defined.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The determinism level of the JavaScript UDF if defined. Possible values: ["DETERMINISM_LEVEL_UNSPECIFIED", "DETERMINISTIC", "NOT_DETERMINISTIC"]
	// +optional
	DeterminismLevel *string `json:"determinismLevel,omitempty" tf:"determinism_level"`
	// Optional. If language = "JAVASCRIPT", this field stores the path of the
	// imported JAVASCRIPT libraries.
	// +optional
	ImportedLibraries []string `json:"importedLibraries,omitempty" tf:"imported_libraries"`
	// The language of the routine. Possible values: ["SQL", "JAVASCRIPT"]
	// +optional
	Language *string `json:"language,omitempty" tf:"language"`
	// The time when this routine was modified, in milliseconds since the
	// epoch.
	// +optional
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty" tf:"last_modified_time"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Optional. Can be set only if routineType = "TABLE_VALUED_FUNCTION".
	//
	// If absent, the return table type is inferred from definitionBody at query time in each query
	// that references this routine. If present, then the columns in the evaluated table result will
	// be cast to match the column types specificed in return table type, at query time.
	// +optional
	ReturnTableType *string `json:"returnTableType,omitempty" tf:"return_table_type"`
	// A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
	// If absent, the return type is inferred from definitionBody at query time in each query
	// that references this routine. If present, then the evaluated result will be cast to
	// the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON
	// string, any changes to the string will create a diff, even if the JSON itself hasn't
	// changed. If the API returns a different value for the same schema, e.g. it switche
	// d the order of values or replaced STRUCT field type with RECORD field type, we currently
	// cannot suppress the recurring diff this causes. As a workaround, we recommend using
	// the schema as returned by the API.
	// +optional
	ReturnType *string `json:"returnType,omitempty" tf:"return_type"`
	// The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.
	RoutineID *string `json:"routineID" tf:"routine_id"`
	// The type of routine. Possible values: ["SCALAR_FUNCTION", "PROCEDURE", "TABLE_VALUED_FUNCTION"]
	// +optional
	RoutineType *string `json:"routineType,omitempty" tf:"routine_type"`
}

type RoutineStatus struct {
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

// RoutineList is a list of Routines
type RoutineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Routine CRD objects
	Items []Routine `json:"items,omitempty"`
}
