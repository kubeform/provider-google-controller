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

type LossPreventionInspectTemplate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LossPreventionInspectTemplateSpec   `json:"spec,omitempty"`
	Status            LossPreventionInspectTemplateStatus `json:"status,omitempty"`
}

type LossPreventionInspectTemplateSpecInspectConfigCustomInfoTypesDictionaryCloudStoragePath struct {
	// A url representing a file or path (no wildcards) in Cloud Storage. Example: 'gs://[BUCKET_NAME]/dictionary.txt'
	Path *string `json:"path" tf:"path"`
}

type LossPreventionInspectTemplateSpecInspectConfigCustomInfoTypesDictionaryWordList struct {
	// Words or phrases defining the dictionary. The dictionary must contain at least one
	// phrase and every phrase must contain at least 2 characters that are letters or digits.
	Words []string `json:"words" tf:"words"`
}

type LossPreventionInspectTemplateSpecInspectConfigCustomInfoTypesDictionary struct {
	// Newline-delimited file of words in Cloud Storage. Only a single file is accepted.
	// +optional
	CloudStoragePath *LossPreventionInspectTemplateSpecInspectConfigCustomInfoTypesDictionaryCloudStoragePath `json:"cloudStoragePath,omitempty" tf:"cloud_storage_path"`
	// List of words or phrases to search for.
	// +optional
	WordList *LossPreventionInspectTemplateSpecInspectConfigCustomInfoTypesDictionaryWordList `json:"wordList,omitempty" tf:"word_list"`
}

type LossPreventionInspectTemplateSpecInspectConfigCustomInfoTypesInfoType struct {
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names
	// listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	Name *string `json:"name" tf:"name"`
}

type LossPreventionInspectTemplateSpecInspectConfigCustomInfoTypesRegex struct {
	// The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included.
	// +optional
	GroupIndexes []int64 `json:"groupIndexes,omitempty" tf:"group_indexes"`
	// Pattern defining the regular expression.
	// Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.
	Pattern *string `json:"pattern" tf:"pattern"`
}

type LossPreventionInspectTemplateSpecInspectConfigCustomInfoTypesStoredType struct {
	// Resource name of the requested StoredInfoType, for example 'organizations/433245324/storedInfoTypes/432452342'
	// or 'projects/project-id/storedInfoTypes/432452342'.
	Name *string `json:"name" tf:"name"`
}

type LossPreventionInspectTemplateSpecInspectConfigCustomInfoTypes struct {
	// Dictionary which defines the rule.
	// +optional
	Dictionary *LossPreventionInspectTemplateSpecInspectConfigCustomInfoTypesDictionary `json:"dictionary,omitempty" tf:"dictionary"`
	// If set to EXCLUSION_TYPE_EXCLUDE this infoType will not cause a finding to be returned. It still can be used for rules matching. Possible values: ["EXCLUSION_TYPE_EXCLUDE"]
	// +optional
	ExclusionType *string `json:"exclusionType,omitempty" tf:"exclusion_type"`
	// CustomInfoType can either be a new infoType, or an extension of built-in infoType, when the name matches one of existing
	// infoTypes and that infoType is specified in 'info_types' field. Specifying the latter adds findings to the
	// one detected by the system. If built-in info type is not specified in 'info_types' list then the name is
	// treated as a custom info type.
	InfoType *LossPreventionInspectTemplateSpecInspectConfigCustomInfoTypesInfoType `json:"infoType" tf:"info_type"`
	// Likelihood to return for this CustomInfoType. This base value can be altered by a detection rule if the finding meets the criteria
	// specified by the rule. Default value: "VERY_LIKELY" Possible values: ["VERY_UNLIKELY", "UNLIKELY", "POSSIBLE", "LIKELY", "VERY_LIKELY"]
	// +optional
	Likelihood *string `json:"likelihood,omitempty" tf:"likelihood"`
	// Regular expression which defines the rule.
	// +optional
	Regex *LossPreventionInspectTemplateSpecInspectConfigCustomInfoTypesRegex `json:"regex,omitempty" tf:"regex"`
	// A reference to a StoredInfoType to use with scanning.
	// +optional
	StoredType *LossPreventionInspectTemplateSpecInspectConfigCustomInfoTypesStoredType `json:"storedType,omitempty" tf:"stored_type"`
}

