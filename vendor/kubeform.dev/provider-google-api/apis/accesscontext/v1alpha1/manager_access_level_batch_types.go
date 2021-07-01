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

type ManagerAccessLevelBatch struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagerAccessLevelBatchSpec   `json:"spec,omitempty"`
	Status            ManagerAccessLevelBatchStatus `json:"status,omitempty"`
}

type ManagerAccessLevelBatchSpecAccessLevelsBasicConditionsDevicePolicyOsConstraints struct {
	// The minimum allowed OS version. If not set, any version
	// of this OS satisfies the constraint.
	// Format: "major.minor.patch" such as "10.5.301", "9.2.1".
	// +optional
	MinimumVersion *string `json:"minimumVersion,omitempty" tf:"minimum_version"`
	// The operating system type of the device. Possible values: ["OS_UNSPECIFIED", "DESKTOP_MAC", "DESKTOP_WINDOWS", "DESKTOP_LINUX", "DESKTOP_CHROME_OS", "ANDROID", "IOS"]
	OsType *string `json:"osType" tf:"os_type"`
}

type ManagerAccessLevelBatchSpecAccessLevelsBasicConditionsDevicePolicy struct {
	// A list of allowed device management levels.
	// An empty list allows all management levels. Possible values: ["MANAGEMENT_UNSPECIFIED", "NONE", "BASIC", "COMPLETE"]
	// +optional
	AllowedDeviceManagementLevels []string `json:"allowedDeviceManagementLevels,omitempty" tf:"allowed_device_management_levels"`
	// A list of allowed encryptions statuses.
	// An empty list allows all statuses. Possible values: ["ENCRYPTION_UNSPECIFIED", "ENCRYPTION_UNSUPPORTED", "UNENCRYPTED", "ENCRYPTED"]
	// +optional
	AllowedEncryptionStatuses []string `json:"allowedEncryptionStatuses,omitempty" tf:"allowed_encryption_statuses"`
	// A list of allowed OS versions.
	// An empty list allows all types and all versions.
	// +optional
	OsConstraints []ManagerAccessLevelBatchSpecAccessLevelsBasicConditionsDevicePolicyOsConstraints `json:"osConstraints,omitempty" tf:"os_constraints"`
	// Whether the device needs to be approved by the customer admin.
	// +optional
	RequireAdminApproval *bool `json:"requireAdminApproval,omitempty" tf:"require_admin_approval"`
	// Whether the device needs to be corp owned.
	// +optional
	RequireCorpOwned *bool `json:"requireCorpOwned,omitempty" tf:"require_corp_owned"`
	// Whether or not screenlock is required for the DevicePolicy
	// to be true. Defaults to false.
	// +optional
	RequireScreenLock *bool `json:"requireScreenLock,omitempty" tf:"require_screen_lock"`
}

type ManagerAccessLevelBatchSpecAccessLevelsBasicConditions struct {
	// Device specific restrictions, all restrictions must hold for
	// the Condition to be true. If not specified, all devices are
	// allowed.
	// +optional
	DevicePolicy *ManagerAccessLevelBatchSpecAccessLevelsBasicConditionsDevicePolicy `json:"devicePolicy,omitempty" tf:"device_policy"`
	// A list of CIDR block IP subnetwork specification. May be IPv4
	// or IPv6.
	// Note that for a CIDR IP address block, the specified IP address
	// portion must be properly truncated (i.e. all the host bits must
	// be zero) or the input is considered malformed. For example,
	// "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
	// for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
	// is not. The originating IP of a request must be in one of the
	// listed subnets in order for this Condition to be true.
	// If empty, all IP addresses are allowed.
	// +optional
	IpSubnetworks []string `json:"ipSubnetworks,omitempty" tf:"ip_subnetworks"`
	// An allowed list of members (users, service accounts).
	// Using groups is not supported yet.
	//
	// The signed-in user originating the request must be a part of one
	// of the provided members. If not specified, a request may come
	// from any user (logged in/not logged in, not present in any
	// groups, etc.).
	// Formats: 'user:{emailid}', 'serviceAccount:{emailid}'
	// +optional
	Members []string `json:"members,omitempty" tf:"members"`
	// Whether to negate the Condition. If true, the Condition becomes
	// a NAND over its non-empty fields, each field must be false for
	// the Condition overall to be satisfied. Defaults to false.
	// +optional
	Negate *bool `json:"negate,omitempty" tf:"negate"`
	// The request must originate from one of the provided
	// countries/regions.
	// Format: A valid ISO 3166-1 alpha-2 code.
	// +optional
	Regions []string `json:"regions,omitempty" tf:"regions"`
	// A list of other access levels defined in the same Policy,
	// referenced by resource name. Referencing an AccessLevel which
	// does not exist is an error. All access levels listed must be
	// granted for the Condition to be true.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	// +optional
	RequiredAccessLevels []string `json:"requiredAccessLevels,omitempty" tf:"required_access_levels"`
}

