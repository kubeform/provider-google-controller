// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"google.golang.org/api/googleapi"
)

func resourceAccessContextManagerAccessLevelCondition() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccessContextManagerAccessLevelConditionCreate,
		Read:   resourceAccessContextManagerAccessLevelConditionRead,
		Delete: resourceAccessContextManagerAccessLevelConditionDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"access_level": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The name of the Access Level to add this condition to.`,
			},
			"device_policy": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `Device specific restrictions, all restrictions must hold for
the Condition to be true. If not specified, all devices are
allowed.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allowed_device_management_levels": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Description: `A list of allowed device management levels.
An empty list allows all management levels. Possible values: ["MANAGEMENT_UNSPECIFIED", "NONE", "BASIC", "COMPLETE"]`,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{"MANAGEMENT_UNSPECIFIED", "NONE", "BASIC", "COMPLETE"}, false),
							},
						},
						"allowed_encryption_statuses": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Description: `A list of allowed encryptions statuses.
An empty list allows all statuses. Possible values: ["ENCRYPTION_UNSPECIFIED", "ENCRYPTION_UNSUPPORTED", "UNENCRYPTED", "ENCRYPTED"]`,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{"ENCRYPTION_UNSPECIFIED", "ENCRYPTION_UNSUPPORTED", "UNENCRYPTED", "ENCRYPTED"}, false),
							},
						},
						"os_constraints": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Description: `A list of allowed OS versions.
An empty list allows all types and all versions.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"os_type": {
										Type:         schema.TypeString,
										Required:     true,
										ForceNew:     true,
										ValidateFunc: validation.StringInSlice([]string{"OS_UNSPECIFIED", "DESKTOP_MAC", "DESKTOP_WINDOWS", "DESKTOP_LINUX", "DESKTOP_CHROME_OS", "ANDROID", "IOS"}, false),
										Description:  `The operating system type of the device. Possible values: ["OS_UNSPECIFIED", "DESKTOP_MAC", "DESKTOP_WINDOWS", "DESKTOP_LINUX", "DESKTOP_CHROME_OS", "ANDROID", "IOS"]`,
									},
									"minimum_version": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Description: `The minimum allowed OS version. If not set, any version
of this OS satisfies the constraint.
Format: "major.minor.patch" such as "10.5.301", "9.2.1".`,
									},
								},
							},
						},
						"require_admin_approval": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Description: `Whether the device needs to be approved by the customer admin.`,
						},
						"require_corp_owned": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Description: `Whether the device needs to be corp owned.`,
						},
						"require_screen_lock": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
							Description: `Whether or not screenlock is required for the DevicePolicy
to be true. Defaults to false.`,
						},
					},
				},
			},
			"ip_subnetworks": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `A list of CIDR block IP subnetwork specification. May be IPv4
or IPv6.
Note that for a CIDR IP address block, the specified IP address
portion must be properly truncated (i.e. all the host bits must
be zero) or the input is considered malformed. For example,
"192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly,
for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32"
is not. The originating IP of a request must be in one of the
listed subnets in order for this Condition to be true.
If empty, all IP addresses are allowed.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"members": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `An allowed list of members (users, service accounts).
Using groups is not supported yet.

The signed-in user originating the request must be a part of one
of the provided members. If not specified, a request may come
from any user (logged in/not logged in, not present in any
groups, etc.).
Formats: 'user:{emailid}', 'serviceAccount:{emailid}'`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"negate": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Description: `Whether to negate the Condition. If true, the Condition becomes
a NAND over its non-empty fields, each field must be false for
the Condition overall to be satisfied. Defaults to false.`,
			},
			"regions": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `The request must originate from one of the provided
countries/regions.
Format: A valid ISO 3166-1 alpha-2 code.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"required_access_levels": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `A list of other access levels defined in the same Policy,
referenced by resource name. Referencing an AccessLevel which
does not exist is an error. All access levels listed must be
granted for the Condition to be true.
Format: accessPolicies/{policy_id}/accessLevels/{short_name}`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		UseJSONNumber: true,
	}
}

func resourceAccessContextManagerAccessLevelConditionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	ipSubnetworksProp, err := expandNestedAccessContextManagerAccessLevelConditionIpSubnetworks(d.Get("ip_subnetworks"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_subnetworks"); !isEmptyValue(reflect.ValueOf(ipSubnetworksProp)) && (ok || !reflect.DeepEqual(v, ipSubnetworksProp)) {
		obj["ipSubnetworks"] = ipSubnetworksProp
	}
	requiredAccessLevelsProp, err := expandNestedAccessContextManagerAccessLevelConditionRequiredAccessLevels(d.Get("required_access_levels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("required_access_levels"); !isEmptyValue(reflect.ValueOf(requiredAccessLevelsProp)) && (ok || !reflect.DeepEqual(v, requiredAccessLevelsProp)) {
		obj["requiredAccessLevels"] = requiredAccessLevelsProp
	}
	membersProp, err := expandNestedAccessContextManagerAccessLevelConditionMembers(d.Get("members"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("members"); !isEmptyValue(reflect.ValueOf(membersProp)) && (ok || !reflect.DeepEqual(v, membersProp)) {
		obj["members"] = membersProp
	}
	negateProp, err := expandNestedAccessContextManagerAccessLevelConditionNegate(d.Get("negate"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("negate"); !isEmptyValue(reflect.ValueOf(negateProp)) && (ok || !reflect.DeepEqual(v, negateProp)) {
		obj["negate"] = negateProp
	}
	devicePolicyProp, err := expandNestedAccessContextManagerAccessLevelConditionDevicePolicy(d.Get("device_policy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("device_policy"); !isEmptyValue(reflect.ValueOf(devicePolicyProp)) && (ok || !reflect.DeepEqual(v, devicePolicyProp)) {
		obj["devicePolicy"] = devicePolicyProp
	}
	regionsProp, err := expandNestedAccessContextManagerAccessLevelConditionRegions(d.Get("regions"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("regions"); !isEmptyValue(reflect.ValueOf(regionsProp)) && (ok || !reflect.DeepEqual(v, regionsProp)) {
		obj["regions"] = regionsProp
	}

	lockName, err := replaceVars(d, config, "{{access_level}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{access_level}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new AccessLevelCondition: %#v", obj)

	obj, err = resourceAccessContextManagerAccessLevelConditionPatchCreateEncoder(d, meta, obj)
	if err != nil {
		return err
	}
	url, err = addQueryParams(url, map[string]string{"updateMask": "basic.conditions"})
	if err != nil {
		return err
	}
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating AccessLevelCondition: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{access_level}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = PollingWaitTime(resourceAccessContextManagerAccessLevelConditionPollRead(d, meta), PollCheckForExistence, "Creating AccessLevelCondition", d.Timeout(schema.TimeoutCreate), 1)
	if err != nil {
		return fmt.Errorf("Error waiting to create AccessLevelCondition: %s", err)
	}

	log.Printf("[DEBUG] Finished creating AccessLevelCondition %q: %#v", d.Id(), res)

	return resourceAccessContextManagerAccessLevelConditionRead(d, meta)
}

func resourceAccessContextManagerAccessLevelConditionPollRead(d *schema.ResourceData, meta interface{}) PollReadFunc {
	return func() (map[string]interface{}, error) {
		config := meta.(*Config)

		url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{access_level}}")
		if err != nil {
			return nil, err
		}

		billingProject := ""

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		userAgent, err := generateUserAgentString(d, config.userAgent)
		if err != nil {
			return nil, err
		}

		res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
		if err != nil {
			return res, err
		}
		res, err = flattenNestedAccessContextManagerAccessLevelCondition(d, meta, res)
		if err != nil {
			return nil, err
		}

		if res == nil {
			// Nested object not found, spoof a 404 error for poll
			return nil, &googleapi.Error{
				Code:    404,
				Message: "nested object AccessContextManagerAccessLevelCondition not found",
			}
		}

		return res, nil
	}
}

func resourceAccessContextManagerAccessLevelConditionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{access_level}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("AccessContextManagerAccessLevelCondition %q", d.Id()))
	}

	res, err = flattenNestedAccessContextManagerAccessLevelCondition(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing AccessContextManagerAccessLevelCondition because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	if err := d.Set("ip_subnetworks", flattenNestedAccessContextManagerAccessLevelConditionIpSubnetworks(res["ipSubnetworks"], d, config)); err != nil {
		return fmt.Errorf("Error reading AccessLevelCondition: %s", err)
	}
	if err := d.Set("required_access_levels", flattenNestedAccessContextManagerAccessLevelConditionRequiredAccessLevels(res["requiredAccessLevels"], d, config)); err != nil {
		return fmt.Errorf("Error reading AccessLevelCondition: %s", err)
	}
	if err := d.Set("members", flattenNestedAccessContextManagerAccessLevelConditionMembers(res["members"], d, config)); err != nil {
		return fmt.Errorf("Error reading AccessLevelCondition: %s", err)
	}
	if err := d.Set("negate", flattenNestedAccessContextManagerAccessLevelConditionNegate(res["negate"], d, config)); err != nil {
		return fmt.Errorf("Error reading AccessLevelCondition: %s", err)
	}
	if err := d.Set("device_policy", flattenNestedAccessContextManagerAccessLevelConditionDevicePolicy(res["devicePolicy"], d, config)); err != nil {
		return fmt.Errorf("Error reading AccessLevelCondition: %s", err)
	}
	if err := d.Set("regions", flattenNestedAccessContextManagerAccessLevelConditionRegions(res["regions"], d, config)); err != nil {
		return fmt.Errorf("Error reading AccessLevelCondition: %s", err)
	}

	return nil
}

func resourceAccessContextManagerAccessLevelConditionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	lockName, err := replaceVars(d, config, "{{access_level}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{access_level}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	obj, err = resourceAccessContextManagerAccessLevelConditionPatchDeleteEncoder(d, meta, obj)
	if err != nil {
		return handleNotFoundError(err, d, "AccessLevelCondition")
	}
	url, err = addQueryParams(url, map[string]string{"updateMask": "basic.conditions"})
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Deleting AccessLevelCondition %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "AccessLevelCondition")
	}

	log.Printf("[DEBUG] Finished deleting AccessLevelCondition %q: %#v", d.Id(), res)
	return nil
}

func flattenNestedAccessContextManagerAccessLevelConditionIpSubnetworks(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerAccessLevelConditionRequiredAccessLevels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerAccessLevelConditionMembers(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerAccessLevelConditionNegate(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerAccessLevelConditionDevicePolicy(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["require_screen_lock"] =
		flattenNestedAccessContextManagerAccessLevelConditionDevicePolicyRequireScreenLock(original["requireScreenlock"], d, config)
	transformed["allowed_encryption_statuses"] =
		flattenNestedAccessContextManagerAccessLevelConditionDevicePolicyAllowedEncryptionStatuses(original["allowedEncryptionStatuses"], d, config)
	transformed["allowed_device_management_levels"] =
		flattenNestedAccessContextManagerAccessLevelConditionDevicePolicyAllowedDeviceManagementLevels(original["allowedDeviceManagementLevels"], d, config)
	transformed["os_constraints"] =
		flattenNestedAccessContextManagerAccessLevelConditionDevicePolicyOsConstraints(original["osConstraints"], d, config)
	transformed["require_admin_approval"] =
		flattenNestedAccessContextManagerAccessLevelConditionDevicePolicyRequireAdminApproval(original["requireAdminApproval"], d, config)
	transformed["require_corp_owned"] =
		flattenNestedAccessContextManagerAccessLevelConditionDevicePolicyRequireCorpOwned(original["requireCorpOwned"], d, config)
	return []interface{}{transformed}
}
func flattenNestedAccessContextManagerAccessLevelConditionDevicePolicyRequireScreenLock(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerAccessLevelConditionDevicePolicyAllowedEncryptionStatuses(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerAccessLevelConditionDevicePolicyAllowedDeviceManagementLevels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerAccessLevelConditionDevicePolicyOsConstraints(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"minimum_version": flattenNestedAccessContextManagerAccessLevelConditionDevicePolicyOsConstraintsMinimumVersion(original["minimumVersion"], d, config),
			"os_type":         flattenNestedAccessContextManagerAccessLevelConditionDevicePolicyOsConstraintsOsType(original["osType"], d, config),
		})
	}
	return transformed
}
func flattenNestedAccessContextManagerAccessLevelConditionDevicePolicyOsConstraintsMinimumVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerAccessLevelConditionDevicePolicyOsConstraintsOsType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerAccessLevelConditionDevicePolicyRequireAdminApproval(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerAccessLevelConditionDevicePolicyRequireCorpOwned(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNestedAccessContextManagerAccessLevelConditionRegions(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandNestedAccessContextManagerAccessLevelConditionIpSubnetworks(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedAccessContextManagerAccessLevelConditionRequiredAccessLevels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedAccessContextManagerAccessLevelConditionMembers(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedAccessContextManagerAccessLevelConditionNegate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedAccessContextManagerAccessLevelConditionDevicePolicy(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRequireScreenLock, err := expandNestedAccessContextManagerAccessLevelConditionDevicePolicyRequireScreenLock(original["require_screen_lock"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequireScreenLock); val.IsValid() && !isEmptyValue(val) {
		transformed["requireScreenlock"] = transformedRequireScreenLock
	}

	transformedAllowedEncryptionStatuses, err := expandNestedAccessContextManagerAccessLevelConditionDevicePolicyAllowedEncryptionStatuses(original["allowed_encryption_statuses"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedEncryptionStatuses); val.IsValid() && !isEmptyValue(val) {
		transformed["allowedEncryptionStatuses"] = transformedAllowedEncryptionStatuses
	}

	transformedAllowedDeviceManagementLevels, err := expandNestedAccessContextManagerAccessLevelConditionDevicePolicyAllowedDeviceManagementLevels(original["allowed_device_management_levels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedDeviceManagementLevels); val.IsValid() && !isEmptyValue(val) {
		transformed["allowedDeviceManagementLevels"] = transformedAllowedDeviceManagementLevels
	}

	transformedOsConstraints, err := expandNestedAccessContextManagerAccessLevelConditionDevicePolicyOsConstraints(original["os_constraints"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOsConstraints); val.IsValid() && !isEmptyValue(val) {
		transformed["osConstraints"] = transformedOsConstraints
	}

	transformedRequireAdminApproval, err := expandNestedAccessContextManagerAccessLevelConditionDevicePolicyRequireAdminApproval(original["require_admin_approval"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequireAdminApproval); val.IsValid() && !isEmptyValue(val) {
		transformed["requireAdminApproval"] = transformedRequireAdminApproval
	}

	transformedRequireCorpOwned, err := expandNestedAccessContextManagerAccessLevelConditionDevicePolicyRequireCorpOwned(original["require_corp_owned"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequireCorpOwned); val.IsValid() && !isEmptyValue(val) {
		transformed["requireCorpOwned"] = transformedRequireCorpOwned
	}

	return transformed, nil
}

func expandNestedAccessContextManagerAccessLevelConditionDevicePolicyRequireScreenLock(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedAccessContextManagerAccessLevelConditionDevicePolicyAllowedEncryptionStatuses(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedAccessContextManagerAccessLevelConditionDevicePolicyAllowedDeviceManagementLevels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedAccessContextManagerAccessLevelConditionDevicePolicyOsConstraints(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedMinimumVersion, err := expandNestedAccessContextManagerAccessLevelConditionDevicePolicyOsConstraintsMinimumVersion(original["minimum_version"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMinimumVersion); val.IsValid() && !isEmptyValue(val) {
			transformed["minimumVersion"] = transformedMinimumVersion
		}

		transformedOsType, err := expandNestedAccessContextManagerAccessLevelConditionDevicePolicyOsConstraintsOsType(original["os_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedOsType); val.IsValid() && !isEmptyValue(val) {
			transformed["osType"] = transformedOsType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNestedAccessContextManagerAccessLevelConditionDevicePolicyOsConstraintsMinimumVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedAccessContextManagerAccessLevelConditionDevicePolicyOsConstraintsOsType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedAccessContextManagerAccessLevelConditionDevicePolicyRequireAdminApproval(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedAccessContextManagerAccessLevelConditionDevicePolicyRequireCorpOwned(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedAccessContextManagerAccessLevelConditionRegions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func flattenNestedAccessContextManagerAccessLevelCondition(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["basic"]
	if !ok || v == nil {
		return nil, nil
	}
	res = v.(map[string]interface{})

	v, ok = res["conditions"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value basic.conditions. Actual value: %v", v)
	}

	_, item, err := resourceAccessContextManagerAccessLevelConditionFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceAccessContextManagerAccessLevelConditionFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedIpSubnetworks, err := expandNestedAccessContextManagerAccessLevelConditionIpSubnetworks(d.Get("ip_subnetworks"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedIpSubnetworks := flattenNestedAccessContextManagerAccessLevelConditionIpSubnetworks(expectedIpSubnetworks, d, meta.(*Config))
	expectedRequiredAccessLevels, err := expandNestedAccessContextManagerAccessLevelConditionRequiredAccessLevels(d.Get("required_access_levels"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedRequiredAccessLevels := flattenNestedAccessContextManagerAccessLevelConditionRequiredAccessLevels(expectedRequiredAccessLevels, d, meta.(*Config))
	expectedMembers, err := expandNestedAccessContextManagerAccessLevelConditionMembers(d.Get("members"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedMembers := flattenNestedAccessContextManagerAccessLevelConditionMembers(expectedMembers, d, meta.(*Config))
	expectedNegate, err := expandNestedAccessContextManagerAccessLevelConditionNegate(d.Get("negate"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedNegate := flattenNestedAccessContextManagerAccessLevelConditionNegate(expectedNegate, d, meta.(*Config))
	expectedDevicePolicy, err := expandNestedAccessContextManagerAccessLevelConditionDevicePolicy(d.Get("device_policy"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedDevicePolicy := flattenNestedAccessContextManagerAccessLevelConditionDevicePolicy(expectedDevicePolicy, d, meta.(*Config))
	expectedRegions, err := expandNestedAccessContextManagerAccessLevelConditionRegions(d.Get("regions"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedRegions := flattenNestedAccessContextManagerAccessLevelConditionRegions(expectedRegions, d, meta.(*Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		itemIpSubnetworks := flattenNestedAccessContextManagerAccessLevelConditionIpSubnetworks(item["ipSubnetworks"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemIpSubnetworks)) && isEmptyValue(reflect.ValueOf(expectedFlattenedIpSubnetworks))) && !reflect.DeepEqual(itemIpSubnetworks, expectedFlattenedIpSubnetworks) {
			log.Printf("[DEBUG] Skipping item with ipSubnetworks= %#v, looking for %#v)", itemIpSubnetworks, expectedFlattenedIpSubnetworks)
			continue
		}
		itemRequiredAccessLevels := flattenNestedAccessContextManagerAccessLevelConditionRequiredAccessLevels(item["requiredAccessLevels"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemRequiredAccessLevels)) && isEmptyValue(reflect.ValueOf(expectedFlattenedRequiredAccessLevels))) && !reflect.DeepEqual(itemRequiredAccessLevels, expectedFlattenedRequiredAccessLevels) {
			log.Printf("[DEBUG] Skipping item with requiredAccessLevels= %#v, looking for %#v)", itemRequiredAccessLevels, expectedFlattenedRequiredAccessLevels)
			continue
		}
		itemMembers := flattenNestedAccessContextManagerAccessLevelConditionMembers(item["members"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemMembers)) && isEmptyValue(reflect.ValueOf(expectedFlattenedMembers))) && !reflect.DeepEqual(itemMembers, expectedFlattenedMembers) {
			log.Printf("[DEBUG] Skipping item with members= %#v, looking for %#v)", itemMembers, expectedFlattenedMembers)
			continue
		}
		itemNegate := flattenNestedAccessContextManagerAccessLevelConditionNegate(item["negate"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemNegate)) && isEmptyValue(reflect.ValueOf(expectedFlattenedNegate))) && !reflect.DeepEqual(itemNegate, expectedFlattenedNegate) {
			log.Printf("[DEBUG] Skipping item with negate= %#v, looking for %#v)", itemNegate, expectedFlattenedNegate)
			continue
		}
		itemDevicePolicy := flattenNestedAccessContextManagerAccessLevelConditionDevicePolicy(item["devicePolicy"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemDevicePolicy)) && isEmptyValue(reflect.ValueOf(expectedFlattenedDevicePolicy))) && !reflect.DeepEqual(itemDevicePolicy, expectedFlattenedDevicePolicy) {
			log.Printf("[DEBUG] Skipping item with devicePolicy= %#v, looking for %#v)", itemDevicePolicy, expectedFlattenedDevicePolicy)
			continue
		}
		itemRegions := flattenNestedAccessContextManagerAccessLevelConditionRegions(item["regions"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemRegions)) && isEmptyValue(reflect.ValueOf(expectedFlattenedRegions))) && !reflect.DeepEqual(itemRegions, expectedFlattenedRegions) {
			log.Printf("[DEBUG] Skipping item with regions= %#v, looking for %#v)", itemRegions, expectedFlattenedRegions)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}

// PatchCreateEncoder handles creating request data to PATCH parent resource
// with list including new object.
func resourceAccessContextManagerAccessLevelConditionPatchCreateEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceAccessContextManagerAccessLevelConditionListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	_, found, err := resourceAccessContextManagerAccessLevelConditionFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}

	// Return error if item already created.
	if found != nil {
		return nil, fmt.Errorf("Unable to create AccessLevelCondition, existing object already found: %+v", found)
	}

	// Return list with the resource to create appended
	res := map[string]interface{}{
		"conditions": append(currItems, obj),
	}
	wrapped := map[string]interface{}{
		"basic": res,
	}
	res = wrapped

	return res, nil
}

// PatchDeleteEncoder handles creating request data to PATCH parent resource
// with list excluding object to delete.
func resourceAccessContextManagerAccessLevelConditionPatchDeleteEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	currItems, err := resourceAccessContextManagerAccessLevelConditionListForPatch(d, meta)
	if err != nil {
		return nil, err
	}

	idx, item, err := resourceAccessContextManagerAccessLevelConditionFindNestedObjectInList(d, meta, currItems)
	if err != nil {
		return nil, err
	}
	if item == nil {
		// Spoof 404 error for proper handling by Delete (i.e. no-op)
		return nil, &googleapi.Error{
			Code:    404,
			Message: "AccessLevelCondition not found in list",
		}
	}

	updatedItems := append(currItems[:idx], currItems[idx+1:]...)
	res := map[string]interface{}{
		"conditions": updatedItems,
	}
	wrapped := map[string]interface{}{
		"basic": res,
	}
	res = wrapped

	return res, nil
}

// ListForPatch handles making API request to get parent resource and
// extracting list of objects.
func resourceAccessContextManagerAccessLevelConditionListForPatch(d *schema.ResourceData, meta interface{}) ([]interface{}, error) {
	config := meta.(*Config)
	url, err := replaceVars(d, config, "{{AccessContextManagerBasePath}}{{access_level}}")
	if err != nil {
		return nil, err
	}

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return nil, err
	}

	res, err := sendRequest(config, "GET", "", url, userAgent, nil)
	if err != nil {
		return nil, err
	}

	var v interface{}
	var ok bool
	if v, ok = res["basic"]; ok && v != nil {
		res = v.(map[string]interface{})
	} else {
		return nil, nil
	}

	v, ok = res["conditions"]
	if ok && v != nil {
		ls, lsOk := v.([]interface{})
		if !lsOk {
			return nil, fmt.Errorf(`expected list for nested field "conditions"`)
		}
		return ls, nil
	}
	return nil, nil
}
