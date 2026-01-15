package advancedsettings

import (
	"fmt"

	"catchpoint-provider/internal/helpers"
	"catchpoint-provider/internal/types"
)

var GetAdvancedSettingsForTestAndMonitorCombinationFunc = GetAdvancedSettingsForTestAndMonitorCombination

func ValidateAdvancedSettingsCombination(advancedSettingsList []map[string]any, test types.TestType, monitor string) error {
	validAdvancedSettings := GetAdvancedSettingsForTestAndMonitorCombinationFunc(test, monitor)
	if validAdvancedSettings == nil {
		return fmt.Errorf("invalid monitor type '%s' for test type '%v'", monitor, test)
	}

	for _, advancedSetting := range validAdvancedSettings {
		found := false
		for _, m := range advancedSettingsList {
			if _, ok := m[advancedSetting]; ok {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("invalid advanced setting '%s' for monitor '%s'", advancedSetting, monitor)
		}
	}

	return nil
}

// Return every possible advanced setting for all test types for the Folder and Product schemas.
// In this way, we never have to change the advanced settings for Folders or Products. As long as they're available
// for a test type, they'll be available for the Folder and Product schemas.
func GetAllAdvancedSettings() []string {
	return helpers.FlatCombineStringSlicesUnique(
		apiAdvancedSettings,
		bgpAdvancedSettings,
		dnsDirectAdvancedSettings,
		dnsExperienceAdvancedSettings,
		pingAdvancedSettings,
		playwrightAdvancedSettings,
		puppeteerAdvancedSettings,
		sslAdvancedSettings,
		tracerouteAdvancedSettings,
		transactionChromeAdvancedSettings,
		transactionEmulatedAdvancedSettings,
		transactionMobileAdvancedSettings,
		webChromeAdvancedSettings,
		webEmulatedHttpAdvancedSettings,
		webMobileAdvancedSettings,
		webPlaybackAdvancedSettings)
}

// For Schema construction, we need to return all possible advanced settings for a given test type.
func GetAdvancedSettingsForTestType(testType types.TestType) []string {
	switch testType {
	case types.APIType:
		return apiAdvancedSettings
	case types.BGPType:
		return bgpAdvancedSettings
	case types.DNSType:
		return helpers.FlatCombineStringSlicesUnique(dnsDirectAdvancedSettings,
			dnsExperienceAdvancedSettings)
	case types.PingType:
		return pingAdvancedSettings
	case types.PlaywrightType:
		return playwrightAdvancedSettings
	case types.PuppeteerType:
		return puppeteerAdvancedSettings
	case types.SSLType:
		return sslAdvancedSettings
	case types.TracerouteType:
		return tracerouteAdvancedSettings
	case types.TransactionType:
		return helpers.FlatCombineStringSlicesUnique(transactionChromeAdvancedSettings,
			transactionEmulatedAdvancedSettings,
			transactionMobileAdvancedSettings)
	case types.WebType:
		return helpers.FlatCombineStringSlicesUnique(webChromeAdvancedSettings,
			webEmulatedHttpAdvancedSettings,
			webMobileAdvancedSettings,
			webPlaybackAdvancedSettings)
	default:
		// This will only happen if someone didn't create the type mapping for a given test.
		panic(fmt.Sprintf("Unsupported test type for advanced settings: %d", testType))
	}
}

// For further validation, we need to return advanced settings for a specific test type and monitor combination.
func GetAdvancedSettingsForTestAndMonitorCombination(testType types.TestType, monitor string) []string {
	switch testType {
	case types.APIType:
		return apiAdvancedSettings
	case types.BGPType:
		return bgpAdvancedSettings // BGP has no advanced settings, so return empty slice.
	case types.DNSType:
		return getDNSAdvancedSettings(monitor)
	case types.PingType:
		return pingAdvancedSettings
	case types.PlaywrightType:
		return playwrightAdvancedSettings
	case types.PuppeteerType:
		return puppeteerAdvancedSettings
	case types.SSLType:
		return sslAdvancedSettings
	case types.TracerouteType:
		return tracerouteAdvancedSettings
	case types.TransactionType:
		return getTransactionAdvancedSettings(monitor)
	case types.WebType:
		return getWebAdvancedSettings(monitor)
	default:
		// This will only happen if someone didn't create the type mapping for a given test.
		panic(fmt.Sprintf("Unsupported test type and monitor combination for advanced settings: %d, %s", testType, monitor))
	}
}

// #region Helper Functions

func getDNSAdvancedSettings(monitor string) []string {
	switch monitor {
	case types.DNSDirectString:
		return dnsDirectAdvancedSettings
	case types.DNSExperienceString:
		return dnsExperienceAdvancedSettings
	default:
		// This will only happen if someone didn't create the type mapping for a given DNS monitor type.
		panic(fmt.Sprintf("Unsupported monitor for DNS advanced settings: %s", monitor))
	}
}

func getTransactionAdvancedSettings(monitor string) []string {
	switch monitor {
	case types.ChromeString:
		return transactionChromeAdvancedSettings
	case types.EmulatedString:
		return transactionEmulatedAdvancedSettings
	case types.MobileString:
		return transactionMobileAdvancedSettings
	default:
		// This will only happen if someone didn't create the type mapping for a given Transaction monitor type.
		panic(fmt.Sprintf("Unsupported monitor for Transaction test advanced settings: %s", monitor))
	}
}

func getWebAdvancedSettings(monitor string) []string {
	switch monitor {
	case types.ChromeString:
		return webChromeAdvancedSettings
	case types.EmulatedString, types.HTTPString:
		return webEmulatedHttpAdvancedSettings
	case types.MobileString:
		return webMobileAdvancedSettings
	case types.PlaybackString:
		return webPlaybackAdvancedSettings
	case types.MobilePlaybackString:
		return webPlaybackAdvancedSettings
	default:
		// This will only happen if someone didn't create the type mapping for a given Web monitor type.
		panic(fmt.Sprintf("Unsupported monitor for Web test advanced settings: %s", monitor))
	}
}

// #endregion
