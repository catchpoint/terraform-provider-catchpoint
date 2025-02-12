terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "1.5.0"
    }
  }
}

provider "catchpoint" {
api_token="5618ABF44CA1117B4286C9572XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

resource "manage_folder" "folder"{
  provider = catchpoint
  division_id = 3063
  product_id=42631
  folder_name = "folder"
  schedule_settings{
      frequency="4 hours"
      node_distribution ="concurrent"
      no_of_subset_nodes = 5
      run_schedule_id = 50429
      maintenance_schedule_id = 3517
      node_ids =[6008,6009]
      node_group_ids =[8095,8158]
    }
    request_settings {
      authentication {
        authentication_type = "basic"
        password_ids = [45696]
      }
      token_ids = [55836]
      http_request_headers {
        user_agent {
          value = "amit"
        }
      }
    }
     insights {
      indicator_ids =[7369,7755]
      tracepoint_ids =[7366,7367]
    }
   advanced_settings {
        verify_test_on_failure = true
        additional_monitor="ping icmp"
      }
     alert_settings{
        alert_rule {
            alert_type="test failure"
            node_threshold_type="runs"
            threshold_number_of_runs=2
            threshold_percentage_of_runs=60
            critical_reminder = "15 minutes"
            omit_scatterplot = false
            enable_consecutive = false
            threshold_interval ="10 minutes"
            notification_group {
          subject = "testing24"
          recipient_email_ids=["vikash@catchpoint.com"]
        }
        }
        alert_rule {
            alert_type="host failure"
            node_threshold_type="runs"
            threshold_number_of_runs=2
            threshold_percentage_of_runs=60
            critical_reminder = "15 minutes"
            omit_scatterplot = false
            threshold_interval ="10 minutes"
            enable_consecutive = "false"
            notification_group {
          subject = "testing24"
          recipient_email_ids=["vikash@catchpoint.com"]
        }
        }
        alert_rule {
            alert_type="content match"
            alert_sub_type="response headers"
            node_threshold_type="runs"
            operation_type = "not equals"
            expression="something"
            critical_reminder = "10 minutes"
            omit_scatterplot = true
            threshold_number_of_runs=5
            notification_group {
          subject = "testing24"
          recipient_email_ids=["vikash@catchpoint.com"]
        }
        }
        notification_group {
          subject = "testing24"
          recipient_email_ids=["vikash@catchpoint.com"]
        }
    }
}