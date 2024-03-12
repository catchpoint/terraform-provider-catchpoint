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
