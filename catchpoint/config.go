package catchpoint

const (
	catchpointTestURIProd  = "https://io.catchpoint.com/api/v2/tests"
	catchpointTestURIStage = "https://iostage.catchpoint.com/api/v2/tests"
	catchpointTestURIQa    = "https://ioqa.catchpoint.com/api/v2/tests"
)

var catchpointTestURI = "https://io.catchpoint.com/api/v2/tests"

func setTestUriByEnv(environment string) {

	switch environment {
	case "prod", "":
		catchpointTestURI = catchpointTestURIProd
	case "stage":
		catchpointTestURI = catchpointTestURIStage
	case "qa":
		catchpointTestURI = catchpointTestURIQa
	default:
		catchpointTestURI = catchpointTestURIProd
	}
}
