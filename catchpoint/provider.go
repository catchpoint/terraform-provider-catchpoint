package catchpoint

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Catchpoint API v2 token",
				DefaultFunc: schema.EnvDefaultFunc("CATCHPOINT_API_TOKEN", nil),
			},
			"log_json": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Enable or disable test json payload logging for debugging. Accepts string and converts to bool using ParseBool function",
				DefaultFunc: schema.EnvDefaultFunc("LOG_JSON", nil),
			},
			"catchpoint_environment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Set the environment to stage, qa or prod. This is for internal use",
				DefaultFunc: schema.EnvDefaultFunc("CATCHPOINT_ENVIRONMENT", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"web_test":         resourceWebTestType(),
			"api_test":         resourceApiTestType(),
			"transaction_test": resourceTransactionTestType(),
			"traceroute_test":  resourceTracerouteTestType(),
			"ping_test":        resourcePingTestType(),
			"bgp_test":         resourceBgpTestType(),
			"dns_test":         resourceDnsTestType(),
			"ssl_test":         resourceSslTestType(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	api_token := d.Get("api_token").(string)
	log_json := d.Get("log_json").(string)
	is_log_json, err := strconv.ParseBool(log_json)
	if err != nil || log_json == "" {
		is_log_json = false
	}
	catchpoint_environment := d.Get("catchpoint_environment").(string)
	setTestUriByEnv(catchpoint_environment)
	return newConfig(api_token, is_log_json, catchpoint_environment), nil
}
