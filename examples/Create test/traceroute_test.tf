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
    start_time = "2024-04-30T04:59:00Z"
    end_time="2024-10-30T04:59:00Z"

    schedule_settings{
      frequency="6 hours"
      node_distribution ="random"
      no_of_subset_nodes = 5
      node_ids =[6388]
      node_group_ids =[9922,9848]
    }
}