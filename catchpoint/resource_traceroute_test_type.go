package catchpoint

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceTracerouteTestType() *schema.Resource {
	return &schema.Resource{
		Create: resourceTracerouteTestCreate,
		Read:   resourceTracerouteTestRead,
		Update: resourceTracerouteTestUpdate,
		Delete: resourceTracerouteTestDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"monitor": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The monitor to use for the Traceroute Test. Supported: 'traceroute icmp','traceroute tcp','traceroute udp'",
				Default:      "traceroute icmp",
				ValidateFunc: validation.StringInSlice([]string{"traceroute icmp", "traceroute tcp", "traceroute udp"}, false),
			},
			"division_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The Division where the Test will be created",
			},
			"product_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The parent Product under which the Test will be created",
			},
			"folder_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Optional. The Folder under which the Test will be created",
			},
			"test_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the Test",
			},
			"test_location": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The domain or IP to be tested. Example: www.catchpoint.com",
			},
			"test_description": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Optional. The Test description",
			},
			"enable_test_data_webhook": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Optional. Switch for enabling test data webhook feature",
			},
			"alerts_paused": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Optional. Switch for pausing Test alerts",
			},
			"start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Optional. Start time for the Test in ISO format",
			},
			"end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional. End time for the Test in ISO format",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional. Test status: active or inactive",
			},
			"label": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Optional. Label with key, values pair",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"values": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"thresholds": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Optional. Test thresholds for test time and availability percentage",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"test_time_warning": {
							Type:     schema.TypeFloat,
							Required: true,
						},
						"test_time_critical": {
							Type:     schema.TypeFloat,
							Required: true,
						},
						"availability_warning": {
							Type:     schema.TypeFloat,
							Required: true,
						},
						"availability_critical": {
							Type:     schema.TypeFloat,
							Required: true,
						},
					},
				},
			},
			"schedule_settings": {
				Type:        schema.TypeSet,
				Optional:    true,
				MaxItems:    1,
				Description: "Optional. Used for overriding the schedule section",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"run_schedule_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The run schedule id to utilize for the test",
						},
						"maintenance_schedule_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The maintenance schedule id to utilize for the test",
						},
						"frequency": {
							Type:         schema.TypeString,
							Required:     true,
							Description:  "Sets the scheduling frequency: '1 minute', '5 minutes', '10 minutes', '15 minutes', '20 minutes', '30 minutes', '60 minutes', '2 hours', '3 hours', '4 hours', '6 hours', '8 hours', '12 hours', '24 hours', '4 minutes', '2 minutes'",
							ValidateFunc: validation.StringInSlice([]string{"1 minute", "5 minutes", "10 minutes", "15 minutes", "20 minutes", "30 minutes", "60 minutes", "2 hours", "3 hours", "4 hours", "6 hours", "8 hours", "12 hours", "24 hours", "4 minutes", "2 minutes"}, false),
						},
						"node_distribution": {
							Type:         schema.TypeString,
							Required:     true,
							Description:  "Node distribution type: 'random' or 'concurrent'",
							ValidateFunc: validation.StringInSlice([]string{"random", "concurrent"}, false),
						},
						"node_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Optional if node_group_ids is used. Node ids in a list",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"node_group_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Optional if node_ids is used. Node group ids in a list",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"no_of_subset_nodes": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Number of subset nodes",
						},
					},
				},
			},
			"alert_settings": {
				Type:        schema.TypeSet,
				Optional:    true,
				MaxItems:    1,
				Description: "Optional. Used for overriding the alert section",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alert_rule": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Optional. Sets the alert rule with attributes such as threshold, trigger type, warning, critical trigger and more",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"node_threshold_type": {
										Type:         schema.TypeString,
										Required:     true,
										Description:  "Sets the node threshold type for alert: 'runs', 'average across node' or 'node'",
										ValidateFunc: validation.StringInSlice([]string{"runs", "average across nodes", "node"}, false),
									},
									"threshold_number_of_runs": {
										Type:        schema.TypeInt,
										Description: "Optional. Sets the threshold for the number of runs or nodes the alert should trigger",
										Optional:    true,
									},
									"threshold_percentage_of_runs": {
										Type:        schema.TypeFloat,
										Description: "Optional. Sets the threshold for the percentage of runs the alert should trigger",
										Optional:    true,
									},
									"number_of_failing_nodes": {
										Type:        schema.TypeInt,
										Description: "Optional. Sets the number of failed nodes the alert should trigger if node_threshold_type is 'average across nodes'",
										Optional:    true,
									},
									"trigger_type": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "Optional. Sets the trigger type: 'specific value', 'trailing value', 'trendshift'",
										ValidateFunc: validation.StringInSlice([]string{"specific value", "trailing value", "trendshift"}, false),
									},
									"operation_type": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "Optional. Sets the operation type:'equals', 'not equals', 'greater than', 'greater than or equals', 'less than', 'less than or equals'",
										ValidateFunc: validation.StringInSlice([]string{"equals", "not equals", "greater than", "greater than or equals", "less than", "less than or equals"}, false),
									},
									"statistical_type": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "Optional. Sets the statistical type for 'trailing value' trigger type. Supports only 'average' for now",
										ValidateFunc: validation.StringInSlice([]string{"average"}, false),
									},
									"historical_interval": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "Optional. Sets the historical interval for 'trailing value' trigger type: '5 minutes', '10 minutes', '15 minutes', '30 minutes', '1 hour', '2 hours', '6 hours', '12 hours', '1 day', '1 week'",
										ValidateFunc: validation.StringInSlice([]string{"5 minutes", "10 minutes", "15 minutes", "30 minutes", "1 hour", "2 hours", "6 hours", "12 hours", "1 day", "1 week"}, false),
									},
									"warning_trigger": {
										Type:        schema.TypeFloat,
										Description: "Optional. Warning trigger value for 'specific value' and 'trailing value' trigger types.",
										Optional:    true,
									},
									"critical_trigger": {
										Type:        schema.TypeFloat,
										Description: "Optional. Critical trigger value for 'specific value' and 'trailing value' trigger types.",
										Optional:    true,
									},
									"enable_consecutive": {
										Type:        schema.TypeBool,
										Description: "Optional. Checks consecutive number of runs or nodes for triggering alerts.",
										Optional:    true,
										Default:     false,
									},
									"consecutive_number_of_runs": {
										Type:        schema.TypeInt,
										Description: "Optional. Sets the number of consecutive runs only if enable_consecutive field is true and node_threshold_type is node",
										Optional:    true,
									},
									"expression": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Optional. Sets trigger expression for ASN alert type ",
									},
									"warning_reminder": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "Optional. Sets alert warning reminder interval: 'none', '1 minute', '5 minutes', '10 minutes', '15 minutes', '30 minutes', '1 hour', 'daily'",
										ValidateFunc: validation.StringInSlice([]string{"none", "1 minute", "5 minutes", "10 minutes", "15 minutes", "30 minutes", "1 hour", "daily"}, false),
									},
									"critical_reminder": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "Optional. Sets alert critical reminder interval: 'none', '1 minute', '5 minutes', '10 minutes', '15 minutes', '30 minutes', '1 hour', 'daily'",
										ValidateFunc: validation.StringInSlice([]string{"none", "1 minute", "5 minutes", "10 minutes", "15 minutes", "30 minutes", "1 hour", "daily"}, false),
									},
									"threshold_interval": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "Optional. Sets the alert time threshold: 'default', '5 minutes', '10 minutes', '15 minutes', '30 minutes', '1 hour', '2 hours', '6 hours', '12 hours'",
										ValidateFunc: validation.StringInSlice([]string{"default", "5 minutes", "10 minutes", "15 minutes", "30 minutes", "1 hour", "2 hours", "6 hours", "12 hours"}, false),
									},
									"use_rolling_window": {
										Type:        schema.TypeBool,
										Description: "Optional. Set to true for using rolling window instead of schedule time threshold",
										Optional:    true,
										Default:     false,
									},
									"notification_type": {
										Type:         schema.TypeString,
										Description:  "Optional. Notification group type to alert. Supports only default contacts for now.",
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"default contacts"}, false),
									},
									"alert_type": {
										Type:         schema.TypeString,
										Description:  "Sets the alert type",
										Required:     true,
										ValidateFunc: validation.StringInSlice([]string{"test failure", "ping", "path", "asn"}, false),
									},
									"alert_sub_type": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "Optional. Sets the sub alert type: '# cities','# asns','# countries','# hops','origin as','path as','origin neighbor','prefix mismatch','ping rtt','ping packet loss'",
										ValidateFunc: validation.StringInSlice([]string{"# cities", "# asns", "# countries", "# hops", "origin as", "path as", "origin neighbor", "prefix mismatch", "ping rtt", "ping packet loss"}, false),
									},
									"enforce_test_failure": {
										Type:        schema.TypeBool,
										Description: "Optional. Sets enforce test failure property for an alert",
										Optional:    true,
										Default:     false,
									},
									"omit_scatterplot": {
										Type:        schema.TypeBool,
										Description: "Optional. Omits scatterplot image from alert emails if set to true",
										Optional:    true,
										Default:     false,
									},
									"notification_group": {
										Type:        schema.TypeSet,
										Optional:    true,
										MaxItems:    1,
										Description: "Notification group for configuring alert notifications, including recipients' email addresses and alert settings. To ensure either recipient_email_ids or contact_groups is provided",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"notify_on_warning": {
													Type:        schema.TypeBool,
													Description: "Optional. Set to true to include warning alerts in notifications. Default is false.",
													Optional:    true,
													Default:     false,
												},
												"notify_on_critical": {
													Type:        schema.TypeBool,
													Description: "Optional. Set to true to include critical alerts in notifications. Default is false.",
													Optional:    true,
													Default:     false,
												},
												"notify_on_improved": {
													Type:        schema.TypeBool,
													Description: "Optional. Set to true to include improved alerts in notifications. Default is false.",
													Optional:    true,
													Default:     false,
												},
												"subject": {
													Type:        schema.TypeString,
													Description: "Email subject for the alert notifications. Required field.",
													Required:    true,
												},
												"recipient_email_ids": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "List of email addresses to receive alert notifications. To ensure either recipient_email_ids or contact_groups is provided",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"contact_groups": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "List of contact groups to receive alert notifications. To ensure either recipient_email_ids or contact_groups is provided",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"notification_group": {
							Type:        schema.TypeSet,
							Required:    true,
							MaxItems:    1,
							Description: "Notification group for setting up alert recipients, adding alert webhook ids. To ensure either recipient_email_ids or contact_groups is provided",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"subject": {
										Type:        schema.TypeString,
										Description: "Email subject for the alert notifications. Required field.",
										Required:    true,
									},
									"alert_webhook_ids": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Optional. Alert webhook ids for the webhook endpoints to associate with this alert setting.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"recipient_email_ids": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Optional. List of emails to alert. To ensure either recipient_email_ids or contact_groups is provided",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"contact_groups": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of contact groups to receive alert notifications. To ensure either recipient_email_ids or contact_groups is provided",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
			"advanced_settings": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Optional. Used for overriding the advanced settings",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_path_mtu_discovery": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables Path MTU Discovery",
							Optional:    true,
							Default:     false,
						},
						"ping_count": {
							Type:         schema.TypeInt,
							Description:  "Optional. Set the ping count from 4 to 20",
							Optional:     true,
							ValidateFunc: validation.IntBetween(4, 20),
						},
						"failure_hop_count": {
							Type:         schema.TypeInt,
							Description:  "Optional. Set the failure hop count from 7 to 20",
							Optional:     true,
							ValidateFunc: validation.IntBetween(7, 20),
						},
						"additional_monitor": {
							Type:         schema.TypeString,
							Description:  "Optional. Set the additional monitor to run along with the test monitor: 'ping icmp', 'ping tcp', 'ping udp','traceroute icmp','traceroute udp','traceroute tcp'",
							ValidateFunc: validation.StringInSlice([]string{"ping icmp", "ping tcp", "ping udp", "traceroute icmp", "traceroute udp", "traceroute tcp"}, false),
							Optional:     true,
						},
						"bandwidth_throttling": {
							Type:         schema.TypeString,
							Description:  "Optional. Set the bandwidth throttling for chrome: 'gprs','regular 2g','good 2g','regular 3g','good 3g','regular 4g','dsl','wifi'",
							ValidateFunc: validation.StringInSlice([]string{"gprs", "regular 2g", "good 2g", "regular 3g", "good 3g", "regular 4g", "dsl", "wifi"}, false),
							Optional:     true,
						},
					},
				},
			},
		},
	}
}

