# Catchpoint Provider

With the Catchpoint Terraform provider, users can manage Web, API, Transaction, DNS, SSL, BGP, Traceroute, and Ping tests with minimum configurations. To explicitly inherit sections - such as Scheduling, Advanced Settings, Insights, and Alerts - set the block 'x_setting_type' to inherit. See example below.
Note: not all test types have all 5 nested blocks. 

## How to run

- Create the main.tf file with the Catchpoint API v2 Key
`provider "catchpoint" {api_token="*YOUR_TOKEN*"}`

- Define your resource e.g.
```
resource "web_test" "test" {
    provider=catchpoint
    division_id=1000
    product_id=15330
    test_name="catchpointTf"
    test_url="https://www.catchpoint.com"
    enable_test_data_webhook=true
    end_time="2023-09-05T04:59:00Z"
    
    insights { insight_setting_type = "inherit" }
    request_settings { request_setting_type = "inherit" }
    schedule_settings { schedule_setting_type = "inherit" }
    advanced_settings { advanced_setting_type = "inherit" }
    alert_settings { alert_setting_type = "inherit" }
}
```

- In the terminal, run terraform init to initialize the provider.
- Run terraform plan to see the changes to be applied.
- Run terraform apply to apply the plan. If Terraform Debug logs are enabled, the API response can be seen from the Catchpoint provider. The logs will indicate whether the test was created successfully.
