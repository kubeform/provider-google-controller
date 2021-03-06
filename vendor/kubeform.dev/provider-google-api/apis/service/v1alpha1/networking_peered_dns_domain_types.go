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

type NetworkingPeeredDNSDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkingPeeredDNSDomainSpec   `json:"spec,omitempty"`
	Status            NetworkingPeeredDNSDomainStatus `json:"status,omitempty"`
}

type NetworkingPeeredDNSDomainSpec struct {
	State *NetworkingPeeredDNSDomainSpecResource `json:"state,omitempty" tf:"-"`

	Resource NetworkingPeeredDNSDomainSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type NetworkingPeeredDNSDomainSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The DNS domain name suffix of the peered DNS domain.
	DnsSuffix *string `json:"dnsSuffix" tf:"dns_suffix"`
	// Name of the peered DNS domain.
	Name *string `json:"name" tf:"name"`
	// Network in the consumer project to peer with.
	Network *string `json:"network" tf:"network"`
	// +optional
	Parent *string `json:"parent,omitempty" tf:"parent"`
	// The ID of the project that the service account will be created in. Defaults to the provider project configuration.
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The name of the service to create a peered DNS domain for, e.g. servicenetworking.googleapis.com
	// +optional
	Service *string `json:"service,omitempty" tf:"service"`
}

type NetworkingPeeredDNSDomainStatus struct {
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

// NetworkingPeeredDNSDomainList is a list of NetworkingPeeredDNSDomains
type NetworkingPeeredDNSDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkingPeeredDNSDomain CRD objects
	Items []NetworkingPeeredDNSDomain `json:"items,omitempty"`
}
