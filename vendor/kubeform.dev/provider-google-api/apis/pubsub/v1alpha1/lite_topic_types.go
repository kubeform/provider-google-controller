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

type LiteTopic struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LiteTopicSpec   `json:"spec,omitempty"`
	Status            LiteTopicStatus `json:"status,omitempty"`
}

type LiteTopicSpecPartitionConfigCapacity struct {
	// Subscribe throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
	PublishMIBPerSec *int64 `json:"publishMIBPerSec" tf:"publish_mib_per_sec"`
	// Publish throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
	SubscribeMIBPerSec *int64 `json:"subscribeMIBPerSec" tf:"subscribe_mib_per_sec"`
}

type LiteTopicSpecPartitionConfig struct {
	// The capacity configuration.
	// +optional
	Capacity *LiteTopicSpecPartitionConfigCapacity `json:"capacity,omitempty" tf:"capacity"`
	// The number of partitions in the topic. Must be at least 1.
	Count *int64 `json:"count" tf:"count"`
}

type LiteTopicSpecRetentionConfig struct {
	// The provisioned storage, in bytes, per partition. If the number of bytes stored
	// in any of the topic's partitions grows beyond this value, older messages will be
	// dropped to make room for newer ones, regardless of the value of period.
	PerPartitionBytes *string `json:"perPartitionBytes" tf:"per_partition_bytes"`
	// How long a published message is retained. If unset, messages will be retained as
	// long as the bytes retained for each partition is below perPartitionBytes.
	// +optional
	Period *string `json:"period,omitempty" tf:"period"`
}

type LiteTopicSpec struct {
	State *LiteTopicSpecResource `json:"state,omitempty" tf:"-"`

	Resource LiteTopicSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type LiteTopicSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the topic.
	Name *string `json:"name" tf:"name"`
	// The settings for this topic's partitions.
	// +optional
	PartitionConfig *LiteTopicSpecPartitionConfig `json:"partitionConfig,omitempty" tf:"partition_config"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The region of the pubsub lite topic.
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// The settings for a topic's message retention.
	// +optional
	RetentionConfig *LiteTopicSpecRetentionConfig `json:"retentionConfig,omitempty" tf:"retention_config"`
	// The zone of the pubsub lite topic.
	// +optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
}

type LiteTopicStatus struct {
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

// LiteTopicList is a list of LiteTopics
type LiteTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LiteTopic CRD objects
	Items []LiteTopic `json:"items,omitempty"`
}
