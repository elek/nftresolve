//go:build mage
// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

//Test executes all unit and integration tests
func Test() error {
	err := sh.RunV("go", "test", "./...")
	return err
}

//Test executes all unit and integration tests
func Coverage() error {
	fmt.Println("Executing tests and generate coverate information")
	err := sh.RunV("go", "test", "-coverprofile=/tmp/coverage.out", "./...")
	if err != nil {
		return err
	}
	return sh.RunV("go", "tool", "cover", "-html=/tmp/coverage.out", "-o", "coverage.html")
}

//Lint executes all the linters with golangci-lint
func Lint() error {
	return sh.RunV("golangci-lint", "run")
}

//Build do a standard go build
func Build() error {
	return sh.RunV("go", "build", "-o", "nftresolve")
}

//Format reformat code automatically
func Format() error {
	err := sh.RunV("gofmt", "-w", ".")
	if err != nil {
		return err
	}
	return sh.RunV("goimports", "-w", ".")

}

//GenVendor updates `./vendor` dependencies
func GenVendor() error {
	err := sh.RunV("go", "mod", "tidy")
	if err != nil {
		return err
	}
	return sh.RunV("go", "mod", "vendor")
}

//GenBuild re-generates `./build` helper binary
func GenBuild() error {
	envs := map[string]string{
		"CGO_ENABLED": "0",
		"GOOS":        "linux",
		"GOARCH":      "amd64",
	}
	return sh.RunWithV(envs, "mage", "-compile", "build")

}
