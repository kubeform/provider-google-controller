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

type Subnetwork struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetworkSpec   `json:"spec,omitempty"`
	Status            SubnetworkStatus `json:"status,omitempty"`
}

type SubnetworkSpecLogConfig struct {
	// Can only be specified if VPC flow logging for this subnetwork is enabled.
	// Toggles the aggregation interval for collecting flow logs. Increasing the
	// interval time will reduce the amount of generated flow logs for long
	// lasting connections. Default is an interval of 5 seconds per connection. Default value: "INTERVAL_5_SEC" Possible values: ["INTERVAL_5_SEC", "INTERVAL_30_SEC", "INTERVAL_1_MIN", "INTERVAL_5_MIN", "INTERVAL_10_MIN", "INTERVAL_15_MIN"]
	// +optional
	AggregationInterval *string `json:"aggregationInterval,omitempty" tf:"aggregation_interval"`
	// Export filter used to define which VPC flow logs should be logged, as as CEL expression. See
	// https://cloud.google.com/vpc/docs/flow-logs#filtering for details on how to format this field.
	// The default value is 'true', which evaluates to include everything.
	// +optional
	FilterExpr *string `json:"filterExpr,omitempty" tf:"filter_expr"`
	// Can only be specified if VPC flow logging for this subnetwork is enabled.
	// The value of the field must be in [0, 1]. Set the sampling rate of VPC
	// flow logs within the subnetwork where 1.0 means all collected logs are
	// reported and 0.0 means no logs are reported. Default is 0.5 which means
	// half of all collected logs are reported.
	// +optional
	FlowSampling *float64 `json:"flowSampling,omitempty" tf:"flow_sampling"`
	// Can only be specified if VPC flow logging for this subnetwork is enabled.
	// Configures whether metadata fields should be added to the reported VPC
	// flow logs. Default value: "INCLUDE_ALL_METADATA" Possible values: ["EXCLUDE_ALL_METADATA", "INCLUDE_ALL_METADATA", "CUSTOM_METADATA"]
	// +optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata"`
	// List of metadata fields that should be added to reported logs.
	// Can only be specified if VPC flow logs for this subnetwork is enabled and "metadata" is set to CUSTOM_METADATA.
	// +optional
	MetadataFields []string `json:"metadataFields,omitempty" tf:"metadata_fields"`
}

type SubnetworkSpecSecondaryIPRange struct {
	// The range of IP addresses belonging to this subnetwork secondary
	// range. Provide this property when you create the subnetwork.
	// Ranges must be unique and non-overlapping with all primary and
	// secondary IP ranges within a network. Only IPv4 is supported.
	IpCIDRRange *string `json:"ipCIDRRange" tf:"ip_cidr_range"`
	// The name associated with this subnetwork secondary range, used
	// when adding an alias IP range to a VM instance. The name must
	// be 1-63 characters long, and comply with RFC1035. The name
	// must be unique within the subnetwork.
	RangeName *string `json:"rangeName" tf:"range_name"`
}

type SubnetworkSpec struct {
	State *SubnetworkSpecResource `json:"state,omitempty" tf:"-"`

	Resource SubnetworkSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SubnetworkSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Creation timestamp in RFC3339 text format.
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp"`
	// An optional description of this resource. Provide this property when
	// you create the resource. This field can be set only at resource
	// creation time.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Fingerprint of this resource. This field is used internally during updates of this resource.
	// +optional
	// Deprecated
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint"`
	// The gateway address for default routes to reach destination addresses
	// outside this subnetwork.
	// +optional
	GatewayAddress *string `json:"gatewayAddress,omitempty" tf:"gateway_address"`
	// The range of internal addresses that are owned by this subnetwork.
	// Provide this property when you create the subnetwork. For example,
	// 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and
	// non-overlapping within a network. Only IPv4 is supported.
	IpCIDRRange *string `json:"ipCIDRRange" tf:"ip_cidr_range"`
	// Denotes the logging options for the subnetwork flow logs. If logging is enabled
	// logs will be exported to Stackdriver. This field cannot be set if the 'purpose' of this
	// subnetwork is 'INTERNAL_HTTPS_LOAD_BALANCER'
	// +optional
	LogConfig *SubnetworkSpecLogConfig `json:"logConfig,omitempty" tf:"log_config"`
	// The name of the resource, provided by the client when initially
	// creating the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `json:"name" tf:"name"`
	// The network this subnet belongs to.
	// Only networks that are in the distributed mode can have subnetworks.
	Network *string `json:"network" tf:"network"`
	// When enabled, VMs in this subnetwork without external IP addresses can
	// access Google APIs and services by using Private Google Access.
	// +optional
	PrivateIPGoogleAccess *bool `json:"privateIPGoogleAccess,omitempty" tf:"private_ip_google_access"`
	// The private IPv6 google access type for the VMs in this subnet.
	// +optional
	PrivateIpv6GoogleAccess *string `json:"privateIpv6GoogleAccess,omitempty" tf:"private_ipv6_google_access"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The GCP region for this subnetwork.
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// An array of configurations for secondary IP ranges for VM instances
	// contained in this subnetwork. The primary IP of such VM must belong
	// to the primary ipCidrRange of the subnetwork. The alias IPs may belong
	// to either primary or secondary ranges.
	//
	// **Note**: This field uses [attr-as-block mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html) to avoid
	// breaking users during the 0.12 upgrade. To explicitly send a list
	// of zero objects you must use the following syntax:
	// 'example=[]'
	// For more details about this behavior, see [this section](https://www.terraform.io/docs/configuration/attr-as-blocks.html#defining-a-fixed-object-collection-value).
	// +optional
	SecondaryIPRange []SubnetworkSpecSecondaryIPRange `json:"secondaryIPRange,omitempty" tf:"secondary_ip_range"`
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
}

type SubnetworkStatus struct {
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

// SubnetworkList is a list of Subnetworks
type SubnetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Subnetwork CRD objects
	Items []Subnetwork `json:"items,omitempty"`
}
