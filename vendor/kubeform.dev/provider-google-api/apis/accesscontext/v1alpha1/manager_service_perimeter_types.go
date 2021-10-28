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

type ManagerServicePerimeter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagerServicePerimeterSpec   `json:"spec,omitempty"`
	Status            ManagerServicePerimeterStatus `json:"status,omitempty"`
}

type ManagerServicePerimeterSpecSpecEgressPoliciesEgressFrom struct {
	// A list of identities that are allowed access through this 'EgressPolicy'.
	// Should be in the format of email address. The email address should
	// represent individual user or service account only.
	// +optional
	Identities []string `json:"identities,omitempty" tf:"identities"`
	// Specifies the type of identities that are allowed access to outside the
	// perimeter. If left unspecified, then members of 'identities' field will
	// be allowed access. Possible values: ["IDENTITY_TYPE_UNSPECIFIED", "ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT"]
	// +optional
	IdentityType *string `json:"identityType,omitempty" tf:"identity_type"`
}

type ManagerServicePerimeterSpecSpecEgressPoliciesEgressToOperationsMethodSelectors struct {
	// Value for 'method' should be a valid method name for the corresponding
	// 'serviceName' in 'ApiOperation'. If '*' used as value for method,
	// then ALL methods and permissions are allowed.
	// +optional
	Method *string `json:"method,omitempty" tf:"method"`
	// Value for permission should be a valid Cloud IAM permission for the
	// corresponding 'serviceName' in 'ApiOperation'.
	// +optional
	Permission *string `json:"permission,omitempty" tf:"permission"`
}

type ManagerServicePerimeterSpecSpecEgressPoliciesEgressToOperations struct {
	// API methods or permissions to allow. Method or permission must belong
	// to the service specified by 'serviceName' field. A single MethodSelector
	// entry with '*' specified for the 'method' field will allow all methods
	// AND permissions for the service specified in 'serviceName'.
	// +optional
	MethodSelectors []ManagerServicePerimeterSpecSpecEgressPoliciesEgressToOperationsMethodSelectors `json:"methodSelectors,omitempty" tf:"method_selectors"`
	// The name of the API whose methods or permissions the 'IngressPolicy' or
	// 'EgressPolicy' want to allow. A single 'ApiOperation' with serviceName
	// field set to '*' will allow all methods AND permissions for all services.
	// +optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name"`
}

type ManagerServicePerimeterSpecSpecEgressPoliciesEgressTo struct {
	// A list of 'ApiOperations' that this egress rule applies to. A request matches
	// if it contains an operation/service in this list.
	// +optional
	Operations []ManagerServicePerimeterSpecSpecEgressPoliciesEgressToOperations `json:"operations,omitempty" tf:"operations"`
	// A list of resources, currently only projects in the form
	// 'projects/<projectnumber>', that match this to stanza. A request matches
	// if it contains a resource in this list. If * is specified for resources,
	// then this 'EgressTo' rule will authorize access to all resources outside
	// the perimeter.
	// +optional
	Resources []string `json:"resources,omitempty" tf:"resources"`
}

type ManagerServicePerimeterSpecSpecEgressPolicies struct {
	// Defines conditions on the source of a request causing this 'EgressPolicy' to apply.
	// +optional
	EgressFrom *ManagerServicePerimeterSpecSpecEgressPoliciesEgressFrom `json:"egressFrom,omitempty" tf:"egress_from"`
	// Defines the conditions on the 'ApiOperation' and destination resources that
	// cause this 'EgressPolicy' to apply.
	// +optional
	EgressTo *ManagerServicePerimeterSpecSpecEgressPoliciesEgressTo `json:"egressTo,omitempty" tf:"egress_to"`
}

type ManagerServicePerimeterSpecSpecIngressPoliciesIngressFromSources struct {
	// An 'AccessLevel' resource name that allow resources within the
	// 'ServicePerimeters' to be accessed from the internet. 'AccessLevels' listed
	// must be in the same policy as this 'ServicePerimeter'. Referencing a nonexistent
	// 'AccessLevel' will cause an error. If no 'AccessLevel' names are listed,
	// resources within the perimeter can only be accessed via Google Cloud calls
	// with request origins within the perimeter.
	// Example 'accessPolicies/MY_POLICY/accessLevels/MY_LEVEL.'
	// If * is specified, then all IngressSources will be allowed.
	// +optional
	AccessLevel *string `json:"accessLevel,omitempty" tf:"access_level"`
	// A Google Cloud resource that is allowed to ingress the perimeter.
	// Requests from these resources will be allowed to access perimeter data.
	// Currently only projects are allowed. Format 'projects/{project_number}'
	// The project may be in any Google Cloud organization, not just the
	// organization that the perimeter is defined in. '*' is not allowed, the case
	// of allowing all Google Cloud resources only is not supported.
	// +optional
	Resource *string `json:"resource,omitempty" tf:"resource"`
}

