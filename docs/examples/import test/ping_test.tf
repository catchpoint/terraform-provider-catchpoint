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

resource "ping_test" "pingTest" {
    provider=catchpoint
    id="2340152"
}

# =========================================================
# Command to run the importing test details:
# terraform import ping_test.pingTest 2340152