func resourceTracerouteTestCreate(d *schema.ResourceData, m interface{}) error {
	api_token := m.(*Config).ApiToken
	monitor := d.Get("monitor").(string)
	monitor_id := getMonitorId(monitor)
	division_id := d.Get("division_id").(int)
	product_id := d.Get("product_id").(int)
	folder_id := d.Get("folder_id").(int)
	test_name := d.Get("test_name").(string)
	test_location := d.Get("test_location").(string)
	test_description := d.Get("test_description").(string)
	enable_test_data_webhook := d.Get("enable_test_data_webhook").(bool)
	alerts_paused := d.Get("alerts_paused").(bool)
	start_time := d.Get("start_time").(string)
	if start_time == "" {
		start_time = getTime()
	}
	end_time := d.Get("end_time").(string)
	status := d.Get("status").(string)
	status_id := getTestStatusTypeId(status)
	test_type := TestType(Traceroute)

	var testConfig = TestConfig{}

	testConfig = TestConfig{
		TestType:              int(test_type),
		TestUrl:               test_location,
		Monitor:               monitor_id,
		DivisionId:            division_id,
		ProductId:             product_id,
		FolderId:              folder_id,
		TestName:              test_name,
		TestDescription:       test_description,
		EnableTestDataWebhook: enable_test_data_webhook,
		AlertsPaused:          alerts_paused,
		StartTime:             start_time,
		EndTime:               end_time,
		TestStatus:            status_id,
	}

	label, labelOk := d.GetOk("label")
	if labelOk {
		label_lists := label.(*schema.Set).List()

		setLabels(int(test_type), label_lists, &testConfig)
	}

	thresholds, thresholdOk := d.GetOk("thresholds")
	if thresholdOk {
		thresholds_lists := thresholds.(*schema.Set).List()
		threshold := thresholds_lists[0].(map[string]interface{})

		setThresholds(int(test_type), threshold, &testConfig)
	}

	schedule_settings, schedule_settingsOk := d.GetOk("schedule_settings")
	if schedule_settingsOk {
		schedule_setting_list := schedule_settings.(*schema.Set).List()
		schedule_setting := schedule_setting_list[0].(map[string]interface{})

		err := setScheduleSettings(int(test_type), schedule_setting, &testConfig)
		if err != nil {
			return err
		}
	}

	alert_settings, alert_settingsOk := d.GetOk("alert_settings")
	if alert_settingsOk {
		alert_setting_list := alert_settings.(*schema.Set).List()
		alert_setting := alert_setting_list[0].(map[string]interface{})

		err := setAlertSettings(int(test_type), alert_setting, &testConfig)
		if err != nil {
			return err
		}
	}

	advanced_settings, advanced_settingsOk := d.GetOk("advanced_settings")
	if advanced_settingsOk {
		advanced_setting_list := advanced_settings.(*schema.Set).List()
		advanced_setting := advanced_setting_list[0].(map[string]interface{})

		setAdvancedSettings(int(test_type), advanced_setting, &testConfig)
	}

	jsonStr := createJson(testConfig)

	if m.(*Config).LogJson {
		log.Printf("[TEST JSON] \n" + jsonStr)
	}

	log.Printf("[DEBUG] Creating test: " + test_name)
	respBody, respStatus, testId, err := createTest(api_token, jsonStr)
	if err != nil {
		log.Fatal(err)
	}
	if respStatus != "200 ok" {
		log.Printf("[ERROR] Error while creating test: " + test_name)
		log.Printf("[ERROR] Error description: " + respBody)
		return errors.New(respStatus)
	}

	log.Printf("[DEBUG] Response Code from Catchpoint API: " + respStatus)
	log.Print(respBody)

	d.SetId(testId)
	return resourceTracerouteTestRead(d, m)
}

