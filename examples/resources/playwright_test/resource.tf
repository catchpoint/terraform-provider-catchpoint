terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "2.0.0"
    }
  }
}

resource "playwright_test" "minimal_playwright_test" {
  product_id  = 2345
  monitor     = "chrome"
  provider    = catchpoint
  division_id = 1234
  test_name   = "A minimal Playwright Test for https://www.example.com"
  test_script = <<-EOT
    // Step - 1
    await page.goto('https://www.example.com');
EOT

  insights { insight_setting_type = "inherit" }
  request_settings { request_setting_type = "inherit" }
  schedule_settings { schedule_setting_type = "inherit" }
  advanced_settings { advanced_setting_type = "inherit" }
  alert_settings { alert_setting_type = "inherit" }
}

resource "playwright_test" "playwright_test" {
  test_name                = "playwright_TF_12"
  provider                 = catchpoint
  monitor                  = "chrome"
  test_script              = <<-EOT
    // Step - 1
    await page.goto('https://www.example.com/');
EOT
  division_id              = 2633
  product_id               = 25232
  test_description         = "A Playwright test for example.com"
  alerts_paused            = true
  enable_test_data_webhook = true
  start_time               = "2024-04-30T04:59:00Z"
  end_time                 = "2024-10-30T04:59:00Z"
  gateway_address_or_host  = "192.168.0.1"

  label {
    key    = "label1"
    values = ["v1", "v2"]
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
    verify_test_on_failure                   = true
    viewport_height                          = 800
    viewport_width                           = 400
    zone_data_collection_enabled             = true
  }

  thresholds {
    test_time_warning     = 47
    test_time_critical    = 56
    availability_critical = 30
    availability_warning  = 67
  }

  schedule_settings {
    frequency          = "6 hours"
    node_distribution  = "random"
    no_of_subset_nodes = 5
    node_ids           = [6388]
    node_group_ids     = [9922, 9848]
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

  alert_settings {
    alert_rule {
      alert_type                 = "timing"
      alert_sub_type             = "response"
      node_threshold_type        = "node"
      threshold_number_of_runs   = 5
      threshold_interval         = "30 minutes"
      trigger_type               = "specific value"
      warning_trigger            = 51
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
      subject          = "testing"
      recipient_emails = ["vikash@catchpoint.com"]
    }
  }
}

resource "playwright_test" "imported_playwrighttest" {
  provider    = catchpoint
  monitor     = "chrome"
  test_name   = "imported playwright test"
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
