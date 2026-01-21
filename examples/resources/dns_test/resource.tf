terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "2.0.0"
    }
  }
}

resource "dns_test" "minimal_dnstest" {
  provider    = catchpoint
  test_name   = "minimal_dnstest"
  division_id = 2633
  product_id  = 23791
  monitor     = "dns direct"
  query_type  = "a"
  test_domain = "https://www.example.com"

  schedule_settings { schedule_setting_type = "inherit" }
  advanced_settings { advanced_setting_type = "inherit" }
  alert_settings { alert_setting_type = "inherit" }
}

resource "dns_test" "testDNS" {
  test_name   = "Terraform DNS Test"
  provider    = catchpoint
  division_id = 2633
  product_id  = 23791
  monitor     = "dns direct"
  query_type  = "a"
  test_domain = "https:www.google.com"
  start_time  = "2024-04-30T04:59:00Z"
  end_time    = "2024-10-30T04:59:00Z"

  label {
    key    = "automation"
    values = ["terraform"]
  }

  advanced_settings {
    debug_primary_host_on_failure = true
    enable_tcp_protocol           = true
    enable_nsid                   = true
    disable_recursive_resolution  = true
    enable_dnssec                 = true
    enable_path_mtu_discovery     = true
    additional_monitor            = "ping icmp"
    edns_subnet                   = "192.168.0.0/24"
  }

  thresholds {
    test_time_critical    = 7000
    test_time_warning     = 3000
    availability_critical = 80
    availability_warning  = 90
  }

  schedule_settings {
    frequency          = "6 hours"
    node_distribution  = "random"
    no_of_subset_nodes = 5
    node_ids           = [6388]
    node_group_ids     = [9922, 9848]
  }

  alert_settings {
    alert_setting_type = "inherit"
  }
}

resource "dns_test" "imported_dnstest" {
  provider    = catchpoint
  test_name   = "imported_dnstest"
  division_id = 2633
  product_id  = 23791
  monitor     = "dns direct"
  query_type  = "a"
  test_domain = "https://www.example.com"

  schedule_settings { schedule_setting_type = "inherit" }
  advanced_settings { advanced_setting_type = "inherit" }
  alert_settings { alert_setting_type = "inherit" }
}