type ManagerServicePerimeterSpecSpecIngressPoliciesIngressFrom struct {
	// A list of identities that are allowed access through this ingress policy.
	// Should be in the format of email address. The email address should represent
	// individual user or service account only.
	// +optional
	Identities []string `json:"identities,omitempty" tf:"identities"`
	// Specifies the type of identities that are allowed access from outside the
	// perimeter. If left unspecified, then members of 'identities' field will be
	// allowed access. Possible values: ["IDENTITY_TYPE_UNSPECIFIED", "ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT"]
	// +optional
	IdentityType *string `json:"identityType,omitempty" tf:"identity_type"`
	// Sources that this 'IngressPolicy' authorizes access from.
	// +optional
	Sources []ManagerServicePerimeterSpecSpecIngressPoliciesIngressFromSources `json:"sources,omitempty" tf:"sources"`
}

type ManagerServicePerimeterSpecSpecIngressPoliciesIngressToOperationsMethodSelectors struct {
	// Value for method should be a valid method name for the corresponding
	// serviceName in 'ApiOperation'. If '*' used as value for 'method', then
	// ALL methods and permissions are allowed.
	// +optional
	Method *string `json:"method,omitempty" tf:"method"`
	// Value for permission should be a valid Cloud IAM permission for the
	// corresponding 'serviceName' in 'ApiOperation'.
	// +optional
	Permission *string `json:"permission,omitempty" tf:"permission"`
}

type ManagerServicePerimeterSpecSpecIngressPoliciesIngressToOperations struct {
	// API methods or permissions to allow. Method or permission must belong to
	// the service specified by serviceName field. A single 'MethodSelector' entry
	// with '*' specified for the method field will allow all methods AND
	// permissions for the service specified in 'serviceName'.
	// +optional
	MethodSelectors []ManagerServicePerimeterSpecSpecIngressPoliciesIngressToOperationsMethodSelectors `json:"methodSelectors,omitempty" tf:"method_selectors"`
	// The name of the API whose methods or permissions the 'IngressPolicy' or
	// 'EgressPolicy' want to allow. A single 'ApiOperation' with 'serviceName'
	// field set to '*' will allow all methods AND permissions for all services.
	// +optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name"`
}

type ManagerServicePerimeterSpecSpecIngressPoliciesIngressTo struct {
	// A list of 'ApiOperations' the sources specified in corresponding 'IngressFrom'
	// are allowed to perform in this 'ServicePerimeter'.
	// +optional
	Operations []ManagerServicePerimeterSpecSpecIngressPoliciesIngressToOperations `json:"operations,omitempty" tf:"operations"`
	// A list of resources, currently only projects in the form
	// 'projects/<projectnumber>', protected by this 'ServicePerimeter'
	// that are allowed to be accessed by sources defined in the
	// corresponding 'IngressFrom'. A request matches if it contains
	// a resource in this list. If '*' is specified for resources,
	// then this 'IngressTo' rule will authorize access to all
	// resources inside the perimeter, provided that the request
	// also matches the 'operations' field.
	// +optional
	Resources []string `json:"resources,omitempty" tf:"resources"`
}

type ManagerServicePerimeterSpecSpecIngressPolicies struct {
	// Defines the conditions on the source of a request causing this 'IngressPolicy'
	// to apply.
	// +optional
	IngressFrom *ManagerServicePerimeterSpecSpecIngressPoliciesIngressFrom `json:"ingressFrom,omitempty" tf:"ingress_from"`
	// Defines the conditions on the 'ApiOperation' and request destination that cause
	// this 'IngressPolicy' to apply.
	// +optional
	IngressTo *ManagerServicePerimeterSpecSpecIngressPoliciesIngressTo `json:"ingressTo,omitempty" tf:"ingress_to"`
}

