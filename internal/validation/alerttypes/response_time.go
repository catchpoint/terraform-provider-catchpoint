package alerttypes

/**
Note: there is a lot of duplication in these maps. This is intentional to make it very clear
which alert subtypes are valid for each monitor type. While we could reduce duplication by
reusing maps, it would make it harder to see at a glance which subtypes are valid for each
monitor type.

There's also some overlap between monitor types, such as Web and Transaction, which share many
of the same subtypes, but not all (e.g., DaysToExpiration is missing for Transaction).
*/

import (
	"catchpoint-provider/internal/types"
)

// Response Time is sometimes used without any subtypes (Ping, DNS).
var validEmptyResponseTimeMatchMap = map[string]*AlertSubTypes{
	types.Timing: {}, // Response-Time in the portal and DB.
}

var validSSLResponseTimeMatchMap = map[string]*AlertSubTypes{
	types.Timing: {
		SubTypes: []string{
			types.Connect,
			types.DaysToExpiration,
			types.HandshakeTime,
		},
	},
}

// Playwright and Puppeteer tests have the same subtypes.
var validPlaywrightPuppeteerResponseTimeMatchMap = map[string]*AlertSubTypes{
	types.Timing: {
		SubTypes: []string{
			types.Connect,
			types.ContentLoad,
			types.DNS,
			types.Load,
			types.Redirect,
			types.Response,
			types.Send,
			types.ServerResponse,
			types.TestTime,
			types.TestTimeWithSuspect,
			// TODO: where is this? types.TimeToFirstByte,
			// TODO: where is this? types.DataTransferTime,
			types.DocumentComplete,
			types.DOMLoad,
			// TODO: where is this? types.CumulativeLayoutShift,
			// TODO: where is this? types.FirstContentfulPaint,
			// TODO: where is this? types.FirstPaint,
			// TODO: where is this? types.LargestContentfulPaint,
		},
	},
}

// API test.
var validAPIResponseTimeMatchMap = map[string]*AlertSubTypes{
	types.Timing: {
		SubTypes: []string{
			types.Connect,
			types.ContentLoad,
			types.DNS,
			types.Load,
			types.Redirect,
			types.Response,
			types.Send,
			types.ServerResponse,
			types.TestTime,
			types.TestTimeWithSuspect,
			// TODO: where is this? types.TimeToFirstByte,
			// TODO: where is this? types.DataTransferTime,
			types.DaysToExpiration,
			types.Wait,
		},
	},
}

// Web Test - Mobile monitor.
var validWebMobileResponseTimeMatchMap = map[string]*AlertSubTypes{
	types.Timing: {
		SubTypes: []string{
			types.Connect,
			types.ContentLoad,
			types.DNS,
			types.Load,
			types.Redirect,
			types.Response,
			types.Send,
			types.ServerResponse,
			types.TestTime,
			types.TestTimeWithSuspect,
			// TODO: where is this? types.TimeToFirstByte,
			types.DaysToExpiration,
			types.DocumentComplete,
			types.DOMLoad,
			// TODO: where is this? types.CumulativeLayoutShift,
			// TODO: where is this? types.FirstContentfulPaint,
			// TODO: where is this? types.FirstPaint,
			// TODO: where is this? types.LargestContentfulPaint,
			types.Wait,
		},
	},
}

// Web - Chrome, Playback, MobilePlayback have all the same subtypes.
var validWebChromePlaybackResponseTimeMatchMap = map[string]*AlertSubTypes{
	types.Timing: {
		SubTypes: []string{
			types.Connect,
			types.ContentLoad,
			types.DNS,
			types.Load,
			types.Redirect,
			types.Response,
			types.Send,
			types.ServerResponse,
			types.TestTime,
			types.TestTimeWithSuspect,
			// TODO: where is this? types.TimeToFirstByte,
			// TODO: where is this? types.DataTransferTime,
			types.DaysToExpiration,
			types.DocumentComplete,
			types.DOMLoad,
			// TODO: where is this? types.CumulativeLayoutShift,
			// TODO: where is this? types.FirstContentfulPaint,
			// TODO: where is this? types.FirstPaint,
			// TODO: where is this? types.LargestContentfulPaint,
			types.Wait,
		},
	},
}

// Web - Emulated and HTTP have all the same subtypes.
var validWebEmulatedResponseTimeMatchMap = map[string]*AlertSubTypes{
	types.Timing: {
		SubTypes: []string{
			types.Connect,
			types.ContentLoad,
			types.DNS,
			types.Load,
			types.Redirect,
			types.Response,
			types.Send,
			types.ServerResponse,
			types.TestTime,
			types.TestTimeWithSuspect,
			// TODO: where is this? types.TimeToFirstByte,
			// TODO: where is this? types.DataTransferTime,
			types.DaysToExpiration,
			types.DocumentComplete,
			types.DOMLoad,
			types.Wait,
		},
	},
}

// Transaction Test - Chrome monitor.
var validTransactionChromeResponseTimeMatchMap = map[string]*AlertSubTypes{
	types.Timing: {
		SubTypes: []string{
			types.Connect,
			types.ContentLoad,
			types.DNS,
			types.Load,
			types.Redirect,
			types.Response,
			types.Send,
			types.ServerResponse,
			types.TestTime,
			types.TestTimeWithSuspect,
			// TODO: where is this? types.TimeToFirstByte,
			// TODO: where is this? types.DataTransferTime,
			types.DocumentComplete,
			types.DOMLoad,
			// TODO: where is this? types.CumulativeLayoutShift,
			// TODO: where is this? types.FirstContentfulPaint,
			// TODO: where is this? types.FirstPaint,
			// TODO: where is this? types.LargestContentfulPaint,
			types.Wait,
		},
	},
}

// Transaction Test - Emulated monitor.
var validTransactionEmulatedResponseTimeMatchMap = map[string]*AlertSubTypes{
	types.Timing: {
		SubTypes: []string{
			types.Connect,
			types.ContentLoad,
			types.DNS,
			types.Load,
			types.Redirect,
			types.Response,
			types.Send,
			types.ServerResponse,
			types.TestTime,
			types.TestTimeWithSuspect,
			// TODO: where is this? types.TimeToFirstByte,
			// TODO: where is this? types.DataTransferTime,
			types.DocumentComplete,
			types.DOMLoad,
			types.Wait,
		},
	},
}

// Transaction Test - Mobile monitor.
var validTransactionMobileResponseTimeMatchMap = map[string]*AlertSubTypes{
	types.Timing: {
		SubTypes: []string{
			types.Connect,
			types.ContentLoad,
			types.DNS,
			types.Load,
			types.Redirect,
			types.Response,
			types.Send,
			types.ServerResponse,
			types.TestTime,
			types.TestTimeWithSuspect,
			// TODO: where is this? types.TimeToFirstByte,
			types.DocumentComplete,
			types.DOMLoad,
			// TODO: where is this? types.CumulativeLayoutShift,
			// TODO: where is this? types.FirstContentfulPaint,
			// TODO: where is this? types.FirstPaint,
			// TODO: where is this? types.LargestContentfulPaint,
			types.Wait,
		},
	},
}
