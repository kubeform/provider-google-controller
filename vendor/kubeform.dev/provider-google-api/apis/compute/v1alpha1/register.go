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
	"kubeform.dev/provider-google-api/apis/compute"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemeGroupVersion = schema.GroupVersion{Group: compute.GroupName, Version: "v1alpha1"}

var (
	// TODO: move SchemeBuilder with zz_generated.deepcopy.go to k8s.io/api.
	// localSchemeBuilder and AddToScheme will stay in k8s.io/kubernetes.
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
	// We only register manually written functions here. The registration of the
	// generated functions takes place in the generated files. The separation
	// makes the code compile even when the generated files are missing.
	localSchemeBuilder.Register(addKnownTypes)
}

// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Address{},
		&AddressList{},
		&AttachedDisk{},
		&AttachedDiskList{},
		&Autoscaler{},
		&AutoscalerList{},
		&BackendBucket{},
		&BackendBucketList{},
		&BackendBucketSignedURLKey{},
		&BackendBucketSignedURLKeyList{},
		&BackendService{},
		&BackendServiceList{},
		&BackendServiceSignedURLKey{},
		&BackendServiceSignedURLKeyList{},
		&Disk{},
		&DiskList{},
		&DiskIamBinding{},
		&DiskIamBindingList{},
		&DiskIamMember{},
		&DiskIamMemberList{},
		&DiskIamPolicy{},
		&DiskIamPolicyList{},
		&DiskResourcePolicyAttachment{},
		&DiskResourcePolicyAttachmentList{},
		&ExternalVPNGateway{},
		&ExternalVPNGatewayList{},
		&Firewall{},
		&FirewallList{},
		&ForwardingRule{},
		&ForwardingRuleList{},
		&GlobalAddress{},
		&GlobalAddressList{},
		&GlobalForwardingRule{},
		&GlobalForwardingRuleList{},
		&GlobalNetworkEndpoint{},
		&GlobalNetworkEndpointList{},
		&GlobalNetworkEndpointGroup{},
		&GlobalNetworkEndpointGroupList{},
		&HaVPNGateway{},
		&HaVPNGatewayList{},
		&HealthCheck{},
		&HealthCheckList{},
		&HttpHealthCheck{},
		&HttpHealthCheckList{},
		&HttpsHealthCheck{},
		&HttpsHealthCheckList{},
		&Image{},
		&ImageList{},
		&ImageIamBinding{},
		&ImageIamBindingList{},
		&ImageIamMember{},
		&ImageIamMemberList{},
		&ImageIamPolicy{},
		&ImageIamPolicyList{},
		&Instance{},
		&InstanceList{},
		&InstanceFromTemplate{},
		&InstanceFromTemplateList{},
		&InstanceGroup{},
		&InstanceGroupList{},
		&InstanceGroupManager{},
		&InstanceGroupManagerList{},
		&InstanceGroupNamedPort{},
		&InstanceGroupNamedPortList{},
		&InstanceIamBinding{},
		&InstanceIamBindingList{},
		&InstanceIamMember{},
		&InstanceIamMemberList{},
		&InstanceIamPolicy{},
		&InstanceIamPolicyList{},
		&InstanceTemplate{},
		&InstanceTemplateList{},
		&InterconnectAttachment{},
		&InterconnectAttachmentList{},
		&ManagedSslCertificate{},
		&ManagedSslCertificateList{},
		&Network{},
		&NetworkList{},
		&NetworkEndpoint{},
		&NetworkEndpointList{},
		&NetworkEndpointGroup{},
		&NetworkEndpointGroupList{},
		&NetworkPeering{},
		&NetworkPeeringList{},
		&NetworkPeeringRoutesConfig{},
		&NetworkPeeringRoutesConfigList{},
		&NodeGroup{},
		&NodeGroupList{},
		&NodeTemplate{},
		&NodeTemplateList{},
		&PacketMirroring{},
		&PacketMirroringList{},
		&PerInstanceConfig{},
		&PerInstanceConfigList{},
		&ProjectDefaultNetworkTier{},
		&ProjectDefaultNetworkTierList{},
		&ProjectMetadata{},
		&ProjectMetadataList{},
		&ProjectMetadataItem{},
		&ProjectMetadataItemList{},
		&RegionAutoscaler{},
		&RegionAutoscalerList{},
		&RegionBackendService{},
		&RegionBackendServiceList{},
		&RegionDisk{},
		&RegionDiskList{},
		&RegionDiskIamBinding{},
		&RegionDiskIamBindingList{},
		&RegionDiskIamMember{},
		&RegionDiskIamMemberList{},
		&RegionDiskIamPolicy{},
		&RegionDiskIamPolicyList{},
		&RegionDiskResourcePolicyAttachment{},
		&RegionDiskResourcePolicyAttachmentList{},
		&RegionHealthCheck{},
		&RegionHealthCheckList{},
		&RegionInstanceGroupManager{},
		&RegionInstanceGroupManagerList{},
		&RegionNetworkEndpointGroup{},
		&RegionNetworkEndpointGroupList{},
		&RegionPerInstanceConfig{},
		&RegionPerInstanceConfigList{},
		&RegionSslCertificate{},
		&RegionSslCertificateList{},
		&RegionTargetHTTPProxy{},
		&RegionTargetHTTPProxyList{},
		&RegionTargetHTTPSProxy{},
		&RegionTargetHTTPSProxyList{},
		&RegionURLMap{},
		&RegionURLMapList{},
		&Reservation{},
		&ReservationList{},
		&ResourcePolicy{},
		&ResourcePolicyList{},
		&Route{},
		&RouteList{},
		&Router{},
		&RouterList{},
		&RouterInterface{},
		&RouterInterfaceList{},
		&RouterNAT{},
		&RouterNATList{},
		&RouterPeer{},
		&RouterPeerList{},
		&SecurityPolicy{},
		&SecurityPolicyList{},
		&SharedVpcHostProject{},
		&SharedVpcHostProjectList{},
		&SharedVpcServiceProject{},
		&SharedVpcServiceProjectList{},
		&Snapshot{},
		&SnapshotList{},
		&SslCertificate{},
		&SslCertificateList{},
		&SslPolicy{},
		&SslPolicyList{},
		&Subnetwork{},
		&SubnetworkList{},
		&SubnetworkIamBinding{},
		&SubnetworkIamBindingList{},
		&SubnetworkIamMember{},
		&SubnetworkIamMemberList{},
		&SubnetworkIamPolicy{},
		&SubnetworkIamPolicyList{},
		&TargetGrpcProxy{},
		&TargetGrpcProxyList{},
		&TargetHTTPProxy{},
		&TargetHTTPProxyList{},
		&TargetHTTPSProxy{},
		&TargetHTTPSProxyList{},
		&TargetInstance{},
		&TargetInstanceList{},
		&TargetPool{},
		&TargetPoolList{},
		&TargetSslProxy{},
		&TargetSslProxyList{},
		&TargetTcpProxy{},
		&TargetTcpProxyList{},
		&UrlMap{},
		&UrlMapList{},
		&VpnGateway{},
		&VpnGatewayList{},
		&VpnTunnel{},
		&VpnTunnelList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
