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

type ServicesGameServerCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicesGameServerClusterSpec   `json:"spec,omitempty"`
	Status            ServicesGameServerClusterStatus `json:"status,omitempty"`
}

type ServicesGameServerClusterSpecConnectionInfoGkeClusterReference struct {
	// The full or partial name of a GKE cluster, using one of the following
	// forms:
	//
	// * 'projects/{project_id}/locations/{location}/clusters/{cluster_id}'
	// * 'locations/{location}/clusters/{cluster_id}'
	// * '{cluster_id}'
	//
	// If project and location are not specified, the project and location of the
	// GameServerCluster resource are used to generate the full name of the
	// GKE cluster.
	Cluster *string `json:"cluster" tf:"cluster"`
}

type ServicesGameServerClusterSpecConnectionInfo struct {
	// Reference of the GKE cluster where the game servers are installed.
	GkeClusterReference *ServicesGameServerClusterSpecConnectionInfoGkeClusterReference `json:"gkeClusterReference" tf:"gke_cluster_reference"`
	// Namespace designated on the game server cluster where the game server
	// instances will be created. The namespace existence will be validated
	// during creation.
	Namespace *string `json:"namespace" tf:"namespace"`
}

type ServicesGameServerClusterSpec struct {
	KubeformOutput *ServicesGameServerClusterSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource ServicesGameServerClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ServicesGameServerClusterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Required. The resource name of the game server cluster
	ClusterID *string `json:"clusterID" tf:"cluster_id"`
	// Game server cluster connection information. This information is used to
	// manage game server clusters.
	ConnectionInfo *ServicesGameServerClusterSpecConnectionInfo `json:"connectionInfo" tf:"connection_info"`
	// Human readable description of the cluster.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// The labels associated with this game server cluster. Each label is a
	// key-value pair.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// Location of the Cluster.
	// +optional
	Location *string `json:"location,omitempty" tf:"location"`
	// The resource id of the game server cluster, eg:
	//
	// 'projects/{project_id}/locations/{location}/realms/{realm_id}/gameServerClusters/{cluster_id}'.
	// For example,
	//
	// 'projects/my-project/locations/{location}/realms/zanzibar/gameServerClusters/my-onprem-cluster'.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// The realm id of the game server realm.
	RealmID *string `json:"realmID" tf:"realm_id"`
}

type ServicesGameServerClusterStatus struct {
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

// ServicesGameServerClusterList is a list of ServicesGameServerClusters
type ServicesGameServerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServicesGameServerCluster CRD objects
	Items []ServicesGameServerCluster `json:"items,omitempty"`
}
