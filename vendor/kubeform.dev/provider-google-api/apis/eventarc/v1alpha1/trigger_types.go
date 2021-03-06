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

type Trigger struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TriggerSpec   `json:"spec,omitempty"`
	Status            TriggerStatus `json:"status,omitempty"`
}

type TriggerSpecDestinationCloudRunService struct {
	// Optional. The relative path on the Cloud Run service the events should be sent to. The value must conform to the definition of URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute".
	// +optional
	Path *string `json:"path,omitempty" tf:"path"`
	// Required. The region the Cloud Run service is deployed in.
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// Required. The name of the Cloud Run service being addressed. See https://cloud.google.com/run/docs/reference/rest/v1/namespaces.services. Only services located in the same project of the trigger object can be addressed.
	Service *string `json:"service" tf:"service"`
}

type TriggerSpecDestination struct {
	// [WARNING] Configuring a Cloud Function in Trigger is not supported as of today. The Cloud Function resource name. Format: projects/{project}/locations/{location}/functions/{function}
	// +optional
	CloudFunction *string `json:"cloudFunction,omitempty" tf:"cloud_function"`
	// Cloud Run fully-managed service that receives the events. The service should be running in the same project of the trigger.
	// +optional
	CloudRunService *TriggerSpecDestinationCloudRunService `json:"cloudRunService,omitempty" tf:"cloud_run_service"`
}

type TriggerSpecMatchingCriteria struct {
	// Required. The name of a CloudEvents attribute. Currently, only a subset of attributes are supported for filtering. All triggers MUST provide a filter for the 'type' attribute.
	Attribute *string `json:"attribute" tf:"attribute"`
	// Required. The value for the attribute.
	Value *string `json:"value" tf:"value"`
}

type TriggerSpecTransportPubsub struct {
	// Output only. The name of the Pub/Sub subscription created and managed by Eventarc system as a transport for the event delivery. Format: `projects/{PROJECT_ID}/subscriptions/{SUBSCRIPTION_NAME}`.
	// +optional
	Subscription *string `json:"subscription,omitempty" tf:"subscription"`
	// Optional. The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: `projects/{PROJECT_ID}/topics/{TOPIC_NAME}. You may set an existing topic for triggers of the type google.cloud.pubsub.topic.v1.messagePublished` only. The topic you provide here will not be deleted by Eventarc at trigger deletion.
	// +optional
	Topic *string `json:"topic,omitempty" tf:"topic"`
}

type TriggerSpecTransport struct {
	// The Pub/Sub topic and subscription used by Eventarc as delivery intermediary.
	// +optional
	Pubsub *TriggerSpecTransportPubsub `json:"pubsub,omitempty" tf:"pubsub"`
}

type TriggerSpec struct {
	State *TriggerSpecResource `json:"state,omitempty" tf:"-"`

	Resource TriggerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type TriggerSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Output only. The creation time.
	// +optional
	CreateTime *string `json:"createTime,omitempty" tf:"create_time"`
	// Required. Destination specifies where the events should be sent to.
	Destination *TriggerSpecDestination `json:"destination" tf:"destination"`
	// Output only. This checksum is computed by the server based on the value of other fields, and may be sent only on create requests to ensure the client has an up-to-date value before proceeding.
	// +optional
	Etag *string `json:"etag,omitempty" tf:"etag"`
	// Optional. User labels attached to the triggers that can be used to group resources.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// The location for the resource
	Location *string `json:"location" tf:"location"`
	// Required. null The list of filters that applies to event attributes. Only events that match all the provided filters will be sent to the destination.
	MatchingCriteria []TriggerSpecMatchingCriteria `json:"matchingCriteria" tf:"matching_criteria"`
	// Required. The resource name of the trigger. Must be unique within the location on the project.
	Name *string `json:"name" tf:"name"`
	// The project for the resource
	// +optional
	Project *string `json:"project,omitempty" tf:"project"`
	// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have `roles/eventarc.eventReceiver` IAM role.
	// +optional
	ServiceAccount *string `json:"serviceAccount,omitempty" tf:"service_account"`
	// Optional. In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
	// +optional
	Transport *TriggerSpecTransport `json:"transport,omitempty" tf:"transport"`
	// Output only. Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
	// +optional
	Uid *string `json:"uid,omitempty" tf:"uid"`
	// Output only. The last-modified time.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time"`
}

type TriggerStatus struct {
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

// TriggerList is a list of Triggers
type TriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Trigger CRD objects
	Items []Trigger `json:"items,omitempty"`
}
