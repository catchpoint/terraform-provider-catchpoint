terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "0.2.1"
    }
  }
}

provider "catchpoint" {
    api_token="5618ABF44CA1117B4286C9572D00B47AD166125E4F88A6DE472AF29562C50595"
}

resource "transaction_test" "testTransaction" {
    test_name  = "Transaction_TF"
    provider=catchpoint
    division_id=2633
    product_id=23791
    monitor="chrome"
    test_script="// Step - 1\r\nopen(\"\\\"https:www.google.com)\")"
    start_time = "2024-04-30T04:59:00Z"
    end_time="2024-10-30T04:59:00Z"

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

    alert_settings {
        alert_rule {
            alert_type="timing"
            alert_sub_type="response"
            node_threshold_type="node"
            threshold_number_of_runs=5
            threshold_interval="30 minutes"
            trigger_type="specific value"
            warning_trigger=50
            critical_trigger=70.0
            operation_type = "less than or equals"
            use_rolling_window=true
            notification_group {
              notify_on_critical = true
              subject = "contact group testing"
              recipient_email_ids = ["vkumar@catchpoint.com"]
              contact_groups = ["Agent deployment"]
            }
        }
        alert_rule {
            alert_type="availability"
            alert_sub_type="test"
            node_threshold_type="average across nodes"
            threshold_number_of_runs=3
            trigger_type="trailing value"
            warning_trigger=50
            critical_trigger=80
            historical_interval="15 minutes"
            operation_type = "greater than"
            notification_group {
              notify_on_critical = true
              subject = "contact group testing"
              recipient_email_ids = ["vkumar@catchpoint.com"]
              contact_groups = ["Agent deployment"]
            }
        }
        notification_group {
              notify_on_critical = true
              subject = "contact group testing"
              recipient_email_ids = ["vkumar@catchpoint.com"]
              contact_groups = ["Agent deployment"]
            }
    }
}