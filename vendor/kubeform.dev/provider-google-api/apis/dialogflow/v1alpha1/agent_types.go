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

type Agent struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AgentSpec   `json:"spec,omitempty"`
	Status            AgentStatus `json:"status,omitempty"`
}

type AgentSpec struct {
	KubeformOutput *AgentSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource AgentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type AgentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// API version displayed in Dialogflow console. If not specified, V2 API is assumed. Clients are free to query
	// different service endpoints for different API versions. However, bots connectors and webhook calls will follow
	// the specified API version.
	// * API_VERSION_V1: Legacy V1 API.
	// * API_VERSION_V2: V2 API.
	// * API_VERSION_V2_BETA_1: V2beta1 API. Possible values: ["API_VERSION_V1", "API_VERSION_V2", "API_VERSION_V2_BETA_1"]
	// +optional
	ApiVersion *string `json:"apiVersion,omitempty" tf:"api_version"`
	// The URI of the agent's avatar, which are used throughout the Dialogflow console. When an image URL is entered
	// into this field, the Dialogflow will save the image in the backend. The address of the backend image returned
	// from the API will be shown in the [avatarUriBackend] field.
	// +optional
	AvatarURI *string `json:"avatarURI,omitempty" tf:"avatar_uri"`
	// The URI of the agent's avatar as returned from the API. Output only. To provide an image URL for the agent avatar,
	// the [avatarUri] field can be used.
	// +optional
	AvatarURIBackend *string `json:"avatarURIBackend,omitempty" tf:"avatar_uri_backend"`
	// To filter out false positive results and still get variety in matched natural language inputs for your agent,
	// you can tune the machine learning classification threshold. If the returned score value is less than the threshold
	// value, then a fallback intent will be triggered or, if there are no fallback intents defined, no intent will be
	// triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the
	// default of 0.3 is used.
	// +optional
	ClassificationThreshold *float64 `json:"classificationThreshold,omitempty" tf:"classification_threshold"`
	// The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/docs/reference/language)
	// for a list of the currently supported language codes. This field cannot be updated after creation.
	DefaultLanguageCode *string `json:"defaultLanguageCode" tf:"default_language_code"`
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The name of this agent.
	DisplayName *string `json:"displayName" tf:"display_name"`
	// Determines whether this agent should log conversation queries.
	// +optional
	EnableLogging *bool `json:"enableLogging,omitempty" tf:"enable_logging"`
	// Determines how intents are detected from user queries.
	// * MATCH_MODE_HYBRID: Best for agents with a small number of examples in intents and/or wide use of templates
	// syntax and composite entities.
	// * MATCH_MODE_ML_ONLY: Can be used for agents with a large number of examples in intents, especially the ones
	// using @sys.any or very large developer entities. Possible values: ["MATCH_MODE_HYBRID", "MATCH_MODE_ML_ONLY"]
	// +optional
	MatchMode *string `json:"matchMode,omitempty" tf:"match_mode"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The list of all languages supported by this agent (except for the defaultLanguageCode).
	// +optional
	SupportedLanguageCodes []string `json:"supportedLanguageCodes,omitempty" tf:"supported_language_codes"`
	// The agent tier. If not specified, TIER_STANDARD is assumed.
	// * TIER_STANDARD: Standard tier.
	// * TIER_ENTERPRISE: Enterprise tier (Essentials).
	// * TIER_ENTERPRISE_PLUS: Enterprise tier (Plus).
	// NOTE: Due to consistency issues, the provider will not read this field from the API. Drift is possible between
	// the Terraform state and Dialogflow if the agent tier is changed outside of Terraform. Possible values: ["TIER_STANDARD", "TIER_ENTERPRISE", "TIER_ENTERPRISE_PLUS"]
	// +optional
	Tier *string `json:"tier,omitempty" tf:"tier"`
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	TimeZone *string `json:"timeZone" tf:"time_zone"`
}

type AgentStatus struct {
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

// AgentList is a list of Agents
type AgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Agent CRD objects
	Items []Agent `json:"items,omitempty"`
}
