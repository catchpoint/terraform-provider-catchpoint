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
