terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "0.2.1"
    }
  }
}

 

provider "catchpoint" {
api_token="5618ABF44CA1117B428XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

resource "ping_test" "test5501" {
  test_name  = "Ping_TF3"
  provider=catchpoint
  division_id=2633
  product_id=23791
  test_location="https:www.google.com"
  monitor ="ping tcp"
  status="active"
  end_time="2023-10-30T04:59:00Z"
  schedule_settings{
      frequency="6 hours"
      node_distribution ="random"
      no_of_subset_nodes = 5
      node_ids =[6388]
      node_group_ids =[9922,9848]
    }
}