resource "ssl_test" "sslTest" {
    test_name  = "SSL_TF 1"
    provider=catchpoint
    division_id=26335
    product_id=252325
    monitor="ssl"
    test_location="ssl://www.catchpoint.com"
    end_time="2024-10-30T04:59:00Z"
}