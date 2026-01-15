terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "2.0.0"
    }
  }
}

resource "puppeteer_test" "minimal_puppeteertest" {
  provider    = catchpoint
  test_name   = "minimal puppeteer test"
  test_script = <<-EOT
    // Step - 1
    await page.goto('https://www.example.com/');
EOT
  division_id = 2633
  product_id  = 25232

  insights { insight_setting_type = "inherit" }
  request_settings { request_setting_type = "inherit" }
  schedule_settings { schedule_setting_type = "inherit" }
  advanced_settings { advanced_setting_type = "inherit" }
  alert_settings { alert_setting_type = "inherit" }
}

resource "puppeteer_test" "test" {
  test_name                = "puppeteer_test"
  provider                 = catchpoint
  test_script              = <<-EOT
    // Step - 1
    await page.goto('https://www.example.com/');
EOT
  division_id              = 2633
  product_id               = 25232
  test_description         = "A puppeteer test for example.com"
  alerts_paused            = false
  enable_test_data_webhook = true
  start_time               = "2024-04-30T04:59:00Z"
  end_time                 = "2024-10-30T04:59:00Z"
  gateway_address_or_host  = "192.168.0.11"

  thresholds {
    test_time_critical    = 250
    test_time_warning     = 50
    availability_critical = 80
    availability_warning  = 90
  }

  advanced_settings {
    additional_monitor                       = "ping icmp"
    allow_test_download_limit_override       = true
    bandwidth_throttling                     = "gprs"
    capture_filmstrip                        = true
    debug_primary_host_on_failure            = true
    debug_referenced_hosts_on_failure        = true
    disable_cross_origin_iframe_access       = true
    enable_path_mtu_discovery                = true
    enable_self_versus_third_party_zones     = true
    enforce_test_failure_if_runs_longer_than = 30
    f40x_or_50x_http_mark_successful         = true
    host_data_collection_enabled             = true
    ignore_ssl_failures                      = true
    stop_test_on_document_complete           = true
    verify_test_on_failure                   = true
    viewport_height                          = 800
    viewport_width                           = 400
    wait_for_no_activity                     = 0
    zone_data_collection_enabled             = true
  }

  label {
    key    = "label1"
    values = ["v1", "v2"]
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
      alert_type                 = "timing"
      alert_sub_type             = "response"
      node_threshold_type        = "node"
      threshold_number_of_runs   = 5
      threshold_interval         = "30 minutes"
      trigger_type               = "specific value"
      warning_trigger            = 50
      critical_trigger           = 70.0
      operation_type             = "less than or equals"
      use_rolling_window         = true
      enable_consecutive         = true
      consecutive_number_of_runs = 5
      notification_group {
        notify_on_critical = true
        subject            = "contact group testing"
        recipient_emails   = ["vkumar@catchpoint.com"]
        contact_group_ids  = [1234]
      }
    }
    notification_group {
      subject           = "testing"
      recipient_emails  = ["vikash@catchpoint.com"]
      contact_group_ids = [1234]
    }
  }
}

resource "puppeteer_test" "imported_puppeteertest" {
  test_name   = "imported puppeteer test"
  provider    = catchpoint
  test_script = <<-EOT
    // Step - 1
    await page.goto('https://www.example.com/');
EOT
  division_id = 2633
  product_id  = 25232

  insights { insight_setting_type = "inherit" }
  request_settings { request_setting_type = "inherit" }
  schedule_settings { schedule_setting_type = "inherit" }
  advanced_settings { advanced_setting_type = "inherit" }
  alert_settings { alert_setting_type = "inherit" }
}
