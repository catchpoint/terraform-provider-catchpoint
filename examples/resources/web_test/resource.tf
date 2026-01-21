terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "2.0.0"
    }
  }
}

resource "web_test" "minimal_web_test" {
  provider    = catchpoint
  monitor     = "http"
  test_name   = "Web minimal test"
  test_url    = "https://www.example.com"
  division_id = 2923
  product_id  = 28335

  insights { insight_setting_type = "inherit" }
  request_settings { request_setting_type = "inherit" }
  schedule_settings { schedule_setting_type = "inherit" }
  advanced_settings { advanced_setting_type = "inherit" }
  alert_settings { alert_setting_type = "inherit" }
}

resource "web_test" "test" {
  test_name                = "Web Test created from Terraform"
  monitor                  = "http"
  provider                 = catchpoint
  division_id              = 2923
  product_id               = 28335
  test_description         = "An http/object test for example.com"
  test_url                 = "https://www.example.com"
  alerts_paused            = false
  enable_test_data_webhook = true
  start_time               = "2024-04-30T04:59:00Z"
  end_time                 = "2024-10-30T04:59:00Z"

  label {
    key    = "label1"
    values = ["v1", "v2"]
  }

  insights { insight_setting_type = "inherit" }

  advanced_settings {
    additional_monitor                       = "ping icmp"
    bandwidth_throttling                     = "gprs"
    capture_filmstrip                        = true
    capture_screenshot                       = true
    disable_cross_origin_iframe_access       = true
    viewport_height                          = 800
    viewport_width                           = 400
    allow_test_download_limit_override       = true
    capture_response_content                 = true
    capture_http_headers                     = true
    debug_primary_host_on_failure            = true
    debug_referenced_hosts_on_failure        = true
    enable_path_mtu_discovery                = true
    enable_self_versus_third_party_zones     = true
    enforce_test_failure_if_runs_longer_than = 120
    f40x_or_50x_http_mark_successful         = true
    host_data_collection_enabled             = true
    ignore_ssl_failures                      = true
    verify_test_on_failure                   = true
    zone_data_collection_enabled             = true
    stop_test_on_document_complete           = true
    wait_for_no_activity                     = 500
  }

  request_settings {
    authentication {
      authentication_type = "basic"
      password_ids        = [2332]
    }
    token_ids = [1096]
    http_request_headers {
      user_agent {
        value = "vikash"
      }
    }
  }

  schedule_settings {
    frequency          = "6 hours"
    node_distribution  = "random"
    no_of_subset_nodes = 5
    node_ids           = [6388]
    node_group_ids     = [9922, 9848]
  }

  alert_settings {
    alert_rule {
      alert_type               = "timing"
      alert_sub_type           = "response"
      node_threshold_type      = "node"
      threshold_number_of_runs = 5
      threshold_interval       = "30 minutes"
      trigger_type             = "specific value"
      warning_trigger          = 50
      critical_trigger         = 70.0
      operation_type           = "less than or equals"
      use_rolling_window       = true
      notification_group {
        notify_on_critical = true
        subject            = "contact group testing"
        recipient_emails   = ["vkumar@catchpoint.com"]
        contact_group_ids  = [1234]
      }
    }
    alert_rule {
      alert_type               = "availability"
      alert_sub_type           = "test"
      node_threshold_type      = "average across nodes"
      threshold_number_of_runs = 3
      trigger_type             = "trailing value"
      warning_trigger          = 50
      critical_trigger         = 80
      historical_interval      = "15 minutes"
      operation_type           = "greater than"
      notification_group {
        notify_on_critical = true
        subject            = "contact group testing"
        recipient_emails   = ["vkumar@catchpoint.com"]
        contact_group_ids  = [1234]
      }
    }
    notification_group {
      notify_on_critical = true
      subject            = "contact group testing"
      recipient_emails   = ["vkumar@catchpoint.com"]
      contact_group_ids  = [1234]
    }
  }
}

resource "web_test" "imported_webtest" {
  test_name   = "Web Test imported into Terraform"
  monitor     = "http"
  provider    = catchpoint
  division_id = 2923
  product_id  = 28335
  test_url    = "https://www.example.com"

  insights { insight_setting_type = "inherit" }
  request_settings { request_setting_type = "inherit" }
  schedule_settings { schedule_setting_type = "inherit" }
  advanced_settings { advanced_setting_type = "inherit" }
  alert_settings { alert_setting_type = "inherit" }
}