type ManagerServicePerimeterSpecSpecVpcAccessibleServices struct {
	// The list of APIs usable within the Service Perimeter.
	// Must be empty unless 'enableRestriction' is True.
	// +optional
	AllowedServices []string `json:"allowedServices,omitempty" tf:"allowed_services"`
	// Whether to restrict API calls within the Service Perimeter to the
	// list of APIs specified in 'allowedServices'.
	// +optional
	EnableRestriction *bool `json:"enableRestriction,omitempty" tf:"enable_restriction"`
}

type ManagerServicePerimeterSpecSpec struct {
	// A list of AccessLevel resource names that allow resources within
	// the ServicePerimeter to be accessed from the internet.
	// AccessLevels listed must be in the same policy as this
	// ServicePerimeter. Referencing a nonexistent AccessLevel is a
	// syntax error. If no AccessLevel names are listed, resources within
	// the perimeter can only be accessed via GCP calls with request
	// origins within the perimeter. For Service Perimeter Bridge, must
	// be empty.
	//
	// Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}
	// +optional
	AccessLevels []string `json:"accessLevels,omitempty" tf:"access_levels"`
	// List of EgressPolicies to apply to the perimeter. A perimeter may
	// have multiple EgressPolicies, each of which is evaluated separately.
	// Access is granted if any EgressPolicy grants it. Must be empty for
	// a perimeter bridge.
	// +optional
	EgressPolicies []ManagerServicePerimeterSpecSpecEgressPolicies `json:"egressPolicies,omitempty" tf:"egress_policies"`
	// List of 'IngressPolicies' to apply to the perimeter. A perimeter may
	// have multiple 'IngressPolicies', each of which is evaluated
	// separately. Access is granted if any 'Ingress Policy' grants it.
	// Must be empty for a perimeter bridge.
	// +optional
	IngressPolicies []ManagerServicePerimeterSpecSpecIngressPolicies `json:"ingressPolicies,omitempty" tf:"ingress_policies"`
	// A list of GCP resources that are inside of the service perimeter.
	// Currently only projects are allowed.
	// Format: projects/{project_number}
	// +optional
	Resources []string `json:"resources,omitempty" tf:"resources"`
	// GCP services that are subject to the Service Perimeter
	// restrictions. Must contain a list of services. For example, if
	// 'storage.googleapis.com' is specified, access to the storage
	// buckets inside the perimeter must meet the perimeter's access
	// restrictions.
	// +optional
	RestrictedServices []string `json:"restrictedServices,omitempty" tf:"restricted_services"`
	// Specifies how APIs are allowed to communicate within the Service
	// Perimeter.
	// +optional
	VpcAccessibleServices *ManagerServicePerimeterSpecSpecVpcAccessibleServices `json:"vpcAccessibleServices,omitempty" tf:"vpc_accessible_services"`
}

type ManagerServicePerimeterSpecStatusEgressPoliciesEgressFrom struct {
	// A list of identities that are allowed access through this 'EgressPolicy'.
	// Should be in the format of email address. The email address should
	// represent individual user or service account only.
	// +optional
	Identities []string `json:"identities,omitempty" tf:"identities"`
	// Specifies the type of identities that are allowed access to outside the
	// perimeter. If left unspecified, then members of 'identities' field will
	// be allowed access. Possible values: ["IDENTITY_TYPE_UNSPECIFIED", "ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT"]
	// +optional
	IdentityType *string `json:"identityType,omitempty" tf:"identity_type"`
}

type ManagerServicePerimeterSpecStatusEgressPoliciesEgressToOperationsMethodSelectors struct {
	// Value for 'method' should be a valid method name for the corresponding
	// 'serviceName' in 'ApiOperation'. If '*' used as value for method,
	// then ALL methods and permissions are allowed.
	// +optional
	Method *string `json:"method,omitempty" tf:"method"`
	// Value for permission should be a valid Cloud IAM permission for the
	// corresponding 'serviceName' in 'ApiOperation'.
	// +optional
	Permission *string `json:"permission,omitempty" tf:"permission"`
}

