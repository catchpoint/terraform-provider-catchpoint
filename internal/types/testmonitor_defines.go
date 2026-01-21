package types

// Chrome Versions.
const (
	Chrome53  = "53"
	Chrome59  = "59"
	Chrome63  = "63"
	Chrome66  = "66"
	Chrome71  = "71"
	Chrome75  = "75"
	Chrome79  = "79"
	Chrome85  = "85"
	Chrome87  = "87"
	Chrome89  = "89"
	Chrome97  = "97"
	Chrome108 = "108"
	Chrome120 = "120"
	Preview   = "preview"
	Stable    = "stable"
)

// Test Types.
const (
	HTTPString           = "http"
	EmulatedString       = "emulated"
	ChromeString         = "chrome"
	EdgeString           = "edge"
	PlaybackString       = "playback"
	MobilePlaybackString = "mobile playback"
	MobileString         = "mobile"
	APIString            = "api"
	PingICMPString       = "ping icmp"
	PingTCPString        = "ping tcp"
	PingUDPString        = "ping udp"
	TracerouteICMPString = "traceroute icmp"
	TracerouteUDPString  = "traceroute udp"
	TracerouteTCPString  = "traceroute tcp"
	DNSExperienceString  = "dns experience"
	DNSDirectString      = "dns direct"
	SSLString            = "ssl"
	BGPString            = "bgp"
	BGPBasicString       = "bgp basic"
)

// API types.
const (
	// APIScriptTypeID 1
	SeleniumString = "selenium"
	// APIScriptTypeID 2
	JavaScriptString = "javascript"
	// APIScriptTypeID 3
	PlaywrightString = "playwright"
	// APIScriptTypeID 4
	PuppeteerString = "puppeteer"
)

// UserAgent types used for device simulation in Web/Transaction/Puppeteer/Playwright tests.
const (
	UserAgentIE            = "ie"
	UserAgentChrome        = "chrome"
	UserAgentAndroid       = "android"
	UserAgentiPhone        = "iphone"
	UserAgentiPad2         = "ipad 2"
	UserAgentKindleFire    = "kindle fire"
	UserAgentGalaxyTab     = "galaxy tab"
	UserAgentiPhone5       = "iphone 5"
	UserAgentiPadMini      = "ipad mini"
	UserAgentGalaxyNote    = "galaxy note"
	UserAgentNexus7        = "nexus 7"
	UserAgentNexus4        = "nexus 4"
	UserAgentNokiaLumia920 = "nokia lumia920"
	UserAgentIPhone6       = "iphone 6"
	UserAgentBlackberryZ30 = "blackberry z30"
	UserAgentGalaxyS4      = "galaxy s4"
	UserAgentHTCOneX       = "htc onex"
	UserAgentLGOptimusG    = "lg optimusg"
	UserAgentDroidRazrHD   = "droid razr hd"
	UserAgentNexus6        = "nexus 6"
	UserAgentIPhone6S      = "iphone 6s"
	UserAgentGalaxyS6      = "galaxy s6"
	UserAgentIPhone7       = "iphone 7"
	UserAgentGooglePixel   = "google pixel"
	UserAgentGalaxyS8      = "galaxy s8"
)
