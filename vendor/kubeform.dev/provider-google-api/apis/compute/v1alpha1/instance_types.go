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

type InstanceSpecAttachedDisk struct {
	// Name with which the attached disk is accessible under /dev/disk/by-id/
	// +optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name"`
	// A 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set.
	// +optional
	DiskEncryptionKeyRaw *string `json:"-" sensitive:"true" tf:"disk_encryption_key_raw"`
	// The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource.
	// +optional
	DiskEncryptionKeySha256 *string `json:"diskEncryptionKeySha256,omitempty" tf:"disk_encryption_key_sha256"`
	// The self_link of the encryption key that is stored in Google Cloud KMS to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set.
	// +optional
	KmsKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link"`
	// Read/write mode for the disk. One of "READ_ONLY" or "READ_WRITE".
	// +optional
	Mode *string `json:"mode,omitempty" tf:"mode"`
	// The name or self_link of the disk attached to this instance.
	Source *string `json:"source" tf:"source"`
}

type InstanceSpecBootDiskInitializeParams struct {
	// The image from which this disk was initialised.
	// +optional
	Image *string `json:"image,omitempty" tf:"image"`
	// A set of key/value label pairs assigned to the disk.
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels"`
	// The size of the image in gigabytes.
	// +optional
	Size *int64 `json:"size,omitempty" tf:"size"`
	// The Google Compute Engine disk type. One of pd-standard, pd-ssd or pd-balanced.
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type InstanceSpecBootDisk struct {
	// Whether the disk will be auto-deleted when the instance is deleted.
	// +optional
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete"`
	// Name with which attached disk will be accessible under /dev/disk/by-id/
	// +optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name"`
	// A 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set.
	// +optional
	DiskEncryptionKeyRaw *string `json:"-" sensitive:"true" tf:"disk_encryption_key_raw"`
	// The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource.
	// +optional
	DiskEncryptionKeySha256 *string `json:"diskEncryptionKeySha256,omitempty" tf:"disk_encryption_key_sha256"`
	// Parameters with which a disk was created alongside the instance.
	// +optional
	InitializeParams *InstanceSpecBootDiskInitializeParams `json:"initializeParams,omitempty" tf:"initialize_params"`
	// The self_link of the encryption key that is stored in Google Cloud KMS to encrypt this disk. Only one of kms_key_self_link and disk_encryption_key_raw may be set.
	// +optional
	KmsKeySelfLink *string `json:"kmsKeySelfLink,omitempty" tf:"kms_key_self_link"`
	// Read/write mode for the disk. One of "READ_ONLY" or "READ_WRITE".
	// +optional
	Mode *string `json:"mode,omitempty" tf:"mode"`
	// The name or self_link of the disk attached to this instance.
	// +optional
	Source *string `json:"source,omitempty" tf:"source"`
}

type InstanceSpecConfidentialInstanceConfig struct {
	// Defines whether the instance should have confidential compute enabled.
	EnableConfidentialCompute *bool `json:"enableConfidentialCompute" tf:"enable_confidential_compute"`
}

type InstanceSpecGuestAccelerator struct {
	// The number of the guest accelerator cards exposed to this instance.
	Count *int64 `json:"count" tf:"count"`
	// The accelerator type resource exposed to this instance. E.g. nvidia-tesla-k80.
	Type *string `json:"type" tf:"type"`
}

type InstanceSpecNetworkInterfaceAccessConfig struct {
	// The IP address that is be 1:1 mapped to the instance's network ip.
	// +optional
	NatIP *string `json:"natIP,omitempty" tf:"nat_ip"`
	// The networking tier used for configuring this instance. One of PREMIUM or STANDARD.
	// +optional
	NetworkTier *string `json:"networkTier,omitempty" tf:"network_tier"`
	// The DNS domain name for the public PTR record.
	// +optional
	PublicPtrDomainName *string `json:"publicPtrDomainName,omitempty" tf:"public_ptr_domain_name"`
}

type InstanceSpecNetworkInterfaceAliasIPRange struct {
	// The IP CIDR range represented by this alias IP range.
	IpCIDRRange *string `json:"ipCIDRRange" tf:"ip_cidr_range"`
	// The subnetwork secondary range name specifying the secondary range from which to allocate the IP CIDR range for this alias IP range.
	// +optional
	SubnetworkRangeName *string `json:"subnetworkRangeName,omitempty" tf:"subnetwork_range_name"`
}

type InstanceSpecNetworkInterface struct {
	// Access configurations, i.e. IPs via which this instance can be accessed via the Internet.
	// +optional
	AccessConfig []InstanceSpecNetworkInterfaceAccessConfig `json:"accessConfig,omitempty" tf:"access_config"`
	// An array of alias IP ranges for this network interface.
	// +optional
	AliasIPRange []InstanceSpecNetworkInterfaceAliasIPRange `json:"aliasIPRange,omitempty" tf:"alias_ip_range"`
	// The name of the interface
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// The name or self_link of the network attached to this interface.
	// +optional
	Network *string `json:"network,omitempty" tf:"network"`
	// The private IP address assigned to the instance.
	// +optional
	NetworkIP *string `json:"networkIP,omitempty" tf:"network_ip"`
	// The type of vNIC to be used on this interface. Possible values:GVNIC, VIRTIO_NET
	// +optional
	NicType *string `json:"nicType,omitempty" tf:"nic_type"`
	// The name or self_link of the subnetwork attached to this interface.
	// +optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork"`
	// The project in which the subnetwork belongs.
	// +optional
	SubnetworkProject *string `json:"subnetworkProject,omitempty" tf:"subnetwork_project"`
}

type InstanceSpecReservationAffinitySpecificReservation struct {
	// Corresponds to the label key of a reservation resource. To target a SPECIFIC_RESERVATION by name, specify compute.googleapis.com/reservation-name as the key and specify the name of your reservation as the only value.
	Key *string `json:"key" tf:"key"`
	// Corresponds to the label values of a reservation resource.
	Values []string `json:"values" tf:"values"`
}

type InstanceSpecReservationAffinity struct {
	// Specifies the label selector for the reservation to use.
	// +optional
	SpecificReservation *InstanceSpecReservationAffinitySpecificReservation `json:"specificReservation,omitempty" tf:"specific_reservation"`
	// The type of reservation from which this instance can consume resources.
	Type *string `json:"type" tf:"type"`
}

type InstanceSpecSchedulingNodeAffinities struct {
	Key      *string  `json:"key" tf:"key"`
	Operator *string  `json:"operator" tf:"operator"`
	Values   []string `json:"values" tf:"values"`
}

type InstanceSpecScheduling struct {
	// Specifies if the instance should be restarted if it was terminated by Compute Engine (not a user).
	// +optional
	AutomaticRestart *bool `json:"automaticRestart,omitempty" tf:"automatic_restart"`
	// +optional
	MinNodeCpus *int64 `json:"minNodeCpus,omitempty" tf:"min_node_cpus"`
	// Specifies node affinities or anti-affinities to determine which sole-tenant nodes your instances and managed instance groups will use as host systems.
	// +optional
	NodeAffinities []InstanceSpecSchedulingNodeAffinities `json:"nodeAffinities,omitempty" tf:"node_affinities"`
	// Describes maintenance behavior for the instance. One of MIGRATE or TERMINATE,
	// +optional
	OnHostMaintenance *string `json:"onHostMaintenance,omitempty" tf:"on_host_maintenance"`
	// Whether the instance is preemptible.
	// +optional
	Preemptible *bool `json:"preemptible,omitempty" tf:"preemptible"`
}

type InstanceSpecScratchDisk struct {
	// The disk interface used for attaching this disk. One of SCSI or NVME.
	Interface *string `json:"interface" tf:"interface"`
}

type InstanceSpecServiceAccount struct {
	// The service account e-mail address.
	// +optional
	Email *string `json:"email,omitempty" tf:"email"`
	// A list of service scopes.
	Scopes []string `json:"scopes" tf:"scopes"`
}

type InstanceSpecShieldedInstanceConfig struct {
	// Whether integrity monitoring is enabled for the instance.
	// +optional
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty" tf:"enable_integrity_monitoring"`
	// Whether secure boot is enabled for the instance.
	// +optional
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty" tf:"enable_secure_boot"`
	// Whether the instance uses vTPM.
	// +optional
	EnableVtpm *bool `json:"enableVtpm,omitempty" tf:"enable_vtpm"`
}

type InstanceSpec struct {
	KubeformOutput *InstanceSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource InstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type InstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires stopping the instance without setting this field, the update will fail.
	// +optional
	AllowStoppingForUpdate *bool `json:"allowStoppingForUpdate,omitempty" tf:"allow_stopping_for_update"`
	// List of disks attached to the instance
	// +optional
	AttachedDisk []InstanceSpecAttachedDisk `json:"attachedDisk,omitempty" tf:"attached_disk"`
	// The boot disk for the instance.
	BootDisk *InstanceSpecBootDisk `json:"bootDisk" tf:"boot_disk"`
	// Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
	// +optional
	CanIPForward *bool `json:"canIPForward,omitempty" tf:"can_ip_forward"`
	// The Confidential VM config being used by the instance.  on_host_maintenance has to be set to TERMINATE or this will fail to create.
	// +optional
	ConfidentialInstanceConfig *InstanceSpecConfidentialInstanceConfig `json:"confidentialInstanceConfig,omitempty" tf:"confidential_instance_config"`
	// The CPU platform used by this instance.
	// +optional
	CpuPlatform *string `json:"cpuPlatform,omitempty" tf:"cpu_platform"`
	// Current status of the instance.
	// +optional
	CurrentStatus *string `json:"currentStatus,omitempty" tf:"current_status"`
	// Whether deletion protection is enabled on this instance.
	// +optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection"`
	// A brief description of the resource.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Desired status of the instance. Either "RUNNING" or "TERMINATED".
	// +optional
	DesiredStatus *string `json:"desiredStatus,omitempty" tf:"desired_status"`
	// Whether the instance has virtual displays enabled.
	// +optional
	EnableDisplay *bool `json:"enableDisplay,omitempty" tf:"enable_display"`
	// List of the type and count of accelerator cards attached to the instance.
	// +optional
	GuestAccelerator []InstanceSpecGuestAccelerator `json:"guestAccelerator,omitempty" tf:"guest_accelerator"`
	// A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
	// +optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname"`
	// The server-assigned unique identifier of this instance.
	// +optional
	InstanceID *string `json:"instanceID,omitempty" tf:"instance_id"`
	// The unique fingerprint of the labels.
	// +optional
	LabelFingerprint *string `json:"labelFingerprint,omitempty" tf:"label_fingerprint"`
	// A set of key/value label pairs assigned to the instance.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// The machine type to create.
	MachineType *string `json:"machineType" tf:"machine_type"`
	// Metadata key/value pairs made available within the instance.
	// +optional
	Metadata *map[string]string `json:"metadata,omitempty" tf:"metadata"`
	// The unique fingerprint of the metadata.
	// +optional
	MetadataFingerprint *string `json:"metadataFingerprint,omitempty" tf:"metadata_fingerprint"`
	// Metadata startup scripts made available within the instance.
	// +optional
	MetadataStartupScript *string `json:"metadataStartupScript,omitempty" tf:"metadata_startup_script"`
	// The minimum CPU platform specified for the VM instance.
	// +optional
	MinCPUPlatform *string `json:"minCPUPlatform,omitempty" tf:"min_cpu_platform"`
	// The name of the instance. One of name or self_link must be provided.
	Name *string `json:"name" tf:"name"`
	// The networks attached to the instance.
	NetworkInterface []InstanceSpecNetworkInterface `json:"networkInterface" tf:"network_interface"`
	// The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither self_link nor project are provided, the provider project is used.
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Specifies the reservations that this instance can consume from.
	// +optional
	ReservationAffinity *InstanceSpecReservationAffinity `json:"reservationAffinity,omitempty" tf:"reservation_affinity"`
	// A list of short names or self_links of resource policies to attach to the instance. Modifying this list will cause the instance to recreate. Currently a max of 1 resource policy is supported.
	// +optional
	ResourcePolicies []string `json:"resourcePolicies,omitempty" tf:"resource_policies"`
	// The scheduling strategy being used by the instance.
	// +optional
	Scheduling *InstanceSpecScheduling `json:"scheduling,omitempty" tf:"scheduling"`
	// The scratch disks attached to the instance.
	// +optional
	ScratchDisk []InstanceSpecScratchDisk `json:"scratchDisk,omitempty" tf:"scratch_disk"`
	// The URI of the created resource.
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
	// The service account to attach to the instance.
	// +optional
	ServiceAccount *InstanceSpecServiceAccount `json:"serviceAccount,omitempty" tf:"service_account"`
	// The shielded vm config being used by the instance.
	// +optional
	ShieldedInstanceConfig *InstanceSpecShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty" tf:"shielded_instance_config"`
	// The list of tags attached to the instance.
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// The unique fingerprint of the tags.
	// +optional
	TagsFingerprint *string `json:"tagsFingerprint,omitempty" tf:"tags_fingerprint"`
	// The zone of the instance. If self_link is provided, this value is ignored. If neither self_link nor zone are provided, the provider zone is used.
	// +optional
	Zone *string `json:"zone,omitempty" tf:"zone"`
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