type ManagerAccessLevelBatchSpecAccessLevelsBasic struct {
	// How the conditions list should be combined to determine if a request
	// is granted this AccessLevel. If AND is used, each Condition in
	// conditions must be satisfied for the AccessLevel to be applied. If
	// OR is used, at least one Condition in conditions must be satisfied
	// for the AccessLevel to be applied. Default value: "AND" Possible values: ["AND", "OR"]
	// +optional
	CombiningFunction *string `json:"combiningFunction,omitempty" tf:"combining_function"`
	// A set of requirements for the AccessLevel to be granted.
	// +kubebuilder:validation:MinItems=1
	Conditions []ManagerAccessLevelBatchSpecAccessLevelsBasicConditions `json:"conditions" tf:"conditions"`
}

type ManagerAccessLevelBatchSpecAccessLevelsCustomExpr struct {
	// Description of the expression
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Textual representation of an expression in Common Expression Language syntax.
	Expression *string `json:"expression" tf:"expression"`
	// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file
	// +optional
	Location *string `json:"location,omitempty" tf:"location"`
	// Title for the expression, i.e. a short string describing its purpose.
	// +optional
	Title *string `json:"title,omitempty" tf:"title"`
}

type ManagerAccessLevelBatchSpecAccessLevelsCustom struct {
	// Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language.
	// This page details the objects and attributes that are used to the build the CEL expressions for
	// custom access levels - https://cloud.google.com/access-context-manager/docs/custom-access-level-spec.
	Expr *ManagerAccessLevelBatchSpecAccessLevelsCustomExpr `json:"expr" tf:"expr"`
}

type ManagerAccessLevelBatchSpecAccessLevels struct {
	// A set of predefined conditions for the access level and a combining function.
	// +optional
	Basic *ManagerAccessLevelBatchSpecAccessLevelsBasic `json:"basic,omitempty" tf:"basic"`
	// Custom access level conditions are set using the Cloud Common Expression Language to represent the necessary conditions for the level to apply to a request.
	// See CEL spec at: https://github.com/google/cel-spec.
	// +optional
	Custom *ManagerAccessLevelBatchSpecAccessLevelsCustom `json:"custom,omitempty" tf:"custom"`
	// Description of the AccessLevel and its use. Does not affect behavior.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Resource name for the Access Level. The short_name component must begin
	// with a letter and only include alphanumeric and '_'.
	// Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	Name *string `json:"name" tf:"name"`
	// Human readable title. Must be unique within the Policy.
	Title *string `json:"title" tf:"title"`
}

type ManagerAccessLevelBatchSpec struct {
	KubeformOutput *ManagerAccessLevelBatchSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource ManagerAccessLevelBatchSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ManagerAccessLevelBatchSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The desired Access Levels that should replace all existing Access Levels in the Access Policy.
	// +optional
	AccessLevels []ManagerAccessLevelBatchSpecAccessLevels `json:"accessLevels,omitempty" tf:"access_levels"`
	// The AccessPolicy this AccessLevel lives in.
	// Format: accessPolicies/{policy_id}
	Parent *string `json:"parent" tf:"parent"`
}

type ManagerAccessLevelBatchStatus struct {
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

// ManagerAccessLevelBatchList is a list of ManagerAccessLevelBatchs
type ManagerAccessLevelBatchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagerAccessLevelBatch CRD objects
	Items []ManagerAccessLevelBatch `json:"items,omitempty"`
}
