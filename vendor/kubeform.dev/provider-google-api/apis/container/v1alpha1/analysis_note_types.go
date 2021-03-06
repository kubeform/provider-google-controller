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

type AnalysisNote struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AnalysisNoteSpec   `json:"spec,omitempty"`
	Status            AnalysisNoteStatus `json:"status,omitempty"`
}

type AnalysisNoteSpecAttestationAuthorityHint struct {
	// The human readable name of this Attestation Authority, for
	// example "qa".
	HumanReadableName *string `json:"humanReadableName" tf:"human_readable_name"`
}

type AnalysisNoteSpecAttestationAuthority struct {
	// This submessage provides human-readable hints about the purpose of
	// the AttestationAuthority. Because the name of a Note acts as its
	// resource reference, it is important to disambiguate the canonical
	// name of the Note (which might be a UUID for security purposes)
	// from "readable" names more suitable for debug output. Note that
	// these hints should NOT be used to look up AttestationAuthorities
	// in security sensitive contexts, such as when looking up
	// Attestations to verify.
	Hint *AnalysisNoteSpecAttestationAuthorityHint `json:"hint" tf:"hint"`
}

type AnalysisNoteSpecRelatedURL struct {
	// Label to describe usage of the URL
	// +optional
	Label *string `json:"label,omitempty" tf:"label"`
	// Specific URL associated with the resource.
	Url *string `json:"url" tf:"url"`
}

type AnalysisNoteSpec struct {
	State *AnalysisNoteSpecResource `json:"state,omitempty" tf:"-"`

	Resource AnalysisNoteSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AnalysisNoteSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Note kind that represents a logical attestation "role" or "authority".
	// For example, an organization might have one AttestationAuthority for
	// "QA" and one for "build". This Note is intended to act strictly as a
	// grouping mechanism for the attached Occurrences (Attestations). This
	// grouping mechanism also provides a security boundary, since IAM ACLs
	// gate the ability for a principle to attach an Occurrence to a given
	// Note. It also provides a single point of lookup to find all attached
	// Attestation Occurrences, even if they don't all live in the same
	// project.
	AttestationAuthority *AnalysisNoteSpecAttestationAuthority `json:"attestationAuthority" tf:"attestation_authority"`
	// The time this note was created.
	// +optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time"`
	// Time of expiration for this note. Leave empty if note does not expire.
	// +optional
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time"`
	// The type of analysis this note describes
	// +optional
	Kind *string `json:"kind,omitempty" tf:"kind"`
	// A detailed description of the note
	// +optional
	LongDescription *string `json:"longDescription,omitempty" tf:"long_description"`
	// The name of the note.
	Name *string `json:"name" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Names of other notes related to this note.
	// +optional
	RelatedNoteNames []string `json:"relatedNoteNames,omitempty" tf:"related_note_names"`
	// URLs associated with this note and related metadata.
	// +optional
	RelatedURL []AnalysisNoteSpecRelatedURL `json:"relatedURL,omitempty" tf:"related_url"`
	// A one sentence description of the note.
	// +optional
	ShortDescription *string `json:"shortDescription,omitempty" tf:"short_description"`
	// The time this note was last updated.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time"`
}

type AnalysisNoteStatus struct {
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

// AnalysisNoteList is a list of AnalysisNotes
type AnalysisNoteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AnalysisNote CRD objects
	Items []AnalysisNote `json:"items,omitempty"`
}
