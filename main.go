package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"catchpoint-provider/catchpoint"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs --provider-name=catchpoint

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary.
	version string = "dev"

	// goreleaser can pass other information to the main package, such as the specific commit
	// https://goreleaser.com/cookbooks/using-main.version/
)

func main() {
	var debug bool
	var versionFlag bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.BoolVar(&versionFlag, "version", false, "print the provider version and exit")
	flag.Parse()

	if versionFlag {
		fmt.Println(version)
		return
	}

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/catchpoint/catchpoint",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), catchpoint.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
