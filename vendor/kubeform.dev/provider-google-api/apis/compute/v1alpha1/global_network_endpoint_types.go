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

type GlobalNetworkEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlobalNetworkEndpointSpec   `json:"spec,omitempty"`
	Status            GlobalNetworkEndpointStatus `json:"status,omitempty"`
}

type GlobalNetworkEndpointSpec struct {
	State *GlobalNetworkEndpointSpecResource `json:"state,omitempty" tf:"-"`

	Resource GlobalNetworkEndpointSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type GlobalNetworkEndpointSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Fully qualified domain name of network endpoint.
	// This can only be specified when network_endpoint_type of the NEG is INTERNET_FQDN_PORT.
	// +optional
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn"`
	// The global network endpoint group this endpoint is part of.
	GlobalNetworkEndpointGroup *string `json:"globalNetworkEndpointGroup" tf:"global_network_endpoint_group"`
	// IPv4 address external endpoint.
	// +optional
	IpAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`
	// Port number of the external endpoint.
	Port *int64 `json:"port" tf:"port"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
}

type GlobalNetworkEndpointStatus struct {
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

// GlobalNetworkEndpointList is a list of GlobalNetworkEndpoints
type GlobalNetworkEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlobalNetworkEndpoint CRD objects
	Items []GlobalNetworkEndpoint `json:"items,omitempty"`
}
