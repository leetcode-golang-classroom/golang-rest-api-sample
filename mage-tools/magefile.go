//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = Build

// clean the build binary
func Clean() error {
	return sh.Rm("bin")
}

// Creates the binary in the current directory.
func Build() error {
	mg.Deps(Clean)
	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}
	return sh.Run("go", "build", "-o", "./bin/server", "./cmd/main.go")
}

// start the server
func Launch() error {
	mg.Deps(Build)
	err := sh.RunV("./bin/server")
	if err != nil {
		return err
	}
	return nil
}

// run the test
func Test() error {
	err := sh.RunV("go", "test", "-v", "./...")
	if err != nil {
		return err
	}
	return nil
}
