terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "2.0.0"
    }
  }
}

resource "api_test" "minimal_api_test" {
  provider         = catchpoint
  division_id      = 2633
  product_id       = 23791
  test_script_type = "selenium"
  test_name        = "A minimal Selenium API Test for https://www.catchpoint.com"
  test_script      = <<-EOT
    // Step - 1
    open("https://www.catchpoint.com")
EOT

  insights { insight_setting_type = "inherit" }
  request_settings { request_setting_type = "inherit" }
  schedule_settings { schedule_setting_type = "inherit" }
  advanced_settings { advanced_setting_type = "inherit" }
  alert_settings { alert_setting_type = "inherit" }
}

resource "api_test" "apitest" {
  test_name        = "APItest_TF"
  provider         = catchpoint
  division_id      = 2633
  product_id       = 23791
  test_script      = <<-EOT
    // Step - 1
    open("https://www.example.com")
EOT
  test_script_type = "selenium"
  start_time       = "2024-04-30T04:59:00Z"
  end_time         = "2024-10-30T04:59:00Z"

  alert_settings { alert_setting_type = "inherit" }

  thresholds {
    test_time_critical    = 7000
    test_time_warning     = 3000
    availability_critical = 80
    availability_warning  = 90
  }

  advanced_settings {
    additional_monitor                       = "ping icmp"
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
    t30x_redirects_do_not_follow             = true
    verify_test_on_failure                   = true
    zone_data_collection_enabled             = true
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
}

resource "api_test" "imported_apitest" {
  provider         = catchpoint
  division_id      = 3
  test_name        = "imported_apitest"
  product_id       = 23791
  test_script      = <<-EOT
    // Step - 1
    open("https://www.example.com")
EOT
  test_script_type = "selenium"

  insights { insight_setting_type = "inherit" }
  request_settings { request_setting_type = "inherit" }
  schedule_settings { schedule_setting_type = "inherit" }
  advanced_settings { advanced_setting_type = "inherit" }
  alert_settings { alert_setting_type = "inherit" }
}
