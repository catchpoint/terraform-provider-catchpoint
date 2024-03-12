terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "0.2.1"
    }
  }
}

 

provider "catchpoint" {
api_token="5618ABF44CA1117B42XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

 

 

resource "api_test" "test55" {
  test_name  = "APItest_TF"
  provider=catchpoint
  division_id=2633
  product_id=23791
  test_script="// Step - 1\r\nopen(\"\\\"https:www.google.com)\")"
  test_script_type="selenium"
  end_time="2023-10-30T04:59:00Z"
  request_settings {
        authentication {
          authentication_type = "basic"
          password_ids = [2332]
        }
        token_ids = [1096]
        http_request_headers {
          user_agent {
            value = "vikash"
          }
        }
      }

      schedule_settings{
      frequency="6 hours"
      node_distribution ="random"
      no_of_subset_nodes = 5
      node_ids =[6388]
      node_group_ids =[9922,9848]
    }
}