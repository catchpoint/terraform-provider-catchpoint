terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "2.0.0"
    }
  }
}

resource "transaction_test" "minimal_chrome_transaction_test" {
  provider       = catchpoint
  monitor        = "chrome"
  test_name      = "Transaction chrome minimal test"
  test_script    = <<-EOT
    // Step - 1
    open("https://www.example.com")
EOT 
  division_id    = 2633
  product_id     = 23791
  chrome_version = "preview"

  insights { insight_setting_type = "inherit" }
  request_settings { request_setting_type = "inherit" }
  schedule_settings { schedule_setting_type = "inherit" }
  advanced_settings { advanced_setting_type = "inherit" }
  alert_settings { alert_setting_type = "inherit" }
}

resource "transaction_test" "minimal_mobile_transaction_test" {
  provider    = catchpoint
  monitor     = "chrome"
  test_name   = "Transaction mobile minimal test"
  test_script = <<-EOT
    // Step - 1
    open("https://www.example.com")
EOT 
  division_id = 2633
  product_id  = 23791
  simulate    = "iphone"

  insights { insight_setting_type = "inherit" }
  request_settings { request_setting_type = "inherit" }
  schedule_settings { schedule_setting_type = "inherit" }
  advanced_settings { advanced_setting_type = "inherit" }
  alert_settings { alert_setting_type = "inherit" }
}

resource "transaction_test" "transactiontest" {
  test_name   = "Transaction_TF"
  provider    = catchpoint
  division_id = 2633
  product_id  = 23791
  monitor     = "chrome"
  test_script = <<-EOT
    // Step - 1
    open("https://www.example.com")
EOT
  start_time  = "2024-04-30T04:59:00Z"
  end_time    = "2024-10-30T04:59:00Z"

  thresholds {
    test_time_critical    = 7000
    test_time_warning     = 3000
    availability_critical = 80
    availability_warning  = 90
  }

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
  }

  insights {
    tracepoint_ids = [12003, 11887]
    indicator_ids  = [12006]
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

resource "transaction_test" "imported_transactiontest" {
  test_name      = "imported transaction test"
  provider       = catchpoint
  division_id    = 2633
  product_id     = 23791
  monitor        = "chrome"
  chrome_version = "stable"
  test_script    = <<-EOT
    // Step - 1
    open("https://www.example.com")
EOT

  insights { insight_setting_type = "inherit" }
  request_settings { request_setting_type = "inherit" }
  schedule_settings { schedule_setting_type = "inherit" }
  advanced_settings { advanced_setting_type = "inherit" }
  alert_settings { alert_setting_type = "inherit" }
}
