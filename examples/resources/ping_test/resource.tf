terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "2.0.0"
    }
  }
}

resource "ping_test" "minimal_pingtest" {
  provider      = catchpoint
  division_id   = 2633
  product_id    = 23791
  test_location = "www.example.com"
  test_name     = "minimal ping test"
  monitor       = "ping tcp"

  schedule_settings { schedule_setting_type = "inherit" }
  advanced_settings { advanced_setting_type = "inherit" }
  alert_settings { alert_setting_type = "inherit" }
}

resource "ping_test" "ping_test" {
  test_name     = "Ping_TF3"
  provider      = catchpoint
  division_id   = 2633
  product_id    = 23791
  test_location = "www.example.com"
  monitor       = "ping tcp"
  status        = "active"
  start_time    = "2024-04-30T04:59:00Z"
  end_time      = "2024-10-30T04:59:00Z"

  advanced_settings {
    additional_monitor            = "traceroute icmp"
    enable_path_mtu_discovery     = true
    verify_test_on_failure        = true
    debug_primary_host_on_failure = true
  }

  thresholds {
    test_time_critical    = 250
    test_time_warning     = 50
    availability_critical = 80
    availability_warning  = 90
  }

  alert_settings {
    alert_setting_type = "inherit"
  }

  schedule_settings {
    frequency          = "6 hours"
    node_distribution  = "random"
    no_of_subset_nodes = 5
    node_ids           = [6388]
    node_group_ids     = [9922, 9848]
  }
}

resource "ping_test" "imported_pingtest" {
  provider      = catchpoint
  division_id   = 2633
  product_id    = 23791
  test_location = "www.example.com"
  test_name     = "imported ping test"
  monitor       = "ping tcp"

  schedule_settings { schedule_setting_type = "inherit" }
  advanced_settings { advanced_setting_type = "inherit" }
  alert_settings { alert_setting_type = "inherit" }
}
