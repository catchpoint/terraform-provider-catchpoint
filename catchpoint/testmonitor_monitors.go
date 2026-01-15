package catchpoint

import (
	testexpand "catchpoint-provider/internal/expand/testmonitor"
	testmonitorResource "catchpoint-provider/internal/models/resource/testmonitor"
	testSchema "catchpoint-provider/internal/schema/testmonitor"
	testTransform "catchpoint-provider/internal/transform/testmonitor"
	cptypes "catchpoint-provider/internal/types"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func NewAPITestResource() resource.Resource {
	return &TestResource[*testmonitorResource.APITestResourceModel]{
		testType:      cptypes.APIType,
		schemaBuilder: testSchema.BuildAPITestSchema,
		expandFunc:    testexpand.ExpandTestConfig[*testmonitorResource.APITestResourceModel],
		transformFunc: testTransform.JSONToTerraformTest[*testmonitorResource.APITestResourceModel],
	}
}

func NewBGPTestResource() resource.Resource {
	return &TestResource[*testmonitorResource.BGPTestResourceModel]{
		testType:      cptypes.BGPType,
		schemaBuilder: testSchema.BuildBGPTestSchema,
		expandFunc:    testexpand.ExpandTestConfig[*testmonitorResource.BGPTestResourceModel],
		transformFunc: testTransform.JSONToTerraformTest[*testmonitorResource.BGPTestResourceModel],
	}
}

func NewDNSTestResource() resource.Resource {
	return &TestResource[*testmonitorResource.DNSTestResourceModel]{
		testType:      cptypes.DNSType,
		schemaBuilder: testSchema.BuildDNSTestSchema,
		expandFunc:    testexpand.ExpandTestConfig[*testmonitorResource.DNSTestResourceModel],
		transformFunc: testTransform.JSONToTerraformTest[*testmonitorResource.DNSTestResourceModel],
	}
}

func NewSSLTestResource() resource.Resource {
	return &TestResource[*testmonitorResource.SSLTestResourceModel]{
		testType:      cptypes.SSLType,
		schemaBuilder: testSchema.BuildSSLTestSchema,
		expandFunc:    testexpand.ExpandTestConfig[*testmonitorResource.SSLTestResourceModel],
		transformFunc: testTransform.JSONToTerraformTest[*testmonitorResource.SSLTestResourceModel],
	}
}

func NewPingTestResource() resource.Resource {
	return &TestResource[*testmonitorResource.PingTestResourceModel]{
		testType:      cptypes.PingType,
		schemaBuilder: testSchema.BuildPingTestSchema,
		expandFunc:    testexpand.ExpandTestConfig[*testmonitorResource.PingTestResourceModel],
		transformFunc: testTransform.JSONToTerraformTest[*testmonitorResource.PingTestResourceModel],
	}
}

func NewPlaywrightTestResource() resource.Resource {
	return &TestResource[*testmonitorResource.PlaywrightTestResourceModel]{
		testType:      cptypes.PlaywrightType,
		schemaBuilder: testSchema.BuildPlaywrightTestSchema,
		expandFunc:    testexpand.ExpandTestConfig[*testmonitorResource.PlaywrightTestResourceModel],
		transformFunc: testTransform.JSONToTerraformTest[*testmonitorResource.PlaywrightTestResourceModel],
	}
}

func NewPuppeteerTestResource() resource.Resource {
	return &TestResource[*testmonitorResource.PuppeteerTestResourceModel]{
		testType:      cptypes.PuppeteerType,
		schemaBuilder: testSchema.BuildPuppeteerTestSchema,
		expandFunc:    testexpand.ExpandTestConfig[*testmonitorResource.PuppeteerTestResourceModel],
		transformFunc: testTransform.JSONToTerraformTest[*testmonitorResource.PuppeteerTestResourceModel],
	}
}

func NewTracerouteTestResource() resource.Resource {
	return &TestResource[*testmonitorResource.TracerouteTestResourceModel]{
		testType:      cptypes.TracerouteType,
		schemaBuilder: testSchema.BuildTracerouteTestSchema,
		expandFunc:    testexpand.ExpandTestConfig[*testmonitorResource.TracerouteTestResourceModel],
		transformFunc: testTransform.JSONToTerraformTest[*testmonitorResource.TracerouteTestResourceModel],
	}
}

func NewTransactionTestResource() resource.Resource {
	return &TestResource[*testmonitorResource.TransactionTestResourceModel]{
		testType:      cptypes.TransactionType,
		schemaBuilder: testSchema.BuildTransactionTestSchema,
		expandFunc:    testexpand.ExpandTestConfig[*testmonitorResource.TransactionTestResourceModel],
		transformFunc: testTransform.JSONToTerraformTest[*testmonitorResource.TransactionTestResourceModel],
	}
}

func NewWebTestResource() resource.Resource {
	return &TestResource[*testmonitorResource.WebTestResourceModel]{
		testType:      cptypes.WebType,
		schemaBuilder: testSchema.BuildWebTestSchema,
		expandFunc:    testexpand.ExpandTestConfig[*testmonitorResource.WebTestResourceModel],
		transformFunc: testTransform.JSONToTerraformTest[*testmonitorResource.WebTestResourceModel],
	}
}

// provider.go calls this to get all resources our provider can manage.
func listResources() []func() resource.Resource {
	return []func() resource.Resource{
		// Product and Folder are not test types so their functions are slightly different.
		NewProductResource,
		NewFolderResource,
		// The rest are all test types so they use the generic TestResource struct with type parameters and
		// schemaBuilders, expandFuncs, and transformFuncs specific to each test type.
		NewAPITestResource,
		NewBGPTestResource,
		NewDNSTestResource,
		NewSSLTestResource,
		NewPingTestResource,
		NewPlaywrightTestResource,
		NewPuppeteerTestResource,
		NewTracerouteTestResource,
		NewTransactionTestResource,
		NewWebTestResource,
	}
}
