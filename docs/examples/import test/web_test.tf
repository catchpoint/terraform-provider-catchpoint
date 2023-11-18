terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "0.2.7"
    }
  }
}
provider "catchpoint" {
api_token="5618ABF44CAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

resource "web_test" "webTest" {
    provider=catchpoint
    id="2340171"
}


# =========================================================
# Command to run the importing test details:
# terraform import web_test.webTest 2340171