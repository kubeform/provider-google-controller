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

type CxAgent struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CxAgentSpec   `json:"spec,omitempty"`
	Status            CxAgentStatus `json:"status,omitempty"`
}

type CxAgentSpecSpeechToTextSettings struct {
	// Whether to use speech adaptation for speech recognition.
	// +optional
	EnableSpeechAdaptation *bool `json:"enableSpeechAdaptation,omitempty" tf:"enable_speech_adaptation"`
}

type CxAgentSpec struct {
	State *CxAgentSpecResource `json:"state,omitempty" tf:"-"`

	Resource CxAgentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type CxAgentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo integration.
	// +optional
	AvatarURI *string `json:"avatarURI,omitempty" tf:"avatar_uri"`
	// The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language)
	// for a list of the currently supported language codes. This field cannot be updated after creation.
	DefaultLanguageCode *string `json:"defaultLanguageCode" tf:"default_language_code"`
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The human-readable name of the agent, unique within the location.
	DisplayName *string `json:"displayName" tf:"display_name"`
	// Indicates if automatic spell correction is enabled in detect intent requests.
	// +optional
	EnableSpellCorrection *bool `json:"enableSpellCorrection,omitempty" tf:"enable_spell_correction"`
	// Determines whether this agent should log conversation queries.
	// +optional
	EnableStackdriverLogging *bool `json:"enableStackdriverLogging,omitempty" tf:"enable_stackdriver_logging"`
	// The name of the location this agent is located in.
	//
	// ~> **Note:** The first time you are deploying an Agent in your project you must configure location settings.
	//  This is a one time step but at the moment you can only [configure location settings](https://cloud.google.com/dialogflow/cx/docs/concept/region#location-settings) via the Dialogflow CX console.
	//  Another options is to use global location so you don't need to manually configure location settings.
	Location *string `json:"location" tf:"location"`
	// The unique identifier of the agent.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.
	// +optional
	SecuritySettings *string `json:"securitySettings,omitempty" tf:"security_settings"`
	// Settings related to speech recognition.
	// +optional
	SpeechToTextSettings *CxAgentSpecSpeechToTextSettings `json:"speechToTextSettings,omitempty" tf:"speech_to_text_settings"`
	// Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	// +optional
	StartFlow *string `json:"startFlow,omitempty" tf:"start_flow"`
	// The list of all languages supported by this agent (except for the default_language_code).
	// +optional
	SupportedLanguageCodes []string `json:"supportedLanguageCodes,omitempty" tf:"supported_language_codes"`
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	TimeZone *string `json:"timeZone" tf:"time_zone"`
}

type CxAgentStatus struct {
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

// CxAgentList is a list of CxAgents
type CxAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CxAgent CRD objects
	Items []CxAgent `json:"items,omitempty"`
}
