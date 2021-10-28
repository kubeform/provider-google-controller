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

type Fulfillment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FulfillmentSpec   `json:"spec,omitempty"`
	Status            FulfillmentStatus `json:"status,omitempty"`
}

type FulfillmentSpecFeatures struct {
	// The type of the feature that enabled for fulfillment.
	// * SMALLTALK: Fulfillment is enabled for SmallTalk. Possible values: ["SMALLTALK"]
	Type *string `json:"type" tf:"type"`
}

type FulfillmentSpecGenericWebService struct {
	// The password for HTTP Basic authentication.
	// +optional
	Password *string `json:"password,omitempty" tf:"password"`
	// The HTTP request headers to send together with fulfillment requests.
	// +optional
	RequestHeaders *map[string]string `json:"requestHeaders,omitempty" tf:"request_headers"`
	// The fulfillment URI for receiving POST requests. It must use https protocol.
	Uri *string `json:"uri" tf:"uri"`
	// The user name for HTTP Basic authentication.
	// +optional
	Username *string `json:"username,omitempty" tf:"username"`
}

type FulfillmentSpec struct {
	State *FulfillmentSpecResource `json:"state,omitempty" tf:"-"`

	Resource FulfillmentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type FulfillmentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The human-readable name of the fulfillment, unique within the agent.
	DisplayName *string `json:"displayName" tf:"display_name"`
	// Whether fulfillment is enabled.
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// The field defines whether the fulfillment is enabled for certain features.
	// +optional
	Features []FulfillmentSpecFeatures `json:"features,omitempty" tf:"features"`
	// Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
	// +optional
	GenericWebService *FulfillmentSpecGenericWebService `json:"genericWebService,omitempty" tf:"generic_web_service"`
	// The unique identifier of the fulfillment.
	// Format: projects/<Project ID>/agent/fulfillment - projects/<Project ID>/locations/<Location ID>/agent/fulfillment
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
}

type FulfillmentStatus struct {
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

// FulfillmentList is a list of Fulfillments
type FulfillmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Fulfillment CRD objects
	Items []Fulfillment `json:"items,omitempty"`
}
