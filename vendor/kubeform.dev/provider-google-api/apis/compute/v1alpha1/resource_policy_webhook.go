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

func (r *ResourcePolicy) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-compute-google-kubeform-com-v1alpha1-resourcepolicy,mutating=false,failurePolicy=fail,groups=compute.google.kubeform.com,resources=resourcepolicies,versions=v1alpha1,name=resourcepolicy.compute.google.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &ResourcePolicy{}

var resourcepolicyForceNewList = map[string]bool{
	"/description": true,
	"/group_placement_policy/*/availability_domain_count":      true,
	"/group_placement_policy/*/collocation":                    true,
	"/group_placement_policy/*/vm_count":                       true,
	"/instance_schedule_policy/*/expiration_time":              true,
	"/instance_schedule_policy/*/start_time":                   true,
	"/instance_schedule_policy/*/time_zone":                    true,
	"/instance_schedule_policy/*/vm_start_schedule/*/schedule": true,
	"/instance_schedule_policy/*/vm_stop_schedule/*/schedule":  true,
	"/name":    true,
	"/project": true,
	"/region":  true,
	"/snapshot_schedule_policy/*/retention_policy/*/max_retention_days":                  true,
	"/snapshot_schedule_policy/*/retention_policy/*/on_source_disk_delete":               true,
	"/snapshot_schedule_policy/*/schedule/*/daily_schedule/*/days_in_cycle":              true,
	"/snapshot_schedule_policy/*/schedule/*/daily_schedule/*/start_time":                 true,
	"/snapshot_schedule_policy/*/schedule/*/hourly_schedule/*/hours_in_cycle":            true,
	"/snapshot_schedule_policy/*/schedule/*/hourly_schedule/*/start_time":                true,
	"/snapshot_schedule_policy/*/schedule/*/weekly_schedule/*/day_of_weeks/*/day":        true,
	"/snapshot_schedule_policy/*/schedule/*/weekly_schedule/*/day_of_weeks/*/start_time": true,
	"/snapshot_schedule_policy/*/snapshot_properties/*/guest_flush":                      true,
	"/snapshot_schedule_policy/*/snapshot_properties/*/labels":                           true,
	"/snapshot_schedule_policy/*/snapshot_properties/*/storage_locations":                true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *ResourcePolicy) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ResourcePolicy) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*ResourcePolicy)
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

	for key := range resourcepolicyForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`resourcepolicy "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ResourcePolicy) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`resourcepolicy "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
