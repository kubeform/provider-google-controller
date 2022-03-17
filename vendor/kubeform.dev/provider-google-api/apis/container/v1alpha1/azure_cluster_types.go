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

type AzureCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzureClusterSpec   `json:"spec,omitempty"`
	Status            AzureClusterStatus `json:"status,omitempty"`
}

type AzureClusterSpecAuthorizationAdminUsers struct {
	// Required. The name of the user, e.g. `my-gcp-id@gmail.com`.
	Username *string `json:"username" tf:"username"`
}

type AzureClusterSpecAuthorization struct {
	// Required. Users that can perform operations as a cluster admin. A new ClusterRoleBinding will be created to grant the cluster-admin ClusterRole to the users. At most one user can be specified. For more info on RBAC, see https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
	AdminUsers []AzureClusterSpecAuthorizationAdminUsers `json:"adminUsers" tf:"admin_users"`
}

type AzureClusterSpecControlPlaneDatabaseEncryption struct {
	// The ARM ID of the Azure Key Vault key to encrypt / decrypt data. For example: `/subscriptions/<subscription-id>/resourceGroups/<resource-group-id>/providers/Microsoft.KeyVault/vaults/<key-vault-id>/keys/<key-name>` Encryption will always take the latest version of the key and hence specific version is not supported.
	KeyID *string `json:"keyID" tf:"key_id"`
}

type AzureClusterSpecControlPlaneMainVolume struct {
	// Optional. The size of the disk, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.
	// +optional
	SizeGib *int64 `json:"sizeGib,omitempty" tf:"size_gib"`
}

type AzureClusterSpecControlPlaneProxyConfig struct {
	// The ARM ID the of the resource group containing proxy keyvault. Resource group ids are formatted as `/subscriptions/<subscription-id>/resourceGroups/<resource-group-name>`
	ResourceGroupID *string `json:"resourceGroupID" tf:"resource_group_id"`
	// The URL the of the proxy setting secret with its version. Secret ids are formatted as `https:<key-vault-name>.vault.azure.net/secrets/<secret-name>/<secret-version>`.
	SecretID *string `json:"secretID" tf:"secret_id"`
}

type AzureClusterSpecControlPlaneReplicaPlacements struct {
	// For a given replica, the Azure availability zone where to provision the control plane VM and the ETCD disk.
	AzureAvailabilityZone *string `json:"azureAvailabilityZone" tf:"azure_availability_zone"`
	// For a given replica, the ARM ID of the subnet where the control plane VM is deployed. Make sure it's a subnet under the virtual network in the cluster configuration.
	SubnetID *string `json:"subnetID" tf:"subnet_id"`
}

type AzureClusterSpecControlPlaneRootVolume struct {
	// Optional. The size of the disk, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.
	// +optional
	SizeGib *int64 `json:"sizeGib,omitempty" tf:"size_gib"`
}

type AzureClusterSpecControlPlaneSshConfig struct {
	// Required. The SSH public key data for VMs managed by Anthos. This accepts the authorized_keys file format used in OpenSSH according to the sshd(8) manual page.
	AuthorizedKey *string `json:"authorizedKey" tf:"authorized_key"`
}

type AzureClusterSpecControlPlane struct {
	// Optional. Configuration related to application-layer secrets encryption.
	// +optional
	DatabaseEncryption *AzureClusterSpecControlPlaneDatabaseEncryption `json:"databaseEncryption,omitempty" tf:"database_encryption"`
	// Optional. Configuration related to the main volume provisioned for each control plane replica. The main volume is in charge of storing all of the cluster's etcd state. When unspecified, it defaults to a 8-GiB Azure Disk.
	// +optional
	MainVolume *AzureClusterSpecControlPlaneMainVolume `json:"mainVolume,omitempty" tf:"main_volume"`
	// Proxy configuration for outbound HTTP(S) traffic.
	// +optional
	ProxyConfig *AzureClusterSpecControlPlaneProxyConfig `json:"proxyConfig,omitempty" tf:"proxy_config"`
	// Configuration for where to place the control plane replicas. Up to three replica placement instances can be specified. If replica_placements is set, the replica placement instances will be applied to the three control plane replicas as evenly as possible.
	// +optional
	ReplicaPlacements []AzureClusterSpecControlPlaneReplicaPlacements `json:"replicaPlacements,omitempty" tf:"replica_placements"`
	// Optional. Configuration related to the root volume provisioned for each control plane replica. When unspecified, it defaults to 32-GiB Azure Disk.
	// +optional
	RootVolume *AzureClusterSpecControlPlaneRootVolume `json:"rootVolume,omitempty" tf:"root_volume"`
	// Required. SSH configuration for how to access the underlying control plane machines.
	SshConfig *AzureClusterSpecControlPlaneSshConfig `json:"sshConfig" tf:"ssh_config"`
	// Required. The ARM ID of the subnet where the control plane VMs are deployed. Example: `/subscriptions//resourceGroups//providers/Microsoft.Network/virtualNetworks//subnets/default`.
	SubnetID *string `json:"subnetID" tf:"subnet_id"`
	// Optional. A set of tags to apply to all underlying control plane Azure resources.
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// Required. The Kubernetes version to run on control plane replicas (e.g. `1.19.10-gke.1000`). You can list all supported versions on a given Google Cloud region by calling GetAzureServerConfig.
	Version *string `json:"version" tf:"version"`
	// Optional. The Azure VM size name. Example: `Standard_DS2_v2`. For available VM sizes, see https://docs.microsoft.com/en-us/azure/virtual-machines/vm-naming-conventions. When unspecified, it defaults to `Standard_DS2_v2`.
	// +optional
	VmSize *string `json:"vmSize,omitempty" tf:"vm_size"`
}

