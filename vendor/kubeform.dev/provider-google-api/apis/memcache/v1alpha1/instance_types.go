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

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpecMemcacheNodes struct {
	// Hostname or IP address of the Memcached node used by the clients to connect to the Memcached server on this node.
	// +optional
	Host *string `json:"host,omitempty" tf:"host"`
	// Identifier of the Memcached node. The node id does not include project or location like the Memcached instance name.
	// +optional
	NodeID *string `json:"nodeID,omitempty" tf:"node_id"`
	// The port number of the Memcached server on this node.
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// Current state of the Memcached node.
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// Location (GCP Zone) for the Memcached node.
	// +optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
}

type InstanceSpecMemcacheParameters struct {
	// This is a unique ID associated with this set of parameters.
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// User-defined set of parameters to use in the memcache process.
	// +optional
	Params *map[string]string `json:"params,omitempty" tf:"params"`
}

type InstanceSpecNodeConfig struct {
	// Number of CPUs per node.
	CpuCount *int64 `json:"cpuCount" tf:"cpu_count"`
	// Memory size in Mebibytes for each memcache node.
	MemorySizeMb *int64 `json:"memorySizeMb" tf:"memory_size_mb"`
}

type InstanceSpec struct {
	State *InstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type InstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The full name of the GCE network to connect the instance to.  If not provided,
	// 'default' will be used.
	// +optional
	AuthorizedNetwork *string `json:"authorizedNetwork,omitempty" tf:"authorized_network"`
	// Creation timestamp in RFC3339 text format.
	// +optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time"`
	// Endpoint for Discovery API
	// +optional
	DiscoveryEndpoint *string `json:"discoveryEndpoint,omitempty" tf:"discovery_endpoint"`
	// A user-visible name for the instance.
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// Resource labels to represent user-provided metadata.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// The full version of memcached server running on this instance.
	// +optional
	MemcacheFullVersion *string `json:"memcacheFullVersion,omitempty" tf:"memcache_full_version"`
	// Additional information about the instance state, if available.
	// +optional
	MemcacheNodes []InstanceSpecMemcacheNodes `json:"memcacheNodes,omitempty" tf:"memcache_nodes"`
	// User-specified parameters for this memcache instance.
	// +optional
	MemcacheParameters *InstanceSpecMemcacheParameters `json:"memcacheParameters,omitempty" tf:"memcache_parameters"`
	// The major version of Memcached software. If not provided, latest supported version will be used.
	// Currently the latest supported major version is MEMCACHE_1_5. The minor version will be automatically
	// determined by our system based on the latest supported minor version. Default value: "MEMCACHE_1_5" Possible values: ["MEMCACHE_1_5"]
	// +optional
	MemcacheVersion *string `json:"memcacheVersion,omitempty" tf:"memcache_version"`
	// The resource name of the instance.
	Name *string `json:"name" tf:"name"`
	// Configuration for memcache nodes.
	NodeConfig *InstanceSpecNodeConfig `json:"nodeConfig" tf:"node_config"`
	// Number of nodes in the memcache instance.
	NodeCount *int64 `json:"nodeCount" tf:"node_count"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The region of the Memcache instance. If it is not provided, the provider region is used.
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// Zones where memcache nodes should be provisioned.  If not
	// provided, all zones will be used.
	// +optional
	Zones []string `json:"zones,omitempty" tf:"zones"`
}

type InstanceStatus struct {
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

// InstanceList is a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Instance CRD objects
	Items []Instance `json:"items,omitempty"`
}
