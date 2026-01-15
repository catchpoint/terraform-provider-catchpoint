package advancedsettings

import (
	"testing"

	"catchpoint-provider/internal/testutil"
	"catchpoint-provider/internal/types"
)

func TestGetAdvancedSettingsForTestTypeSSL(t *testing.T) {
	advancedSettings := GetAdvancedSettingsForTestType(types.SSLType)

	expectedKeys := []string{
		"additional_monitor",
		"enable_path_mtu_discovery",
		"certificate_revocation_disabled",
		"verify_test_on_failure",
	}

	testutil.AssertElementsMatch(t, "AdvancedSettingsSSL", advancedSettings, expectedKeys)
}

func TestGetAdvancedSettingsForTestAndMonitorCombinationSSL(t *testing.T) {
	advancedSettings := GetAdvancedSettingsForTestAndMonitorCombination(types.SSLType, types.SSLString)

	expectedKeys := []string{
		"additional_monitor",
		"enable_path_mtu_discovery",
		"certificate_revocation_disabled",
		"verify_test_on_failure",
	}

	testutil.AssertElementsMatch(t, "AdvancedSettingsSSL", advancedSettings, expectedKeys)
}
