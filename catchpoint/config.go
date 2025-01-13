package catchpoint

const (
	catchpointTestURIProd     = "https://io.catchpoint.com/api/v2/tests"
	catchpointTestURIStage    = "https://iostage.catchpoint.com/api/v2/tests"
	catchpointTestURIQa       = "https://ioqa.catchpoint.com/api/v2/tests"
	catchpointProductURIProd  = "https://io.catchpoint.com/api/v2/products"
	catchpointProductURIStage = "https://iostage.catchpoint.com/api/v2/products"
	catchpointProductURIQa    = "https://ioqa.catchpoint.com/api/v2/products"
	catchpointFolderURIProd   = "https://io.catchpoint.com/api/v3.3/folders"
	catchpointFolderURIStage  = "https://iostage.catchpoint.com/api/v3.3/folders"
	catchpointFolderURIQa     = "https://ioqa.catchpoint.com/api/v3.3/folders"
)

var catchpointTestURI = "https://io.catchpoint.com/api/v2/tests"
var catchpointProductURI = "https://io.catchpoint.com/api/v2/products"
var catchpointFolderURI = "https://io.catchpoint.com/api/v3.3/folders"

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
