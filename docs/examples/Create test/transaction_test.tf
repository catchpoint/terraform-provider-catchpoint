terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "0.2.1"
    }
  }
}

provider "catchpoint" {
    api_token="REMOVED"
}

resource "transaction_test" "testTransaction" {
    test_name  = "Transaction_TF"
    provider=catchpoint
    division_id=2633
    product_id=23791
    monitor="chrome"
    test_script="//Step-1\r\nopen(\"https:www.google.com)"
    end_time="2023-10-30T04:59:00Z"
}