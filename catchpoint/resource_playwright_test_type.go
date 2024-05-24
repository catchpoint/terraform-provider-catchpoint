package catchpoint

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePlaywrightTestType() *schema.Resource {
	return &schema.Resource{
		Create: resourcePlaywrightTestCreate,
		Read:   resourcePlaywrightTestRead,
		Update: resourcePlaywrightTestUpdate,
		Delete: resourcePlaywrightTestDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"monitor": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The monitor to use for the Playwright Test. Supported: 'playwright', 'chrome'",
				Default:      "playwright",
				ValidateFunc: validation.StringInSlice([]string{"playwright", "chrome"}, false),
			},
			"simulate": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "The device to simulate for mobile monitor",
				ValidateFunc: validation.StringInSlice([]string{"android", "iphone", "ipad 2", "kindle fire", "galaxy tab", "iphone 5", "ipad mini", "galaxy note", "nexus 7", "nexus 4", "nokia lumia920", "iphone 6", "blackberry z30", "galaxy s4", "htc onex", "lg optimusg", "droid razr hd", "nexus 6", "iphone 6s", "galaxy s6", "iphone 7", "google pixel", "galaxy s8"}, false),
			},
			"chrome_version": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Chrome version to use. Supported: 'preview', 'stable', '108', '89', '87', '85', '75', '71', '66', '63', '59', '53'",
				ValidateFunc: validation.StringInSlice([]string{"preview", "stable", "108", "89", "87", "85", "75", "71", "66", "63", "59", "53"}, false),
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
				Computed:    true,
				Description: "Optional. The Folder under which the Test will be created",
			},
			"test_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the Test",
			},
			"test_description": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Optional. The Test description",
			},
			"test_script": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Script that will simulate user workflow",
			},
			"test_script_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The type of script. Supported: 'playwright'",
				ValidateFunc: validation.StringInSlice([]string{"playwright"}, false),
				Default:      "playwright",
			},
			"gateway_address_or_host": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Host/IP to use for network troubleshooting and monitoring",
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
				Computed:    true,
				Description: "Start time for the Test in ISO format like 2024-12-30T04:59:00Z",
			},
			"end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "End time for the Test in ISO format like 2024-12-30T04:59:00Z",
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Test status: active or inactive",
				ValidateFunc: validation.StringInSlice([]string{"active", "inactive"}, false),
			},
			"label": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
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
			"request_settings": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "Optional. Used for overriding authentication and HTTP request headers",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:     schema.TypeSet,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication_type": {
										Type:         schema.TypeString,
										Required:     true,
										Description:  "Type of authentication to use 'basic', 'ntlm', 'digest', 'login'",
										ValidateFunc: validation.StringInSlice([]string{"basic", "ntlm", "digest", "login"}, false),
									},
									"password_ids": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Optional. Password ids in a list",
										Sensitive:   true,
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
								},
							},
						},
						"token_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Optional. Token ids in a list",
							Sensitive:   true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"http_request_headers": {
							Type:     schema.TypeSet,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"user_agent": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the user agent header for test url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"accept": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the accept header for test url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"accept_encoding": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the user accept encoding header for test url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"accept_language": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the accept language header for test url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"accept_charset": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the accept charset header for test url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"cookie": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the cookie header for test url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"cache_control": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the cache control header for test url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"pragma": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the pragma header for test url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"referer": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the referer header for test url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"host": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the host header for test url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"dns_override": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the dns override header for the given child_host_pattern",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},
									"request_override": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the request override header for test url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"request_block": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the request block header for test url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"request_delay": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the request delay header for test url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"insights": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "Optional. Used for overriding the insights section",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tracepoint_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Optional. Tracepoint ids in a list",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"indicator_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Optional. Indicator ids in a list",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
					},
				},
			},
			"schedule_settings": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
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
										Description:  "Optional. Sets the operation type: 'not equals', 'greater than', 'greater than or equals', 'less than', 'less than or equals'",
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
										Description: "Optional. Sets trigger expression for content match alert type ",
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
										Description:  "Sets the alert type: 'test failure', 'timing', 'availability'",
										Required:     true,
										ValidateFunc: validation.StringInSlice([]string{"test failure", "timing", "availability", "host failure", "requests", "content match", "byte length"}, false),
									},
									"alert_sub_type": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "Optional. Sets the sub alert type: 'dns', 'connect', 'send', 'wait', 'load', 'ttfb', 'content load', 'response', 'test time', 'dom load', 'test time with suspect', 'server response', 'document complete', 'redirect', 'test', 'content', '% downtime'",
										ValidateFunc: validation.StringInSlice([]string{"dns", "connect", "send", "wait", "load", "ttfb", "content load", "response", "test time", "dom load", "test time with suspect", "server response", "document complete", "redirect", "test", "content", "% downtime", "# requests", "# hosts", "# connections", "# redirects", "# other", "# images", "# scripts", "# html", "# css", "# xml", "# flash", "# media", "regular expression", "response code", "response headers", "byte length", "page", "file size"}, false),
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
										Required:    true,
										MaxItems:    5,
										Description: "List of Notification groups for configuring alert notifications, including recipients' email addresses and alert settings. To ensure either recipient_email_ids or contact_groups is provided ",
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
				Computed:    true,
				Description: "Optional. Used for overriding the advanced settings",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"verify_test_on_failure": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables verify on test failure setting",
							Optional:    true,
							Default:     false,
						},
						"debug_primary_host_on_failure": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables debug primary host on failure setting",
							Optional:    true,
							Default:     false,
						},
						"enable_http2": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables enable http/2 setting",
							Optional:    true,
							Default:     false,
						},
						"debug_referenced_hosts_on_failure": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables debug referenced hosts on failure setting",
							Optional:    true,
							Default:     false,
						},
						"capture_http_headers": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables capture http headers setting for all runs",
							Optional:    true,
							Default:     false,
						},
						"capture_response_content": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables capture response content setting for all runs",
							Optional:    true,
							Default:     false,
						},
						"ignore_ssl_failures": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables ignore SSL failures setting",
							Optional:    true,
							Default:     false,
						},
						"host_data_collection_enabled": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables host data collection setting",
							Optional:    true,
							Default:     false,
						},
						"zone_data_collection_enabled": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables zone data collection setting",
							Optional:    true,
							Default:     false,
						},
						"f40x_or_50x_http_mark_successful": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables 40x or 50x error mark successful setting",
							Optional:    true,
							Default:     false,
						},
						"enable_self_versus_third_party_zones": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables self versus third party zones setting and matches self zone by test URL",
							Optional:    true,
							Default:     false,
						},
						"allow_test_download_limit_override": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables test download limit override setting",
							Optional:    true,
							Default:     false,
						},
						"capture_filmstrip": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables capture filmstrip setting",
							Optional:    true,
							Default:     false,
						},
						"capture_screenshot": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables capture screenshot setting for all runs",
							Optional:    true,
							Default:     false,
						},
						"stop_test_on_document_complete": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables stop test on document complete setting",
							Optional:    true,
							Default:     false,
						},
						"disable_cross_origin_iframe_access": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables disable cross origin iframe access setting for chrome monitor",
							Optional:    true,
							Default:     false,
						},
						"stop_test_on_dom_content_load": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables stop test on DOM content load setting",
							Optional:    true,
							Default:     false,
						},
						"enforce_test_failure_if_runs_longer_than": {
							Type:         schema.TypeInt,
							Description:  "Optional. Set the time value in seconds post which the test will be marked as failure.",
							ValidateFunc: validation.IntInSlice([]int{5, 10, 15, 20, 30, 60, 90, 120}),
							Optional:     true,
						},
						"wait_for_no_activity": {
							Type:         schema.TypeInt,
							Description:  "Optional. Set the time value in ms to stop the test after no network activity on document complete. Use with stop_test_on_document_complete flag",
							ValidateFunc: validation.IntInSlice([]int{0, 500, 1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000}),
							Optional:     true,
						},
						"viewport_height": {
							Type:        schema.TypeInt,
							Description: "Optional. Set the viewport height. Use with viewport_width attribute",
							Optional:    true,
						},
						"viewport_width": {
							Type:        schema.TypeInt,
							Description: "Optional. Set the viewport width. Use with viewport_height attribute",
							Optional:    true,
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
						"enable_path_mtu_discovery": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables Path MTU Discovery",
							Optional:    true,
							Default:     false,
						},
					},
				},
			},
		},
	}
}

