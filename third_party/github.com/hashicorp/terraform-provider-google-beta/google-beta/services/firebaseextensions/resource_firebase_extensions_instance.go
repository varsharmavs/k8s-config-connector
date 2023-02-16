// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

package firebaseextensions

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func ResourceFirebaseExtensionsInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirebaseExtensionsInstanceCreate,
		Read:   resourceFirebaseExtensionsInstanceRead,
		Update: resourceFirebaseExtensionsInstanceUpdate,
		Delete: resourceFirebaseExtensionsInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirebaseExtensionsInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"config": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `The current Config of the Extension Instance.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"extension_ref": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `The ref of the Extension from the Registry (e.g. publisher-id/awesome-extension)`,
						},
						"params": {
							Type:        schema.TypeMap,
							Required:    true,
							Description: `Environment variables that may be configured for the Extension`,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"allowed_event_types": {
							Type:     schema.TypeList,
							Optional: true,
							Description: `List of extension events selected by consumer that extension is allowed to
emit, identified by their types.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"eventarc_channel": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: `Fully qualified Eventarc resource name that consumers should use for event triggers.`,
						},
						"extension_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: `The version of the Extension from the Registry (e.g. 1.0.3). If left blank, latest is assumed.`,
						},
						"system_params": {
							Type:     schema.TypeMap,
							Computed: true,
							Optional: true,
							Description: `Params whose values are only available at deployment time.
Unlike other params, these will not be set as environment variables on
functions.`,
							Elem: &schema.Schema{Type: schema.TypeString},
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The time at which the Extension Instance Config was created.`,
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The unique identifier for this configuration.`,
						},
						"populated_postinstall_content": {
							Type:     schema.TypeString,
							Computed: true,
							Description: `Postinstall instructions to be shown for this Extension, with
template strings representing function and parameter values substituted
with actual values. These strings include: ${param:FOO},
${function:myFunc.url},
${function:myFunc.name}, and ${function:myFunc.location}`,
						},
					},
				},
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The ID to use for the Extension Instance, which will become the final
component of the instance's name.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time at which the Extension Instance was created.`,
			},
			"error_status": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `If this Instance has 'state: ERRORED', the error messages
will be found here.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: `The status code, which should be an enum value of google.rpc.Code.`,
						},
						"details": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `A list of messages that carry the error details.`,
							Elem: &schema.Schema{
								Type: schema.TypeMap,
							},
						},
						"message": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `A developer-facing error message, which should be in English.`,
						},
					},
				},
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `A weak etag that is computed by the server based on other configuration
values and may be sent on update and delete requests to ensure the
client has an up-to-date value before proceeding.`,
			},
			"last_operation_name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The name of the last operation that acted on this Extension
Instance`,
			},
			"last_operation_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The type of the last operation that acted on the Extension Instance.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The fully-qualified resource name of the Extension Instance.`,
			},
			"runtime_data": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Data set by the extension instance at runtime.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fatal_error": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `The fatal error state for the extension instance`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"error_message": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `The error message. This is set by the extension developer to give
more detail on why the extension is unusable and must be re-installed
or reconfigured.`,
									},
								},
							},
						},
						"processing_state": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: `The processing state for the extension instance`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"detail_message": {
										Type:     schema.TypeString,
										Optional: true,
										Description: `Details about the processing. e.g. This could include the type of
processing in progress or it could list errors or failures.
This information will be shown in the console on the detail page
for the extension instance.`,
									},
									"state": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The processing state of the extension instance.`,
									},
								},
							},
						},
						"state_update_time": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `The time of the last state update.`,
						},
					},
				},
			},
			"service_account_email": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The email of the service account to be used at runtime by compute resources