type ManagerServicePerimeterSpecStatusEgressPoliciesEgressToOperations struct {
	// API methods or permissions to allow. Method or permission must belong
	// to the service specified by 'serviceName' field. A single MethodSelector
	// entry with '*' specified for the 'method' field will allow all methods
	// AND permissions for the service specified in 'serviceName'.
	// +optional
	MethodSelectors []ManagerServicePerimeterSpecStatusEgressPoliciesEgressToOperationsMethodSelectors `json:"methodSelectors,omitempty" tf:"method_selectors"`
	// The name of the API whose methods or permissions the 'IngressPolicy' or
	// 'EgressPolicy' want to allow. A single 'ApiOperation' with serviceName
	// field set to '*' will allow all methods AND permissions for all services.
	// +optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name"`
}

type ManagerServicePerimeterSpecStatusEgressPoliciesEgressTo struct {
	// A list of 'ApiOperations' that this egress rule applies to. A request matches
	// if it contains an operation/service in this list.
	// +optional
	Operations []ManagerServicePerimeterSpecStatusEgressPoliciesEgressToOperations `json:"operations,omitempty" tf:"operations"`
	// A list of resources, currently only projects in the form
	// 'projects/<projectnumber>', that match this to stanza. A request matches
	// if it contains a resource in this list. If * is specified for resources,
	// then this 'EgressTo' rule will authorize access to all resources outside
	// the perimeter.
	// +optional
	Resources []string `json:"resources,omitempty" tf:"resources"`
}

type ManagerServicePerimeterSpecStatusEgressPolicies struct {
	// Defines conditions on the source of a request causing this 'EgressPolicy' to apply.
	// +optional
	EgressFrom *ManagerServicePerimeterSpecStatusEgressPoliciesEgressFrom `json:"egressFrom,omitempty" tf:"egress_from"`
	// Defines the conditions on the 'ApiOperation' and destination resources that
	// cause this 'EgressPolicy' to apply.
	// +optional
	EgressTo *ManagerServicePerimeterSpecStatusEgressPoliciesEgressTo `json:"egressTo,omitempty" tf:"egress_to"`
}

type ManagerServicePerimeterSpecStatusIngressPoliciesIngressFromSources struct {
	// An 'AccessLevel' resource name that allow resources within the
	// 'ServicePerimeters' to be accessed from the internet. 'AccessLevels' listed
	// must be in the same policy as this 'ServicePerimeter'. Referencing a nonexistent
	// 'AccessLevel' will cause an error. If no 'AccessLevel' names are listed,
	// resources within the perimeter can only be accessed via Google Cloud calls
	// with request origins within the perimeter.
	// Example 'accessPolicies/MY_POLICY/accessLevels/MY_LEVEL.'
	// If * is specified, then all IngressSources will be allowed.
	// +optional
	AccessLevel *string `json:"accessLevel,omitempty" tf:"access_level"`
	// A Google Cloud resource that is allowed to ingress the perimeter.
	// Requests from these resources will be allowed to access perimeter data.
	// Currently only projects are allowed. Format 'projects/{project_number}'
	// The project may be in any Google Cloud organization, not just the
	// organization that the perimeter is defined in. '*' is not allowed, the case
	// of allowing all Google Cloud resources only is not supported.
	// +optional
	Resource *string `json:"resource,omitempty" tf:"resource"`
}

type ManagerServicePerimeterSpecStatusIngressPoliciesIngressFrom struct {
	// A list of identities that are allowed access through this ingress policy.
	// Should be in the format of email address. The email address should represent
	// individual user or service account only.
	// +optional
	Identities []string `json:"identities,omitempty" tf:"identities"`
	// Specifies the type of identities that are allowed access from outside the
	// perimeter. If left unspecified, then members of 'identities' field will be
	// allowed access. Possible values: ["IDENTITY_TYPE_UNSPECIFIED", "ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT"]
	// +optional
	IdentityType *string `json:"identityType,omitempty" tf:"identity_type"`
	// Sources that this 'IngressPolicy' authorizes access from.
	// +optional
	Sources []ManagerServicePerimeterSpecStatusIngressPoliciesIngressFromSources `json:"sources,omitempty" tf:"sources"`
}

type ManagerServicePerimeterSpecStatusIngressPoliciesIngressToOperationsMethodSelectors struct {
	// Value for method should be a valid method name for the corresponding
	// serviceName in 'ApiOperation'. If '*' used as value for 'method', then
	// ALL methods and permissions are allowed.
	// +optional
	Method *string `json:"method,omitempty" tf:"method"`
	// Value for permission should be a valid Cloud IAM permission for the
	// corresponding 'serviceName' in 'ApiOperation'.
	// +optional
	Permission *string `json:"permission,omitempty" tf:"permission"`
}

