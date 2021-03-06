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
	"fmt"
	"strings"

	base "kubeform.dev/apimachinery/api/v1alpha1"
	"kubeform.dev/apimachinery/pkg/util"

	jsoniter "github.com/json-iterator/go"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *Cluster) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-container-google-kubeform-com-v1alpha1-cluster,mutating=false,failurePolicy=fail,groups=container.google.kubeform.com,resources=clusters,versions=v1alpha1,name=cluster.container.google.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &Cluster{}

var clusterForceNewList = map[string]bool{
	"/authenticator_groups_config/*/security_group":         true,
	"/cluster_ipv4_cidr":                                    true,
	"/confidential_nodes/*/enabled":                         true,
	"/default_max_pods_per_node":                            true,
	"/description":                                          true,
	"/enable_autopilot":                                     true,
	"/enable_kubernetes_alpha":                              true,
	"/enable_tpu":                                           true,
	"/initial_node_count":                                   true,
	"/ip_allocation_policy/*/cluster_ipv4_cidr_block":       true,
	"/ip_allocation_policy/*/cluster_secondary_range_name":  true,
	"/ip_allocation_policy/*/services_ipv4_cidr_block":      true,
	"/ip_allocation_policy/*/services_secondary_range_name": true,
	"/location": true,
	"/master_auth/*/client_certificate_config/*/issue_client_certificate": true,
	"/name":                                    true,
	"/network":                                 true,
	"/networking_mode":                         true,
	"/node_config/*/boot_disk_kms_key":         true,
	"/node_config/*/disk_size_gb":              true,
	"/node_config/*/disk_type":                 true,
	"/node_config/*/gcfs_config/*/enabled":     true,
	"/node_config/*/guest_accelerator/*/count": true,
	"/node_config/*/guest_accelerator/*/gpu_partition_size":                             true,
	"/node_config/*/guest_accelerator/*/type":                                           true,
	"/node_config/*/labels":                                                             true,
	"/node_config/*/local_ssd_count":                                                    true,
	"/node_config/*/machine_type":                                                       true,
	"/node_config/*/metadata":                                                           true,
	"/node_config/*/min_cpu_platform":                                                   true,
	"/node_config/*/node_group":                                                         true,
	"/node_config/*/oauth_scopes":                                                       true,
	"/node_config/*/preemptible":                                                        true,
	"/node_config/*/service_account":                                                    true,
	"/node_config/*/shielded_instance_config/*/enable_integrity_monitoring":             true,
	"/node_config/*/shielded_instance_config/*/enable_secure_boot":                      true,
	"/node_config/*/tags":                                                               true,
	"/node_config/*/taint/*/effect":                                                     true,
	"/node_config/*/taint/*/key":                                                        true,
	"/node_config/*/taint/*/value":                                                      true,
	"/node_pool/*/initial_node_count":                                                   true,
	"/node_pool/*/max_pods_per_node":                                                    true,
	"/node_pool/*/name":                                                                 true,
	"/node_pool/*/name_prefix":                                                          true,
	"/node_pool/*/node_config/*/boot_disk_kms_key":                                      true,
	"/node_pool/*/node_config/*/disk_size_gb":                                           true,
	"/node_pool/*/node_config/*/disk_type":                                              true,
	"/node_pool/*/node_config/*/gcfs_config/*/enabled":                                  true,
	"/node_pool/*/node_config/*/guest_accelerator/*/count":                              true,
	"/node_pool/*/node_config/*/guest_accelerator/*/gpu_partition_size":                 true,
	"/node_pool/*/node_config/*/guest_accelerator/*/type":                               true,
	"/node_pool/*/node_config/*/labels":                                                 true,
	"/node_pool/*/node_config/*/local_ssd_count":                                        true,
	"/node_pool/*/node_config/*/machine_type":                                           true,
	"/node_pool/*/node_config/*/metadata":                                               true,
	"/node_pool/*/node_config/*/min_cpu_platform":                                       true,
	"/node_pool/*/node_config/*/node_group":                                             true,
	"/node_pool/*/node_config/*/oauth_scopes":                                           true,
	"/node_pool/*/node_config/*/preemptible":                                            true,
	"/node_pool/*/node_config/*/service_account":                                        true,
	"/node_pool/*/node_config/*/shielded_instance_config/*/enable_integrity_monitoring": true,
	"/node_pool/*/node_config/*/shielded_instance_config/*/enable_secure_boot":          true,
	"/node_pool/*/node_config/*/tags":                                                   true,
	"/node_pool/*/node_config/*/taint/*/effect":                                         true,
	"/node_pool/*/node_config/*/taint/*/key":                                            true,
	"/node_pool/*/node_config/*/taint/*/value":                                          true,
	"/private_cluster_config/*/enable_private_endpoint":                                 true,
	"/private_cluster_config/*/enable_private_nodes":                                    true,
	"/private_cluster_config/*/master_ipv4_cidr_block":                                  true,
	"/project":    true,
	"/subnetwork": true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Cluster) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Cluster) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*Cluster)
	oldObj := res.Spec.Resource

	jsnitr := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		TagKey:                 "tf",
		ValidateJsonRawMessage: true,
		TypeEncoders:           GetEncoder(),
		TypeDecoders:           GetDecoder(),
	}.Froze()

	byteNew, err := jsnitr.Marshal(newObj)
	if err != nil {
		return err
	}
	tempNew := make(map[string]interface{})
	err = jsnitr.Unmarshal(byteNew, &tempNew)
	if err != nil {
		return err
	}

	byteOld, err := jsnitr.Marshal(oldObj)
	if err != nil {
		return err
	}
	tempOld := make(map[string]interface{})
	err = jsnitr.Unmarshal(byteOld, &tempOld)
	if err != nil {
		return err
	}

	for key, _ := range clusterForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`cluster "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Cluster) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`cluster "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