type LossPreventionInspectTemplateSpecInspectConfigInfoTypes struct {
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	Name *string `json:"name" tf:"name"`
}

type LossPreventionInspectTemplateSpecInspectConfigLimitsMaxFindingsPerInfoTypeInfoType struct {
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	Name *string `json:"name" tf:"name"`
}

type LossPreventionInspectTemplateSpecInspectConfigLimitsMaxFindingsPerInfoType struct {
	// Type of information the findings limit applies to. Only one limit per infoType should be provided. If InfoTypeLimit does
	// not have an infoType, the DLP API applies the limit against all infoTypes that are found but not
	// specified in another InfoTypeLimit.
	InfoType *LossPreventionInspectTemplateSpecInspectConfigLimitsMaxFindingsPerInfoTypeInfoType `json:"infoType" tf:"info_type"`
	// Max findings limit for the given infoType.
	MaxFindings *int64 `json:"maxFindings" tf:"max_findings"`
}

type LossPreventionInspectTemplateSpecInspectConfigLimits struct {
	// Configuration of findings limit given for specified infoTypes.
	// +optional
	MaxFindingsPerInfoType []LossPreventionInspectTemplateSpecInspectConfigLimitsMaxFindingsPerInfoType `json:"maxFindingsPerInfoType,omitempty" tf:"max_findings_per_info_type"`
	// Max number of findings that will be returned for each item scanned. The maximum returned is 2000.
	MaxFindingsPerItem *int64 `json:"maxFindingsPerItem" tf:"max_findings_per_item"`
	// Max number of findings that will be returned per request/job. The maximum returned is 2000.
	MaxFindingsPerRequest *int64 `json:"maxFindingsPerRequest" tf:"max_findings_per_request"`
}

type LossPreventionInspectTemplateSpecInspectConfigRuleSetInfoTypes struct {
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	Name *string `json:"name" tf:"name"`
}

type LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath struct {
	// A url representing a file or path (no wildcards) in Cloud Storage. Example: 'gs://[BUCKET_NAME]/dictionary.txt'
	Path *string `json:"path" tf:"path"`
}

type LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesExclusionRuleDictionaryWordList struct {
	// Words or phrases defining the dictionary. The dictionary must contain at least one
	// phrase and every phrase must contain at least 2 characters that are letters or digits.
	Words []string `json:"words" tf:"words"`
}

type LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesExclusionRuleDictionary struct {
	// Newline-delimited file of words in Cloud Storage. Only a single file is accepted.
	// +optional
	CloudStoragePath *LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath `json:"cloudStoragePath,omitempty" tf:"cloud_storage_path"`
	// List of words or phrases to search for.
	// +optional
	WordList *LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesExclusionRuleDictionaryWordList `json:"wordList,omitempty" tf:"word_list"`
}

type LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes struct {
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	Name *string `json:"name" tf:"name"`
}

type LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes struct {
	// If a finding is matched by any of the infoType detectors listed here, the finding will be excluded from the scan results.
	InfoTypes []LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes `json:"infoTypes" tf:"info_types"`
}

type LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesExclusionRuleRegex struct {
	// The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included.
	// +optional
	GroupIndexes []int64 `json:"groupIndexes,omitempty" tf:"group_indexes"`
	// Pattern defining the regular expression.
	// Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.
	Pattern *string `json:"pattern" tf:"pattern"`
}

type LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesExclusionRule struct {
	// Dictionary which defines the rule.
	// +optional
	Dictionary *LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesExclusionRuleDictionary `json:"dictionary,omitempty" tf:"dictionary"`
	// Set of infoTypes for which findings would affect this rule.
	// +optional
	ExcludeInfoTypes *LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes `json:"excludeInfoTypes,omitempty" tf:"exclude_info_types"`
	// How the rule is applied. See the documentation for more information: https://cloud.google.com/dlp/docs/reference/rest/v2/InspectConfig#MatchingType Possible values: ["MATCHING_TYPE_FULL_MATCH", "MATCHING_TYPE_PARTIAL_MATCH", "MATCHING_TYPE_INVERSE_MATCH"]
	MatchingType *string `json:"matchingType" tf:"matching_type"`
	// Regular expression which defines the rule.
	// +optional
	Regex *LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesExclusionRuleRegex `json:"regex,omitempty" tf:"regex"`
}

type LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesHotwordRuleHotwordRegex struct {
	// The index of the submatch to extract as findings. When not specified,
	// the entire match is returned. No more than 3 may be included.
	// +optional
	GroupIndexes []int64 `json:"groupIndexes,omitempty" tf:"group_indexes"`
	// Pattern defining the regular expression. Its syntax
	// (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.
	Pattern *string `json:"pattern" tf:"pattern"`
}

type LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment struct {
	// Set the likelihood of a finding to a fixed value. Either this or relative_likelihood can be set. Possible values: ["VERY_UNLIKELY", "UNLIKELY", "POSSIBLE", "LIKELY", "VERY_LIKELY"]
	// +optional
	FixedLikelihood *string `json:"fixedLikelihood,omitempty" tf:"fixed_likelihood"`
	// Increase or decrease the likelihood by the specified number of levels. For example,
	// if a finding would be POSSIBLE without the detection rule and relativeLikelihood is 1,
	// then it is upgraded to LIKELY, while a value of -1 would downgrade it to UNLIKELY.
	// Likelihood may never drop below VERY_UNLIKELY or exceed VERY_LIKELY, so applying an
	// adjustment of 1 followed by an adjustment of -1 when base likelihood is VERY_LIKELY
	// will result in a final likelihood of LIKELY. Either this or fixed_likelihood can be set.
	// +optional
	RelativeLikelihood *int64 `json:"relativeLikelihood,omitempty" tf:"relative_likelihood"`
}

type LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesHotwordRuleProximity struct {
	// Number of characters after the finding to consider. Either this or window_before must be specified
	// +optional
	WindowAfter *int64 `json:"windowAfter,omitempty" tf:"window_after"`
	// Number of characters before the finding to consider. Either this or window_after must be specified
	// +optional
	WindowBefore *int64 `json:"windowBefore,omitempty" tf:"window_before"`
}

type LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesHotwordRule struct {
	// Regular expression pattern defining what qualifies as a hotword.
	HotwordRegex *LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesHotwordRuleHotwordRegex `json:"hotwordRegex" tf:"hotword_regex"`
	// Likelihood adjustment to apply to all matching findings.
	LikelihoodAdjustment *LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment `json:"likelihoodAdjustment" tf:"likelihood_adjustment"`
	// Proximity of the finding within which the entire hotword must reside. The total length of the window cannot
	// exceed 1000 characters. Note that the finding itself will be included in the window, so that hotwords may be
	// used to match substrings of the finding itself. For example, the certainty of a phone number regex
	// '(\\d{3}) \\d{3}-\\d{4}' could be adjusted upwards if the area code is known to be the local area code of a company
	// office using the hotword regex '(xxx)', where 'xxx' is the area code in question.
	Proximity *LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesHotwordRuleProximity `json:"proximity" tf:"proximity"`
}

type LossPreventionInspectTemplateSpecInspectConfigRuleSetRules struct {
	// The rule that specifies conditions when findings of infoTypes specified in InspectionRuleSet are removed from results.
	// +optional
	ExclusionRule *LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesExclusionRule `json:"exclusionRule,omitempty" tf:"exclusion_rule"`
	// Hotword-based detection rule.
	// +optional
	HotwordRule *LossPreventionInspectTemplateSpecInspectConfigRuleSetRulesHotwordRule `json:"hotwordRule,omitempty" tf:"hotword_rule"`
}

