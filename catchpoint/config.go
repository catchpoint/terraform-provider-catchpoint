package catchpoint

const (
	catchpointTestURIProd     = "https://io.catchpoint.com/api/v2/tests"
	catchpointTestURIStage    = "https://iostage.catchpoint.com/api/v2/tests"
	catchpointTestURIQa       = "https://ioqa.catchpoint.com/api/v2/tests"
	catchpointProductURIProd  = "https://io.catchpoint.com/api/v2/products"
	catchpointProductURIStage = "https://iostage.catchpoint.com/api/v2/products"
	catchpointProductURIQa    = "https://ioqa.catchpoint.com/api/v2/products"
)

var catchpointTestURI = "https://io.catchpoint.com/api/v2/tests"
var catchpointProductURI = "https://io.catchpoint.com/api/v2/products"

func setTestUriByEnv(environment string) {

	switch environment {
	case "prod", "":
		catchpointTestURI = catchpointTestURIProd
		catchpointProductURI = catchpointProductURIProd
	case "stage":
		catchpointTestURI = catchpointTestURIStage
		catchpointProductURI = catchpointProductURIStage
	case "qa":
		catchpointTestURI = catchpointTestURIQa
		catchpointProductURI = catchpointProductURIQa
	default:
		catchpointTestURI = catchpointTestURIProd
		catchpointProductURI = catchpointProductURIProd
	}
}
