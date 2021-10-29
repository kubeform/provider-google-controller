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

func (r *WorkflowTemplate) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-dataproc-google-kubeform-com-v1alpha1-workflowtemplate,mutating=false,failurePolicy=fail,groups=dataproc.google.kubeform.com,resources=workflowtemplates,versions=v1alpha1,name=workflowtemplate.dataproc.google.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &WorkflowTemplate{}

var workflowtemplateForceNewList = map[string]bool{
	"/dag_timeout":                                                                      true,
	"/jobs/*/hadoop_job/*/archive_uris":                                                 true,
	"/jobs/*/hadoop_job/*/args":                                                         true,
	"/jobs/*/hadoop_job/*/file_uris":                                                    true,
	"/jobs/*/hadoop_job/*/jar_file_uris":                                                true,
	"/jobs/*/hadoop_job/*/logging_config/*/driver_log_levels":                           true,
	"/jobs/*/hadoop_job/*/main_class":                                                   true,
	"/jobs/*/hadoop_job/*/main_jar_file_uri":                                            true,
	"/jobs/*/hadoop_job/*/properties":                                                   true,
	"/jobs/*/hive_job/*/continue_on_failure":                                            true,
	"/jobs/*/hive_job/*/jar_file_uris":                                                  true,
	"/jobs/*/hive_job/*/properties":                                                     true,
	"/jobs/*/hive_job/*/query_file_uri":                                                 true,
	"/jobs/*/hive_job/*/query_list/*/queries":                                           true,
	"/jobs/*/hive_job/*/script_variables":                                               true,
	"/jobs/*/labels":                                                                    true,
	"/jobs/*/pig_job/*/continue_on_failure":                                             true,
	"/jobs/*/pig_job/*/jar_file_uris":                                                   true,
	"/jobs/*/pig_job/*/logging_config/*/driver_log_levels":                              true,
	"/jobs/*/pig_job/*/properties":                                                      true,
	"/jobs/*/pig_job/*/query_file_uri":                                                  true,
	"/jobs/*/pig_job/*/query_list/*/queries":                                            true,
	"/jobs/*/pig_job/*/script_variables":                                                true,
	"/jobs/*/prerequisite_step_ids":                                                     true,
	"/jobs/*/presto_job/*/client_tags":                                                  true,
	"/jobs/*/presto_job/*/continue_on_failure":                                          true,
	"/jobs/*/presto_job/*/logging_config/*/driver_log_levels":                           true,
	"/jobs/*/presto_job/*/output_format":                                                true,
	"/jobs/*/presto_job/*/properties":                                                   true,
	"/jobs/*/presto_job/*/query_file_uri":                                               true,
	"/jobs/*/presto_job/*/query_list/*/queries":                                         true,
	"/jobs/*/pyspark_job/*/archive_uris":                                                true,
	"/jobs/*/pyspark_job/*/args":                                                        true,
	"/jobs/*/pyspark_job/*/file_uris":                                                   true,
	"/jobs/*/pyspark_job/*/jar_file_uris":                                               true,
	"/jobs/*/pyspark_job/*/logging_config/*/driver_log_levels":                          true,
	"/jobs/*/pyspark_job/*/main_python_file_uri":                                        true,
	"/jobs/*/pyspark_job/*/properties":                                                  true,
	"/jobs/*/pyspark_job/*/python_file_uris":                                            true,
	"/jobs/*/scheduling/*/max_failures_per_hour":                                        true,
	"/jobs/*/scheduling/*/max_failures_total":                                           true,
	"/jobs/*/spark_job/*/archive_uris":                                                  true,
	"/jobs/*/spark_job/*/args":                                                          true,
	"/jobs/*/spark_job/*/file_uris":                                                     true,
	"/jobs/*/spark_job/*/jar_file_uris":                                                 true,
	"/jobs/*/spark_job/*/logging_config/*/driver_log_levels":                            true,
	"/jobs/*/spark_job/*/main_class":                                                    true,
	"/jobs/*/spark_job/*/main_jar_file_uri":                                             true,
	"/jobs/*/spark_job/*/properties":                                                    true,
	"/jobs/*/spark_r_job/*/archive_uris":                                                true,
	"/jobs/*/spark_r_job/*/args":                                                        true,
	"/jobs/*/spark_r_job/*/file_uris":                                                   true,
	"/jobs/*/spark_r_job/*/logging_config/*/driver_log_levels":                          true,
	"/jobs/*/spark_r_job/*/main_r_file_uri":                                             true,
	"/jobs/*/spark_r_job/*/properties":                                                  true,
	"/jobs/*/spark_sql_job/*/jar_file_uris":                                             true,
	"/jobs/*/spark_sql_job/*/logging_config/*/driver_log_levels":                        true,
	"/jobs/*/spark_sql_job/*/properties":                                                true,
	"/jobs/*/spark_sql_job/*/query_file_uri":                                            true,
	"/jobs/*/spark_sql_job/*/query_list/*/queries":                                      true,
	"/jobs/*/spark_sql_job/*/script_variables":                                          true,
	"/jobs/*/step_id":                                                                   true,
	"/labels":                                                                           true,
	"/location":                                                                         true,
	"/name":                                                                             true,
	"/parameters/*/description":                                                         true,
	"/parameters/*/fields":                                                              true,
	"/parameters/*/name":                                                                true,
	"/parameters/*/validation/*/regex/*/regexes":                                        true,
	"/parameters/*/validation/*/values/*/values":                                        true,
	"/placement/*/cluster_selector/*/cluster_labels":                                    true,
	"/placement/*/cluster_selector/*/zone":                                              true,
	"/placement/*/managed_cluster/*/cluster_name":                                       true,
	"/placement/*/managed_cluster/*/config/*/autoscaling_config/*/policy":               true,
	"/placement/*/managed_cluster/*/config/*/encryption_config/*/gce_pd_kms_key_name":   true,
	"/placement/*/managed_cluster/*/config/*/endpoint_config/*/enable_http_port_access": true,
	"/placement/*/managed_cluster/*/config/*/gce_cluster_config/*/internal_ip_only":     true,
	"/placement/*/managed_cluster/*/config/*/gce_cluster_config/*/metadata":             true,
	"/placement/*/managed_cluster/*/config/*/gce_cluster_config/*/network":              true,
	"/placement/*/managed_cluster/*/config/*/gce_cluster_config/*/node_group_affinity/*/node_group":                 true,
	"/placement/*/managed_cluster/*/config/*/gce_cluster_config/*/private_ipv6_google_access":                       true,
	"/placement/*/managed_cluster/*/config/*/gce_cluster_config/*/reservation_affinity/*/consume_reservation_type":  true,
	"/placement/*/managed_cluster/*/config/*/gce_cluster_config/*/reservation_affinity/*/key":                       true,
	"/placement/*/managed_cluster/*/config/*/gce_cluster_config/*/reservation_affinity/*/values":                    true,
	"/placement/*/managed_cluster/*/config/*/gce_cluster_config/*/service_account":                                  true,
	"/placement/*/managed_cluster/*/config/*/gce_cluster_config/*/service_account_scopes":                           true,
	"/placement/*/managed_cluster/*/config/*/gce_cluster_config/*/subnetwork":                                       true,
	"/placement/*/managed_cluster/*/config/*/gce_cluster_config/*/tags":                                             true,
	"/placement/*/managed_cluster/*/config/*/gce_cluster_config/*/zone":                                             true,
	"/placement/*/managed_cluster/*/config/*/initialization_actions/*/executable_file":                              true,
	"/placement/*/managed_cluster/*/config/*/initialization_actions/*/execution_timeout":                            true,
	"/placement/*/managed_cluster/*/config/*/lifecycle_config/*/auto_delete_time":                                   true,
	"/placement/*/managed_cluster/*/config/*/lifecycle_config/*/auto_delete_ttl":                                    true,
	"/placement/*/managed_cluster/*/config/*/lifecycle_config/*/idle_delete_ttl":                                    true,
	"/placement/*/managed_cluster/*/config/*/master_config/*/accelerators/*/accelerator_count":                      true,
	"/placement/*/managed_cluster/*/config/*/master_config/*/accelerators/*/accelerator_type":                       true,
	"/placement/*/managed_cluster/*/config/*/master_config/*/disk_config/*/boot_disk_size_gb":                       true,
	"/placement/*/managed_cluster/*/config/*/master_config/*/disk_config/*/boot_disk_type":                          true,
	"/placement/*/managed_cluster/*/config/*/master_config/*/disk_config/*/num_local_ssds":                          true,
	"/placement/*/managed_cluster/*/config/*/master_config/*/image":                                                 true,
	"/placement/*/managed_cluster/*/config/*/master_config/*/machine_type":                                          true,
	"/placement/*/managed_cluster/*/config/*/master_config/*/min_cpu_platform":                                      true,
	"/placement/*/managed_cluster/*/config/*/master_config/*/num_instances":                                         true,
	"/placement/*/managed_cluster/*/config/*/master_config/*/preemptibility":                                        true,
	"/placement/*/managed_cluster/*/config/*/secondary_worker_config/*/accelerators/*/accelerator_count":            true,
	"/placement/*/managed_cluster/*/config/*/secondary_worker_config/*/accelerators/*/accelerator_type":             true,
	"/placement/*/managed_cluster/*/config/*/secondary_worker_config/*/disk_config/*/boot_disk_size_gb":             true,
	"/placement/*/managed_cluster/*/config/*/secondary_worker_config/*/disk_config/*/boot_disk_type":                true,
	"/placement/*/managed_cluster/*/config/*/secondary_worker_config/*/disk_config/*/num_local_ssds":                true,
	"/placement/*/managed_cluster/*/config/*/secondary_worker_config/*/image":                                       true,
	"/placement/*/managed_cluster/*/config/*/secondary_worker_config/*/machine_type":                                true,
	"/placement/*/managed_cluster/*/config/*/secondary_worker_config/*/min_cpu_platform":                            true,
	"/placement/*/managed_cluster/*/config/*/secondary_worker_config/*/num_instances":                               true,
	"/placement/*/managed_cluster/*/config/*/secondary_worker_config/*/preemptibility":                              true,
	"/placement/*/managed_cluster/*/config/*/security_config/*/kerberos_config/*/cross_realm_trust_admin_server":    true,
	"/placement/*/managed_cluster/*/config/*/security_config/*/kerberos_config/*/cross_realm_trust_kdc":             true,
	"/placement/*/managed_cluster/*/config/*/security_config/*/kerberos_config/*/cross_realm_trust_realm":           true,
	"/placement/*/managed_cluster/*/config/*/security_config/*/kerberos_config/*/cross_realm_trust_shared_password": true,
	"/placement/*/managed_cluster/*/config/*/security_config/*/kerberos_config/*/enable_kerberos":                   true,
	"/placement/*/managed_cluster/*/config/*/security_config/*/kerberos_config/*/kdc_db_key":                        true,
	"/placement/*/managed_cluster/*/config/*/security_config/*/kerberos_config/*/key_password":                      true,
	"/placement/*/managed_cluster/*/config/*/security_config/*/kerberos_config/*/keystore":                          true,
	"/placement/*/managed_cluster/*/config/*/security_config/*/kerberos_config/*/keystore_password":                 true,
	"/placement/*/managed_cluster/*/config/*/security_config/*/kerberos_config/*/kms_key":                           true,
	"/placement/*/managed_cluster/*/config/*/security_config/*/kerberos_config/*/realm":                             true,
	"/placement/*/managed_cluster/*/config/*/security_config/*/kerberos_config/*/root_principal_password":           true,
	"/placement/*/managed_cluster/*/config/*/security_config/*/kerberos_config/*/tgt_lifetime_hours":                true,
	"/placement/*/managed_cluster/*/config/*/security_config/*/kerberos_config/*/truststore":                        true,
	"/placement/*/managed_cluster/*/config/*/security_config/*/kerberos_config/*/truststore_password":               true,
	"/placement/*/managed_cluster/*/config/*/software_config/*/image_version":                                       true,
	"/placement/*/managed_cluster/*/config/*/software_config/*/optional_components":                                 true,
	"/placement/*/managed_cluster/*/config/*/software_config/*/properties":                                          true,
	"/placement/*/managed_cluster/*/config/*/staging_bucket":                                                        true,
	"/placement/*/managed_cluster/*/config/*/temp_bucket":                                                           true,
	"/placement/*/managed_cluster/*/config/*/worker_config/*/accelerators/*/accelerator_count":                      true,
	"/placement/*/managed_cluster/*/config/*/worker_config/*/accelerators/*/accelerator_type":                       true,
	"/placement/*/managed_cluster/*/config/*/worker_config/*/disk_config/*/boot_disk_size_gb":                       true,
	"/placement/*/managed_cluster/*/config/*/worker_config/*/disk_config/*/boot_disk_type":                          true,
	"/placement/*/managed_cluster/*/config/*/worker_config/*/disk_config/*/num_local_ssds":                          true,
	"/placement/*/managed_cluster/*/config/*/worker_config/*/image":                                                 true,
	"/placement/*/managed_cluster/*/config/*/worker_config/*/machine_type":                                          true,
	"/placement/*/managed_cluster/*/config/*/worker_config/*/min_cpu_platform":                                      true,
	"/placement/*/managed_cluster/*/config/*/worker_config/*/num_instances":                                         true,
	"/placement/*/managed_cluster/*/config/*/worker_config/*/preemptibility":                                        true,
	"/placement/*/managed_cluster/*/labels":                                                                         true,
	"/project":                                                                                                      true,
	"/version":                                                                                                      true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *WorkflowTemplate) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *WorkflowTemplate) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*WorkflowTemplate)
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

	for key := range workflowtemplateForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`workflowtemplate "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *WorkflowTemplate) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`workflowtemplate "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
