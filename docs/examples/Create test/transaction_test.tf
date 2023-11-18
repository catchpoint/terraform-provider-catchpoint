terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "0.2.1"
    }
  }
}

provider "catchpoint" {
    api_token="5618ABF44CA1117B4286C9572D00B47AD166125E4F88A6DE472AF29562C50595"
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