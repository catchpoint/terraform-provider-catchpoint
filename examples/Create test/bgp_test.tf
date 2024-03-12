terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "0.2.1"
    }
  }
}

 

provider "catchpoint" {
  api_token="5618ABF44CA1117B4286C9572XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

 

resource "bgp_test" "test" {
    test_name  = "BGP test created from Terraform"
    provider=catchpoint
    division_id=2633
    product_id=23791
    prefix="101.188.67.134/8"
    end_time="2023-09-30T04:59:00Z"
}