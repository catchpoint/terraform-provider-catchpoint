terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "2.0.0"
    }
  }
}

resource "bgp_test" "minimal_bgp_standard" {
  provider    = catchpoint
  monitor     = "bgp"
  division_id = 2633
  product_id  = 23791
  test_name   = "BGP Standard Test - 192.168.0.0/24"
  prefix      = "192.168.0.0/24"

  alert_settings { alert_setting_type = "inherit" }
}

resource "bgp_test" "test" {
  test_name                = "BGP test created from Terraform"
  monitor                  = "bgp"
  provider                 = catchpoint
  division_id              = 2633
  product_id               = 23791
  prefix                   = "101.188.67.134/8"
  start_time               = "2024-04-30T04:59:00Z"
  end_time                 = "2024-10-30T04:59:00Z"
  enable_test_data_webhook = true
  alerts_paused            = false

  alert_settings {
    alert_setting_type = "inherit & add"
    alert_rule {
      alert_type = "availability"
      alert_sub_type = "% downtime"
      node_threshold_type = "node"
      all_match_records = false
      critical_reminder = "none"
      warning_reminder = "none"
      enable_consecutive = false
      enforce_test_failure = false
      notification_type = "default contacts"
      omit_scatterplot = false
      threshold_interval = "2 hours"
      threshold_number_of_runs = 1
      trigger_type = "specific value"
      use_rolling_window = true
      warning_trigger = 10
    }
    notification_group {
      subject          = "$${NotificationLevel}:  test=#$${TestId} - $${TestName}, alert=$${AlertType}"
      recipient_emails = ["jmitchell@catchpoint.com"]
    }
  }
  label {
    key    = "automation"
    values = ["terraform"]
  }
}

resource "bgp_test" "imported_bgptest" {
  provider    = catchpoint
  test_name   = "BGP test imported into Terraform"
  monitor     = "bgp"
  division_id = 2633
  product_id  = 23791
  prefix      = "192.168.0.0/24"

  alert_settings { alert_setting_type = "inherit" }
}
