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

type Policy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicySpec   `json:"spec,omitempty"`
	Status            PolicyStatus `json:"status,omitempty"`
}

type PolicySpecBooleanPolicy struct {
	// If true, then the Policy is enforced. If false, then any configuration is acceptable.
	Enforced *bool `json:"enforced" tf:"enforced"`
}

type PolicySpecListPolicyAllow struct {
	// The policy allows or denies all values.
	// +optional
	All *bool `json:"all,omitempty" tf:"all"`
	// The policy can define specific values that are allowed or denied.
	// +optional
	Values []string `json:"values,omitempty" tf:"values"`
}

type PolicySpecListPolicyDeny struct {
	// The policy allows or denies all values.
	// +optional
	All *bool `json:"all,omitempty" tf:"all"`
	// The policy can define specific values that are allowed or denied.
	// +optional
	Values []string `json:"values,omitempty" tf:"values"`
}

type PolicySpecListPolicy struct {
	// One or the other must be set.
	// +optional
	Allow *PolicySpecListPolicyAllow `json:"allow,omitempty" tf:"allow"`
	// One or the other must be set.
	// +optional
	Deny *PolicySpecListPolicyDeny `json:"deny,omitempty" tf:"deny"`
	// If set to true, the values from the effective Policy of the parent resource are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy.
	// +optional
	InheritFromParent *bool `json:"inheritFromParent,omitempty" tf:"inherit_from_parent"`
	// The Google Cloud Console will try to default to a configuration that matches the value specified in this field.
	// +optional
	SuggestedValue *string `json:"suggestedValue,omitempty" tf:"suggested_value"`
}

type PolicySpecRestorePolicy struct {
	// May only be set to true. If set, then the default Policy is restored.
	Default *bool `json:"default" tf:"default"`
}

type PolicySpec struct {
	State *PolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource PolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type PolicySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// A boolean policy is a constraint that is either enforced or not.
	// +optional
	BooleanPolicy *PolicySpecBooleanPolicy `json:"booleanPolicy,omitempty" tf:"boolean_policy"`
	// The name of the Constraint the Policy is configuring, for example, serviceuser.services.
	Constraint *string `json:"constraint" tf:"constraint"`
	// The etag of the organization policy. etag is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.
	// +optional
	Etag *string `json:"etag,omitempty" tf:"etag"`
	// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values.
	// +optional
	ListPolicy *PolicySpecListPolicy `json:"listPolicy,omitempty" tf:"list_policy"`
	OrgID      *string               `json:"orgID" tf:"org_id"`
	// A restore policy is a constraint to restore the default policy.
	// +optional
	RestorePolicy *PolicySpecRestorePolicy `json:"restorePolicy,omitempty" tf:"restore_policy"`
	// The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z".
	// +optional
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time"`
	// Version of the Policy. Default version is 0.
	// +optional
	Version *int64 `json:"version,omitempty" tf:"version"`
}

type PolicyStatus struct {
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

// PolicyList is a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Policy CRD objects
	Items []Policy `json:"items,omitempty"`
}
