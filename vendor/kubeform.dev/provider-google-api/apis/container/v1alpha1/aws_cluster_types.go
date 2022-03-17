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

type AwsCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsClusterSpec   `json:"spec,omitempty"`
	Status            AwsClusterStatus `json:"status,omitempty"`
}

type AwsClusterSpecAuthorizationAdminUsers struct {
	// Required. The name of the user, e.g. `my-gcp-id@gmail.com`.
	Username *string `json:"username" tf:"username"`
}

type AwsClusterSpecAuthorization struct {
	// Required. Users to perform operations as a cluster admin. A managed ClusterRoleBinding will be created to grant the `cluster-admin` ClusterRole to the users. At most one user can be specified. For more info on RBAC, see https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
	AdminUsers []AwsClusterSpecAuthorizationAdminUsers `json:"adminUsers" tf:"admin_users"`
}

type AwsClusterSpecControlPlaneAwsServicesAuthentication struct {
	// Required. The Amazon Resource Name (ARN) of the role that the Anthos Multi-Cloud API will assume when managing AWS resources on your account.
	RoleArn *string `json:"roleArn" tf:"role_arn"`
	// Optional. An identifier for the assumed role session. When unspecified, it defaults to `multicloud-service-agent`.
	// +optional
	RoleSessionName *string `json:"roleSessionName,omitempty" tf:"role_session_name"`
}

type AwsClusterSpecControlPlaneConfigEncryption struct {
	// Required. The ARN of the AWS KMS key used to encrypt cluster configuration.
	KmsKeyArn *string `json:"kmsKeyArn" tf:"kms_key_arn"`
}

type AwsClusterSpecControlPlaneDatabaseEncryption struct {
	// Required. The ARN of the AWS KMS key used to encrypt cluster secrets.
	KmsKeyArn *string `json:"kmsKeyArn" tf:"kms_key_arn"`
}

type AwsClusterSpecControlPlaneMainVolume struct {
	// Optional. The number of I/O operations per second (IOPS) to provision for GP3 volume.
	// +optional
	Iops *int64 `json:"iops,omitempty" tf:"iops"`
	// Optional. The Amazon Resource Name (ARN) of the Customer Managed Key (CMK) used to encrypt AWS EBS volumes. If not specified, the default Amazon managed key associated to the AWS region where this cluster runs will be used.
	// +optional
	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`
	// Optional. The size of the volume, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.
	// +optional
	SizeGib *int64 `json:"sizeGib,omitempty" tf:"size_gib"`
	// Optional. Type of the EBS volume. When unspecified, it defaults to GP2 volume. Possible values: VOLUME_TYPE_UNSPECIFIED, GP2, GP3
	// +optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type"`
}

type AwsClusterSpecControlPlaneProxyConfig struct {
	// The ARN of the AWS Secret Manager secret that contains the HTTP(S) proxy configuration.
	SecretArn *string `json:"secretArn" tf:"secret_arn"`
	// The version string of the AWS Secret Manager secret that contains the HTTP(S) proxy configuration.
	SecretVersion *string `json:"secretVersion" tf:"secret_version"`
}

type AwsClusterSpecControlPlaneRootVolume struct {
	// Optional. The number of I/O operations per second (IOPS) to provision for GP3 volume.
	// +optional
	Iops *int64 `json:"iops,omitempty" tf:"iops"`
	// Optional. The Amazon Resource Name (ARN) of the Customer Managed Key (CMK) used to encrypt AWS EBS volumes. If not specified, the default Amazon managed key associated to the AWS region where this cluster runs will be used.
	// +optional
	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`
	// Optional. The size of the volume, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.
	// +optional
	SizeGib *int64 `json:"sizeGib,omitempty" tf:"size_gib"`
	// Optional. Type of the EBS volume. When unspecified, it defaults to GP2 volume. Possible values: VOLUME_TYPE_UNSPECIFIED, GP2, GP3
	// +optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type"`
}

type AwsClusterSpecControlPlaneSshConfig struct {
	// Required. The name of the EC2 key pair used to login into cluster machines.
	Ec2KeyPair *string `json:"ec2KeyPair" tf:"ec2_key_pair"`
}

