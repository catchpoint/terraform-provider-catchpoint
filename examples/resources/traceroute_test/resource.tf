terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "2.0.0"
    }
  }
}

resource "traceroute_test" "minimal_traceroute_icmp" {
  provider      = catchpoint
  monitor       = "traceroute icmp"
  division_id   = 2633
  test_name     = "Traceroute ICMP minimal test for google.com"
  test_location = "google.com"
  product_id    = 23791

  schedule_settings { schedule_setting_type = "inherit" }
  advanced_settings { advanced_setting_type = "inherit" }
  alert_settings { alert_setting_type = "inherit" }
}

resource "traceroute_test" "traceroutetest" {
  test_name     = "traceroute test"
  provider      = catchpoint
  division_id   = 2633
  product_id    = 23791
  monitor       = "traceroute icmp"
  test_location = "https://www.google.com"
  start_time    = "2024-04-30T04:59:00Z"
  end_time      = "2024-10-30T04:59:00Z"

  advanced_settings { advanced_setting_type = "inherit" }
  
  thresholds {
    test_time_critical    = 250
    test_time_warning     = 50
    availability_critical = 80
    availability_warning  = 90
  }

  label {
    key    = "automation"
    values = ["terraform"]
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

resource "traceroute_test" "imported_traceroutetest" {
  provider    = catchpoint
  test_name   = "imported traceroute test"
  division_id = 2633
  product_id  = 23791
  monitor     = "traceroute icmp"
  test_location = "https://www.google.com"

  schedule_settings { schedule_setting_type = "inherit" }
  advanced_settings { advanced_setting_type = "inherit" }
  alert_settings { alert_setting_type = "inherit" }
}
