package internal

import "strings"

type Config struct {
	APIToken    string
	LogJSON     bool
	Environment string
}

func NewConfig(apiToken string, logJSON bool, cpEnvironment string) *Config {
	return &Config{
		APIToken:    apiToken,
		LogJSON:     logJSON,
		Environment: cpEnvironment,
	}
}

const (
	catchpointURIProd      = "https://io.catchpoint.com/api"
	catchpointURIStage     = "https://iostage.catchpoint.com/api"
	catchpointURIQA        = "https://ioqa.catchpoint.com/api"
	catchpointTestsPath    = "/tests"
	catchpointProductsPath = "/products"
	catchpointFoldersPath  = "/folders"
)

// Global variables. Default to prod.
var CatchpointTestURI string
var CatchpointProductURI string
var CatchpointFolderURI string
var CatchpointEnvironment string

func SetEnvironment(environment string, apiVersion string) {
	switch strings.ToLower(environment) {
	case "prod", "":
		CatchpointTestURI = catchpointURIProd + "/" + apiVersion + catchpointTestsPath
		CatchpointProductURI = catchpointURIProd + "/" + apiVersion + catchpointProductsPath
		CatchpointFolderURI = catchpointURIProd + "/" + apiVersion + catchpointFoldersPath
	case "stage":
		CatchpointTestURI = catchpointURIStage + "/" + apiVersion + catchpointTestsPath
		CatchpointProductURI = catchpointURIStage + "/" + apiVersion + catchpointProductsPath
		CatchpointFolderURI = catchpointURIStage + "/" + apiVersion + catchpointFoldersPath
	case "qa":
		CatchpointTestURI = catchpointURIQA + "/" + apiVersion + catchpointTestsPath
		CatchpointProductURI = catchpointURIQA + "/" + apiVersion + catchpointProductsPath
		CatchpointFolderURI = catchpointURIQA + "/" + apiVersion + catchpointFoldersPath
	}

	CatchpointEnvironment = environment
}
