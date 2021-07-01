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

type RegionSslCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegionSslCertificateSpec   `json:"spec,omitempty"`
	Status            RegionSslCertificateStatus `json:"status,omitempty"`
}

type RegionSslCertificateSpec struct {
	KubeformOutput *RegionSslCertificateSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource RegionSslCertificateSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type RegionSslCertificateSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The certificate in PEM format.
	// The certificate chain must be no greater than 5 certs long.
	// The chain must include at least one intermediate cert.
	Certificate *string `json:"-" sensitive:"true" tf:"certificate"`
	// The unique identifier for the resource.
	// +optional
	CertificateID *int64 `json:"certificateID,omitempty" tf:"certificate_id"`
	// Creation timestamp in RFC3339 text format.
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp"`
	// An optional description of this resource.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	//
	// These are in the same namespace as the managed SSL certificates.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// Creates a unique name beginning with the specified prefix. Conflicts with name.
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	// The write-only private key in PEM format.
	PrivateKey *string `json:"-" sensitive:"true" tf:"private_key"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The Region in which the created regional ssl certificate should reside.
	// If it is not provided, the provider region is used.
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
}

type RegionSslCertificateStatus struct {
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

// RegionSslCertificateList is a list of RegionSslCertificates
type RegionSslCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RegionSslCertificate CRD objects
	Items []RegionSslCertificate `json:"items,omitempty"`
}
