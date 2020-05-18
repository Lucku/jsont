// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Build

var name = "jsont"

// A build step that requires additional params, or platform specific steps for example
func Build() error {

	mg.Deps(InstallDeps)
	fmt.Println("Building...")

	filename := name

	if runtime.GOOS == "windows" {
		filename += ".exe"
	}

	return sh.RunV("go", "build", "-ldflags=\"-s -w\"", "-o", filename, ".")
}

// A custom install step if you need your bin someplace other than go/bin
func Install() error {
	mg.Deps(Build, Test)
	fmt.Println("Installing...")

	gobin, err := sh.Output("go", "env", "GOBIN")

	if err != nil {
		return fmt.Errorf("can't determine GOBIN: %v", err)
	}

	return os.Rename(filepath.Join(".", name), filepath.Join(gobin, name))
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	return sh.RunV("go", "get", "-u", "./...")
}

func Test() error {
	fmt.Println("Running tests")
	return sh.RunV("go", "test", "./...")
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll(name)
}

// tag returns the git tag for the current branch or "" if none.
func tag() string {
	s, _ := sh.Output("git", "describe", "--tags")
	return s
}
