//go:build tools

package tools

//This file ensures go mod tidy does not remove tfplugindocs from the dependencies on its next run.
import (
	// document generation
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
)
