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

type ManagerDeployment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagerDeploymentSpec   `json:"spec,omitempty"`
	Status            ManagerDeploymentStatus `json:"status,omitempty"`
}

type ManagerDeploymentSpecLabels struct {
	// Key for label.
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// Value of label.
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type ManagerDeploymentSpecTargetConfig struct {
	// The full YAML contents of your configuration file.
	Content *string `json:"content" tf:"content"`
}

type ManagerDeploymentSpecTargetImports struct {
	// The full contents of the template that you want to import.
	// +optional
	Content *string `json:"content,omitempty" tf:"content"`
	// The name of the template to import, as declared in the YAML
	// configuration.
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type ManagerDeploymentSpecTarget struct {
	// The root configuration file to use for this deployment.
	Config *ManagerDeploymentSpecTargetConfig `json:"config" tf:"config"`
	// Specifies import files for this configuration. This can be
	// used to import templates or other files. For example, you might
	// import a text file in order to use the file in a template.
	// +optional
	Imports []ManagerDeploymentSpecTargetImports `json:"imports,omitempty" tf:"imports"`
}

type ManagerDeploymentSpec struct {
	State *ManagerDeploymentSpecResource `json:"state,omitempty" tf:"-"`

	Resource ManagerDeploymentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ManagerDeploymentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Set the policy to use for creating new resources. Only used on
	// create and update. Valid values are 'CREATE_OR_ACQUIRE' (default) or
	// 'ACQUIRE'. If set to 'ACQUIRE' and resources do not already exist,
	// the deployment will fail. Note that updating this field does not
	// actually affect the deployment, just how it is updated. Default value: "CREATE_OR_ACQUIRE" Possible values: ["ACQUIRE", "CREATE_OR_ACQUIRE"]
	// +optional
	CreatePolicy *string `json:"createPolicy,omitempty" tf:"create_policy"`
	// Set the policy to use for deleting new resources on update/delete.
	// Valid values are 'DELETE' (default) or 'ABANDON'. If 'DELETE',
	// resource is deleted after removal from Deployment Manager. If
	// 'ABANDON', the resource is only removed from Deployment Manager
	// and is not actually deleted. Note that updating this field does not
	// actually change the deployment, just how it is updated. Default value: "DELETE" Possible values: ["ABANDON", "DELETE"]
	// +optional
	DeletePolicy *string `json:"deletePolicy,omitempty" tf:"delete_policy"`
	// Unique identifier for deployment. Output only.
	// +optional
	DeploymentID *string `json:"deploymentID,omitempty" tf:"deployment_id"`
	// Optional user-provided description of deployment.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Key-value pairs to apply to this labels.
	// +optional
	Labels []ManagerDeploymentSpecLabels `json:"labels,omitempty" tf:"labels"`
	// Output only. URL of the manifest representing the last manifest that
	// was successfully deployed.
	// +optional
	Manifest *string `json:"manifest,omitempty" tf:"manifest"`
	// Unique name for the deployment
	Name *string `json:"name" tf:"name"`
	// If set to true, a deployment is created with "shell" resources
	// that are not actually instantiated. This allows you to preview a
	// deployment. It can be updated to false to actually deploy
	// with real resources.
	//  ~>**NOTE:** Deployment Manager does not allow update
	// of a deployment in preview (unless updating to preview=false). Thus,
	// Terraform will force-recreate deployments if either preview is updated
	// to true or if other fields are updated while preview is true.
	// +optional
	Preview *bool `json:"preview,omitempty" tf:"preview"`
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Output only. Server defined URL for the resource.
	// +optional
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link"`
	// Parameters that define your deployment, including the deployment
	// configuration and relevant templates.
	Target *ManagerDeploymentSpecTarget `json:"target" tf:"target"`
}

type ManagerDeploymentStatus struct {
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

// ManagerDeploymentList is a list of ManagerDeployments
type ManagerDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagerDeployment CRD objects
	Items []ManagerDeployment `json:"items,omitempty"`
}
