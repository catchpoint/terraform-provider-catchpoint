terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "0.2.1"
    }
  }
}

 

provider "catchpoint" {
api_token="5618ABF44CA1117B42XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

 

 

resource "traceroute_test" "testTraceroute" {
    test_name  = "Tracetest_TF"
    provider=catchpoint
    division_id=2633
    product_id=23791
    monitor="traceroute icmp"
    test_location="https://www.google.com"
    end_time="2023-10-30T04:59:00Z"
}