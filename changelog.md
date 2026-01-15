# v2.0.0

## Breaking Changes
	- Removed "contact_groups" from "alert_settings,notification_group" and "alert_rule,notification_group". The new field is "contact_group_ids" and expects an array of int (like 'node_ids').
	- Renamed "recipient_email_ids" to "recipient_emails" as the field takes strings, not ids.
	- The "schedule_settings" for a Product are required by the API so they are also now required by the user's Terraform config.
	- AlertSubType "handshake_time" changed to "handshake time" for consistency with other types.
	- AlertSubType "days_to_expiration" changed to "days to expiration" for consistency with other types.
	- For 'web_test' resources, the monitor 'object' has been renamed to 'http'.
	- For 'playwright_test' resources, the 'playwright' monitor type was renamed to 'edge' for clarity.
	- 'monitor' is now required only when the test resource can set more than one 'monitor' type, i.e. for the following monitors:
	    - BGP ('bgp', 'bgp basic')
	    - DNS ('dns experience', 'dns direct')
	    - Ping ('ping icmp', 'ping tcp', 'ping udp')
	    - Playwright ('edge', 'chrome')
	    - Traceroute ('traceroute icmp', 'traceroute tcp', 'traceroute udp')
	    - Transaction ('chrome', 'mobile', 'emulated')
	    - Web ('chrome', 'emulated', 'http', 'playback', 'mobile playback', 'mobile')
    - For 'playwright_test', 'puppeteer_test', the 'test_script_type' field was made read-only. There is no reason to set this as there is only one option for each.
    - The various `x_setting_type` attributes in nested blocks can now be set to enforce inheritance. e.g. `insights { insight_setting_type = "inherit" }` will ensure that the resource will always inherit the insight settings from the folder or product.

## Other Changes

### general
	- Updated: Enhanced API errors with more details.
	- Updated: now defaulting to latest API version 4.0.
	- Updated: Can override the API version by setting `CATCHPOINT_API_VERSION`.
	- Updated: Examples are now part of main documentation.
	- Fixed: upstream deletion of an object no longer causes an error but allows the object to be created again.
	- Fixed: Schedules, AdvancedSettings, RequestSettings, and InsightSettings should all properly calculate the inherit/override setting.
	- Fixed: nested contact_groups now correctly use an int64[] instead of string[].

### manage_folder
    - Fixed: 'schedule_settings' may now be overridden from a folder or will inherit if omitted.

### manage_product
	- Updated: 'manage_product' resource now can be set with _either_ node_group_ids and/or node_ids).
	- Fixed: 'manage_product' always showing as updating when Insights are not set (not yet fixed: resources showing as updating when alerts are defined).
    - Fixed: 'manage_product' resource type always requires a schedule and validation has been moved up front.
	- Fixed: 'manage_product' can now set DNS AlertTypes.

### test resources
	- Updated: Tests may now be moved between different folder_ids and product_ids.
	- Updated: the values in the 'thresholds' block are no longer required.
	- Fixed: 'traceroute_test' schema was allowing unsupported additional settings (additional_monitor, bandwidth_throttling).
	- Fixed: 'web_test' could not update userAgentTypeID changes ('simulate' field).
	- Fixed: perpetual diff in testscript inputs.
	- Fixed: added missing "gateway_address_or_host" setting to SSL schema.
	- Fixed: 'ssl' test resource now validates the test_location field as part of the schema.
	- Fixed: Playwright and Puppeteer test types errantly offered 'chrome_version' and 'simulate' fields.
	- Fixed: 'end_date' is no longer a required field. Note: sometimes the API may require it, but the Terraform Provider does not.
	- Fixed: node_group_ids were not updating for already created tests.

# v1.5.0

ENHANCEMENTS

* Create and update Products and Folders.
* Configure “Inherit and Add” for alerts.
* Configure DNS Answer alerts for DNS tests.
* Configure custom headers.
* Support for chrome version 120 in web test and transaction test.

BUG FIXES

* Unable to add additional monitor advanced settings while creating and updating a ping test.
* Contact group is empty when we are importing test details.
* Certificate field is missing if enforce_certificate_pinning and enforce_certificate_key_pinning are set to true hence unable to create SSL test.
* Unable to update schedule settings and insights for a tests.
* Username is displaying as null for playwright test.
* HTTP headers is displaying as an empty array for all test types.

# v1.4.0

FIX

* Added support to import additional settings for a test. [Known Issue]: The fix done in v1.2.0 to correctly detect changes for optional fields may not work for some additional settings.

# v1.3.0

ENHANCEMENT

* A minor change to handle bulk requests by limiting the processing of number of requests to 7 per second has context menu.

# v1.2.0

BUG FIXES

* Fixed a minor bug to detect changes correctly for optional fields on running "terraform apply" command.


# v1.0.1

Catchpoint internal testing:

* import command fails when the alert settings are set to Override
* Update operation in terraform without any changes, displays 1 change applied message

# v1.0.0

BUG FIXES:

* Not able to update the status of a test using terraform..
* Unable to add request settings to a test
* Unable to add alerts to DNS(Direct and Experience), Ping test and Traceroute while creating
* Not able to add node groups to create a test using terraform.
* advanced settings, request settings, scheduling, and labels returning empty arrays empty arrays when import a test
* Query type is displaying as null when trying to import the DNS test details.
* Add Subject field for default notification group
* Support for default notification group for alert setting
* Support for notification group for each alert rule
* Support for contact groups in each notification group
* Support for 40x_or_50x_http_mark_successful not 30x_redirects_do_not_follow advance setting.
* Support for adding "enable_consecutive=true"for nodes in alerts
* Support for subset of nodes in schedule setting when creating a test

NEW TEST TYPE SUPPORTED: 

* Support for playwright test
* Support for puppeteer test


# v0.2.9

Catchpoint internal testing:

* fixed couple of bugs.
* Document updates.

# v0.2.8

Catchpoint internal testing:

* fixed couple of bugs.
* Document updates.

# v0.2.7

BUG FIXES:

* Web test request setting issue.
* frequency issue fixed.

# v0.2.6

BUG FIXES:

* Web test request setting issue.
* frequency issue fixed.


# v0.2.5

BUG FIXES:

* catchpoint test import issues fixed.
* plugin crashed issue fixed

# v0.2.4

BUG FIXES:

* catchpoint test import issues fixed.
* plugin crashed issue fixed

# v0.2.3

BUG FIXES:

* catchpoint test import issues fixed.
* plugin crashed issue fixed

# v0.2.2

BUG FIXES:

* catchpoint test import issues fixed.

# v0.2.1

FEATURES:

* **New Resource:** `catchpoint`