created for the operation of the Extension instance.`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The current state of the Extension Instance.
  DEPLOYING:
    The Extension Instance is waiting on an Operation to complete. Could
    resolve to 'ACTIVE', 'PAUSED', 'ERRORED'.
  UNINSTALLING:
    The Extension Instance is being removed from the project. Could resolve
    to 'ERRORED', but more likely the instance will soon cease to exist.
  ACTIVE:
    The Extension Instance is installed and ready.
  ERRORED:
    The Extension Instance encountered an error while 'DEPLOYING' or
    'UNINSTALLING'.
  PAUSED:
    The Extension's resources have been removed from the project, but the
    Config remains so the Instance can be restored.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time at which the Extension Instance was updated.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceFirebaseExtensionsInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	configProp, err := expandFirebaseExtensionsInstanceConfig(d.Get("config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("config"); !tpgresource.IsEmptyValue(reflect.ValueOf(configProp)) && (ok || !reflect.DeepEqual(v, configProp)) {
		obj["config"] = configProp
	}
	etagProp, err := expandFirebaseExtensionsInstanceEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseExtensionsBasePath}}projects/{{project}}/instances?instanceId={{instance_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Instance: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating Instance: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/instances/{{instance_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = FirebaseExtensionsOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Instance", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create Instance: %s", err)
	}

	if err := d.Set("name", flattenFirebaseExtensionsInstanceName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "projects/{{project}}/instances/{{instance_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Instance %q: %#v", d.Id(), res)

	return resourceFirebaseExtensionsInstanceRead(d, meta)
}

func resourceFirebaseExtensionsInstanceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseExtensionsBasePath}}projects/{{project}}/instances/{{instance_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("FirebaseExtensionsInstance %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	if err := d.Set("name", flattenFirebaseExtensionsInstanceName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("create_time", flattenFirebaseExtensionsInstanceCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("update_time", flattenFirebaseExtensionsInstanceUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("config", flattenFirebaseExtensionsInstanceConfig(res["config"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("state", flattenFirebaseExtensionsInstanceState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("error_status", flattenFirebaseExtensionsInstanceErrorStatus(res["errorStatus"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("service_account_email", flattenFirebaseExtensionsInstanceServiceAccountEmail(res["serviceAccountEmail"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("last_operation_name", flattenFirebaseExtensionsInstanceLastOperationName(res["lastOperationName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("last_operation_type", flattenFirebaseExtensionsInstanceLastOperationType(res["lastOperationType"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("etag", flattenFirebaseExtensionsInstanceEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("runtime_data", flattenFirebaseExtensionsInstanceRuntimeData(res["runtimeData"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	return nil
}

func resourceFirebaseExtensionsInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	configProp, err := expandFirebaseExtensionsInstanceConfig(d.Get("config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, configProp)) {
		obj["config"] = configProp
	}
	etagProp, err := expandFirebaseExtensionsInstanceEtag(d.Get("etag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseExtensionsBasePath}}projects/{{project}}/instances/{{instance_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Instance %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("config") {
		updateMask = append(updateMask, "config.params",
			"config.system_params",
			"config.extension_ref",
			"config.extension_version",
			"config.allowed_event_types",
			"config.eventarc_channel")
	}

	if d.HasChange("etag") {
		updateMask = append(updateMask, "etag")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating Instance %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Instance %q: %#v", d.Id(), res)
	}

	err = FirebaseExtensionsOperationWaitTime(
		config, res, project, "Updating Instance", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceFirebaseExtensionsInstanceRead(d, meta)
}

func resourceFirebaseExtensionsInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Instance: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{FirebaseExtensionsBasePath}}projects/{{project}}/instances/{{instance_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Instance %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Instance")
	}

	err = FirebaseExtensionsOperationWaitTime(
		config, res, project, "Deleting Instance", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Instance %q: %#v", d.Id(), res)
	return nil
}

func resourceFirebaseExtensionsInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/instances/(?P<instance_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<instance_id>[^/]+)",
		"(?P<instance_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/instances/{{instance_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenFirebaseExtensionsInstanceName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["name"] =
		flattenFirebaseExtensionsInstanceConfigName(original["name"], d, config)
	transformed["create_time"] =
		flattenFirebaseExtensionsInstanceConfigCreateTime(original["createTime"], d, config)
	transformed["params"] =
		flattenFirebaseExtensionsInstanceConfigParams(original["params"], d, config)
	transformed["system_params"] =
		flattenFirebaseExtensionsInstanceConfigSystemParams(original["systemParams"], d, config)
	transformed["extension_ref"] =
		flattenFirebaseExtensionsInstanceConfigExtensionRef(original["extensionRef"], d, config)
	transformed["extension_version"] =
		flattenFirebaseExtensionsInstanceConfigExtensionVersion(original["extensionVersion"], d, config)
	transformed["allowed_event_types"] =
		flattenFirebaseExtensionsInstanceConfigAllowedEventTypes(original["allowedEventTypes"], d, config)
	transformed["eventarc_channel"] =
		flattenFirebaseExtensionsInstanceConfigEventarcChannel(original["eventarcChannel"], d, config)
	transformed["populated_postinstall_content"] =
		flattenFirebaseExtensionsInstanceConfigPopulatedPostinstallContent(original["populatedPostinstallContent"], d, config)
	return []interface{}{transformed}
}
func flattenFirebaseExtensionsInstanceConfigName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceConfigCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceConfigParams(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceConfigSystemParams(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceConfigExtensionRef(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceConfigExtensionVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceConfigAllowedEventTypes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceConfigEventarcChannel(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceConfigPopulatedPostinstallContent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceErrorStatus(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["code"] =
		flattenFirebaseExtensionsInstanceErrorStatusCode(original["code"], d, config)
	transformed["message"] =
		flattenFirebaseExtensionsInstanceErrorStatusMessage(original["message"], d, config)
	transformed["details"] =
		flattenFirebaseExtensionsInstanceErrorStatusDetails(original["details"], d, config)
	return []interface{}{transformed}
}
func flattenFirebaseExtensionsInstanceErrorStatusCode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenFirebaseExtensionsInstanceErrorStatusMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceErrorStatusDetails(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceServiceAccountEmail(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceLastOperationName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceLastOperationType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceEtag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceRuntimeData(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["state_update_time"] =
		flattenFirebaseExtensionsInstanceRuntimeDataStateUpdateTime(original["stateUpdateTime"], d, config)
	transformed["processing_state"] =
		flattenFirebaseExtensionsInstanceRuntimeDataProcessingState(original["processingState"], d, config)
	transformed["fatal_error"] =
		flattenFirebaseExtensionsInstanceRuntimeDataFatalError(original["fatalError"], d, config)
	return []interface{}{transformed}
}
func flattenFirebaseExtensionsInstanceRuntimeDataStateUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceRuntimeDataProcessingState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["state"] =
		flattenFirebaseExtensionsInstanceRuntimeDataProcessingStateState(original["state"], d, config)
	transformed["detail_message"] =
		flattenFirebaseExtensionsInstanceRuntimeDataProcessingStateDetailMessage(original["detailMessage"], d, config)
	return []interface{}{transformed}
}
func flattenFirebaseExtensionsInstanceRuntimeDataProcessingStateState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceRuntimeDataProcessingStateDetailMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenFirebaseExtensionsInstanceRuntimeDataFatalError(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["error_message"] =
		flattenFirebaseExtensionsInstanceRuntimeDataFatalErrorErrorMessage(original["errorMessage"], d, config)
	return []interface{}{transformed}
}
func flattenFirebaseExtensionsInstanceRuntimeDataFatalErrorErrorMessage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandFirebaseExtensionsInstanceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandFirebaseExtensionsInstanceConfigName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["name"] = transformedName
	}

	transformedCreateTime, err := expandFirebaseExtensionsInstanceConfigCreateTime(original["create_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCreateTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["createTime"] = transformedCreateTime
	}

	transformedParams, err := expandFirebaseExtensionsInstanceConfigParams(original["params"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedParams); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["params"] = transformedParams
	}

	transformedSystemParams, err := expandFirebaseExtensionsInstanceConfigSystemParams(original["system_params"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSystemParams); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["systemParams"] = transformedSystemParams
	}

	transformedExtensionRef, err := expandFirebaseExtensionsInstanceConfigExtensionRef(original["extension_ref"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExtensionRef); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["extensionRef"] = transformedExtensionRef
	}

	transformedExtensionVersion, err := expandFirebaseExtensionsInstanceConfigExtensionVersion(original["extension_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExtensionVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["extensionVersion"] = transformedExtensionVersion
	}

	transformedAllowedEventTypes, err := expandFirebaseExtensionsInstanceConfigAllowedEventTypes(original["allowed_event_types"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedEventTypes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowedEventTypes"] = transformedAllowedEventTypes
	}

	transformedEventarcChannel, err := expandFirebaseExtensionsInstanceConfigEventarcChannel(original["eventarc_channel"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEventarcChannel); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["eventarcChannel"] = transformedEventarcChannel
	}

	transformedPopulatedPostinstallContent, err := expandFirebaseExtensionsInstanceConfigPopulatedPostinstallContent(original["populated_postinstall_content"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPopulatedPostinstallContent); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["populatedPostinstallContent"] = transformedPopulatedPostinstallContent
	}

	return transformed, nil
}

func expandFirebaseExtensionsInstanceConfigName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseExtensionsInstanceConfigCreateTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseExtensionsInstanceConfigParams(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandFirebaseExtensionsInstanceConfigSystemParams(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandFirebaseExtensionsInstanceConfigExtensionRef(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseExtensionsInstanceConfigExtensionVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseExtensionsInstanceConfigAllowedEventTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseExtensionsInstanceConfigEventarcChannel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseExtensionsInstanceConfigPopulatedPostinstallContent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseExtensionsInstanceEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}