package expand

import (
	"context"
	"fmt"

	"catchpoint-provider/internal/models"
	cptypes "catchpoint-provider/internal/types"
	"catchpoint-provider/internal/validation"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// GetStatusFromStringOrActive returns the status from the string or defaults to "Active".
func GetStatusFromStringOrActive(status attr.Value) models.IDNameConfig {
	// Use 'status' if set, otherwise default to "Active".
	if status != nil && !status.IsNull() && !status.IsUnknown() {
		statusStr := status.(types.String).ValueString()
		return validation.GetStatusTypeOrDefault(statusStr)
	}
	return models.IDNameConfig{ID: 0, Name: cptypes.Active}
}

// expandStringSettingFromAttrs expands String attributes only if they exist, does nothing otherwise.
func expandStringSettingFromAttrs(attrs map[string]attr.Value, key string, setter func(string)) {
	if attrExistsAndNotNull(attrs, key) {
		setter(attrs[key].(types.String).ValueString())
	}
}

// expandIntSettingFromAttrs expands Int attributes only if they exist, does nothing otherwise.
func expandIntSettingFromAttrs(attrs map[string]attr.Value, key string, setter func(int)) {
	if attrExistsAndNotNull(attrs, key) {
		setter(int(attrs[key].(types.Int64).ValueInt64()))
	}
}

// expandIntListSettingFromAttrs expands Int List attributes only if they exist, does nothing otherwise.
func expandIntListSettingFromAttrs(attrs map[string]attr.Value, key string, setter func([]int)) (diags diag.Diagnostics) {
	if v, ok := attrs[key]; ok && !v.IsNull() && !v.IsUnknown() {
		intList := v.(types.List)

		int64List := []types.Int64{}
		elementsDiags := intList.ElementsAs(context.TODO(), &int64List, false)
		diags.Append(elementsDiags...)
		if diags.HasError() {
			return
		}

		// Convert []types.Int64 to []int
		result := []int{}

		for _, id := range int64List {
			result = append(result, int(id.ValueInt64()))
		}
		setter(result)
	}

	return
}

// getIntListFromAttr extracts a list of integers from the attribute map.
func getIntListFromAttr(key string, attr map[string]attr.Value) ([]int, diag.Diagnostics) {
	if attrExistsAndNotNull(attr, key) {
		intList := attr[key].(types.List)

		var int64List []types.Int64
		diags := intList.ElementsAs(context.TODO(), &int64List, false)
		if diags.HasError() {
			return nil, diags
		}

		// Convert []types.Int64 to []int
		var result []int
		for _, id := range int64List {
			result = append(result, int(id.ValueInt64()))
		}
		return result, nil
	}
	return []int{}, nil
}

func attrExistsAndNotNull(attrs map[string]attr.Value, fieldName string) bool {
	if attr, exists := attrs[fieldName]; exists {
		return !attr.IsNull() && !attr.IsUnknown()
	}
	return false
}

func isBlockEmpty(block types.Object) bool {
	if block.IsNull() || block.IsUnknown() {
		return true
	}

	attrs := block.Attributes()

	for _, attr := range attrs {
		if !attr.IsNull() && !attr.IsUnknown() && hasNonDefaultValue(attr) {
			return false
		}
	}

	return true
}

func hasNonDefaultValue(attr attr.Value) bool {
	switch v := attr.(type) {
	case types.Bool:
		return v.ValueBool()
	case types.String:
		return v.ValueString() != ""
	case types.Int64:
		return v.ValueInt64() != 0
	case types.Float64:
		return v.ValueFloat64() != 0.0
	case types.List:
		return len(v.Elements()) > 0
	case types.Set:
		return len(v.Elements()) > 0
	case types.Object:
		return !isBlockEmpty(v) // Recursive check for nested objects
	default:
		return false
	}
}

func expandStringSettingWithLookup(setting types.String, lookupFunc func(key string) (int, bool), destination *models.IDNameConfig) (diags diag.Diagnostics) {
	if !setting.IsNull() && !setting.IsUnknown() {
		lookupValue, found := lookupFunc(setting.ValueString())
		if found {
			*destination = models.IDNameConfig{
				ID:   lookupValue,
				Name: setting.ValueString(),
			}
		} else {
			diags.Append(diag.NewWarningDiagnostic("Invalid Value", fmt.Sprintf("The provided value for setting %s is not recognized: %s", setting, setting.ValueString())))
		}
	}
	return
}

func expandBoolSetting(setting types.Bool, destination *bool) {
	if !setting.IsNull() && !setting.IsUnknown() {
		*destination = setting.ValueBool()
	}
}

func expandStringSetting(setting types.String, destination *string) {
	if !setting.IsNull() && !setting.IsUnknown() {
		*destination = setting.ValueString()
	}
}

func expandFloat64Setting(setting types.Float64, destination *float64) {
	if !setting.IsNull() && !setting.IsUnknown() {
		*destination = setting.ValueFloat64()
	}
}

func expandIntSetting(setting types.Int64, destination *int) {
	if !setting.IsNull() && !setting.IsUnknown() {
		*destination = int(setting.ValueInt64())
	}
}

func expandIntListSetting(setting types.List, destination *[]int) (diags diag.Diagnostics) {
	if !setting.IsNull() && !setting.IsUnknown() {
		var int64List []types.Int64
		elementsDiags := setting.ElementsAs(context.TODO(), &int64List, false)
		diags.Append(elementsDiags...)
		if diags.HasError() {
			return
		}

		// Convert []types.Int64 to []int
		var result []int
		for _, id := range int64List {
			result = append(result, int(id.ValueInt64()))
		}
		*destination = result
	} else {
		*destination = []int{}
	}
	return
}

func expandStringListSetting(setting types.List, destination *[]string) (diags diag.Diagnostics) {
	if !setting.IsNull() && !setting.IsUnknown() {
		var stringList []types.String
		elementsDiags := setting.ElementsAs(context.TODO(), &stringList, false)
		diags.Append(elementsDiags...)
		if diags.HasError() {
			return
		}

		// Convert []types.String to []string
		var result []string
		for _, str := range stringList {
			result = append(result, str.ValueString())
		}
		*destination = result
	} else {
		*destination = []string{}
	}
	return
}
