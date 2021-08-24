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

func (r *Instance) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-notebooks-google-kubeform-com-v1alpha1-instance,mutating=false,failurePolicy=fail,groups=notebooks.google.kubeform.com,resources=instances,versions=v1alpha1,name=instance.notebooks.google.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &Instance{}

var instanceForceNewList = map[string]bool{
	"/accelerator_config/*/core_count": true,
	"/accelerator_config/*/type":       true,
	"/boot_disk_size_gb":               true,
	"/boot_disk_type":                  true,
	"/container_image/*/repository":    true,
	"/container_image/*/tag":           true,
	"/custom_gpu_driver_path":          true,
	"/data_disk_size_gb":               true,
	"/data_disk_type":                  true,
	"/disk_encryption":                 true,
	"/install_gpu_driver":              true,
	"/instance_owners":                 true,
	"/kms_key":                         true,
	"/location":                        true,
	"/machine_type":                    true,
	"/metadata":                        true,
	"/name":                            true,
	"/network":                         true,
	"/no_proxy_access":                 true,
	"/no_public_ip":                    true,
	"/no_remove_data_disk":             true,
	"/post_startup_script":             true,
	"/project":                         true,
	"/service_account":                 true,
	"/service_account_scopes":          true,
	"/shielded_instance_config/*/enable_integrity_monitoring": true,
	"/shielded_instance_config/*/enable_secure_boot":          true,
	"/shielded_instance_config/*/enable_vtpm":                 true,
	"/subnet":                  true,
	"/tags":                    true,
	"/vm_image/*/image_family": true,
	"/vm_image/*/image_name":   true,
	"/vm_image/*/project":      true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Instance) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Instance) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*Instance)
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

	for key := range instanceForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`instance "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Instance) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`instance "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