func resourceTracerouteTestRead(d *schema.ResourceData, m interface{}) error {
	testId := d.Id()
	api_token := m.(*Config).ApiToken

	log.Printf("[DEBUG] Fetching test: %v", testId)

	test, respStatus, err := getTest(api_token, testId)
	if err != nil {
		return err
	}
	if respStatus != "200 ok" {
		log.Printf("[ERROR] Error while reading test: %v", testId)
		return errors.New(respStatus)
	}
	if test == nil {
		d.SetId("")
		log.Printf("[DEBUG] Test not found %v", testId)
		return nil
	}
	log.Printf("[DEBUG] Response Code from Catchpoint API: " + respStatus)

	testNew := flattenTest(test)

	d.Set("monitor", testNew["monitor"])
	d.Set("division_id", testNew["division_id"])
	d.Set("product_id", testNew["product_id"])
	d.Set("folder_id", testNew["folder_id"])
	d.Set("test_name", testNew["test_name"])
	d.Set("test_description", testNew["test_description"])
	d.Set("enable_test_data_webhook", testNew["enable_test_data_webhook"])
	d.Set("alerts_paused", testNew["alerts_paused"])
	d.Set("start_time", testNew["start_time"])
	d.Set("end_time", testNew["end_time"])
	d.Set("status", testNew["status"])
	d.Set("test_location", testNew["test_url"])
	d.Set("label", testNew["label"])
	d.Set("thresholds", testNew["thresholds"])
	d.Set("schedule_settings", testNew["schedule_settings"])
	d.Set("alert_settings", testNew["alert_settings"])
	d.Set("advanced_settings", testNew["advanced_settings"])

	return nil
}