func resourcePlaywrightTestCreate(d *schema.ResourceData, m interface{}) error {
	api_token := m.(*Config).ApiToken
	monitor := d.Get("monitor").(string)
	monitor_id := getMonitorId(monitor)
	simulate_device := d.Get("simulate").(string)
	simulate_device_id := getUserAgentTypeId(simulate_device)
	chrome_version := d.Get("chrome_version").(string)
	chrome_version_id, chrome_version_name := getChromeVersionId(chrome_version)
	var application_version_id int
	var application_version_name string
	if chrome_version_id == 3 {
		application_version_id, application_version_name = getChromeApplicationVersionId(chrome_version)
	}
	if monitor == "chrome" && chrome_version == "" {
		//default id 1 : stable for chrome monitor if chrome version attribute is not set
		chrome_version_id = 1
	}
	if monitor == "mobile" && simulate_device == "" {
		//default id 3 : android for mobile monitor if simulate device attribute is not set
		simulate_device_id = 3
	}
	division_id := d.Get("division_id").(int)
	product_id := d.Get("product_id").(int)
	folder_id := d.Get("folder_id").(int)
	test_name := d.Get("test_name").(string)
	test_script := d.Get("test_script").(string)
	test_script_type := d.Get("test_script_type").(string)
	test_script_type_id := getApiScriptTypeId(test_script_type)
	test_description := d.Get("test_description").(string)
	gateway_address_or_host := d.Get("gateway_address_or_host").(string)
	enable_test_data_webhook := d.Get("enable_test_data_webhook").(bool)
	alerts_paused := d.Get("alerts_paused").(bool)
	start_time := d.Get("start_time").(string)
	if start_time == "" {
		start_time = getTime()
	}
	end_time := d.Get("end_time").(string)
	status := d.Get("status").(string)
	status_id := getTestStatusTypeId(status)
	test_type := TestType(Playwright)

	var testConfig = TestConfig{}

	testConfig = TestConfig{
		TestType:                 int(test_type),
		Monitor:                  monitor_id,
		SimulateDevice:           simulate_device_id,
		ChromeVersion:            IdName{Id: chrome_version_id, Name: chrome_version_name},
		ChromeApplicationVersion: IdName{Id: application_version_id, Name: application_version_name},
		DivisionId:               division_id,
		ProductId:                product_id,
		FolderId:                 folder_id,
		TestName:                 test_name,
		TestDescription:          test_description,
		GatewayAddressOrHost:     gateway_address_or_host,
		EnableTestDataWebhook:    enable_test_data_webhook,
		AlertsPaused:             alerts_paused,
		StartTime:                start_time,
		EndTime:                  end_time,
		TestStatus:               status_id,
	}

	setRequestData(int(test_type), test_script, monitor_id, test_script_type_id, &testConfig)
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

	request_settings, request_settingsOk := d.GetOk("request_settings")
	if request_settingsOk {
		request_settings_list := request_settings.(*schema.Set).List()
		request_setting := request_settings_list[0].(map[string]interface{})

		err := setRequestSettings(int(test_type), request_setting, &testConfig)
		if err != nil {
			return err
		}
	}

	insight_settings, insight_settingsOk := d.GetOk("insights")
	if insight_settingsOk {
		insight_setting_list := insight_settings.(*schema.Set).List()
		insight_setting := insight_setting_list[0].(map[string]interface{})

		setInsightSettings(int(test_type), insight_setting, &testConfig)
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
	return resourcePlaywrightTestRead(d, m)
}

func resourcePlaywrightTestRead(d *schema.ResourceData, m interface{}) error {
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
	d.Set("simulate", testNew["simulate"])
	d.Set("chrome_version", testNew["chrome_version"])
	d.Set("division_id", testNew["division_id"])
	d.Set("product_id", testNew["product_id"])
	d.Set("folder_id", testNew["folder_id"])
	d.Set("test_name", testNew["test_name"])
	d.Set("test_description", testNew["test_description"])
	d.Set("gateway_address_or_host", testNew["gateway_address_or_host"])
	d.Set("enable_test_data_webhook", testNew["enable_test_data_webhook"])
	d.Set("alerts_paused", testNew["alerts_paused"])
	d.Set("start_time", testNew["start_time"])
	d.Set("end_time", testNew["end_time"])
	d.Set("status", testNew["status"])
	d.Set("test_script", testNew["test_script"])
	d.Set("test_script_type", testNew["test_script_type"])
	d.Set("label", testNew["label"])
	d.Set("thresholds", testNew["thresholds"])
	d.Set("request_settings", testNew["request_settings"])
	d.Set("insights", testNew["insights"])
	d.Set("schedule_settings", testNew["schedule_settings"])
	d.Set("alert_settings", testNew["alert_settings"])
	d.Set("advanced_settings", testNew["advanced_settings"])

	return nil
}

func resourcePlaywrightTestUpdate(d *schema.ResourceData, m interface{}) error {
	testId := d.Id()
	api_token := m.(*Config).ApiToken
	test_type := TestType(Playwright)
	var testConfig = TestConfig{}
	var jsonPatchDocs = []string{}

	if d.HasChange("test_name") {
		testConfigUpdate := TestConfigUpdate{
			UpdatedFieldValue: d.Get("test_name").(string),
		}
		jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, "/name", true))
	}
	if d.HasChange("test_description") {
		testConfigUpdate := TestConfigUpdate{
			UpdatedFieldValue: d.Get("test_description").(string),
		}
		jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, "/description", true))
	}
	if d.HasChange("chrome_version") {
		chrome_version := d.Get("chrome_version").(string)
		chrome_version_id, _ := getChromeVersionId(chrome_version)
		// Specific chrome version was provided
		if chrome_version_id == 3 {
			application_version_id, _ := getChromeApplicationVersionId(chrome_version)
			testConfigUpdate := TestConfigUpdate{
				UpdatedFieldValue: strconv.Itoa(application_version_id),
			}
			jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, "/chromeMonitorVersion/applicationVersionId", true))

		} else {
			testConfigUpdate := TestConfigUpdate{
				UpdatedFieldValue: strconv.Itoa(chrome_version_id),
			}
			jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, "/chromeMonitorVersion/applicationVersionType", true))
		}
	}
	if d.HasChange("monitor") {
		monitor := d.Get("monitor").(string)
		monitor_id := getMonitorId(monitor)
		testConfigUpdate := TestConfigUpdate{
			UpdatedFieldValue: strconv.Itoa(monitor_id),
		}
		jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, "/monitor", true))
	}
	if d.HasChange("gateway_address_or_host") {
		testConfigUpdate := TestConfigUpdate{
			UpdatedFieldValue: d.Get("gateway_address_or_host").(string),
		}
		jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, "/gatewayAddressOrHost", true))
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
		jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, "/status", true))
	}

	if d.HasChange("test_script") {
		test_script_type_id := getApiScriptTypeId(d.Get("test_script_type").(string))
		monitor_id := getMonitorId(d.Get("monitor").(string))
		setRequestData(int(test_type), d.Get("test_script").(string), monitor_id, test_script_type_id, &testConfig)

		testConfigUpdate := TestConfigUpdate{
			UpdatedFieldValue:      d.Get("test_script").(string),
			UpdatedTestRequestData: setTestRequestData(&testConfig),
			SectionToUpdate:        "/testRequestData",
		}
		jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, testConfigUpdate.SectionToUpdate, false))
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

	if d.HasChange("request_settings") {
		request_settings, request_settingsOk := d.GetOk("request_settings")
		if request_settingsOk {
			request_settings_list := request_settings.(*schema.Set).List()
			request_setting := request_settings_list[0].(map[string]interface{})

			err := setRequestSettings(int(test_type), request_setting, &testConfig)
			if err != nil {
				return err
			}
			testConfigUpdate := TestConfigUpdate{
				UpdatedRequestSettingsSection: setTestRequestSettings(&testConfig),
				SectionToUpdate:               "/requestSettings",
			}
			jsonPatchDocs = append(jsonPatchDocs, createJsonPatchDocument(testConfigUpdate, testConfigUpdate.SectionToUpdate, false))
		}
	}

	if d.HasChange("insights") {
		insight_settings, insight_settingsOk := d.GetOk("insights")
		if insight_settingsOk {
			insight_setting_list := insight_settings.(*schema.Set).List()
			insight_setting := insight_setting_list[0].(map[string]interface{})

			setInsightSettings(int(test_type), insight_setting, &testConfig)

			testConfigUpdate := TestConfigUpdate{
				UpdatedInsightSettingsSection: setTestInsightSettings(&testConfig),
				SectionToUpdate:               "/insightData",
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

		return resourcePlaywrightTestRead(d, m)
	} else {
		return errors.New("no changes. Your infrastructure matches the configuration")
	}
}

func resourcePlaywrightTestDelete(d *schema.ResourceData, m interface{}) error {
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
