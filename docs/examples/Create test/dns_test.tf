terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "0.2.5"
    }
  }
}

provider "catchpoint" {
    api_token="5618ABF44CA1117B4286CXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

resource "dns_test" "testDNS" {
    test_name  = "Terraform DNS Test 0.2.5"
    provider=catchpoint
    division_id=2633
    product_id=23791
    monitor="dns direct"
    query_type="a"
    test_domain ="https:www.google.com"
    end_time="2023-11-30T04:59:00Z"
}