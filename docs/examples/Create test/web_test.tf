terraform {
  required_providers {
    catchpoint = {
      source = "catchpoint/catchpoint"
      version = "0.2.6"
    }
  }
}

provider "catchpoint" {
  api_token="ABAA8C66AE593EDCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}


resource "web_test" "test" {
    test_name  = "Web_test1 from terraform "
    monitor="object"
    provider=catchpoint
    division_id=2923
    product_id=28335
    test_description="An object test for catchpoint.com"
    test_url="https://www.catchpoint.com"
    alerts_paused=false
    enable_test_data_webhook=true
    end_time="2023-12-30T04:59:00Z"

    label {
        key="label1"
        values=["v1","v2"]
    }

    advanced_settings {
        verify_test_on_failure = false
        additional_monitor="ping icmp"
    }

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
        }
        alert_rule {
            alert_type="test failure"
            node_threshold_type="runs"
            threshold_number_of_runs=2
            threshold_percentage_of_runs=60
        }
        alert_rule {
            alert_type="host failure"
            node_threshold_type="runs"
            threshold_number_of_runs=5
            threshold_percentage_of_runs=90
        }
        alert_rule {
            alert_type="content match"
            alert_sub_type="response headers"
            node_threshold_type="runs"
            operation_type = "not equals"
            expression="something"
            threshold_number_of_runs=5
        }
        notification_group {
            recipient_email_ids=["vikash@catchpoint.com"]
        }
    }
}
