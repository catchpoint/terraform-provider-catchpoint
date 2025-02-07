package catchpoint

const (
	catchpointTestURIProd     = "https://io.catchpoint.com/api/v3.3/tests"
	catchpointTestURIStage    = "https://iostage.catchpoint.com/api/v3.3/tests"
	catchpointTestURIQa       = "https://ioqa.catchpoint.com/api/v3.3/tests"
	catchpointProductURIProd  = "https://io.catchpoint.com/api/v3.3/products"
	catchpointProductURIStage = "https://iostage.catchpoint.com/api/v3.3/products"
	catchpointProductURIQa    = "https://ioqa.catchpoint.com/api/v3.3/products"
	catchpointFolderURIProd   = "https://io.catchpoint.com/api/v3.3/folders"
	catchpointFolderURIStage  = "https://iostage.catchpoint.com/api/v3.3/folders"
	catchpointFolderURIQa     = "https://ioqa.catchpoint.com/api/v3.3/folders"
)

var catchpointTestURI = "https://io.catchpoint.com/api/v3.3/tests"
var catchpointProductURI = "https://io.catchpoint.com/api/v3.3/products"
var catchpointFolderURI = "https://io.catchpoint.com/api/v3.3/folders"
var catchpointEnvironment string

func setTestUriByEnv(environment string) {

	switch environment {
	case "prod", "":
		catchpointTestURI = catchpointTestURIProd
		catchpointProductURI = catchpointProductURIProd
		catchpointFolderURI = catchpointFolderURIProd
	case "stage":
		catchpointTestURI = catchpointTestURIStage
		catchpointProductURI = catchpointProductURIStage
		catchpointFolderURI = catchpointFolderURIStage
	case "qa":
		catchpointTestURI = catchpointTestURIQa
		catchpointProductURI = catchpointProductURIQa
		catchpointFolderURI = catchpointFolderURIQa
	default:
		catchpointTestURI = catchpointTestURIProd
		catchpointProductURI = catchpointProductURIProd
		catchpointFolderURI = catchpointFolderURIProd
	}
}

func setEnvVariable(environment string) {
	catchpointEnvironment = environment
}
