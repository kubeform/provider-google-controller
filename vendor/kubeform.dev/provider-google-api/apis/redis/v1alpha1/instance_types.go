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

type InstanceSpecServerCaCerts struct {
	// Serial number, as extracted from the certificate.
	// +optional
	Cert *string `json:"cert,omitempty" tf:"cert"`
	// The time when the certificate was created.
	// +optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time"`
	// The time when the certificate expires.
	// +optional
	ExpireTime *string `json:"expireTime,omitempty" tf:"expire_time"`
	// Serial number, as extracted from the certificate.
	// +optional
	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number"`
	// Sha1 Fingerprint of the certificate.
	// +optional
	Sha1Fingerprint *string `json:"sha1Fingerprint,omitempty" tf:"sha1_fingerprint"`
}

type InstanceSpec struct {
	State *InstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type InstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Only applicable to STANDARD_HA tier which protects the instance
	// against zonal failures by provisioning it across two zones.
	// If provided, it must be a different zone from the one provided in
	// [locationId].
	// +optional
	AlternativeLocationID *string `json:"alternativeLocationID,omitempty" tf:"alternative_location_id"`
	// Optional. Indicates whether OSS Redis AUTH is enabled for the
	// instance. If set to "true" AUTH is enabled on the instance.
	// Default value is "false" meaning AUTH is disabled.
	// +optional
	AuthEnabled *bool `json:"authEnabled,omitempty" tf:"auth_enabled"`
	// AUTH String set on the instance. This field will only be populated if auth_enabled is true.
	// +optional
	AuthString *string `json:"-" sensitive:"true" tf:"auth_string"`
	// The full name of the Google Compute Engine network to which the
	// instance is connected. If left unspecified, the default network
	// will be used.
	// +optional
	AuthorizedNetwork *string `json:"authorizedNetwork,omitempty" tf:"authorized_network"`
	// The connection mode of the Redis instance. Default value: "DIRECT_PEERING" Possible values: ["DIRECT_PEERING", "PRIVATE_SERVICE_ACCESS"]
	// +optional
	ConnectMode *string `json:"connectMode,omitempty" tf:"connect_mode"`
	// The time the instance was created in RFC3339 UTC "Zulu" format,
	// accurate to nanoseconds.
	// +optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time"`
	// The current zone where the Redis endpoint is placed.
	// For Basic Tier instances, this will always be the same as the
	// [locationId] provided by the user at creation time. For Standard Tier
	// instances, this can be either [locationId] or [alternativeLocationId]
	// and can change after a failover event.
	// +optional
	CurrentLocationID *string `json:"currentLocationID,omitempty" tf:"current_location_id"`
	// An arbitrary and optional user-provided name for the instance.
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// Hostname or IP address of the exposed Redis endpoint used by clients
	// to connect to the service.
	// +optional
	Host *string `json:"host,omitempty" tf:"host"`
	// Resource labels to represent user provided metadata.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// The zone where the instance will be provisioned. If not provided,
	// the service will choose a zone for the instance. For STANDARD_HA tier,
	// instances will be created across two zones for protection against
	// zonal failures. If [alternativeLocationId] is also provided, it must
	// be different from [locationId].
	// +optional
	LocationID *string `json:"locationID,omitempty" tf:"location_id"`
	// Redis memory size in GiB.
	MemorySizeGb *int64 `json:"memorySizeGb" tf:"memory_size_gb"`
	// The ID of the instance or a fully qualified identifier for the instance.
	Name *string `json:"name" tf:"name"`
	// Output only. Cloud IAM identity used by import / export operations
	// to transfer data to/from Cloud Storage. Format is "serviceAccount:".
	// The value may change over time for a given instance so should be
	// checked before each import/export operation.
	// +optional
	PersistenceIamIdentity *string `json:"persistenceIamIdentity,omitempty" tf:"persistence_iam_identity"`
	// The port number of the exposed Redis endpoint.
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Redis configuration parameters, according to http://redis.io/topics/config.
	// Please check Memorystore documentation for the list of supported parameters:
	// https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs
	// +optional
	RedisConfigs *map[string]string `json:"redisConfigs,omitempty" tf:"redis_configs"`
	// The version of Redis software. If not provided, latest supported
	// version will be used. Please check the API documentation linked
	// at the top for the latest valid values.
	// +optional
	RedisVersion *string `json:"redisVersion,omitempty" tf:"redis_version"`
	// The name of the Redis region of the instance.
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// The CIDR range of internal addresses that are reserved for this
	// instance. If not provided, the service will choose an unused /29
	// block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be
	// unique and non-overlapping with existing subnets in an authorized
	// network.
	// +optional
	ReservedIPRange *string `json:"reservedIPRange,omitempty" tf:"reserved_ip_range"`
	// List of server CA certificates for the instance.
	// +optional
	ServerCaCerts []InstanceSpecServerCaCerts `json:"serverCaCerts,omitempty" tf:"server_ca_certs"`
	// The service tier of the instance. Must be one of these values:
	//
	// - BASIC: standalone instance
	// - STANDARD_HA: highly available primary/replica instances Default value: "BASIC" Possible values: ["BASIC", "STANDARD_HA"]
	// +optional
	Tier *string `json:"tier,omitempty" tf:"tier"`
	// The TLS mode of the Redis instance, If not provided, TLS is disabled for the instance.
	//
	// - SERVER_AUTHENTICATION: Client to Server traffic encryption enabled with server authentcation Default value: "DISABLED" Possible values: ["SERVER_AUTHENTICATION", "DISABLED"]
	// +optional
	TransitEncryptionMode *string `json:"transitEncryptionMode,omitempty" tf:"transit_encryption_mode"`
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
