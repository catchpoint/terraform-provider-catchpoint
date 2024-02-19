func flattenAdvancedSetting(advancedSetting AdvancedSetting) []interface{} {

	advSettingMap := make(map[string]interface{})

	if advancedSetting.MaxStepRuntimeSecOverride != 0 {
		advSettingMap["enforce_failure_test_runs_longer_than"] = advancedSetting.MaxStepRuntimeSecOverride
	}

	if advancedSetting.WaitForNoActivity != nil {
		advSettingMap["wait_for_no_activity"] = *advancedSetting.WaitForNoActivity
	}

	if advancedSetting.ViewportHeight != 0 {
		advSettingMap["viewport_height"] = advancedSetting.ViewportHeight
	}

	if advancedSetting.ViewportWidth != 0 {
		advSettingMap["viewport_width"] = advancedSetting.ViewportWidth
	}

	if advancedSetting.FailureHopCount != 0 {
		advSettingMap["failure_hop_count"] = advancedSetting.FailureHopCount
	}

	if advancedSetting.PingCount != 0 {
		advSettingMap["ping_count"] = advancedSetting.PingCount
	}

	if advancedSetting.EdnsSubnet != "" {
		advSettingMap["edns_subnet"] = advancedSetting.EdnsSubnet
	}

	additionalMonitor := ""
	if advancedSetting.AdditionalMonitor != nil {
		additionalMonitor = getAdditionalMonitorTypeName(advancedSetting.AdditionalMonitor.Id)
		if additionalMonitor != "" {
			advSettingMap["additional_monitor"] = additionalMonitor
		}
	}

	testBandwidthThrottling := ""
	if advancedSetting.TestBandwidthThrottling != nil {
		testBandwidthThrottling = getBandwidthThrottlingTypeName(advancedSetting.TestBandwidthThrottling.Id)
		if testBandwidthThrottling != "" {
			advSettingMap["bandwidth_throttling"] = testBandwidthThrottling
		}
	}

	// advancedSettingTypeName := ""
	// advancedSettingTypeName = getAdvancdSettingTypeName(advancedSetting.AdvancedSettingType.Id)
	// advSettingMap["advanced_setting_type"] = advancedSettingTypeName

	if len(advancedSetting.AppliedTestFlags) > 0 {
		var flagNames []string
		for _, flag := range advancedSetting.AppliedTestFlags {
			flagName := getTestFlagName(flag.Id)
			if flagName != "" {
				flagNames = append(flagNames, flagName)
			}
		}
		if len(flagNames) > 0 {
			advSettingMap["applied_test_flags"] = flagNames
		}
	}

	return []interface{}{advSettingMap}
}