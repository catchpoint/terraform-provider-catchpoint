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

resource "transaction_test" "transactionTest" {
    provider=catchpoint
    id="2340152"
}

# =========================================================
# Command to run the importing test details:
# terraform import transaction_test.transactionTest 2340152