type AzureClusterSpecFleet struct {
	// The name of the managed Hub Membership resource associated to this cluster. Membership names are formatted as projects/<project-number>/locations/global/membership/<cluster-id>.
	// +optional
	Membership *string `json:"membership,omitempty" tf:"membership"`
	// The number of the Fleet host project where this cluster will be registered.
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
}

type AzureClusterSpecNetworking struct {
	// Required. The IP address range of the pods in this cluster, in CIDR notation (e.g. `10.96.0.0/14`). All pods in the cluster get assigned a unique RFC1918 IPv4 address from these ranges. Only a single range is supported. This field cannot be changed after creation.
	PodAddressCIDRBlocks []string `json:"podAddressCIDRBlocks" tf:"pod_address_cidr_blocks"`
	// Required. The IP address range for services in this cluster, in CIDR notation (e.g. `10.96.0.0/14`). All services in the cluster get assigned a unique RFC1918 IPv4 address from these ranges. Only a single range is supported. This field cannot be changed after creating a cluster.
	ServiceAddressCIDRBlocks []string `json:"serviceAddressCIDRBlocks" tf:"service_address_cidr_blocks"`
	// Required. The Azure Resource Manager (ARM) ID of the VNet associated with your cluster. All components in the cluster (i.e. control plane and node pools) run on a single VNet. Example: `/subscriptions/*/resourceGroups/*/providers/Microsoft.Network/virtualNetworks/*` This field cannot be changed after creation.
	VirtualNetworkID *string `json:"virtualNetworkID" tf:"virtual_network_id"`
}

type AzureClusterSpecWorkloadIdentityConfig struct {
	// The ID of the OIDC Identity Provider (IdP) associated to the Workload Identity Pool.
	// +optional
	IdentityProvider *string `json:"identityProvider,omitempty" tf:"identity_provider"`
	// The OIDC issuer URL for this cluster.
	// +optional
	IssuerURI *string `json:"issuerURI,omitempty" tf:"issuer_uri"`
	// The Workload Identity Pool associated to the cluster.
	// +optional
	WorkloadPool *string `json:"workloadPool,omitempty" tf:"workload_pool"`
}

type AzureClusterSpec struct {
	State *AzureClusterSpecResource `json:"state,omitempty" tf:"-"`

	Resource AzureClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AzureClusterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
	// +optional
	Annotations *map[string]string `json:"annotations,omitempty" tf:"annotations"`
	// Required. Configuration related to the cluster RBAC settings.
	Authorization *AzureClusterSpecAuthorization `json:"authorization" tf:"authorization"`
	// Required. The Azure region where the cluster runs. Each Google Cloud region supports a subset of nearby Azure regions. You can call to list all supported Azure regions within a given Google Cloud region.
	AzureRegion *string `json:"azureRegion" tf:"azure_region"`
	// Required. Name of the AzureClient. The `AzureClient` resource must reside on the same GCP project and region as the `AzureCluster`. `AzureClient` names are formatted as `projects/<project-number>/locations/<region>/azureClients/<client-id>`. See Resource Names (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.
	Client *string `json:"client" tf:"client"`
	// Required. Configuration related to the cluster control plane.
	ControlPlane *AzureClusterSpecControlPlane `json:"controlPlane" tf:"control_plane"`
	// Output only. The time at which this cluster was created.
	// +optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time"`
	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Output only. The endpoint of the cluster's API server.
	// +optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint"`
	// Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	// +optional
	Etag *string `json:"etag,omitempty" tf:"etag"`
	// Fleet configuration.
	Fleet *AzureClusterSpecFleet `json:"fleet" tf:"fleet"`
	// The location for the resource
	Location *string `json:"location" tf:"location"`
	// The name of this resource.
	Name *string `json:"name" tf:"name"`
	// Required. Cluster-wide networking configuration.
	Networking *AzureClusterSpecNetworking `json:"networking" tf:"networking"`
	// The project for the resource
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Output only. If set, there are currently changes in flight to the cluster.
	// +optional
	Reconciling *bool `json:"reconciling,omitempty" tf:"reconciling"`
	// Required. The ARM ID of the resource group where the cluster resources are deployed. For example: `/subscriptions/*/resourceGroups/*`
	ResourceGroupID *string `json:"resourceGroupID" tf:"resource_group_id"`
	// Output only. The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR, DEGRADED
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// Output only. A globally unique identifier for the cluster.
	// +optional
	Uid *string `json:"uid,omitempty" tf:"uid"`
	// Output only. The time at which this cluster was last updated.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time"`
	// Output only. Workload Identity settings.
	// +optional
	WorkloadIdentityConfig []AzureClusterSpecWorkloadIdentityConfig `json:"workloadIdentityConfig,omitempty" tf:"workload_identity_config"`
}

type AzureClusterStatus struct {
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

// AzureClusterList is a list of AzureClusters
type AzureClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzureCluster CRD objects
	Items []AzureCluster `json:"items,omitempty"`
}