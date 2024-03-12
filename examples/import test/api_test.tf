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

resource "api_test" "apiTest" {
    provider=catchpoint
    id="2340152"
}

# =========================================================
# Command to run the importing test details:
# terraform import api_test.apiTest 2340152