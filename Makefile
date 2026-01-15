.PHONY: build build-windows test clean document
VERSION ?= 2.0.0

# Build will only work *if* the tests pass, this is to ensure that the provider is in a good state before building.
# The output will be a binary named terraform-provider-catchpoint_v$(VERSION) (as expected by terraform override).
build:
	go test ./...
	go build -ldflags "-X main.version=$(VERSION)" -o terraform-provider-catchpoint_v$(VERSION)

build-windows:
	go test ./...
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o terraform-provider-catchpoint_v$(VERSION).exe

test:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o ./coverage.html

clean:
	rm -f terraform-provider-catchpoint*

document:
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs --provider-name=catchpoint --examples-dir=./examples

snapshot:
	goreleaser build --snapshot --clean

release:
	goreleaser release --skip=publish --clean

vuln-check:
	govulncheck ./...