type ManagerServicePerimeterSpecStatusIngressPoliciesIngressToOperations struct {
	// API methods or permissions to allow. Method or permission must belong to
	// the service specified by serviceName field. A single 'MethodSelector' entry
	// with '*' specified for the method field will allow all methods AND
	// permissions for the service specified in 'serviceName'.
	// +optional
	MethodSelectors []ManagerServicePerimeterSpecStatusIngressPoliciesIngressToOperationsMethodSelectors `json:"methodSelectors,omitempty" tf:"method_selectors"`
	// The name of the API whose methods or permissions the 'IngressPolicy' or
	// 'EgressPolicy' want to allow. A single 'ApiOperation' with 'serviceName'
	// field set to '*' will allow all methods AND permissions for all services.
	// +optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name"`
}

type ManagerServicePerimeterSpecStatusIngressPoliciesIngressTo struct {
	// A list of 'ApiOperations' the sources specified in corresponding 'IngressFrom'
	// are allowed to perform in this 'ServicePerimeter'.
	// +optional
	Operations []ManagerServicePerimeterSpecStatusIngressPoliciesIngressToOperations `json:"operations,omitempty" tf:"operations"`
	// A list of resources, currently only projects in the form
	// 'projects/<projectnumber>', protected by this 'ServicePerimeter'
	// that are allowed to be accessed by sources defined in the
	// corresponding 'IngressFrom'. A request matches if it contains
	// a resource in this list. If '*' is specified for resources,
	// then this 'IngressTo' rule will authorize access to all
	// resources inside the perimeter, provided that the request
	// also matches the 'operations' field.
	// +optional
	Resources []string `json:"resources,omitempty" tf:"resources"`
}

type ManagerServicePerimeterSpecStatusIngressPolicies struct {
	// Defines the conditions on the source of a request causing this 'IngressPolicy'
	// to apply.
	// +optional
	IngressFrom *ManagerServicePerimeterSpecStatusIngressPoliciesIngressFrom `json:"ingressFrom,omitempty" tf:"ingress_from"`
	// Defines the conditions on the 'ApiOperation' and request destination that cause
	// this 'IngressPolicy' to apply.
	// +optional
	IngressTo *ManagerServicePerimeterSpecStatusIngressPoliciesIngressTo `json:"ingressTo,omitempty" tf:"ingress_to"`
}

type ManagerServicePerimeterSpecStatusVpcAccessibleServices struct {
	// The list of APIs usable within the Service Perimeter.
	// Must be empty unless 'enableRestriction' is True.
	// +optional
	AllowedServices []string `json:"allowedServices,omitempty" tf:"allowed_services"`
	// Whether to restrict API calls within the Service Perimeter to the
	// list of APIs specified in 'allowedServices'.
	// +optional
	EnableRestriction *bool `json:"enableRestriction,omitempty" tf:"enable_restriction"`
}

type ManagerServicePerimeterSpecStatus struct {
	// A list of AccessLevel resource names that allow resources within
	// the ServicePerimeter to be accessed from the internet.
	// AccessLevels listed must be in the same policy as this
	// ServicePerimeter. Referencing a nonexistent AccessLevel is a
	// syntax error. If no AccessLevel names are listed, resources within
	// the perimeter can only be accessed via GCP calls with request
	// origins within the perimeter. For Service Perimeter Bridge, must
	// be empty.
	//
	// Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}
	// +optional
	AccessLevels []string `json:"accessLevels,omitempty" tf:"access_levels"`
	// List of EgressPolicies to apply to the perimeter. A perimeter may
	// have multiple EgressPolicies, each of which is evaluated separately.
	// Access is granted if any EgressPolicy grants it. Must be empty for
	// a perimeter bridge.
	// +optional
	EgressPolicies []ManagerServicePerimeterSpecStatusEgressPolicies `json:"egressPolicies,omitempty" tf:"egress_policies"`
	// List of 'IngressPolicies' to apply to the perimeter. A perimeter may
	// have multiple 'IngressPolicies', each of which is evaluated
	// separately. Access is granted if any 'Ingress Policy' grants it.
	// Must be empty for a perimeter bridge.
	// +optional
	IngressPolicies []ManagerServicePerimeterSpecStatusIngressPolicies `json:"ingressPolicies,omitempty" tf:"ingress_policies"`
	// A list of GCP resources that are inside of the service perimeter.
	// Currently only projects are allowed.
	// Format: projects/{project_number}
	// +optional
	Resources []string `json:"resources,omitempty" tf:"resources"`
	// GCP services that are subject to the Service Perimeter
	// restrictions. Must contain a list of services. For example, if
	// 'storage.googleapis.com' is specified, access to the storage
	// buckets inside the perimeter must meet the perimeter's access
	// restrictions.
	// +optional
	RestrictedServices []string `json:"restrictedServices,omitempty" tf:"restricted_services"`
	// Specifies how APIs are allowed to communicate within the Service
	// Perimeter.
	// +optional
	VpcAccessibleServices *ManagerServicePerimeterSpecStatusVpcAccessibleServices `json:"vpcAccessibleServices,omitempty" tf:"vpc_accessible_services"`
}

