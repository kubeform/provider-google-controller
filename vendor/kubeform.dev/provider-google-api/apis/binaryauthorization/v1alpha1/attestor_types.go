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

type Attestor struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AttestorSpec   `json:"spec,omitempty"`
	Status            AttestorStatus `json:"status,omitempty"`
}

type AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey struct {
	// A PEM-encoded public key, as described in
	// 'https://tools.ietf.org/html/rfc7468#section-13'
	// +optional
	PublicKeyPem *string `json:"publicKeyPem,omitempty" tf:"public_key_pem"`
	// The signature algorithm used to verify a message against
	// a signature using this key. These signature algorithm must
	// match the structure and any object identifiers encoded in
	// publicKeyPem (i.e. this algorithm must match that of the
	// public key).
	// +optional
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty" tf:"signature_algorithm"`
}

type AttestorSpecAttestationAuthorityNotePublicKeys struct {
	// ASCII-armored representation of a PGP public key, as the
	// entire output by the command
	// 'gpg --export --armor foo@example.com' (either LF or CRLF
	// line endings). When using this field, id should be left
	// blank. The BinAuthz API handlers will calculate the ID
	// and fill it in automatically. BinAuthz computes this ID
	// as the OpenPGP RFC4880 V4 fingerprint, represented as
	// upper-case hex. If id is provided by the caller, it will
	// be overwritten by the API-calculated ID.
	// +optional
	AsciiArmoredPgpPublicKey *string `json:"asciiArmoredPgpPublicKey,omitempty" tf:"ascii_armored_pgp_public_key"`
	// A descriptive comment. This field may be updated.
	// +optional
	Comment *string `json:"comment,omitempty" tf:"comment"`
	// The ID of this public key. Signatures verified by BinAuthz
	// must include the ID of the public key that can be used to
	// verify them, and that ID must match the contents of this
	// field exactly. Additional restrictions on this field can
	// be imposed based on which public key type is encapsulated.
	// See the documentation on publicKey cases below for details.
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// A raw PKIX SubjectPublicKeyInfo format public key.
	//
	// NOTE: id may be explicitly provided by the caller when using this
	// type of public key, but it MUST be a valid RFC3986 URI. If id is left
	// blank, a default one will be computed based on the digest of the DER
	// encoding of the public key.
	// +optional
	PkixPublicKey *AttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey `json:"pkixPublicKey,omitempty" tf:"pkix_public_key"`
}

type AttestorSpecAttestationAuthorityNote struct {
	// This field will contain the service account email address that
	// this Attestor will use as the principal when querying Container
	// Analysis. Attestor administrators must grant this service account
	// the IAM role needed to read attestations from the noteReference in
	// Container Analysis (containeranalysis.notes.occurrences.viewer).
	// This email address is fixed for the lifetime of the Attestor, but
	// callers should not make any other assumptions about the service
	// account email; future versions may use an email based on a
	// different naming pattern.
	// +optional
	DelegationServiceAccountEmail *string `json:"delegationServiceAccountEmail,omitempty" tf:"delegation_service_account_email"`
	// The resource name of a ATTESTATION_AUTHORITY Note, created by the
	// user. If the Note is in a different project from the Attestor, it
	// should be specified in the format 'projects/*/notes/*' (or the legacy
	// 'providers/*/notes/*'). This field may not be updated.
	// An attestation by this attestor is stored as a Container Analysis
	// ATTESTATION_AUTHORITY Occurrence that names a container image
	// and that links to this Note.
	NoteReference *string `json:"noteReference" tf:"note_reference"`
	// Public keys that verify attestations signed by this attestor. This
	// field may be updated.
	// If this field is non-empty, one of the specified public keys must
	// verify that an attestation was signed by this attestor for the
	// image specified in the admission request.
	// If this field is empty, this attestor always returns that no valid
	// attestations exist.
	// +optional
	PublicKeys []AttestorSpecAttestationAuthorityNotePublicKeys `json:"publicKeys,omitempty" tf:"public_keys"`
}

type AttestorSpec struct {
	State *AttestorSpecResource `json:"state,omitempty" tf:"-"`

	Resource AttestorSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AttestorSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// A Container Analysis ATTESTATION_AUTHORITY Note, created by the user.
	AttestationAuthorityNote *AttestorSpecAttestationAuthorityNote `json:"attestationAuthorityNote" tf:"attestation_authority_note"`
	// A descriptive comment. This field may be updated. The field may be
	// displayed in chooser dialogs.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The resource name.
	Name *string `json:"name" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
}

type AttestorStatus struct {
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

// AttestorList is a list of Attestors
type AttestorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Attestor CRD objects
	Items []Attestor `json:"items,omitempty"`
}