type AwsClusterSpecControlPlane struct {
	// Required. Authentication configuration for management of AWS resources.
	AwsServicesAuthentication *AwsClusterSpecControlPlaneAwsServicesAuthentication `json:"awsServicesAuthentication" tf:"aws_services_authentication"`
	// Required. The ARN of the AWS KMS key used to encrypt cluster configuration.
	ConfigEncryption *AwsClusterSpecControlPlaneConfigEncryption `json:"configEncryption" tf:"config_encryption"`
	// Required. The ARN of the AWS KMS key used to encrypt cluster secrets.
	DatabaseEncryption *AwsClusterSpecControlPlaneDatabaseEncryption `json:"databaseEncryption" tf:"database_encryption"`
	// Required. The name of the AWS IAM instance pofile to assign to each control plane replica.
	IamInstanceProfile *string `json:"iamInstanceProfile" tf:"iam_instance_profile"`
	// Optional. The AWS instance type. When unspecified, it defaults to `t3.medium`.
	// +optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`
	// Optional. Configuration related to the main volume provisioned for each control plane replica. The main volume is in charge of storing all of the cluster's etcd state. Volumes will be provisioned in the availability zone associated with the corresponding subnet. When unspecified, it defaults to 8 GiB with the GP2 volume type.
	// +optional
	MainVolume *AwsClusterSpecControlPlaneMainVolume `json:"mainVolume,omitempty" tf:"main_volume"`
	// Proxy configuration for outbound HTTP(S) traffic.
	// +optional
	ProxyConfig *AwsClusterSpecControlPlaneProxyConfig `json:"proxyConfig,omitempty" tf:"proxy_config"`
	// Optional. Configuration related to the root volume provisioned for each control plane replica. Volumes will be provisioned in the availability zone associated with the corresponding subnet. When unspecified, it defaults to 32 GiB with the GP2 volume type.
	// +optional
	RootVolume *AwsClusterSpecControlPlaneRootVolume `json:"rootVolume,omitempty" tf:"root_volume"`
	// Optional. The IDs of additional security groups to add to control plane replicas. The Anthos Multi-Cloud API will automatically create and manage security groups with the minimum rules needed for a functioning cluster.
	// +optional
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids"`
	// Optional. SSH configuration for how to access the underlying control plane machines.
	// +optional
	SshConfig *AwsClusterSpecControlPlaneSshConfig `json:"sshConfig,omitempty" tf:"ssh_config"`
	// Required. The list of subnets where control plane replicas will run. A replica will be provisioned on each subnet and up to three values can be provided. Each subnet must be in a different AWS Availability Zone (AZ).
	SubnetIDS []string `json:"subnetIDS" tf:"subnet_ids"`
	// Optional. A set of AWS resource tags to propagate to all underlying managed AWS resources. Specify at most 50 pairs containing alphanumerics, spaces, and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters. Values can be up to 255 Unicode characters.
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// Required. The Kubernetes version to run on control plane replicas (e.g. `1.19.10-gke.1000`). You can list all supported versions on a given Google Cloud region by calling .
	Version *string `json:"version" tf:"version"`
}

type AwsClusterSpecFleet struct {
	// The name of the managed Hub Membership resource associated to this cluster. Membership names are formatted as projects/<project-number>/locations/global/membership/<cluster-id>.
	// +optional
	Membership *string `json:"membership,omitempty" tf:"membership"`
	// The number of the Fleet host project where this cluster will be registered.
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
}

type AwsClusterSpecNetworking struct {
	// Required. All pods in the cluster are assigned an RFC1918 IPv4 address from these ranges. Only a single range is supported. This field cannot be changed after creation.
	PodAddressCIDRBlocks []string `json:"podAddressCIDRBlocks" tf:"pod_address_cidr_blocks"`
	// Required. All services in the cluster are assigned an RFC1918 IPv4 address from these ranges. Only a single range is supported. This field cannot be changed after creation.
	ServiceAddressCIDRBlocks []string `json:"serviceAddressCIDRBlocks" tf:"service_address_cidr_blocks"`
	// Required. The VPC associated with the cluster. All component clusters (i.e. control plane and node pools) run on a single VPC. This field cannot be changed after creation.
	VpcID *string `json:"vpcID" tf:"vpc_id"`
}

type AwsClusterSpecWorkloadIdentityConfig struct {
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

type AwsClusterSpec struct {
	State *AwsClusterSpecResource `json:"state,omitempty" tf:"-"`

	Resource AwsClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AwsClusterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
	// +optional
	Annotations *map[string]string `json:"annotations,omitempty" tf:"annotations"`
	// Required. Configuration related to the cluster RBAC settings.
	Authorization *AwsClusterSpecAuthorization `json:"authorization" tf:"authorization"`
	// Required. The AWS region where the cluster runs. Each Google Cloud region supports a subset of nearby AWS regions. You can call to list all supported AWS regions within a given Google Cloud region.
	AwsRegion *string `json:"awsRegion" tf:"aws_region"`
	// Required. Configuration related to the cluster control plane.
	ControlPlane *AwsClusterSpecControlPlane `json:"controlPlane" tf:"control_plane"`
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
	Fleet *AwsClusterSpecFleet `json:"fleet" tf:"fleet"`
	// The location for the resource
	Location *string `json:"location" tf:"location"`
	// The name of this resource.
	Name *string `json:"name" tf:"name"`
	// Required. Cluster-wide networking configuration.
	Networking *AwsClusterSpecNetworking `json:"networking" tf:"networking"`
	// The project for the resource
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Output only. If set, there are currently changes in flight to the cluster.
	// +optional
	Reconciling *bool `json:"reconciling,omitempty" tf:"reconciling"`
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
	WorkloadIdentityConfig []AwsClusterSpecWorkloadIdentityConfig `json:"workloadIdentityConfig,omitempty" tf:"workload_identity_config"`
}

type AwsClusterStatus struct {
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

// AwsClusterList is a list of AwsClusters
type AwsClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCluster CRD objects
	Items []AwsCluster `json:"items,omitempty"`
}