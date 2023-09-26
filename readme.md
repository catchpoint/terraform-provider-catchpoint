# Catchpoint Provider

With the Catchpoint Terraform provider, users can manage Web, API, Transaction, DNS, SSL, BGP, Traceroute, and Ping tests with minimum configurations. To inherit sections such as Scheduling, Advanced Settings, Insights, and Alerts, simply omit the corresponding blocks.

## How to run

- Create the main.tf file with the Catchpoint API v2 Key
`provider "catchpoint" {api_token="8D7E2B17C339AD3EDCDD0761D29AXXXXXXXXXXXXX"}`
- Define your resource e.g.
`resource "web_test" "test" {
    provider=catchpoint
    division_id=1000
    product_id=15330
    test_name="catchpointTf"
    test_url="https://www.catchpoint.com"
    enable_test_data_webhook=true
    end_time="2023-09-05T04:59:00Z"
}`
- In the terminal, run terraform init to initialize the provider.
- Run terraform plan to see the changes to be applied.
- Run terraform apply to apply the plan. If Terraform Debug logs are enabled, the API response can be seen from the Catchpoint provider. The logs will indicate whether the test was created successfully.