type ManagerServicePerimeterSpec struct {
	State *ManagerServicePerimeterSpecResource `json:"state,omitempty" tf:"-"`

	Resource ManagerServicePerimeterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ManagerServicePerimeterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Time the AccessPolicy was created in UTC.
	// +optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time"`
	// Description of the ServicePerimeter and its use. Does not affect
	// behavior.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Resource name for the ServicePerimeter. The short_name component must
	// begin with a letter and only include alphanumeric and '_'.
	// Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}
	Name *string `json:"name" tf:"name"`
	// The AccessPolicy this ServicePerimeter lives in.
	// Format: accessPolicies/{policy_id}
	Parent *string `json:"parent" tf:"parent"`
	// Specifies the type of the Perimeter. There are two types: regular and
	// bridge. Regular Service Perimeter contains resources, access levels,
	// and restricted services. Every resource can be in at most
	// ONE regular Service Perimeter.
	//
	// In addition to being in a regular service perimeter, a resource can also
	// be in zero or more perimeter bridges. A perimeter bridge only contains
	// resources. Cross project operations are permitted if all effected
	// resources share some perimeter (whether bridge or regular). Perimeter
	// Bridge does not contain access levels or services: those are governed
	// entirely by the regular perimeter that resource is in.
	//
	// Perimeter Bridges are typically useful when building more complex
	// topologies with many independent perimeters that need to share some data
	// with a common perimeter, but should not be able to share data among
	// themselves. Default value: "PERIMETER_TYPE_REGULAR" Possible values: ["PERIMETER_TYPE_REGULAR", "PERIMETER_TYPE_BRIDGE"]
	// +optional
	PerimeterType *string `json:"perimeterType,omitempty" tf:"perimeter_type"`
	// Proposed (or dry run) ServicePerimeter configuration.
	// This configuration allows to specify and test ServicePerimeter configuration
	// without enforcing actual access restrictions. Only allowed to be set when
	// the 'useExplicitDryRunSpec' flag is set.
	// +optional
	Spec *ManagerServicePerimeterSpecSpec `json:"spec,omitempty" tf:"spec"`
	// ServicePerimeter configuration. Specifies sets of resources,
	// restricted services and access levels that determine
	// perimeter content and boundaries.
	// +optional
	Status *ManagerServicePerimeterSpecStatus `json:"status,omitempty" tf:"status"`
	// Human readable title. Must be unique within the Policy.
	Title *string `json:"title" tf:"title"`
	// Time the AccessPolicy was updated in UTC.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time"`
	// Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
	// for all Service Perimeters, and that spec is identical to the status for those
	// Service Perimeters. When this flag is set, it inhibits the generation of the
	// implicit spec, thereby allowing the user to explicitly provide a
	// configuration ("spec") to use in a dry-run version of the Service Perimeter.
	// This allows the user to test changes to the enforced config ("status") without
	// actually enforcing them. This testing is done through analyzing the differences
	// between currently enforced and suggested restrictions. useExplicitDryRunSpec must
	// bet set to True if any of the fields in the spec are set to non-default values.
	// +optional
	UseExplicitDryRunSpec *bool `json:"useExplicitDryRunSpec,omitempty" tf:"use_explicit_dry_run_spec"`
}

type ManagerServicePerimeterStatus struct {
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

// ManagerServicePerimeterList is a list of ManagerServicePerimeters
type ManagerServicePerimeterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagerServicePerimeter CRD objects
	Items []ManagerServicePerimeter `json:"items,omitempty"`
}
