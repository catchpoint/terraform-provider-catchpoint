terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "2.0.0"
    }
  }
}

resource "ssl_test" "minimal_ssltest" {
  test_name     = "minimal SSL test"
  provider      = catchpoint
  division_id   = 26335
  product_id    = 252325
  test_location = "ssl://www.example.com"

  schedule_settings { schedule_setting_type = "inherit" }
  advanced_settings { advanced_setting_type = "inherit" }
  alert_settings { alert_setting_type = "inherit" }
}

resource "ssl_test" "ssltest" {
  test_name        = "SSL test"
  test_description = "An SSL test for example.com"
  status           = "active"
  provider         = catchpoint
  division_id      = 26335
  product_id       = 252325
  test_location    = "ssl://www.example.com"
  end_time         = "2024-10-30T04:59:00Z"

  advanced_settings {
    certificate_revocation_disabled = true
    additional_monitor              = "ping icmp"
    enable_path_mtu_discovery       = true
    verify_test_on_failure          = true
  }

  schedule_settings {
    frequency         = "12 hours"
    node_distribution = "random"
    node_ids          = [1234, 5678]
    node_group_ids    = [23456]
  }

  label {
    key    = "automation"
    values = ["terraform"]
  }

  thresholds {
    test_time_critical    = 1000
    test_time_warning     = 500
    availability_critical = 90.0
    availability_warning  = 95.0
  }
  alert_settings {
    alert_setting_type = "inherit"
  }
}

resource "ssl_test" "imported_ssltest" {
  test_name     = "imported ssl test"
  provider      = catchpoint
  division_id   = 26335
  product_id    = 252325
  test_location = "ssl://www.example.com"

  schedule_settings { schedule_setting_type = "inherit" }
  advanced_settings { advanced_setting_type = "inherit" }
  alert_settings { alert_setting_type = "inherit" }
}