func resourceTracerouteTestUpdate(d *schema.ResourceData, m interface{}) error {
	testId := d.Id()
	api_token := m.(*Config).ApiToken
	test_type := TestType(Traceroute)
	var testConfig = TestConfig{}
	var jsonPatchDocs = []string{}

	if d.HasChange("test_name") {
		testConfigUpdate := TestConfigUpdate{
			UpdatedFieldValue: d.Get("test_name").(string),
		}
		jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, "/name", true))
	}
	if d.HasChange("test_location") {
		testConfigUpdate := TestConfigUpdate{
			UpdatedFieldValue: d.Get("test_location").(string),
		}
		jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, "/url", true))
	}
	if d.HasChange("monitor") {
		monitor := d.Get("monitor").(string)
		monitor_id := getMonitorId(monitor)
		testConfigUpdate := TestConfigUpdate{
			UpdatedFieldValue: strconv.Itoa(monitor_id),
		}
		jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, "/monitor", true))
	}
	if d.HasChange("test_description") {
		testConfigUpdate := TestConfigUpdate{
			UpdatedFieldValue: d.Get("test_description").(string),
		}
		jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, "/description", true))
	}
	if d.HasChange("enable_test_data_webhook") {
		testConfigUpdate := TestConfigUpdate{
			UpdatedFieldValue: strconv.FormatBool(d.Get("enable_test_data_webhook").(bool)),
		}
		jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, "/enableTestDataWebhook", true))
	}
	if d.HasChange("alerts_paused") {
		testConfigUpdate := TestConfigUpdate{
			UpdatedFieldValue: strconv.FormatBool(d.Get("alerts_paused").(bool)),
		}
		jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, "/alertsPaused", true))
	}
	if d.HasChange("start_time") {
		testConfigUpdate := TestConfigUpdate{
			UpdatedFieldValue: d.Get("start_time").(string),
		}
		jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, "/startTime", true))
	}
	if d.HasChange("end_time") {
		testConfigUpdate := TestConfigUpdate{
			UpdatedFieldValue: d.Get("end_time").(string),
		}
		jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, "/endTime", true))
	}
	if d.HasChange("status") {
		updated_status_id := getTestStatusTypeId(d.Get("status").(string))
		testConfigUpdate := TestConfigUpdate{
			UpdatedFieldValue: strconv.Itoa(updated_status_id),
		}
		jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, "/status/id", true))
	}

	if d.HasChange("thresholds") {
		thresholds, thresholdOk := d.GetOk("thresholds")
		if thresholdOk {
			thresholds_lists := thresholds.(*schema.Set).List()
			threshold := thresholds_lists[0].(map[string]interface{})

			setThresholds(int(test_type), threshold, &testConfig)

			testConfigUpdate := TestConfigUpdate{
				UpdatedTestThresholds: setTestThresholds(&testConfig),
				SectionToUpdate:       "/thresholdRestModel",
			}
			jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, testConfigUpdate.SectionToUpdate, false))
		}
	}

	if d.HasChange("label") {
		label, labelOk := d.GetOk("label")
		if labelOk {
			label_lists := label.(*schema.Set).List()

			setLabels(int(test_type), label_lists, &testConfig)

			testConfigUpdate := TestConfigUpdate{
				UpdatedLabels:   setTestLabels(&testConfig),
				SectionToUpdate: "/labels",
			}
			jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, testConfigUpdate.SectionToUpdate, false))
		}
	}

	if d.HasChange("advanced_settings") {
		advanced_settings, advanced_settingsOk := d.GetOk("advanced_settings")
		if advanced_settingsOk {
			advanced_setting_list := advanced_settings.(*schema.Set).List()
			advanced_setting := advanced_setting_list[0].(map[string]interface{})

			setAdvancedSettings(int(test_type), advanced_setting, &testConfig)

			testConfigUpdate := TestConfigUpdate{
				UpdatedAdvancedSettingsSection: setTestAdvancedSettings(&testConfig),
				SectionToUpdate:                "/advancedSettings",
			}
			jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, testConfigUpdate.SectionToUpdate, false))
		}
	}

	if d.HasChange("schedule_settings") {
		schedule_settings, schedule_settingsOk := d.GetOk("schedule_settings")
		if schedule_settingsOk {
			schedule_setting_list := schedule_settings.(*schema.Set).List()
			schedule_setting := schedule_setting_list[0].(map[string]interface{})

			err := setScheduleSettings(int(test_type), schedule_setting, &testConfig)
			if err != nil {
				return err
			}

			testConfigUpdate := TestConfigUpdate{
				UpdatedScheduleSettingsSection: setTestScheduleSettings(&testConfig),
				SectionToUpdate:                "/scheduleSettings",
			}
			jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, testConfigUpdate.SectionToUpdate, false))
		}
	}

	if d.HasChange("alert_settings") {
		alert_settings, alert_settingsOk := d.GetOk("alert_settings")
		if alert_settingsOk {
			alert_setting_list := alert_settings.(*schema.Set).List()
			alert_setting := alert_setting_list[0].(map[string]interface{})

			err := setAlertSettings(int(test_type), alert_setting, &testConfig)
			if err != nil {
				return err
			}

			testConfigUpdate := TestConfigUpdate{
				UpdatedAlertSettingsSection: setTestAlertSettings(&testConfig),
				SectionToUpdate:             "/alertGroup",
			}
			jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, testConfigUpdate.SectionToUpdate, false))
		}
	}

	jsonPatchDoc := "[" + strings.Join(jsonPatchDocs, ",") + "]"

	if jsonPatchDoc != "[]" {
		log.Printf("[DEBUG] Updating test: %v", testId)
		if m.(*Config).LogJson {
			log.Printf("[DEBUG] Updating test with JSON PATCH: %v", jsonPatchDoc)
		}
		respBody, respStatus, completed, err := updateTest(api_token, testId, jsonPatchDoc)
		if err != nil {
			log.Fatal(err)
		}
		if !completed {
			log.Printf("[ERROR] Error while Updating test: %v", testId)
			log.Printf("[ERROR] Error description: " + respBody)
			return errors.New(respBody)
		}
		log.Printf("[DEBUG] Response Code from Catchpoint API: " + respStatus)
		log.Print(respBody)
	}

	return resourceTracerouteTestRead(d, m)
}

func resourceTracerouteTestDelete(d *schema.ResourceData, m interface{}) error {
	testId := d.Id()
	api_token := m.(*Config).ApiToken

	log.Printf("[DEBUG] Deleting test: %v", testId)
	respBody, respStatus, completed, err := deleteTest(api_token, testId)
	if err != nil {
		log.Fatal(err)
	}
	if !completed {
		log.Printf("[ERROR] Error while deleting test: %v", testId)
		log.Printf("[ERROR] Error description: " + respBody)
		return errors.New(respBody)
	}
	log.Printf("[DEBUG] Response Code from Catchpoint API: " + respStatus)
	log.Print(respBody)

	return nil
}
