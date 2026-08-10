package catchpoint

import (
	"strings"
	"testing"

	modelresource "catchpoint-provider/internal/models/resource"
	testmonitorresource "catchpoint-provider/internal/models/resource/testmonitor"

	tfresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestGeneratePatchDocumentsIncludesTestScriptChanges(t *testing.T) {
	r := TestResource[*testmonitorresource.PlaywrightTestResourceModel]{}
	resp := &tfresource.UpdateResponse{}

	plan := &testmonitorresource.PlaywrightTestResourceModel{
		BaseTestResourceModel: modelresource.BaseTestResourceModel{
			DivisionID: types.Int64Value(123),
			ProductID:  types.Int64Value(456),
			Name:       types.StringValue("Identity Portal"),
			Monitor:    types.StringValue("edge"),
		},
		TestScriptResourceModel: modelresource.TestScriptResourceModel{
			Script:     types.StringValue("console.log('new script');\n"),
			ScriptType: types.StringValue("playwright"),
		},
	}

	state := &testmonitorresource.PlaywrightTestResourceModel{
		BaseTestResourceModel: modelresource.BaseTestResourceModel{
			DivisionID: types.Int64Value(123),
			ProductID:  types.Int64Value(456),
			Name:       types.StringValue("Identity Portal"),
			Monitor:    types.StringValue("edge"),
		},
		TestScriptResourceModel: modelresource.TestScriptResourceModel{
			Script:     types.StringValue("console.log('old script');\n"),
			ScriptType: types.StringValue("playwright"),
		},
	}

	patches := r.generatePatchDocuments(plan, state, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("unexpected diagnostics: %v", resp.Diagnostics.Errors())
	}

	if len(patches) != 1 {
		t.Fatalf("expected one patch document, got %d: %v", len(patches), patches)
	}

	if !strings.Contains(patches[0], "\"path\":\"/testRequestData\"") {
		t.Fatalf("expected testRequestData patch, got %s", patches[0])
	}

	if !strings.Contains(patches[0], "console.log('new script');") {
		t.Fatalf("expected updated script in patch, got %s", patches[0])
	}
	if !strings.Contains(patches[0], "\"name\":\"playwright\"") {
		t.Fatalf("expected script type in patch, got %s", patches[0])
	}
}