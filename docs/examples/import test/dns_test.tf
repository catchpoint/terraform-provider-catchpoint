terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "0.2.7"
    }
  }
}
provider "catchpoint" {
api_token="5618ABF44CA1117B4286C9572XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

resource "dns_test" "dnsTest" {
    provider=catchpoint
    id="2340152"
}

# =========================================================
# Command to run the importing test details:
# terraform import dns_test.dnsTest 2340152