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

type Firewall struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallSpec   `json:"spec,omitempty"`
	Status            FirewallStatus `json:"status,omitempty"`
}

type FirewallSpecAllow struct {
	// An optional list of ports to which this rule applies. This field
	// is only applicable for UDP or TCP protocol. Each entry must be
	// either an integer or a range. If not specified, this rule
	// applies to connections through any port.
	//
	// Example inputs include: ["22"], ["80","443"], and
	// ["12345-12349"].
	// +optional
	Ports []string `json:"ports,omitempty" tf:"ports"`
	// The IP protocol to which this rule applies. The protocol type is
	// required when creating a firewall rule. This value can either be
	// one of the following well known protocol strings (tcp, udp,
	// icmp, esp, ah, sctp, ipip, all), or the IP protocol number.
	Protocol *string `json:"protocol" tf:"protocol"`
}

type FirewallSpecDeny struct {
	// An optional list of ports to which this rule applies. This field
	// is only applicable for UDP or TCP protocol. Each entry must be
	// either an integer or a range. If not specified, this rule
	// applies to connections through any port.
	//
	// Example inputs include: ["22"], ["80","443"], and
	// ["12345-12349"].
	// +optional
	Ports []string `json:"ports,omitempty" tf:"ports"`
	// The IP protocol to which this rule applies. The protocol type is
	// required when creating a firewall rule. This value can either be
	// one of the following well known protocol strings (tcp, udp,
	// icmp, esp, ah, sctp, ipip, all), or the IP protocol number.
	Protocol *string `json:"protocol" tf:"protocol"`
}

type FirewallSpecLogConfig struct {
	// This field denotes whether to include or exclude metadata for firewall logs. Possible values: ["EXCLUDE_ALL_METADATA", "INCLUDE_ALL_METADATA"]
	Metadata *string `json:"metadata" tf:"metadata"`
}

type FirewallSpec struct {
	KubeformOutput *FirewallSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource FirewallSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type FirewallSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The list of ALLOW rules specified by this firewall. Each rule
	// specifies a protocol and port-range tuple that describes a permitted
	// connection.
	// +optional
	Allow []FirewallSpecAllow `json:"allow,omitempty" tf:"allow"`
	// Creation timestamp in RFC3339 text format.
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp"`
	// The list of DENY rules specified by this firewall. Each rule specifies
	// a protocol and port-range tuple that describes a denied connection.
	// +optional
	Deny []FirewallSpecDeny `json:"deny,omitempty" tf:"deny"`
	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// If destination ranges are specified, the firewall will apply only to
	// traffic that has destination IP address in these ranges. These ranges
	// must be expressed in CIDR format. Only IPv4 is supported.
	// +optional
	DestinationRanges []string `json:"destinationRanges,omitempty" tf:"destination_ranges"`
	// Direction of traffic to which this firewall applies; default is
	// INGRESS. Note: For INGRESS traffic, it is NOT supported to specify
	// destinationRanges; For EGRESS traffic, it is NOT supported to specify
	// sourceRanges OR sourceTags. Possible values: ["INGRESS", "EGRESS"]
	// +optional
	Direction *string `json:"direction,omitempty" tf:"direction"`
	// Denotes whether the firewall rule is disabled, i.e not applied to the
	// network it is associated with. When set to true, the firewall rule is
	// not enforced and the network behaves as if it did not exist. If this
	// is unspecified, the firewall rule will be enabled.
	// +optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled"`
	// This field denotes whether to enable logging for a particular firewall rule. If logging is enabled, logs will be exported to Stackdriver.
	// +optional
	// Deprecated
	EnableLogging *bool `json:"enableLogging,omitempty" tf:"enable_logging"`
	// This field denotes the logging options for a particular firewall rule.
	// If defined, logging is enabled, and logs will be exported to Cloud Logging.
	// +optional
	LogConfig *FirewallSpecLogConfig `json:"logConfig,omitempty" tf:"log_config"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `json:"name" tf:"name"`
	// The name or self_link of the network to attach this firewall to.
	Network *string `json:"network" tf:"network"`
	// Priority for this rule. This is an integer between 0 and 65535, both
	// inclusive. When not specified, the value assumed is 1000. Relative
	// priorities determine precedence of conflicting rules. Lower value of
	// priority implies higher precedence (eg, a rule with priority 0 has
	// higher precedence than a rule with priority 1). DENY rules take
	// precedence over ALLOW rules having equal priority.
	// +optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
	// If source ranges are specified, the firewall will apply only to
	// traffic that has source IP address in these ranges. These ranges must
	// be expressed in CIDR format. One or both of sourceRanges and
	// sourceTags may be set. If both properties are set, the firewall will
	// apply to traffic that has source IP address within sourceRanges OR the
	// source IP that belongs to a tag listed in the sourceTags property. The
	// connection does not need to match both properties for the firewall to
	// apply. Only IPv4 is supported.
	// +optional
	SourceRanges []string `json:"sourceRanges,omitempty" tf:"source_ranges"`
	// If source service accounts are specified, the firewall will apply only
	// to traffic originating from an instance with a service account in this
	// list. Source service accounts cannot be used to control traffic to an
	// instance's external IP address because service accounts are associated
	// with an instance, not an IP address. sourceRanges can be set at the
	// same time as sourceServiceAccounts. If both are set, the firewall will
	// apply to traffic that has source IP address within sourceRanges OR the
	// source IP belongs to an instance with service account listed in
	// sourceServiceAccount. The connection does not need to match both
	// properties for the firewall to apply. sourceServiceAccounts cannot be
	// used at the same time as sourceTags or targetTags.
	// +optional
	// +kubebuilder:validation:MaxItems=10
	SourceServiceAccounts []string `json:"sourceServiceAccounts,omitempty" tf:"source_service_accounts"`
	// If source tags are specified, the firewall will apply only to traffic
	// with source IP that belongs to a tag listed in source tags. Source
	// tags cannot be used to control traffic to an instance's external IP
	// address. Because tags are associated with an instance, not an IP
	// address. One or both of sourceRanges and sourceTags may be set. If
	// both properties are set, the firewall will apply to traffic that has
	// source IP address within sourceRanges OR the source IP that belongs to
	// a tag listed in the sourceTags property. The connection does not need
	// to match both properties for the firewall to apply.
	// +optional
	SourceTags []string `json:"sourceTags,omitempty" tf:"source_tags"`
	// A list of service accounts indicating sets of instances located in the
	// network that may make network connections as specified in allowed[].
	// targetServiceAccounts cannot be used at the same time as targetTags or
	// sourceTags. If neither targetServiceAccounts nor targetTags are
	// specified, the firewall rule applies to all instances on the specified
	// network.
	// +optional
	// +kubebuilder:validation:MaxItems=10
	TargetServiceAccounts []string `json:"targetServiceAccounts,omitempty" tf:"target_service_accounts"`
	// A list of instance tags indicating sets of instances located in the
	// network that may make network connections as specified in allowed[].
	// If no targetTags are specified, the firewall rule applies to all
	// instances on the specified network.
	// +optional
	TargetTags []string `json:"targetTags,omitempty" tf:"target_tags"`
}

type FirewallStatus struct {
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

// FirewallList is a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Firewall CRD objects
	Items []Firewall `json:"items,omitempty"`
}
