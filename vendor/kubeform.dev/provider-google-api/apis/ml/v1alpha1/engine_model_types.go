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

type EngineModel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EngineModelSpec   `json:"spec,omitempty"`
	Status            EngineModelStatus `json:"status,omitempty"`
}

type EngineModelSpecDefaultVersion struct {
	// The name specified for the version when it was created.
	Name *string `json:"name" tf:"name"`
}

type EngineModelSpec struct {
	State *EngineModelSpecResource `json:"state,omitempty" tf:"-"`

	Resource EngineModelSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type EngineModelSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The default version of the model. This version will be used to handle
	// prediction requests that do not specify a version.
	// +optional
	DefaultVersion *EngineModelSpecDefaultVersion `json:"defaultVersion,omitempty" tf:"default_version"`
	// The description specified for the model when it was created.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// One or more labels that you can add, to organize your models.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// The name specified for the model.
	Name *string `json:"name" tf:"name"`
	// If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging
	// +optional
	OnlinePredictionConsoleLogging *bool `json:"onlinePredictionConsoleLogging,omitempty" tf:"online_prediction_console_logging"`
	// If true, online prediction access logs are sent to StackDriver Logging.
	// +optional
	OnlinePredictionLogging *bool `json:"onlinePredictionLogging,omitempty" tf:"online_prediction_logging"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The list of regions where the model is going to be deployed.
	// Currently only one region per model is supported
	// +optional
	Regions []string `json:"regions,omitempty" tf:"regions"`
}

type EngineModelStatus struct {
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

// EngineModelList is a list of EngineModels
type EngineModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EngineModel CRD objects
	Items []EngineModel `json:"items,omitempty"`
}
