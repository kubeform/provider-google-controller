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

type Topic struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TopicSpec   `json:"spec,omitempty"`
	Status            TopicStatus `json:"status,omitempty"`
}

type TopicSpecMessageStoragePolicy struct {
	// A list of IDs of GCP regions where messages that are published to
	// the topic may be persisted in storage. Messages published by
	// publishers running in non-allowed GCP regions (or running outside
	// of GCP altogether) will be routed for storage in one of the
	// allowed regions. An empty list means that no regions are allowed,
	// and is not a valid configuration.
	AllowedPersistenceRegions []string `json:"allowedPersistenceRegions" tf:"allowed_persistence_regions"`
}

type TopicSpecSchemaSettings struct {
	// The encoding of messages validated against schema. Default value: "ENCODING_UNSPECIFIED" Possible values: ["ENCODING_UNSPECIFIED", "JSON", "BINARY"]
	// +optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding"`
	// The name of the schema that messages published should be
	// validated against. Format is projects/{project}/schemas/{schema}.
	// The value of this field will be _deleted-schema_
	// if the schema has been deleted.
	Schema *string `json:"schema" tf:"schema"`
}

type TopicSpec struct {
	State *TopicSpecResource `json:"state,omitempty" tf:"-"`

	Resource TopicSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type TopicSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource name of the Cloud KMS CryptoKey to be used to protect access
	// to messages published on this topic. Your project's PubSub service account
	// ('service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com') must have
	// 'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this feature.
	// The expected format is 'projects/*/locations/*/keyRings/*/cryptoKeys/*'
	// +optional
	KmsKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name"`
	// A set of key/value label pairs to assign to this Topic.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// Indicates the minimum duration to retain a message after it is published
	// to the topic. If this field is set, messages published to the topic in
	// the last messageRetentionDuration are always available to subscribers.
	// For instance, it allows any attached subscription to seek to a timestamp
	// that is up to messageRetentionDuration in the past. If this field is not
	// set, message retention is controlled by settings on individual subscriptions.
	// Cannot be more than 7 days or less than 10 minutes.
	// +optional
	MessageRetentionDuration *string `json:"messageRetentionDuration,omitempty" tf:"message_retention_duration"`
	// Policy constraining the set of Google Cloud Platform regions where
	// messages published to the topic may be stored. If not present, then no
	// constraints are in effect.
	// +optional
	MessageStoragePolicy *TopicSpecMessageStoragePolicy `json:"messageStoragePolicy,omitempty" tf:"message_storage_policy"`
	// Name of the topic.
	Name *string `json:"name" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Settings for validating messages published against a schema.
	// +optional
	SchemaSettings *TopicSpecSchemaSettings `json:"schemaSettings,omitempty" tf:"schema_settings"`
}

type TopicStatus struct {
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

// TopicList is a list of Topics
type TopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Topic CRD objects
	Items []Topic `json:"items,omitempty"`
}