type LossPreventionInspectTemplateSpecInspectConfigRuleSet struct {
	// List of infoTypes this rule set is applied to.
	InfoTypes []LossPreventionInspectTemplateSpecInspectConfigRuleSetInfoTypes `json:"infoTypes" tf:"info_types"`
	// Set of rules to be applied to infoTypes. The rules are applied in order.
	Rules []LossPreventionInspectTemplateSpecInspectConfigRuleSetRules `json:"rules" tf:"rules"`
}

type LossPreventionInspectTemplateSpecInspectConfig struct {
	// List of options defining data content to scan. If empty, text, images, and other content will be included. Possible values: ["CONTENT_TEXT", "CONTENT_IMAGE"]
	// +optional
	ContentOptions []string `json:"contentOptions,omitempty" tf:"content_options"`
	// Custom info types to be used. See https://cloud.google.com/dlp/docs/creating-custom-infotypes to learn more.
	// +optional
	CustomInfoTypes []LossPreventionInspectTemplateSpecInspectConfigCustomInfoTypes `json:"customInfoTypes,omitempty" tf:"custom_info_types"`
	// When true, excludes type information of the findings.
	// +optional
	ExcludeInfoTypes *bool `json:"excludeInfoTypes,omitempty" tf:"exclude_info_types"`
	// When true, a contextual quote from the data that triggered a finding is included in the response.
	// +optional
	IncludeQuote *bool `json:"includeQuote,omitempty" tf:"include_quote"`
	// Restricts what infoTypes to look for. The values must correspond to InfoType values returned by infoTypes.list
	// or listed at https://cloud.google.com/dlp/docs/infotypes-reference.
	//
	// When no InfoTypes or CustomInfoTypes are specified in a request, the system may automatically choose what detectors to run.
	// By default this may be all types, but may change over time as detectors are updated.
	// +optional
	InfoTypes []LossPreventionInspectTemplateSpecInspectConfigInfoTypes `json:"infoTypes,omitempty" tf:"info_types"`
	// Configuration to control the number of findings returned.
	// +optional
	Limits *LossPreventionInspectTemplateSpecInspectConfigLimits `json:"limits,omitempty" tf:"limits"`
	// Only returns findings equal or above this threshold. See https://cloud.google.com/dlp/docs/likelihood for more info Default value: "POSSIBLE" Possible values: ["VERY_UNLIKELY", "UNLIKELY", "POSSIBLE", "LIKELY", "VERY_LIKELY"]
	// +optional
	MinLikelihood *string `json:"minLikelihood,omitempty" tf:"min_likelihood"`
	// Set of rules to apply to the findings for this InspectConfig. Exclusion rules, contained in the set are executed in the end,
	// other rules are executed in the order they are specified for each info type.
	// +optional
	RuleSet []LossPreventionInspectTemplateSpecInspectConfigRuleSet `json:"ruleSet,omitempty" tf:"rule_set"`
}

type LossPreventionInspectTemplateSpec struct {
	State *LossPreventionInspectTemplateSpecResource `json:"state,omitempty" tf:"-"`

	Resource LossPreventionInspectTemplateSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type LossPreventionInspectTemplateSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// A description of the inspect template.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// User set display name of the inspect template.
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// The core content of the template.
	// +optional
	InspectConfig *LossPreventionInspectTemplateSpecInspectConfig `json:"inspectConfig,omitempty" tf:"inspect_config"`
	// The resource name of the inspect template. Set by the server.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The parent of the inspect template in any of the following formats:
	//
	// * 'projects/{{project}}'
	// * 'projects/{{project}}/locations/{{location}}'
	// * 'organizations/{{organization_id}}'
	// * 'organizations/{{organization_id}}/locations/{{location}}'
	Parent *string `json:"parent" tf:"parent"`
}

type LossPreventionInspectTemplateStatus struct {
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

// LossPreventionInspectTemplateList is a list of LossPreventionInspectTemplates
type LossPreventionInspectTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LossPreventionInspectTemplate CRD objects
	Items []LossPreventionInspectTemplate `json:"items,omitempty"`
}
